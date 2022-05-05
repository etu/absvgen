package modules

import (
	"fmt"

	"github.com/etu/absvgen/src/log"
	"github.com/etu/absvgen/src/spec"
)

type Solid struct{}

func (m Solid) Render(specLayer spec.SpecLayer, specFile spec.SpecFile) string {
	log.Print(fmt.Sprintf("[module: solid] Solid module has rendered a <rect> with the fill color %s", specLayer.Colors[0]))

	return fmt.Sprintf(
		"<rect width=\"%d\" height=\"%d\" style=\"fill: %s\" />",
		specFile.Width,
		specFile.Height,
		specLayer.Colors[0],
	)
}
