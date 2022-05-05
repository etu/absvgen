package modules

import "github.com/etu/absvgen/src/spec"

type Solid struct{}

func (m Solid) Render(specLayer spec.SpecLayer, specFile spec.SpecFile) string {
	return "<!-- dummy value from solid object -->"
}
