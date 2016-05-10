package ri

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/mae-global/rigo/ri/rib"
)

type RIBTokenIO struct {
	tokens   []rib.Token
	position int
}

func (w *RIBTokenIO) Write(t rib.Token) {
	if w.tokens == nil {
		w.tokens = make([]rib.Token, 0)
	}
	w.tokens = append(w.tokens, t)
}

func (w *RIBTokenIO) Read() (rib.Token, error) {
	if w.tokens == nil || len(w.tokens) == 0 {
		return rib.EmptyToken, io.EOF
	}

	if w.position >= len(w.tokens) {
		return rib.EmptyToken, io.EOF
	}
	t := w.tokens[w.position]
	w.position++
	return t, nil
}

func (w *RIBTokenIO) Print() string {
	out := ""
	for _, token := range w.tokens {
		tag := "content"
		if token.Type == rib.Tokeniser {
			tag = "tokeniser"
		}

		ritype := token.RiType
		if token.Error != nil {
			ritype = token.Error.Error()
		}
		out += fmt.Sprintf("%04d:%03d --%30s\t(%s)\tL:%10s\tRi:%10s\n",
			token.Line, token.Pos, token.Word, tag, token.Lex, ritype)
	}

	return out
}

func ParseString(content string, writer RterWriter) error {

	tw := new(RIBTokenIO)
	if err := rib.Tokenise(strings.NewReader(content), tw); err != nil {
		return err
	}

	tw1 := new(RIBTokenIO)
	if err := rib.Lexer(tw, tw1, RiBloomFilter()); err != nil {
		return err
	}

	tw2 := new(RIBTokenIO)
	if err := rib.Parser(tw1, tw2); err != nil {
		return err
	}

	/* create the lookup table for the function information */
	lookup := make(map[RtName]RtInt, 0)
	for i := 0; i < len(bloomFilterKeysData); i++ {
		name := RtName(bloomFilterKeysData[i])
		args := RtInt(RiArgumentsData[i])

		lookup[name] = args
	}

	fmt.Printf("\n%s\n\n", tw2.Print())

	currentfunc := ""
	args := make([]Rter, 0)
	tokens := make([]Rter, 0)
	values := make([]Rter, 0)
	params := -1 /* parameterlist starts at? */

	farray := make([]RtFloat, 0)
	isarray := false

	for {
		token, err := tw2.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if token.Type != rib.Content {
			continue
		}

		switch token.RiType {
		case "func":
			if currentfunc != "" { /* write to out with the current */

				/* FIXME: remove the debug printing */
				fmt.Printf("%s (%d) %d args, %d tokens & %d values\n",
					currentfunc, params, len(args), len(tokens), len(values))

				if writer != nil {
					if err := writer.WriteTo(RtName(currentfunc), args, tokens, values); err != nil {
						return err
					}
				}

				currentfunc = ""
				args = make([]Rter, 0)
				tokens = make([]Rter, 0)
				values = make([]Rter, 0)
				farray = make([]RtFloat, 0)
				params = -1

			}

			/* start new func, lookup the information required */
			currentfunc = token.Word
			if v, ok := lookup[RtName(currentfunc)]; ok {
				params = int(v)
			}
			break

		case "float":
			if f, err := strconv.ParseFloat(token.Word, 64); err != nil {
				return err
			} else {
				farray = append(farray, RtFloat(f))
				if !isarray {
					if len(args) >= params {
						values = append(values, RtFloatArray(farray))
					} else {
						args = append(args, RtFloatArray(farray))
					}
				}
			}
			break

		case "token":
			if len(args) >= params {
				tokens = append(tokens, RtToken(token.Word))
			} else {
				args = append(args, RtToken(token.Word))
			}
			break

		case "array_begin":
			farray = make([]RtFloat, 0)
			isarray = true
			break

		case "array_end":
			isarray = false
			if len(args) >= params {
				values = append(values, RtFloatArray(farray))
			} else {
				args = append(args, RtFloatArray(farray))
			}
			break
		}

	}

	return nil
}


