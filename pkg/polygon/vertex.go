package polygon

import (
	"errors"
	"fmt"
	"image"
)

var (
	ErrVertexCoordinateOutOfRange = errors.New("vertex coordinates are out of defined range")
)

type Vertex struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (v Vertex) ValidateCoordinates(min, max Vertex) error {
	if v.X < min.X || v.Y < min.Y || v.X > max.X || v.Y > max.Y {
		return ErrVertexCoordinateOutOfRange
	}
	return nil
}

type Vertices []Vertex

func (v Vertices) Validate(min, max Vertex) error {
	for _, vertex := range v {
		if err := vertex.ValidateCoordinates(min, max); err != nil {
			return fmt.Errorf("invalid coordinates of vertex %+v: %w", vertex, err)
		}
	}

	return nil
}

func ImageBoundsToMinMaxVertex(bounds image.Rectangle) (min Vertex, max Vertex) {
	min = Vertex{
		X: bounds.Min.X,
		Y: bounds.Min.Y,
	}
	max = Vertex{
		X: bounds.Max.X,
		Y: bounds.Max.Y,
	}

	return min, max
}
