package geometry_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/twpayne/go-geos"
)

func TestBounds(t *testing.T) {
	for _, tc := range []struct {
		name            string
		bounds          *geos.Box2D
		expectedEmpty   bool
		expectedGeomWKT string
	}{
		{
			name:            "NewBoundsEmpty",
			bounds:          geos.NewBox2DEmpty(),
			expectedEmpty:   true,
			expectedGeomWKT: "POINT EMPTY",
		},
		{
			name:            "NewBoundsFromGeometry_empty_point",
			bounds:          mustNewGeometryFromWKT(t, "POINT EMPTY").Bounds(),
			expectedEmpty:   true,
			expectedGeomWKT: "POINT EMPTY",
		},
		{
			name:            "NewBoundsFromGeometry_point",
			bounds:          mustNewGeometryFromWKT(t, "POINT (0 1)").Bounds(),
			expectedEmpty:   false,
			expectedGeomWKT: "POINT (0 1)",
		},
		{
			name:            "NewBoundsFromGeometry_line_string",
			bounds:          mustNewGeometryFromWKT(t, "LINESTRING (0 1, 2 3)").Bounds(),
			expectedEmpty:   false,
			expectedGeomWKT: "POLYGON ((0 1, 2 1, 2 3, 0 3, 0 1))",
		},
		{
			name:            "NewBoundsFromGeometry_line_string_empty",
			bounds:          mustNewGeometryFromWKT(t, "LINESTRING EMPTY").Bounds(),
			expectedEmpty:   true,
			expectedGeomWKT: "POINT EMPTY",
		},
		{
			name:            "NewBoundsFromGeometry_polygon_empty",
			bounds:          mustNewGeometryFromWKT(t, "POLYGON EMPTY").Bounds(),
			expectedEmpty:   true,
			expectedGeomWKT: "POINT EMPTY",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			//nolint:gocritic
			assert.True(t, tc.bounds.Equals(tc.bounds))
			assert.Equal(t, tc.expectedEmpty, tc.bounds.IsEmpty())
			expectedGeom, err := geos.NewGeomFromWKT(tc.expectedGeomWKT)
			assert.NoError(t, err)
			assert.True(t, expectedGeom.Equals(tc.bounds.Geom()))
		})
	}
}
