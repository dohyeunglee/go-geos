package geometry_test

import (
	"database/sql"
	"database/sql/driver"

	"github.com/dohyeunglee/go-geos/geometry"
)

var (
	_ driver.Value = &geometry.Geometry{}
	_ sql.Scanner  = &geometry.Geometry{}
)
