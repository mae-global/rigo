/* rigo/ri/rib/parser.go */
package rib

import (
	"fmt"
	"io"
	"strconv"
)

type BloomFilterer interface {
	IsMember(...string) bool
}

const (
	DefaultBufferSize int = 512
)

type TokenType byte

const (
	Tokeniser TokenType = 0
	Content   TokenType = 1
)

type TokenLex byte

func (l TokenLex) String() string {
	switch l {
	case Command:
		return "command"
		break
	case ArgToken:
		return "token"
		break
	case ArgOp:
		return "op"
		break
	}
	return "unknown"
}

const (
	Unknown  TokenLex = 0
	Command  TokenLex = 1
	ArgToken TokenLex = 2
	ArgOp    TokenLex = 3
)

/* tokeniser -> lexer -> parser ~~> run through ri */

type Token struct {
	Word string
	Line int
	Pos  int

	Type   TokenType
	RiType string

	Lex TokenLex

	Error error
	/* TODO: add lexical information here */
	/* TODO: add parser information here */
}

var EmptyToken = Token{Word: "", Line: -1, Pos: -1, Type: Tokeniser}

type TokenWriter interface {
	Write(Token)
}

type TokenReader interface {
	Read() (Token, error)
}

func Tokenise(reader io.Reader, writer TokenWriter) error {

	buf := make([]byte, DefaultBufferSize)
	line := 0
	pos := 0
	word := ""
	withinliteral := false

	for {
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		for _, c := range buf[:n] {
			pos++

			if c == '"' {
				if len(word) > 0 {
					writer.Write(Token{Word: word, Line: line, Pos: pos, Type: Content})
					word = ""
				}

				withinliteral = !withinliteral
				if !withinliteral {
					writer.Write(Token{Word: "_end-lit_", Line: line, Pos: pos, Type: Tokeniser})
				} else {
					writer.Write(Token{Word: "_begin-lit_", Line: line, Pos: pos, Type: Tokeniser})
				}

				continue
			}

			if c == '#' && !withinliteral {
				if len(word) > 0 {
					writer.Write(Token{Word: word, Line: line, Pos: pos, Type: Content})
					word = ""
				}
				writer.Write(Token{Word: "_comment_", Line: line, Pos: pos, Type: Tokeniser})
				continue
			}

			if c == ' ' || c == '\t' && !withinliteral {
				if len(word) > 0 {
					writer.Write(Token{Word: word, Line: line, Pos: pos, Type: Content})
					writer.Write(Token{Word: "_space_", Line: line, Pos: pos, Type: Tokeniser})
					word = ""
				}
				continue
			}
			if c == '\n' && !withinliteral {
				line++
				pos = 0
				if len(word) > 0 {
					writer.Write(Token{Word: word, Line: line, Pos: pos, Type: Content})
					writer.Write(Token{Word: "_newline_", Line: line, Pos: pos, Type: Tokeniser})
					word = ""
				}
				continue
			}
			if c == '[' || c == ']' && !withinliteral {
				if len(word) > 0 {
					writer.Write(Token{Word: word, Line: line, Pos: pos, Type: Content})
					writer.Write(Token{Word: "_space_", Line: line, Pos: pos, Type: Tokeniser})
					word = ""
				}
				writer.Write(Token{Word: string(c), Line: line, Pos: pos, Type: Content})
				writer.Write(Token{Word: "_space_", Line: line, Pos: pos, Type: Tokeniser})
				continue
			}

			word += string(c)
		}
		//	if len(word) > 0 {
		//	writer.Write(Token{Word:word,Line:line,Pos:pos,Type:Content})
		//}

		writer.Write(Token{Word: "_block_", Line: line, Pos: pos, Type: Tokeniser})
	}
	if len(word) > 0 {
		writer.Write(Token{Word: word, Line: line, Pos: pos, Type: Content})
	}

	return nil
}

func Lexer(reader TokenReader, writer TokenWriter, filter BloomFilterer) error {

	iscomment := false
	isliteral := false

	for {
		token, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if token.Type == Tokeniser {

			switch token.Word {
			case "_space_":

				break
			case "_newline_":
				iscomment = false
				break
			case "_begin-lit_":
				isliteral = true
				break
			case "_end-lit_":
				isliteral = false
				break
			case "_block_":

				break
			case "_comment_":
				iscomment = true
				break
			}

			//	fmt.Printf("iscomment=%v, isliteral=%v\n",iscomment,isliteral)
			writer.Write(token)

		}

		if token.Type == Content {
			/* check if a member of ri */
			if filter.IsMember(token.Word) {
				//fmt.Printf("Command found -- %s\n",token.Word)

				token.Lex = Command
				token.RiType = "func"

			} else {
				token.Lex = ArgOp

				if iscomment {
					token.RiType = "comment"
				} else {

					switch token.Word {
					case "[":
						token.RiType = "array_begin"
						break
					case "]":
						token.RiType = "array_end"
						break
					default:
						if isliteral {
							token.RiType = "token"
						} else {
							token.RiType = "number"
						}
						break
					}
					/* -- */
				}
			}
			writer.Write(token)
		}
	}
	return nil
}

func Parser(reader TokenReader, writer TokenWriter) error {

	/* TODO: add a function lookup -- check the variables, all the types need to be passed on etc */

	currentfunc := ""
	variables := 0
	inarray := false

	for {
		token, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if token.Type == Tokeniser {
			continue
		}

		/* start of a new Ri function call */
		if token.RiType == "func" {
			if variables > 0 && currentfunc != "" {
				writer.Write(Token{Word: fmt.Sprintf("%d", variables), Line: 0, Pos: 0, Type: Tokeniser, RiType: "counter"})
			}

			currentfunc = token.Word
			variables = 0
		}

		/* pass all the arguments in the function call */
		if token.Type == Content {

			switch token.RiType {
			case "number":
				if _, err := strconv.ParseFloat(token.Word, 64); err != nil {
					token.RiType = "number"
					token.Error = err
				} else {
					token.RiType = "float"
				}
				if !inarray {
					variables++
				}
				break
			case "array_begin":
				inarray = true
				break
			case "array_end":
				inarray = false
				variables++
				break
			}
		}

		writer.Write(token)
	}
	writer.Write(Token{Word: fmt.Sprintf("%d", variables), Line: 0, Pos: 0, Type: Tokeniser, RiType: "counter"})

	return nil
}
