package polygon

import (
	"image"
	"image/color"
)

// Draw - draws a polygon within given vertices and fills the space using the fillColor
// algorithm used: scanline filling with linear interpolation
// linear function `y = i` (straight, horizontal line) goes from the top of an image
// using linear interpolation to find the x of given a pair of vertices (edge),
// finds the interception between `y = i` and the edge of the polygon
// x1 and x2 represents the horizontal pixels range to color with fillColor
// more about the used algorithm in the README.md
func Draw(img *image.RGBA, vertices Vertices, fillColor color.Color) {
	for y := 0; y < img.Bounds().Dy(); y++ {
		x1 := img.Bounds().Dx()
		x2 := 0

		for i := 0; i < len(vertices); i++ {
			j := (i + 1) % len(vertices)
			if (vertices[i].Y > y && vertices[j].Y <= y) || (vertices[j].Y > y && vertices[i].Y <= y) {
				x := vertices[i].X + (y-vertices[i].Y)*(vertices[j].X-vertices[i].X)/(vertices[j].Y-vertices[i].Y)

				if x < x1 {
					x1 = x
				}

				if x > x2 {
					x2 = x
				}
			}
		}
		for x := x1; x <= x2; x++ {
			img.Set(x, y, fillColor)
		}
	}
}
