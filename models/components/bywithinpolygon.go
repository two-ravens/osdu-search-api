// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type ByWithinPolygon struct {
	Points []Point `json:"points"`
}

func (o *ByWithinPolygon) GetPoints() []Point {
	if o == nil {
		return []Point{}
	}
	return o.Points
}
