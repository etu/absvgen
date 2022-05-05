package main // import "github.com/etu/absvgen"

import (
	"flag"
	"fmt"

	"github.com/etu/absvgen/src/doc"
	"github.com/etu/absvgen/src/log"
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

	// Wether to emit something on stderr or not
	log.Enable(debug)

	// Parse the specification
	log.Print(fmt.Sprintf("Parsing specification file %s", specFile))
	spec := spec.New(specFile)

	// Slice to store render results in
	docObjects := []string{}

	log.Print(fmt.Sprintf("Going through specification layers"))

	for i, v := range spec.Layers {
		log.Print(fmt.Sprintf("[layer: %d, module: %s] Parsing layer", i, v.Module))

		// Get the module we want to use
		module := modules.Get(v.Module)

		log.Print(fmt.Sprintf("[layer: %d, module: %s] Rendering layer", i, v.Module))

		// Render the module
		moduleResult := module.Render(v, spec)

		log.Print(fmt.Sprintf("[layer: %d, module: %s] Result: %s", i, v.Module, moduleResult))

		// Store the result in the slice
		docObjects = append(docObjects, moduleResult)
	}

	log.Print(fmt.Sprintf("Completed specification layers"))
	log.Print(fmt.Sprintf("Wrapping layers with SVG object and header"))

	// Wrap the SVG objects from the different modules in SVG header/footer
	svg := doc.WrapDocObjects(spec, docObjects)

	fmt.Println(svg)
}
