package modules

import (
	"fmt"

	"github.com/etu/absvgen/src/log"
	"github.com/etu/absvgen/src/spec"
)

type Module interface {
	Render(spec.SpecLayer, spec.SpecFile) string
}

func Get(moduleName string) Module {
	if moduleName == "solid" {
		log.Print(fmt.Sprintf("[get module: %s] Requested module found, returning Solid{}", moduleName))
		return Solid{}
	}

	if moduleName == "squares" {
		log.Print(fmt.Sprintf("[get module: %s] Requested module found, returning Squares{}", moduleName))
		return Squares{}
	}

	if moduleName == "triangles" {
		log.Print(fmt.Sprintf("[get module: %s] Requested module found, returning Triangles{}", moduleName))
		return Triangles{}
	}

	log.Print(fmt.Sprintf("[get module: %s] Requested module not found, returning Dummy{}", moduleName))

	return Dummy{}
}
