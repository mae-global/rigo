package ri

import (
	"fmt"
)

type ErrorIgnore int

func (e ErrorIgnore) Name() RtErrorHandler { return RtErrorHandler("ignore") }
func (e ErrorIgnore) Handle(err RtError) error { return nil }

type ErrorPrint int

func (e ErrorPrint) Name() RtErrorHandler { return RtErrorHandler("print") }
func (e ErrorPrint) Handle(err RtError) error {
	fmt.Printf("%s\n",err)
	return err.Error()
}

type ErrorAbort int 

func (e ErrorAbort) Name() RtErrorHandler { return RtErrorHandler("abort") }
func (e ErrorAbort) Handle(err RtError) error {
	fmt.Printf("%s\n",err)
	panic(err.Msg)
	return err.Error()
}

/* ErrorHandler */
func (r *Ri) ErrorHandler(handler RtErrorHandlerFuncer) error {
	r.errorhandler = handler
	return r.writef("ErrorHandler",handler.Name())
}

