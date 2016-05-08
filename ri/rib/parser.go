/* rigo/ri/rib/parser.go */
package rib

import (
	"io"
)

const (
	DefaultBufferSize int = 256
)

type TokenType byte

const (
	Tokeniser TokenType = 0
	Content   TokenType = 1
) 



/* tokeniser -> lexer -> parser ~~> run through ri */

type Token struct {
	Word string
	Line int
	Pos  int

	Type TokenType

	/* TODO: add lexical information here */
	/* TODO: add parser information here */
}

var EmptyToken = Token{Word:"",Line:-1,Pos:-1,Type:Tokeniser}

type TokenWriter interface {
	Write(Token)
}

type TokenReader interface {
	Read() (Token,error)
}

func Tokenise(reader io.Reader,writer TokenWriter) error {

	buf := make([]byte,DefaultBufferSize)
	line := 0
	pos := 0
	for {
		n,err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		word := ""

		for _,c := range buf[:n] {
			pos ++
			
			if c == ' ' || c == '\t' {
				if len(word) > 0 {
					writer.Write(Token{Word:word,Line:line,Pos:pos,Type:Content})
					writer.Write(Token{Word:"_space_",Line:line,Pos:pos,Type:Tokeniser})
					word = ""
				}
				continue
			}
			if c == '\n' {
				line++
				pos = 0
				if len(word) > 0 {
					writer.Write(Token{Word:word,Line:line,Pos:pos,Type:Content})
					writer.Write(Token{Word:"_newline_",Line:line,Pos:pos,Type:Tokeniser})
					word = ""
				}
				continue
			}
			if c == '[' || c == ']' {
				if len(word) > 0 {
					writer.Write(Token{Word:word,Line:line,Pos:pos,Type:Content})
					writer.Write(Token{Word:"_space_",Line:line,Pos:pos,Type:Tokeniser})
					word = ""
				}
				writer.Write(Token{Word:string(c),Line:line,Pos:pos,Type:Content})
				writer.Write(Token{Word:"_space_",Line:line,Pos:pos,Type:Tokeniser})
				continue
			}
						
			word += string(c)
		}
		if len(word) > 0 {
			writer.Write(Token{Word:word,Line:line,Pos:pos,Type:Content})
		}

		writer.Write(Token{Word:"_block_",Line:line,Pos:pos,Type:Tokeniser})
	}

	return nil
}

func Lexer(reader TokenReader,writer TokenWriter) error {

	for {
		token,err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		if token.Type == Tokeniser {


		}

		if token.Type == Content {



		}
	}
	return nil
}

