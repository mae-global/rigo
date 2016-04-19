package bxdf

import (
	. "github.com/mae-global/rigo/ri"
)
/* https://renderman.pixar.com/resources/current/RenderMan/devExamples.html */

type Widget interface {
	Name() RtToken
	NameSpec() RtToken
	Label() RtString
	SetValue(value Rter) error
	GetValue() Rter
	Help() RtString
	Bounds() (Rter,Rter) /* min and max if set */
	Default() error

	Next() Widget
	Prev() Widget
}

type Bxdfer interface {
	Write() []Rter
	Name() RtToken
	NodeId() RtToken
	Classifiation() RtString
	
	Widget(name RtToken) Widget 

	Names() []RtToken
	NamesSpec() []RtToken

	SetValue(name RtToken,value Rter) error 
	Value(name RtToken) Rter
}

/* old implemention ;;

type PxrDiffuse struct {
	DiffuseColor RtColor
	TransmissionColor RtColor
}

func (bxdf *PxrDiffuse) Write(ri *Ri) error {
	return ri.Bxdf("PxrDiffuse",RtToken("color diffuseColor"),bxdf.DiffuseColor,RtToken("color transmissionColor"),bxdf.TransmissionColor)
}

func NewPxrDiffuse() *PxrDiffuse {
	bxdf := &PxrDiffuse{}
	bxdf.DiffuseColor = RtColor{1,1,1}
	return bxdf
}


type PxrConstant struct {
	EmitColor RtColor
}

func (bxdf *PxrConstant) Write(ri *Ri) error {
	return ri.Bxdf("PxrConstant",RtToken("color emitColor"),bxdf.EmitColor)
}

func NewPxrConstant() *PxrConstant {
	bxdf := &PxrConstant{}
	bxdf.EmitColor = RtColor{1,1,1}
	return bxdf
}

type PxrASSurface struct {
	DiffuseColor RtColor
	SpecularColor RtColor
	Ks RtFloat
	Roughness RtFloat
	AnisoRatio RtFloat
}

func (bxdf *PxrASSurface) Write(ri *Ri) error {
	return ri.Bxdf("PxrASSurface",RtToken("color diffuseColor"),bxdf.DiffuseColor,RtToken("color specularColor"),bxdf.SpecularColor,
																RtToken("float ks"),bxdf.Ks,RtToken("float roughness"),bxdf.Roughness,RtToken("float anisoRatio"),bxdf.AnisoRatio)
}

func NewPxrASSurface() *PxrASSurface {
	bxdf := &PxrASSurface{}
	bxdf.DiffuseColor = RtColor{1,1,1}
	bxdf.SpecularColor = RtColor{1,1,1}
	bxdf.Ks = 0.04
	bxdf.Roughness = 0.001
	bxdf.AnisoRatio = 1.0
	return bxdf
}


type PxrSubsurface struct {
	Albedo RtColor
	DiffuseMeanFreePath RtColor
	UnitLength RtFloat
	IndirectAtSss RtFloat
}

func (bxdf *PxrSubsurface) Write(ri *Ri) error {
	return ri.Bxdf("PxrSubsurface",RtToken("color albedo"),bxdf.Albedo,RtToken("color diffuseMeanFreePath"),bxdf.DiffuseMeanFreePath,
																 RtToken("float unitLength"),bxdf.UnitLength,RtToken("float indirectAtSss"),bxdf.IndirectAtSss)
}


func NewPxrSubsurface() *PxrSubsurface {
	bxdf := &PxrSubsurface{}
	bxdf.Albedo = RtColor{0.830,0.791,0.753}
	bxdf.DiffuseMeanFreePath = RtColor{8.51,5.57,3.95}
	bxdf.UnitLength = 0.1
	bxdf.IndirectAtSss = 0 
	return bxdf
}

type PxrBeerGlass struct {
	Eta RtFloat
	Absorption RtColor
}

func (bxdf *PxrBeerGlass) Write(ri *Ri) error {
	return ri.Bxdf("PxrBeerGlass",RtToken("float eta"),bxdf.Eta,RtToken("color absorption"),bxdf.Absorption)
}

func NewPxrBeerGlass() *PxrBeerGlass {
	bxdf := &PxrBeerGlass{}
	bxdf.Eta = 1.5
	return bxdf
}

type PxrSmoothDielectric struct {
	DiffuseColor RtColor
	SpecularColor RtColor
}

func (bxdf *PxrSmoothDielectric) Write(ri *Ri) error {
	return ri.Bxdf("PxrSmoothDielectric",RtToken("color diffuseColor"),bxdf.DiffuseColor,RtToken("color specularColor"),bxdf.SpecularColor)
}

func NewPxrSmoothDielectric() *PxrSmoothDielectric {
	bxdf := &PxrSmoothDielectric{}
	bxdf.DiffuseColor = RtColor{0.15,0.15,0.45}
	bxdf.SpecularColor = RtColor{0.04,0.04,0.04}
	return bxdf
}


type PxrConductor struct {
	SpecularColor RtColor
	Roughness RtFloat
}

func (bxdf *PxrConductor) Write(ri *Ri) error {
	return ri.Bxdf("PxrConductor",RtToken("color specularColor"),bxdf.SpecularColor,RtToken("float roughness"),bxdf.Roughness)
}

func NewPxrConductor() *PxrConductor {
	bxdf := &PxrConductor{}
	bxdf.SpecularColor = RtColor{1,1,1}
	bxdf.Roughness = 0.1
	return bxdf
}

*/






