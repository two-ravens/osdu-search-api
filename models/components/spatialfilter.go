// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type SpatialFilter struct {
	Field           string           `json:"field"`
	ByBoundingBox   *ByBoundingBox   `json:"byBoundingBox,omitempty"`
	ByDistance      *ByDistance      `json:"byDistance,omitempty"`
	ByGeoPolygon    *ByGeoPolygon    `json:"byGeoPolygon,omitempty"`
	ByIntersection  *ByIntersection  `json:"byIntersection,omitempty"`
	ByWithinPolygon *ByWithinPolygon `json:"byWithinPolygon,omitempty"`
}

func (o *SpatialFilter) GetField() string {
	if o == nil {
		return ""
	}
	return o.Field
}

func (o *SpatialFilter) GetByBoundingBox() *ByBoundingBox {
	if o == nil {
		return nil
	}
	return o.ByBoundingBox
}

func (o *SpatialFilter) GetByDistance() *ByDistance {
	if o == nil {
		return nil
	}
	return o.ByDistance
}

func (o *SpatialFilter) GetByGeoPolygon() *ByGeoPolygon {
	if o == nil {
		return nil
	}
	return o.ByGeoPolygon
}

func (o *SpatialFilter) GetByIntersection() *ByIntersection {
	if o == nil {
		return nil
	}
	return o.ByIntersection
}

func (o *SpatialFilter) GetByWithinPolygon() *ByWithinPolygon {
	if o == nil {
		return nil
	}
	return o.ByWithinPolygon
}
