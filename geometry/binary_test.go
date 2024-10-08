package geometry_test

import (
	"encoding"
	"encoding/hex"
	"strconv"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/dohyeunglee/go-geos"
	"github.com/dohyeunglee/go-geos/geometry"
)

var (
	_ encoding.BinaryMarshaler   = &geometry.Geometry{}
	_ encoding.BinaryUnmarshaler = &geometry.Geometry{}
)

func TestBinary(t *testing.T) {
	for i, tc := range []struct {
		geom      *geometry.Geometry
		binaryStr string
	}{
		{
			geom:      geometry.NewGeometry(geos.NewPoint([]float64{1, 2})),
			binaryStr: "0101000000000000000000f03f0000000000000040",
		},
	} {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actualBinary, err := tc.geom.MarshalBinary()
			assert.NoError(t, err)
			assert.Equal(t, tc.binaryStr, hex.EncodeToString(actualBinary))

			var geom geometry.Geometry
			binary, err := hex.DecodeString(tc.binaryStr)
			assert.NoError(t, err)
			assert.NoError(t, geom.UnmarshalBinary(binary))
			assert.True(t, tc.geom.Equals(geom.Geom))
		})
	}
}
