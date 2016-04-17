/* rigo/ri/rib/parser.go */
package rib

import (
	"fmt"
)

/* Test Parser */
func Parse(content []byte) error {

	/* tokenise the content into a stream of tokens : */
	token := ""
	start := 0
	for i,c := range content {
		if c == ' ' || c == '\t' { 
			if len(token) > 0 {
				fmt.Printf("token %d-%d [%s]\n",start,i,token)
			}	
			token = ""
			start = i
			continue
		}
		if c == '\n' {
			if len(token) > 0 {
				fmt.Printf("token %d-%d [%s]\n",start,i,token)
			}
			fmt.Printf("newline\n")
			token = ""
			start = i
			continue
		}

		token += string(c)
	}
	

	return nil
}
