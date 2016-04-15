package ri

import (
	"fmt"
)

/* MakeTexture Convert an image in a standard picture file whise name is picturename */
func (r *Ri) MakeTexture(picturename, texturename RtString, swrap, twrap RtToken, filterfunc RtFilterFunc, swidth, twidth RtFloat, parameterlist ...Rter) error {

	var out = []Rter{picturename, texturename, swrap, twrap, filterfunc, swidth, twidth}
	out = append(out, parameterlist...)

	return r.writef("MakeTexture", out...)
}

/* MakeLatLongEnvironment */
func (r *Ri) MakeLatLongEnvironment(picturename, texturename RtString, filterfunc RtFilterFunc, swidth, twidth RtFloat, parameterlist ...Rter) error {

	var out = []Rter{picturename, texturename, filterfunc, swidth, twidth}
	out = append(out, parameterlist...)

	return r.writef("MakeLatLongEnvironment", out...)
}

/* MakeCubeFaceEnviroment */
func (r *Ri) MakeCubeFaceEnvironment(px, nx, py, ny, pz, nz, texturename RtString, fov RtFloat, filterfunc RtFilterFunc, swidth, twidth RtFloat, parameterlist ...Rter) error {

	var out = []Rter{px, nx, py, ny, pz, nz, texturename, fov, filterfunc, swidth, twidth}
	out = append(out, parameterlist...)

	return r.writef("MakeCubeFaceEnvironment", out...)
}

/* MakeShadow */
func (r *Ri) MakeShadow(picturename, texturename RtString, parameterlist ...Rter) error {

	var out = []Rter{picturename, texturename}
	out = append(out, parameterlist...)

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


/* ReadArchive */
func (r *Ri) ReadArchive(name RtToken, callback RtArchiveCallbackFunc, parameterlist ...Rter) error {

	var out = []Rter{name}
	out = append(out, parameterlist...)

	return r.writef("ReadArchive", out...)
}
