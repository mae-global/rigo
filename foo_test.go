/* rigo/foo_test.go */
package ri

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
)


func FooDefaultTxPipe() *Pipe {
	pipe := NewPipe()
	return pipe.Append(&PipeTimer{}).Append(&PipeToStats{}).Append(&PipeToFile{})
}



type FooTxPipeEnd struct {	
	operations []*FooTransactionOperation	
	
}

func (p FooTxPipeEnd) Name() string {
	return "default-tx-pipe-end"
}

func (p *FooTxPipeEnd) Write(name RtName,list []Rter,info Info) *Result {

	p.operations = append(p.operations,&FooTransactionOperation{name,list})
	return Done()
}

func (p *FooTxPipeEnd) String() string {
	return fmt.Sprintf("recorded %d operations",len(p.operations))
}






func NewFooTransaction(pipe *Pipe) *FooTransaction {
	
	if pipe == nil {
		pipe = FooDefaultTxPipe()
	}
	
	last := pipe.Last()

	add := false
	if last == nil {
		add = true
	} else {
		name := FooTxPipeEnd{}.Name()
		if last.Name() != name {
			add = true
		}
	}
		
	/* check the last pipe element, if not default-tx-pipe-end then add */
	if add {
		end := &FooTxPipeEnd{}
		end.operations = make([]*FooTransactionOperation,0)

		pipe.Append(end)
	} 		

	tx := &FooTransaction{}
	tx.pipe = pipe
	tx.Ri = New(tx.pipe,nil)
	return tx
}

type FooTransactionOperation struct {
	Name RtName
	List []Rter
}

type FooTransaction struct {
	pipe *Pipe
	*Ri
}

func Test_Foo(t *testing.T) {

	pipe := NewPipe()

	tx := NewFooTransaction(pipe)

	Convey("All Foo", t, func() {

		tx.Begin("foo.rib")

		tx.End()

		if end := pipe.GetByName(FooTxPipeEnd{}.Name()); end != nil {
			fmt.Printf("%s\n",end)
		}	else {
			fmt.Printf("FooTxPipeEnd error\n")
		}	

	})
}
