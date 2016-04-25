/* rigo/handles/handles.go */
package handles

import (
	"crypto/rand"
	"io"
	"encoding/hex"
)

func read(n int) (string,error) {

	b := make([]byte,n)
	n,err := io.ReadFull(rand.Reader,b)
	if err != nil {
		return "",err
	}
	return hex.EncodeToString(b[:n]),nil
}

func readExample(n int) string {
	example := []byte("abcdefghijklnmopqrstuvw123456789") /* FIXME, should use a sha512 string instead */
	if n >= len(example) {
		n = len(example) - 1
	}
	return hex.EncodeToString(example[:n])
}
	
	
	










 








	


