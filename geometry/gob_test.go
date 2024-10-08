package geometry_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/dohyeunglee/go-geos"
	"github.com/dohyeunglee/go-geos/geometry"
)

func TestGob(t *testing.T) {
	g := geometry.NewGeometry(geos.NewPoint([]float64{1, 2}))
	data, err := g.GobEncode()
	assert.NoError(t, err)
	var geom geometry.Geometry
	assert.NoError(t, geom.GobDecode(data))
	assert.True(t, g.Geom.Equals(geom.Geom))
}
