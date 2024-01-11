package geos_test

import (
	"fmt"
	"math"
	"runtime"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/go-geos"
)

func TestCoordSeqAliasing(t *testing.T) {
	coords := geos.NewContext().NewCoordSeqFromCoords([][]float64{{0, 1}, {2, 3}}).ToCoords()
	coords[0] = append(coords[0], 4)
	assert.Equal(t, []float64{2, 3}, coords[1])
}

func TestCoordSeqEmpty(t *testing.T) {
	defer runtime.GC() // Exercise finalizers.
	c := geos.NewContext()
	s := c.NewCoordSeq(0, 2)
	assert.Equal(t, 0, s.Size())
	assert.Equal(t, 2, s.Dimensions())
	assert.Equal(t, nil, s.ToCoords())
}

func TestCoordSeqIsCCW(t *testing.T) {
	for _, tc := range []struct {
		name               string
		coords             [][]float64
		expected           bool
		expectedErrPre13_2 bool
	}{
		{
			name:     "ccw",
			coords:   [][]float64{{0, 0}, {1, 0}, {1, 1}, {0, 0}},
			expected: true,
		},
		{
			name:     "cw",
			coords:   [][]float64{{0, 0}, {0, 1}, {1, 1}, {0, 0}},
			expected: false,
		},
		{
			name:               "short",
			coords:             [][]float64{{0, 0}, {1, 0}, {1, 1}},
			expected:           false,
			expectedErrPre13_2: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer runtime.GC() // Exercise finalizers.
			s := geos.NewContext().NewCoordSeqFromCoords(tc.coords)
			if geos.VersionCompare(3, 12, 0) < 0 && tc.expectedErrPre13_2 {
				assert.Panics(t, func() {
					s.IsCCW()
				})
			} else {
				assert.Equal(t, tc.expected, s.IsCCW())
			}
		})
	}
}

func TestCoordSeqMethods(t *testing.T) {
	defer runtime.GC() // Exercise finalizers.
	c := geos.NewContext()
	s := c.NewCoordSeq(2, 3)
	assert.Equal(t, 2, s.Size())
	assert.Equal(t, 3, s.Dimensions())
	assert.Equal(t, 0.0, s.X(0))
	assert.Equal(t, 0.0, s.Y(0))
	assert.True(t, math.IsNaN(s.Z(0)))
	s.SetZ(0, 0)
	s.SetX(1, 1)
	s.SetY(1, 2)
	s.SetZ(1, 3)
	assert.Equal(t, 1.0, s.X(1))
	assert.Equal(t, 2.0, s.Y(1))
	assert.Equal(t, 3.0, s.Z(1))
	assert.Equal(t, 1.0, s.Ordinate(1, 0))
	assert.Equal(t, 2.0, s.Ordinate(1, 1))
	assert.Equal(t, 3.0, s.Ordinate(1, 2))
	assert.Equal(t, [][]float64{{0, 0, 0}, {1, 2, 3}}, s.ToCoords())

	clone := s.Clone()
	assert.Equal(t, 1.0, clone.X(1))
	assert.Equal(t, 2.0, clone.Y(1))
	clone.SetOrdinate(0, 0, -1.0)
	clone.SetOrdinate(0, 1, -2.0)
	assert.Equal(t, -1.0, clone.X(0))
	assert.Equal(t, -2.0, clone.Y(0))
	assert.Equal(t, [][]float64{{-1, -2, 0}, {1, 2, 3}}, clone.ToCoords())

	assert.Equal(t, 3, clone.Dimensions())
	assert.Equal(t, 3.0, clone.Z(1))
	clone.SetOrdinate(0, 2, -3.0)
	assert.Equal(t, -3.0, clone.Z(0))
}

func TestCoordSeqPanics(t *testing.T) {
	c := geos.NewContext()
	s := c.NewCoordSeq(1, 2)

	assert.Panics(t, func() { s.X(-1) })
	assert.NotPanics(t, func() { s.X(0) })
	assert.Panics(t, func() { s.X(1) })

	assert.Panics(t, func() { s.Y(-1) })
	assert.NotPanics(t, func() { s.Y(0) })
	assert.Panics(t, func() { s.Y(1) })

	assert.Panics(t, func() { s.Z(-1) })
	assert.Panics(t, func() { s.Z(0) })
	assert.Panics(t, func() { s.Z(1) })

	assert.Panics(t, func() { s.SetX(-1, 0) })
	assert.NotPanics(t, func() { s.SetX(0, 0) })
	assert.Panics(t, func() { s.SetX(1, 0) })

	assert.Panics(t, func() { s.SetY(-1, 0) })
	assert.NotPanics(t, func() { s.SetY(0, 0) })
	assert.Panics(t, func() { s.SetY(1, 0) })

	assert.Panics(t, func() { s.SetZ(-1, 0) })
	assert.Panics(t, func() { s.SetZ(0, 0) })
	assert.Panics(t, func() { s.SetZ(1, 0) })

	for idx := -1; idx <= 1; idx++ {
		for dim := -1; dim <= 4; dim++ {
			t.Run(fmt.Sprintf("idx_%d_dim_%d", idx, dim), func(t *testing.T) {
				if idx == 0 && 0 <= dim && dim < 2 {
					assert.NotPanics(t, func() { s.Ordinate(idx, dim) })
					assert.NotPanics(t, func() { s.SetOrdinate(idx, dim, 0) })
				} else {
					assert.Panics(t, func() { s.Ordinate(idx, dim) })
					assert.Panics(t, func() { s.SetOrdinate(idx, dim, 0) })
				}
			})
		}
	}
}

func TestCoordSeqCoordsMethods(t *testing.T) {
	for _, tc := range []struct {
		name   string
		coords [][]float64
	}{
		{
			name:   "point_2d",
			coords: [][]float64{{1, 2}},
		},
		{
			name:   "point_3d",
			coords: [][]float64{{1, 2, 3}},
		},
		{
			name:   "linestring_2d",
			coords: [][]float64{{1, 2}, {3, 4}},
		},
		{
			name:   "linestring_3d",
			coords: [][]float64{{1, 2, 3}, {4, 5, 6}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer runtime.GC() // Exercise finalizers.
			c := geos.NewContext()
			s := c.NewCoordSeqFromCoords(tc.coords)
			assert.Equal(t, tc.coords, s.ToCoords())
		})
	}
}
