package main // import "github.com/etu/absvgen"

import (
	"flag"
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
}
