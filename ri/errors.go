package ri


/* ErrorHandler */
func (r *Ri) ErrorHandler(handler RtErrorHandler) error {
	
	return r.writef("ErrorHandler",handler)
}

