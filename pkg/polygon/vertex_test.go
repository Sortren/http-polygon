package polygon

import (
	"testing"
)

func TestShapeBounds(t *testing.T) {
	// given
	vertices := Vertices{
		{
			X: 50,
			Y: 100,
		},
		{
			X: 100,
			Y: 300,
		},
		{
			X: 300,
			Y: 200,
		},
		{
			X: 25,
			Y: 500,
		},
	}

	// when
	bounds := vertices.PolygonBound()

	// then
	expected := ShapeBound{
		MinX: 25,
		MinY: 100,
		MaxX: 300,
		MaxY: 500,
	}

	if expected != bounds {
		t.Fatalf("expected: %+v, result: %+v", expected, bounds)
	}
}
