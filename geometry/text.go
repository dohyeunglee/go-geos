package geometry

import "github.com/dohyeunglee/go-geos"

// NewGeometryFromWKT returns a new Geometry from wkt.
func NewGeometryFromWKT(wkt string) (*Geometry, error) {
	geom, err := geos.NewGeomFromWKT(wkt)
	if err != nil {
		return nil, err
	}
	return &Geometry{Geom: geom}, nil
}

// MarshalText implements encoding.TextMarshaler.
func (g *Geometry) MarshalText() ([]byte, error) {
	return []byte(g.Geom.ToWKT()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (g *Geometry) UnmarshalText(data []byte) error {
	geom, err := geos.NewGeomFromWKT(string(data))
	if err != nil {
		return err
	}
	g.Geom = geom
	return nil
}
