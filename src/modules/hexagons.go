package modules

import (
	"fmt"
	"math"
	"strings"

	"github.com/etu/absvgen/src/log"
	"github.com/etu/absvgen/src/spec"
)

type Hexagons struct{}

func (m Hexagons) Render(specLayer spec.SpecLayer, specFile spec.SpecFile) string {
	// Calculate size of squares
	edgeSize := float64(specFile.Width) / float64(100) * float64(specLayer.Size)
	shortDiagonalSize := math.Sqrt(3) * float64(edgeSize)
	triangleHeight := math.Sqrt(float64(edgeSize*edgeSize) - (0.5*float64(shortDiagonalSize))*(0.5*float64(shortDiagonalSize)))

	offsetStart := false
	x := float64(0)
	y := float64(0)
	colorSelector := 0

	var hexagons []string
	hexagons = append(hexagons, "<g>")

	for y < float64(specFile.Height)+triangleHeight {
		//     B
		//    / \
		//  A/   \C
		//   |   |
		//   |   |
		//  F\   /D
		//    \ /
		//     E

		// Coordinates for point A (in the ascii diagram above)
		ax := x
		ay := y

		// Coordinates for point B (in the ascii diagram above)
		bx := x + (shortDiagonalSize / 2)
		by := y - triangleHeight
		if y < triangleHeight {
			by = y
		}

		// Coordinates for point C (in the ascii diagram above)
		cx := x + shortDiagonalSize
		cy := y

		// Coordinates for point D (in the ascii diagram above)
		dx := x + shortDiagonalSize
		dy := y + edgeSize

		// Coordinates for point E (in the ascii diagram above)
		ex := x + (shortDiagonalSize / 2)
		ey := y + edgeSize + triangleHeight

		// Coordinates for point F (in the ascii diagram above)
		fx := x
		fy := y + edgeSize

		// If we're at the beginning of the X side and we should have
		// an offset start (really a - coordinate, but that's not
		// allowed due to unsigned numbers), we just "flatten" the
		// hexagon against the X wall and move the other points to
		// make a half hexagon.
		if offsetStart && x == 0 {
			bx = 0
			cx = shortDiagonalSize / 2
			dx = shortDiagonalSize / 2
			ex = 0
		}

		// Generate a square
		hexagons = append(hexagons, fmt.Sprintf(
			"<polygon points=\"%f,%f %f,%f %f,%f %f,%f %f,%f %f,%f\" style=\"fill: %s\" />",
			ax,
			ay,
			bx,
			by,
			cx,
			cy,
			dx,
			dy,
			ex,
			ey,
			fx,
			fy,
			specLayer.Colors[colorSelector],
		))

		// Progress on x axis
		//
		// If we have an offset start, we move half an hexagon,
		// otherwise we move a full one.
		if offsetStart && x == 0 {
			x += shortDiagonalSize / 2
		} else {
			x += shortDiagonalSize
		}

		// Progress in color selection
		colorSelector++

		// Reset color selector if we go out of range
		if colorSelector >= len(specLayer.Colors) {
			colorSelector = 0
		}

		// Jump to next row when width is filled
		if x > float64(specFile.Width) {
			offsetStart = !offsetStart
			x = 0
			y += edgeSize + triangleHeight
		}
	}

	hexagons = append(hexagons, "</g>")

	log.Print(fmt.Sprintf("[module: hexagons] Hexagons module has rendered %d hexagons", len(hexagons)-1))

	return strings.Join(hexagons, "")
}
