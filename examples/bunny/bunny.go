/* rigo/examples/bunny/bunny.go
 * render the Standford Bunny just using RiGO 
 */
package main

import (
	"fmt"
	"os"
	"os/user"
	"time"
	"flag"

	. "github.com/mae-global/rigo/ri"
	. "github.com/mae-global/rigo/ris"

	"github.com/mae-global/rigo"
)

const (
	Author = rigo.Author
	Scene  = "Standford Rabbit and company"
)

func main() {

	optProduction := flag.Bool("production",false,"use production settings")
	optDevelopment := flag.Bool("development",false,"use development settings")
	optDebug := flag.Bool("debug",false,"use debug settings")
	optWidth := flag.Float64("width",2100,"width of output")
	optHeight := flag.Float64("height",900.0,"height of the output")
	optScale := flag.Float64("scale",0.5,"scale of the output, [0.0,1.0)")
	optPrettyPrint := flag.Bool("pretty",true,"pretty print RIB output")
	optFine := flag.Bool("fine",false,"fine render, adjusts the max samples")
	optEffects := flag.Bool("effects",false,"use camera effects")
	optBunnies := flag.Bool("bunnies",false,"show other bunnies")
	flag.Parse()

	production := *optProduction
	development := *optDevelopment
	debug := *optDebug
	fine := *optFine
	effects := *optEffects
	width := *optWidth
	height := *optHeight
	scale := *optScale
	prettyPrint := *optPrettyPrint
	bunnies := *optBunnies

	if !production && !development && !debug {
		fmt.Fprintf(os.Stderr,"need to specify at least one rendering option\n\n")
		/* TODO: add flag output */
		os.Exit(1)
	}

	/* Rendering Options :
   *
	 * Production - PxrPathTracer
   * Development - PxrDirectLighting | PxrDefault
   * Debug - PxrVisualizer | PxrDebugShadingContext
	 */

	pipe := rigo.DefaultFilePipe()


	shaders := NewPrefixShaderUniqueGenerator("shader_")
	lights  := NewPrefixLightUniqueGenerator("light_")

	mgr := rigo.NewHandleManager(nil,lights,shaders)

	ctx := rigo.NewContext(pipe,mgr,&rigo.Configuration{PrettyPrint: prettyPrint})

	ri := rigo.RI(ctx)
	ris := rigo.RIS(ctx)

	/* Bxdf Shaders */
	pattern,err := ris.Pattern("PxrFractal","-")
	if err != nil {
		fmt.Fprintf(os.Stderr,"error reading pattern shader -- %v\n",err)
		os.Exit(1)
	}

	hero,err := ris.Bxdf("PxrDisney","-")
	if err != nil {
		fmt.Fprintf(os.Stderr,"error reading bxdf shader -- %v\n",err)
		os.Exit(1)
	}

	hero.SetValue("roughness",RtFloat(0.1))
	hero.SetValue("specular",RtFloat(0.2))
	hero.SetValue("clearcoatGloss",RtFloat(0.5))
	hero.SetValue("metallic",RtFloat(0.1))
	hero.SetValue("clearcoat",RtFloat(0.2))
	hero.SetValue("baseColor",RtColor{0.9,0.9,0.9})
	hero.SetValue("sheenTint",RtFloat(0.2))

	if err := hero.SetReferencedValue("specular",pattern.ReferenceOutput("resultF")); err != nil {
		fmt.Fprintf(os.Stderr,"error referencing resultF value -- %v\n",err)
		os.Exit(1)
	}

	bunny,err := ris.Bxdf("PxrDisney","-")
	if err != nil {
		fmt.Fprintf(os.Stderr,"error reading bxdf shader -- %v\n",err)
		os.Exit(1)
	}

	bunny.SetValue("baseColor",RtColor{0.95,0.95,0.95})


	minSamples := RtInt(16)
	maxSamples := RtInt(32)
	heroModel := RtString("bunny_res4.rib")
	bunnyModel := RtString("bunny_res4.rib")

	if production {
		minSamples = RtInt(32)
		maxSamples = RtInt(64)
		heroModel = RtString("bunny_res2.rib")
		bunnyModel = RtString("bunny_res2.rib")

		if fine {
			minSamples = RtInt(64)
			maxSamples = RtInt(128)
		}
	}

	if err := ri.Begin("bunny.rib"); err != nil {
		fmt.Fprintf(os.Stderr,"error beginning -- %v\n",err)
		os.Exit(1)
	}

	curuser, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr,"error reading current user -- %v\n",err)
		os.Exit(1)
	}

	ri.ArchiveRecord("structure", "Scene %s", Scene)
	ri.ArchiveRecord("structure", "Creator %s", Author)
	ri.ArchiveRecord("structure", "CreationDate %s", time.Now())
	ri.ArchiveRecord("structure", "For %s", curuser.Username)
	ri.ArchiveRecord("structure", "Frames %d", 1)

	ri.Option("searchpath",RtToken("string shader"),RtString("./shaders:@"))
	ri.Option("searchpath",RtToken("string texture"),RtString("./textures:@"))
	ri.Option("searchpath",RtToken("string archive"),RtString("./archives:@"))
	ri.Option("shading",RtToken("int directlightinglocalizedsampling"),RtInt(3))
	ri.Option("bucket",RtToken("string order"),RtString("horizontal"))

	ri.Clipping(0.1,100)
	ri.Display("rabbit.tiff","file","rgba")

	fmt.Printf("Format %dx%d\n",int(width * scale),int(height * scale))
	ri.Format(RtInt(int(width * scale)),RtInt(int(height *scale)),1)

	ri.ShadingRate(1) 
	ri.ShadingInterpolation("smooth")
		
	ri.Attribute("trace",RtToken("int maxspeculardepth"),RtInt(3),RtToken("int displacements"),RtInt(1),
											 RtToken("int maxdiffusedepth"),RtInt(3))

	ri.PixelFilter(GaussianFilter,RtFloat(5),RtFloat(5)) 
	ri.PixelVariance(0.001)

	/* declare some of the parameters */
	ri.Declare("minsamples","int")
	ri.Declare("maxsamples","int")
	ri.Declare("incremental","int")
	ri.Declare("aperture","float[4]")
	ri.Declare("integrationmode","string")

	if err := ri.Hider("raytrace",RtToken("maxsamples"),maxSamples,RtToken("minsamples"),minSamples,
											RtToken("incremental"),RtInt(12),RtToken("aperture"),RtFloatArray{0,0,0,0},
											RtToken("integrationmode"),RtString("path")); err != nil {

		fmt.Fprintf(os.Stderr,"hider error -- %v\n",err)
		os.Exit(1)
	}



	if production {
		fmt.Printf("Setting Production -- PxrPathTracer\n")
		

		ri.Integrator("PxrPathTracer","production",RtToken("int maxPathLength"),RtInt(20),
									RtToken("string sampleMod"),RtString("bxdf"),RtToken("int rouletteDepth"),RtInt(12),
									RtToken("float rouletteThreshold"),RtFloat(0.2),RtToken("int clampDepth"),RtInt(2),
									RtToken("int clampLuminance"),RtInt(10)) 
	}

	if development {
		fmt.Printf("Setting Development -- PxrDirectLighting\n")
			
		ri.Integrator("PxrDirectLighting","development") /* TODO */
	}

	if debug {
		fmt.Printf("Setting Debug -- PxrVisualizer\n")

		ri.Integrator("PxrVisualizer","debug") /* TODO */
	}

		
	ri.Imager("background",RtToken("color color"),RtColor{.45,.45,.45},RtToken("float alpha"),RtFloat(1))
	if effects {		
		ri.Projection("PxrCamera",RtToken("float fov"),RtFloat(40),
															RtToken("float natural"),RtFloat(1),
															RtToken("float radial1"),RtFloat(-0.01))
	} else {
		ri.Projection("PxrCamera",RtToken("float fov"),RtFloat(40))
	}

	ri.Translate(0,-2,20)
	ri.Rotate(-50,1,0,0)
	ri.Rotate(-20,0,1,0)

	ri.WorldBegin()
		ri.AttributeBegin()
			ri.Attribute("identifier",RtToken("string name"),RtString("sky"))
			ri.Attribute("visibility",RtToken("int camera"),RtInt(0))			
			ri.Translate(0,20,0)
			ri.ShadingRate(1)
				
			sky, err := ri.AreaLightSource("PxrStdEnvDayLight",RtToken("float importance"),RtFloat(2),
																		 RtToken("float exposure"),RtFloat(1),RtToken("vector directionVector"),
																		 RtVector{0,1,0},RtToken("color specAmount"),RtColor{.5,.5,.5},
																		 RtToken("float haziness"),RtFloat(1.7),RtToken("float enableShadows"),
																		 RtFloat(1))
			
			if err != nil {
				fmt.Fprintf(os.Stderr,"unable to create sky light -- %v\n",err)
				os.Exit(1)
			}


			ri.Bxdf("PxrLightEmission",RtShaderHandle(sky))
			ri.Geometry("envsphere",RtToken("constant float[2] resolution"),RtFloatArray{1024,1024})
		ri.AttributeEnd()

		ri.Illuminate(sky,true)

		ri.AttributeBegin()
			ri.Attribute("identifier",RtToken("name"),RtString("hero"))
			ri.Attribute("visibility",RtToken("int camera"),RtInt(1),RtToken("int transmission"),RtInt(1))
			ri.Translate(0,-2,0)
			ri.Rotate(-125,0,1,0)

			ri.Pattern(pattern.Name(),pattern.Handle())
			ri.Bxdf(hero.Name(),hero.Handle())
			ri.Scale(50,50,50)

			if err := ri.Procedural2("DelayedReadArchive2","SimpleBound",RtToken("string filename"),heroModel,
										 RtToken("float[6] bound"),RtFloatArray{-1,1,-1,1,-1,1}); err != nil {
				panic(err.Error())
			}

		ri.AttributeEnd()
	
		if bunnies {
		ri.AttributeBegin()
			ri.Attribute("identifier",RtToken("string name"),RtString("bunny_1"))
			ri.Attribute("visibility",RtToken("int camera"),RtInt(1),RtToken("int transmission"),RtInt(1),
																RtToken("int indirect"),RtInt(1))

			ri.Translate(-5,-1,0)
			ri.Rotate(-75,0,1,0)				
		
			ri.Bxdf(bunny.Name(),bunny.Handle())
			ri.Scale(20,20,20)

			ri.Procedural2("DelayedReadArchive2","SimpleBound",RtToken("string filename"),bunnyModel,
										 RtToken("float[6] bound"),RtBound{-1,1,-1,1,-1,1})

		ri.AttributeEnd()
		ri.AttributeBegin()
			ri.Attribute("identifier",RtToken("string name"),RtString("bunny_2"))
			ri.Attribute("visibility",RtToken("int camera"),RtInt(1),RtToken("int transmission"),RtInt(1),
																RtToken("int indirect"),RtInt(1))

			ri.Translate(-5,-1,3)
			ri.Rotate(25,0,1,0)
				
			ri.Bxdf(bunny.Name(),bunny.Handle())
			ri.Scale(20,20,20)

			ri.Procedural2("DelayedReadArchive2","SimpleBound",RtToken("string filename"),bunnyModel,
										 RtToken("float[6] bound"),RtFloatArray{-1,1,-1,1,-1,1})

		ri.AttributeEnd()
		} /* end bunnies */

		ri.AttributeBegin()
			ri.Attribute("identifier",RtToken("string name"),RtString("painted_floor"))
			ri.Attribute("visibility",RtToken("int camera"),RtInt(1),RtToken("int transmission"),RtInt(1))
			ri.Rotate(90,1,0,0)
			ri.Pattern("PxrTexture","floor_pattern",RtToken("string filename"),RtString("ratGrid.tex"),
								 RtToken("int invertT"),RtInt(0))
			ri.Bxdf("PxrDiffuse","floor",RtToken("reference color diffuseColor"),RtString("floor_pattern:resultRGB"))
			ri.Scale(20,20,20)
			ri.Patch("bilinear",RtToken("vertex point P"),RtFloatArray{-1,1,0,1,1,0,-1,-1,0,1,-1,0})
		ri.AttributeEnd()
	ri.WorldEnd()

	if err := ri.End(); err != nil {
		fmt.Fprintf(os.Stderr,"error ending -- %v\n",err)
		os.Exit(1)
	}

	p := pipe.GetByName(rigo.PipeToStats{}.Name())
	if p == nil {
		fmt.Fprintf(os.Stderr,"Pipe Stats not found!\n")
		os.Exit(1)
	}

	s, ok := p.(*rigo.PipeToStats)
	if !ok {
		fmt.Fprintf(os.Stderr, "Pipe Stats not found\n")
		os.Exit(1)
	}

	p = pipe.GetByName(rigo.PipeTimer{}.Name())
	if p == nil {
		fmt.Fprintf(os.Stderr, "Pipe Timer not found\n")
		os.Exit(1)
	}
	t, ok := p.(*rigo.PipeTimer)
	if !ok {
		fmt.Fprintf(os.Stderr, "Pipe Timer not found\n")
		os.Exit(1)
	}

	fmt.Printf("%s\n\n%s\n\n", s, t)
}


	

		
	
