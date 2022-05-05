package modules

import "github.com/etu/absvgen/src/spec"

type Dummy struct{}

func (m Dummy) Render(specLayer spec.SpecLayer, specFile spec.SpecFile) string {
	return "<!-- dummy value from dummy module -->"
}
