package modules

import (
	"fmt"
	"strings"

	"github.com/etu/absvgen/src/spec"
)

type Stripes struct{}

func (m Stripes) Render(specLayer spec.SpecLayer, specFile spec.SpecFile) string {
	// Calculate size of paths
	size := uint16(float32(specFile.Width) / float32(100) * float32(specLayer.Size))
	steps := 0

	var stripes []string
	stripes = append(stripes, "<g>")

	for steps < len(specLayer.Colors) {
		// Calculate X axises for stripes Here we want them evenly
		// spaced based on width, so the first stripe is 1 stripes
		// width from the left, then it should be 1 full stripe
		// between each stripe.
		x := uint16(float32(size)*1.5) + size*2*uint16(steps)

		stripes = append(stripes, fmt.Sprintf(
			"<line x1=\"%d\" y1=\"%d\" x2=\"%d\" y2=\"%d\" style=\"stroke: %s; stroke-width: %d\" />",
			x,
			0,
			x,
			specFile.Height,
			specLayer.Colors[steps],
			size,
		))

		steps++
	}
	stripes = append(stripes, "</g>")

	return strings.Join(stripes, "")
}
