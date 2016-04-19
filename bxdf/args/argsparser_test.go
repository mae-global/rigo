package args 

import (
	"testing"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"

	. "github.com/mae-global/rigo/ri"
	
)

func Test_Parser(t *testing.T) {

	Convey("Parser",t,func() {

		args,err := ParseArgs([]byte(constant_args))
		So(err,ShouldBeNil)
		So(args,ShouldNotBeNil)

		So(len(args.Params),ShouldEqual,2)
		So(args.Params[0].Label,ShouldEqual,"Emit Color")
		So(args.Params[0].Name,ShouldEqual,"emitColor")

		So(args.Params[1].Label,ShouldEqual,"Presence")
		So(args.Params[1].Name,ShouldEqual,"presence")
	})

	Convey("Parse Bxdf Spec",t,func() {

		bxdf,err := Parse("PxrConstant",[]byte(constant_args))
		So(err,ShouldBeNil)
		So(bxdf,ShouldNotBeNil)

		names := bxdf.Names()
		So(names,ShouldNotBeNil)
		So(len(names),ShouldEqual,2)
		So(names[0],ShouldEqual,RtToken("emitColor"))
		So(names[1],ShouldEqual,RtToken("presence"))

		fmt.Printf("names = %v\n",names)

		names = bxdf.NamesSpec()
		So(names,ShouldNotBeNil)
		So(len(names),ShouldEqual,2)
		So(names[0],ShouldEqual,RtToken("color emitColor"))
		So(names[1],ShouldEqual,RtToken("float presence"))

		fmt.Printf("namesSpec = %v\n",names)
		
		r := bxdf.Value("bad")
		So(r,ShouldBeNil)

		r = bxdf.Value("emitColor")
		So(r,ShouldNotBeNil)
		c,ok := r.(RtColor)
		So(ok,ShouldBeTrue)
		So(c.Equal(RtColor{1,1,1}),ShouldBeTrue)

		r = bxdf.Value("presence")
		So(r,ShouldNotBeNil)
		f,ok := r.(RtFloat)
		So(ok,ShouldBeTrue)
		So(f,ShouldEqual,RtFloat(1))

		err = bxdf.SetValue("bad",RtFloat(0.1))
		So(err,ShouldNotBeNil)
		So(err.Error(),ShouldEqual,"Unknown parameter \"bad\"")

		err = bxdf.SetValue("emitColor",RtFloat(0.1))
		So(err,ShouldNotBeNil)
		So(err.Error(),ShouldEqual,"Type mismatch, setting with \"float\", wants \"color\"")

		So(bxdf.SetValue("presence",RtFloat(0.12345)),ShouldBeNil)

		r = bxdf.Value("presence")
		So(r,ShouldNotBeNil)
		f,ok = r.(RtFloat)
		So(ok,ShouldBeTrue)
		So(f,ShouldEqual,RtFloat(0.12345))
	

		widget := bxdf.Widget("presence")
		So(widget,ShouldNotBeNil)

		So(widget.Name(),ShouldEqual,RtToken("presence"))
		So(widget.NameSpec(),ShouldEqual,RtToken("float presence"))
		So(widget.Label(),ShouldEqual,RtString("Presence"))
		So(widget.Help(),ShouldEqual,RtString("help text was here"))
		r = widget.GetValue()
		So(r,ShouldNotBeNil)
		f,ok = r.(RtFloat)
		So(ok,ShouldBeTrue)
		So(f,ShouldEqual,RtFloat(0.12345))
	
		fw,ok := widget.(*RtFloatWidget)
		So(ok,ShouldBeTrue)
		So(fw,ShouldNotBeNil)

		So(fw.Set(4.5),ShouldBeNil)
		So(fw.Value(),ShouldEqual,RtFloat(4.5))


		widget = widget.Next()
		So(widget,ShouldNotBeNil)

		So(widget.Name(),ShouldEqual,RtToken("emitColor"))
		So(widget.NameSpec(),ShouldEqual,RtToken("color emitColor"))
		So(widget.Label(),ShouldEqual,RtString("Emit Color"))
	})

	Convey("Parse PxrConstant.args",t,func() {

		constant,err := ParseArgsFile("PxrConstant")
		So(err,ShouldBeNil)
		So(constant,ShouldNotBeNil)
		names := constant.Names()
		So(len(names),ShouldEqual,2)
	})

	Convey("Parse PxrDiffuse.args",t,func() {

		diffuse,err := ParseArgsFile("PxrDiffuse")
		So(err,ShouldBeNil)
		So(diffuse,ShouldNotBeNil)

		w := diffuse.Widget("bumpNormal")
		So(w,ShouldNotBeNil)
		wn,ok := w.(*RtNormalWidget)
		So(ok,ShouldBeTrue)
		So(wn,ShouldNotBeNil)
		So(wn.Value(),ShouldEqual,RtNormal{0,0,0})

	})
}


const constant_args = `
<args format="1.0">
    <shaderType>
        <tag value="bxdf"/>
    </shaderType>
    <param label="Emit Color" name="emitColor" type="color" 
           default="1. 1. 1." widget="color">
        <tags>
            <tag value="color"/>
        </tags>
    </param>
    <param label="Presence" 
           name="presence" 
           type="float" 
           default="1" min="0" max="1"
           widget="default">
        <tags>
            <tag value="float"/>
        </tags>
        <help>
           help text was here
        </help>
    </param>
    <rfmdata nodeid="1053405" 
             classification="shader/surface:rendernode/RenderMan/bxdf:swatch/rmanSwatch"/>
</args>
`

