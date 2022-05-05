package modules

import (
	"github.com/etu/absvgen/src/log"
	"github.com/etu/absvgen/src/spec"
)

type Dummy struct{}

func (m Dummy) Render(specLayer spec.SpecLayer, specFile spec.SpecFile) string {
	log.Print("[module: dummy] Dummy module has rendered a dummy comment as response")

	return "<!-- dummy value from dummy module -->"
}
