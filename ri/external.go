package ri

import (
	"fmt"
)

/* MakeTexture Convert an image in a standard picture file whise name is picturename */
func (r *Ri) MakeTexture(picturename, texturename RtString, swrap, twrap RtToken, filterfunc RtFilterFunc, swidth, twidth RtFloat, parameterlist ...Rter) error {

	var out = []Rter{picturename, texturename, swrap, twrap, filterfunc, swidth, twidth, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("MakeTexture", out...)
}

func (r *Ri) MakeTextureV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef("MakeTexture", out...)
}

/* MakeLatLongEnvironment */
func (r *Ri) MakeLatLongEnvironment(picturename, texturename RtString, filterfunc RtFilterFunc, swidth, twidth RtFloat, parameterlist ...Rter) error {

	var out = []Rter{picturename, texturename, filterfunc, swidth, twidth, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("MakeLatLongEnvironment", out...)
}

func (r *Ri) MakeLatLongEnvironmentV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef("MakeLatLongEnvironment", out...)
}

/* MakeCubeFaceEnviroment */
func (r *Ri) MakeCubeFaceEnvironment(px, nx, py, ny, pz, nz, texturename RtString, fov RtFloat, filterfunc RtFilterFunc, swidth, twidth RtFloat, parameterlist ...Rter) error {

	var out = []Rter{px, nx, py, ny, pz, nz, texturename, fov, filterfunc, swidth, twidth, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("MakeCubeFaceEnvironment", out...)
}

func (r *Ri) MakeCubeFaceEnvironmentV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef("MakeCubeFaceEnvironment", out...)
}

/* MakeShadow */
func (r *Ri) MakeShadow(picturename, texturename RtString, parameterlist ...Rter) error {

	var out = []Rter{picturename, texturename, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("MakeShadow", out...)
}

func (r *Ri) MakeShadowV(args, tokens, values []Rter) error {

	out := make([]Rter, 0)
	out = append(out, args...)
	out = append(out, PARAMETERLIST)
	out = append(out, Mix(tokens, values)...)

	return r.writef("MakeShadow", out...)
}

/* ArchiveRecord */
func (r *Ri) ArchiveRecord(typeof RtToken, format RtString, args ...interface{}) error {

	var err error

	switch string(typeof) {
	case "comment":

		err = r.writef("#", RtName(fmt.Sprintf(string(format), args...)))
		break
	case "structure":

		err = r.writef("##", RtName(fmt.Sprintf(string(format), args...)))
		break
	case "verbatim":

		err = r.writef("Verbatim", RtName(fmt.Sprintf(string(format), args...)))
		break
	}

	return err
}

func (r *Ri) ArchiveRecordV(args, tokens, values []Rter) error {
	/* FIXME */
	return nil
}

/* ReadArchive */
func (r *Ri) ReadArchive(name RtToken, callback RtArchiveCallbackFunc, parameterlist ...Rter) error {
	/* FIXME, take care of callback ? */
	var out = []Rter{name, PARAMETERLIST}
	out = append(out, parameterlist...)

	return r.writef("ReadArchive", out...)
}

func (r *Ri) ReadArchiveV(args, tokens, values []Rter) error {
	/* FIXME */
	return nil
}
