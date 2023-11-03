package polygon

import (
	"cmp"
	"errors"
	"fmt"
	"image"
	"slices"
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

type ShapeBound struct {
	MinX, MinY, MaxX, MaxY int
}

func (v Vertices) PolygonBound() ShapeBound {
	return ShapeBound{
		MinX: slices.MinFunc(v, func(v1, v2 Vertex) int {
			return cmp.Compare(v1.X, v2.X)
		}).X,
		MinY: slices.MinFunc(v, func(v1, v2 Vertex) int {
			return cmp.Compare(v1.Y, v2.Y)
		}).Y,
		MaxX: slices.MaxFunc(v, func(v1, v2 Vertex) int {
			return cmp.Compare(v1.X, v2.X)
		}).X,
		MaxY: slices.MaxFunc(v, func(v1, v2 Vertex) int {
			return cmp.Compare(v1.Y, v2.Y)
		}).Y,
	}
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
