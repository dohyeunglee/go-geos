package geometry

import "github.com/twpayne/go-geos"

// initialStringBufferSize is the initial size of strings.Buffers used for
// building GeoJSON and KML representations.
const initialStringBufferSize = 1024

// A Geometry is a geometry.
type Geometry struct {
	*geos.Geom
}

// Must panics with err if err is non-nil, otherwise it returns g.
func Must(g *Geometry, err error) *Geometry {
	if err != nil {
		panic(err)
	}
	return g
}

// NewGeometryFromWKB returns a new Geometry from wkb.
func NewGeometryFromWKB(wkb []byte) (*Geometry, error) {
	geom, err := geos.NewGeomFromWKB(wkb)
	if err != nil {
		return nil, err
	}
	return &Geometry{Geom: geom}, nil
}

// NewGeometryFromWKT returns a new Geometry from wkt.
func NewGeometryFromWKT(wkt string) (*Geometry, error) {
	geom, err := geos.NewGeomFromWKT(wkt)
	if err != nil {
		return nil, err
	}
	return &Geometry{Geom: geom}, nil
}

// Bounds returns g's bounds.
func (g *Geometry) Bounds() *geos.Bounds {
	return g.Geom.Bounds()
}

func (g *Geometry) Destroy() {
	g.Geom.Destroy()
}
