package main

import (
	"fmt"
	"os"

	. "github.com/mae-global/rigo/ri"
	"github.com/mae-global/rigo"
)

func main() {

	ri,_ := rigo.DefaultPipeline(&rigo.Configuration{PrettyPrint: true})
	ri.Begin("./001.rib")
	ri.ArchiveRecord("structure","Scene regressions/001")
	ri.ArchiveRecord("structure","Regression 001")
	ri.Display("./001.tif","file","rgb")
	ri.Format(300,300,1)
	ri.Translate(0,0,6)
	ri.WorldBegin()
		ri.Color(RtColor{1,0,0})
		ri.Sphere(1,-1,1,360)
	ri.WorldEnd()

	if err := ri.End(); err != nil {
		fmt.Fprintf(os.Stderr,"End() error -- %v\n",err)
		os.Exit(1)
	} 
}
