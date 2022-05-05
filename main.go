package main // import "github.com/etu/absvgen"

import (
	"flag"
	"fmt"

	"github.com/etu/absvgen/src/doc"
	"github.com/etu/absvgen/src/modules"
	"github.com/etu/absvgen/src/spec"
)

func main() {
	var specFile string
	var debug bool

	//
	// Parse command line flags
	//
	flag.BoolVar(&debug, "debug", false, "Enable or disable debug output")
	flag.StringVar(&specFile, "spec", "spec.yaml", "Specify path to the spec.yaml file")
	flag.Parse()

	// Parse the specification
	spec := spec.New(specFile)

	// Slice to store render results in
	docObjects := []string{}

	for _, v := range spec.Layers {
		// Get the module we want to use
		module := modules.Get(v.Module)

		// Render the module
		moduleResult := module.Render(v, spec)

		// Store the result in the slice
		docObjects = append(docObjects, moduleResult)
	}

	// Wrap the SVG objects from the different modules in SVG header/footer
	svg := doc.WrapDocObjects(spec, docObjects)

	fmt.Println(svg)
}
