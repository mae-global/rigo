package ri

type ArchiveWriter interface {
	Write([]byte) (int,error)
}


func (r *Ri) ArchiveBegin(id RtToken, parameterlist ...Rter) (ArchiveWriter,error) {
	
	aw,err := r.OpenRaw(id)
	if err != nil {
		return nil,err
	}

	list := []Rter{id,PARAMETERLIST}
	list = append(list,parameterlist...)	

	return aw,r.writef("ArchiveBegin",list...)
}

func (r *Ri) ArchiveEnd(id RtToken) error {

	if err := r.CloseRaw(id); err != nil {
		return err
	}

	return r.writef("ArchiveEnd",id)
}

func (r *Ri) ArchiveInstance(id RtToken,parameterlist ...Rter) error {

	list := []Rter{id,PARAMETERLIST}
	list = append(list,parameterlist...)

	return r.writef("ArchiveInstance",list...)
}
