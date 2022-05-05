package modules

import (
	"fmt"
	"strings"

	"github.com/etu/absvgen/src/log"
	"github.com/etu/absvgen/src/spec"
)

type Squares struct{}

func (m Squares) Render(specLayer spec.SpecLayer, specFile spec.SpecFile) string {
	// Calculate size of squares
	squareSize := uint16(float32(specFile.Width) / float32(100) * float32(specLayer.Size))

	x := uint16(0)
	y := uint16(0)
	colorSelector := 0

	var squares []string
	squares = append(squares, "<g>")

	for y < specFile.Height {
		// Generate a square
		squares = append(squares, fmt.Sprintf(
			"<rect width=\"%d\" height=\"%d\" x=\"%d\" y=\"%d\" style=\"fill: %s\" />",
			squareSize,
			squareSize,
			x,
			y,
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

	squares = append(squares, "</g>")

	log.Print(fmt.Sprintf("[module: squares] Squares module has rendered %d", len(squares)-1))

	return strings.Join(squares, "")
}
