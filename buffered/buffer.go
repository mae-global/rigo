/* rigo/buffered/buffer.go */
package buffered

import (
	"bytes"
)

type Buffered struct {
	buf *bytes.Buffer
}

func (buf *Buffered) Bytes() []byte {
	return buf.buf.Bytes()
}

func (buf *Buffered) Close() error {
	return nil
}

func (buf *Buffered) Write(content []byte) (int,error) {
	return buf.buf.Write(content)
}

func New() *Buffered {
	return &Buffered{bytes.NewBuffer(nil)}
}

