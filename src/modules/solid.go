package modules

import (
	"fmt"

	"github.com/etu/absvgen/src/spec"
)

type Solid struct{}

func (m Solid) Render(specLayer spec.SpecLayer, specFile spec.SpecFile) string {
	return fmt.Sprintf(
		"<rect width=\"%d\" height=\"%d\" style=\"fill: %s\" />",
		specFile.Width,
		specFile.Height,
		specLayer.Colors[0],
	)
}
