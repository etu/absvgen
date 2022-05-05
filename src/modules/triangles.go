package modules

import (
	"fmt"
	"strings"

	"github.com/etu/absvgen/src/log"
	"github.com/etu/absvgen/src/spec"
)

type Triangles struct{}

func (m Triangles) Render(specLayer spec.SpecLayer, specFile spec.SpecFile) string {
	// Calculate size of squares *to put triangles in*
	squareSize := uint16(float32(specFile.Width) / float32(100) * float32(specLayer.Size))

	x := uint16(0)
	y := uint16(0)
	colorSelector := 0

	var triangles []string
	triangles = append(triangles, "<g>")

	for y < specFile.Height {
		// Generate a triangle half of the square
		triangles = append(triangles, fmt.Sprintf(
			"<polygon points=\"%d,%d %d,%d %d,%d\" style=\"fill: %s\" />",
			// First point
			x,
			y,
			// Second point
			x+squareSize,
			y,
			// Third point
			x,
			y+squareSize,
			specLayer.Colors[colorSelector],
		))

		colorSelector++

		// Reset color selector if we go out of range
		if colorSelector >= len(specLayer.Colors) {
			colorSelector = 0
		}

		// Generate a triangle half of the square
		triangles = append(triangles, fmt.Sprintf(
			"<polygon points=\"%d,%d %d,%d %d,%d\" style=\"fill: %s\" />",
			// First point
			x+squareSize,
			y,
			// Second point
			x,
			y+squareSize,
			// Third point
			x+squareSize,
			y+squareSize,
			specLayer.Colors[colorSelector],
		))

		// Progress on x axis
		x += squareSize

		// Progress in color selection
		colorSelector++

		// Reset color selector if we go out of range
		if colorSelector >= len(specLayer.Colors) {
			colorSelector = 0
		}

		// Jump to next row when width is filled
		if x > specFile.Width {
			x = 0
			y += squareSize
		}
	}

	triangles = append(triangles, "</g>")

	log.Print(fmt.Sprintf("[module: Triangles] Triangles module has rendered %d", len(triangles)-1))

	return strings.Join(triangles, "")
}
