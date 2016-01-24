# rigo
Implementation of the RenderMan Interface for the Go programming language. 

[Online Documentation](https://godoc.org/github.com/mae-gloab/rigo)

Install with:

    go get github.com/mae-global/rigo

Quick example usage; outputting to a RIB Entity file. 

		pipe := ri.DefaultFilePipe()

		ctx := ri.NewEntity(pipe)
		ctx.Begin("output/exampleD21.rib")
		ctx.AttributeBegin("begin unit cube")
		ctx.Attribute("identifier", RtToken("name"), RtToken("unitcube"))
		ctx.Bound(RtBound{-.5, .5, -.5, .5, -.5, .5})
		ctx.TransformBegin()

		ctx.Comment("far face")
		ctx.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})
		ctx.Rotate(90, 0, 1, 0)

		ctx.Comment("right face")
		ctx.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})
		ctx.Rotate(90, 0, 1, 0)

		ctx.Comment("near face")
		ctx.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})
		ctx.Rotate(90, 0, 1, 0)

		ctx.Comment("left face")
		ctx.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})

		ctx.TransformEnd()
		ctx.TransformBegin()

		ctx.Comment("bottom face")
		ctx.Rotate(90, 1, 0, 0)
		ctx.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})

		ctx.TransformEnd()
		ctx.TransformBegin()

		ctx.Comment("top face")
		ctx.Rotate(-90, 1, 0, 0)
		ctx.Polygon(4, RtToken("P"), RtFloatArray{.5, .5, .5, -.5, .5, .5, -.5, -.5, .5, .5, -.5, .5})

		ctx.TransformEnd()
		ctx.AttributeEnd("end unit cube")
	
		p := pipe.GetByName(PipeToStats{}.Name())
		s, _ := p.(*PipeToStats)
	
		p = pipe.GetByName(PipeTimer{}.Name())
		t,_ := p.(*PipeTimer)
	
		fmt.Printf("%s%s", s,t)
	
The implementation is still under active development, so expect holes and bugs. 
