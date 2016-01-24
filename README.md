# RiGO
Implementation of the RenderMan Interface for the Go programming language. This is currently 
based on Pixar's RenderMan Specification version 3.2.1 (November 2005). This implementation 
is still under *active development*, so *expect* holes and bugs. 

[Online Documentation](https://godoc.org/github.com/mae-global/rigo)

Install with:

    go get github.com/mae-global/rigo

Quick example usage; outputting a Unit Cube to a RIB Entity file. 

```go
/* create a function to record the duration between RiBegin and RiEnd calls */
type MyTimer struct {
	start time.Time
	finish time.Time
}

func (t MyTimer) Name() string {
	return "mytimer"
}

func (t *MyTimer) Took() time.Duration {
	return t.finish.Sub(t.start)
}

func (t *MyTimer) Write(name RtName,list []Rter,info Info) *Result {
	switch string(name) {
		case "Begin","RiBegin":
			t.start = time.Now()
			t.finish = t.start
		break
		case "End","RiEnd":
			t.finish = time.Now()
		break
	}
	return Done()
}

/* Construct a pipeline, including our timer, piping RIB output to file */
pipe := NewPipe()
pipe.Append(&MyTimer{}).Append(&PipeToFile{})

ctx := NewEntity(pipe)

/* Do all our Ri calls */
ctx.Begin("unitcube.rib")
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
ctx.End()	
		
/* grab our timer back and print the duration */
p = pipe.GetByName(MyTimer{}.Name())
t,_ := p.(*MyTimer)
	
fmt.Printf("took %s\n",t.Took())
```	

##Roadmap

- [x] Basic RIB pipe
- [ ] Complete RenderMan Interface
- [ ] Stdout/buffer wrapper around io.Writer interface
- [ ] Complete Error checking for each Ri Call
  - [x] Basic Error checking
	- [ ] Sanity checking
	- [ ] Per call checking
	- [ ] Parameterlist checking
- [ ] RIB parser
- [ ] Call wrapping for Ri[call]Begin/Ri[call]End pairs
- [ ] Call Fragments 
- [ ] Documentation/Examples


###Information

RenderMan Interface Specification is Copyright © 2005 Pixar.
RenderMan © is a registered trademark of Pixar.

