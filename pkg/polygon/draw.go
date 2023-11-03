package polygon

import (
	"image"
	"image/color"
	"sync"
)

type DrawService interface {
	Draw(img *image.RGBA, vertices Vertices, fillColor color.Color)
}

type ConcurrentDrawService struct {
	MaxGoroutines int
}

func NewConcurrentDrawService(maxGoroutines int) *ConcurrentDrawService {
	return &ConcurrentDrawService{MaxGoroutines: maxGoroutines}
}

type StandardDrawService struct{}

func NewStandardDrawService() *StandardDrawService {
	return &StandardDrawService{}
}

// Draw - draws a polygon within given vertices and fills the space using the fillColor
// algorithm used: scanline filling with linear interpolation
// linear function `y = i` (straight, horizontal line) goes from the highest vertex (max Y coordinate)
// using linear interpolation to find the x of given a pair of vertices (edge),
// finds the interception between `y = i` and the edge of the polygon
// x1 and x2 represents the horizontal pixels range to color with fillColor
// more about the used algorithm in the README.md
func (s *StandardDrawService) Draw(img *image.RGBA, vertices Vertices, fillColor color.Color) {
	bounds := vertices.PolygonBound()

	for y := bounds.MinY; y <= bounds.MaxY; y++ {
		x1, x2 := horizontalRangeToFillColor(y, bounds, vertices)
		for x := x1; x <= x2; x++ {
			img.Set(x, y, fillColor)
		}
	}
}

func (c *ConcurrentDrawService) Draw(img *image.RGBA, vertices Vertices, fillColor color.Color) {
	bounds := vertices.PolygonBound()

	lines := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for y := bounds.MinY; y <= bounds.MaxY; y++ {
			lines <- y
		}
		close(lines)
	}()

	for i := 0; i < c.MaxGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for y := range lines {
				x1, x2 := horizontalRangeToFillColor(y, bounds, vertices)

				for x := x1; x <= x2; x++ {
					img.Set(x, y, fillColor)
				}
			}
		}()
	}

	wg.Wait()
}

// horizontalRangeToFillColor - for given linear function y, returns the range from x1 to x2 that should be painted with color
func horizontalRangeToFillColor(y int, bounds ShapeBound, vertices Vertices) (x1, x2 int) {
	x1 = bounds.MaxX
	x2 = bounds.MinX

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

	return x1, x2
}
