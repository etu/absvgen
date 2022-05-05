package modules

import "github.com/etu/absvgen/src/spec"

type Module interface {
	Render(spec.SpecLayer, spec.SpecFile) string
}

func Get(moduleName string) Module {
	if moduleName == "solid" {
		return Solid{}
	}

	return Dummy{}
}
