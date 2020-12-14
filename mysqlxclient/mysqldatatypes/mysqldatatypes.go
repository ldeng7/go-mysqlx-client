package mysqldatatypes

import (
	"fmt"

	"github.com/ldeng7/go-mysql-datatypes/mysqldatatypes/spatial"
)

type Time struct {
	Positive bool
	Hour     uint16
	Minute   uint8
	Second   uint8
}

func (t *Time) String() string {
	if t.Positive {
		return fmt.Sprintf("%d:%02d:%02d", t.Hour, t.Minute, t.Second)
	}
	return fmt.Sprintf("-%d:%02d:%02d", t.Hour, t.Minute, t.Second)
}

type Point = spatial.Point
type LineString = spatial.LineString
type Polygon = spatial.Polygon
type MultiPoint = spatial.MultiPoint
type MultiLineString = spatial.MultiLineString
type MultiPolygon = spatial.MultiPolygon
type GeometryCollection = spatial.GeometryCollection
type GenericGeometry = spatial.GenericGeometry
