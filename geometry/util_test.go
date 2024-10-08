package geometry_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/dohyeunglee/go-geos"
	"github.com/dohyeunglee/go-geos/geometry"
)

func mustNewGeometryFromWKT(t *testing.T, wkt string) *geometry.Geometry {
	t.Helper()
	geom, err := geos.NewGeomFromWKT(wkt)
	assert.NoError(t, err)
	return &geometry.Geometry{Geom: geom}
}
