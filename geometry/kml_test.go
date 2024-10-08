package geometry_test

import (
	"encoding/xml"

	"github.com/dohyeunglee/go-geos/geometry"
)

var _ xml.Marshaler = &geometry.Geometry{}
