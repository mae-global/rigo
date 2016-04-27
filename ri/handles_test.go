/* rigo/ri/handles_test.go */
package ri

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"	
	"fmt"
)

func Test_NumberGenerators(t *testing.T) {

	Convey("Light Number Generator",t,func() {
	
		gen := NewLightNumberGenerator()
		So(gen,ShouldNotBeNil)

		h,err := gen.Generate()
		So(err,ShouldBeNil)
		So(h.String(),ShouldEqual,"\"0\"")

		So(gen.Example().String(),ShouldEqual,"\"0\"")

		for i := 1; i < 10; i++ {
			h,err := gen.Generate()
			So(err,ShouldBeNil)
			So(h.String(),ShouldEqual,fmt.Sprintf("\"%d\"",i))
		}
	})

	Convey("Object Number Generator",t,func() {

		gen := NewObjectNumberGenerator()
		So(gen,ShouldNotBeNil)
		
		h,err := gen.Generate()
		So(err,ShouldBeNil)
		So(h.String(),ShouldEqual,"\"0\"")

		So(gen.Example().String(),ShouldEqual,"\"0\"")

		for i := 1; i < 10; i++ {
			h,err := gen.Generate()
			So(err,ShouldBeNil)
			So(h.String(),ShouldEqual,fmt.Sprintf("\"%d\"",i))
		}
	})
}

func Test_PrefixNumberGenerators(t *testing.T) {

	Convey("Prefix Light Number Generator",t,func() {
	
		gen := NewPrefixLightNumberGenerator("light_")
		So(gen,ShouldNotBeNil)

		h,err := gen.Generate()
		So(err,ShouldBeNil)
		So(h.String(),ShouldEqual,"\"light_0\"")

		So(gen.Example().String(),ShouldEqual,"\"light_0\"")

		for i := 1; i < 10; i++ {
			h,err := gen.Generate()
			So(err,ShouldBeNil)
			So(h.String(),ShouldEqual,fmt.Sprintf("\"light_%d\"",i))
		}
	})

	Convey("Preix Object Number Generator",t,func() {

		gen := NewPrefixObjectNumberGenerator("object_")
		So(gen,ShouldNotBeNil)
		
		h,err := gen.Generate()
		So(err,ShouldBeNil)
		So(h.String(),ShouldEqual,"\"object_0\"")

		So(gen.Example().String(),ShouldEqual,"\"object_0\"")

		for i := 1; i < 10; i++ {
			h,err := gen.Generate()
			So(err,ShouldBeNil)
			So(h.String(),ShouldEqual,fmt.Sprintf("\"object_%d\"",i))
		}
	})
}


func Test_UniqueGenerators(t *testing.T) {

	Convey("Light Unique Generator",t,func() {
	
		gen := NewLightUniqueGenerator()
		So(gen,ShouldNotBeNil)

		h,err := gen.Generate()
		So(err,ShouldBeNil)
		So(h,ShouldNotBeEmpty)
		fmt.Printf("LightUnique : %s\n",h)

		So(gen.Example().String(),ShouldEqual,"\"61626364\"")

		all := make(map[RtLightHandle]int,0)

		for i := 1; i < 1000; i++ {
			h,err := gen.Generate()
			So(err,ShouldBeNil)
			if _,exists := all[h]; exists {
				So(fmt.Errorf("[%s] not unique",h),ShouldBeNil)
			}	
			all[h] = 1		
		}
	})

	Convey("Object Unique Generator",t,func() {

		gen := NewObjectUniqueGenerator()
		So(gen,ShouldNotBeNil)
		
		h,err := gen.Generate()
		So(err,ShouldBeNil)
		So(h,ShouldNotBeEmpty)
		fmt.Printf("ObjectUnique : %s\n",h)

		So(gen.Example().String(),ShouldEqual,"\"61626364\"")
	
		all := make(map[RtObjectHandle]int,0)

		for i := 1; i < 1000; i++ {
			h,err := gen.Generate()
			So(err,ShouldBeNil)
			if _,exists := all[h]; exists {
				So(fmt.Errorf("[%s] not unique",h),ShouldBeNil)
			}
			all[h] = 1 
			
		}
	})
}

func Test_PrefixUniqueGenerators(t *testing.T) {

	Convey("Prefix Light Unique Generator",t,func() {
	
		gen := NewPrefixLightUniqueGenerator("light_")
		So(gen,ShouldNotBeNil)

		h,err := gen.Generate()
		So(err,ShouldBeNil)
		So(h,ShouldNotBeEmpty)
	})

	Convey("Preix Object Number Generator",t,func() {

		gen := NewPrefixObjectUniqueGenerator("object_")
		So(gen,ShouldNotBeNil)
		
		h,err := gen.Generate()
		So(err,ShouldBeNil)
		So(h,ShouldNotBeEmpty)

	})
}





