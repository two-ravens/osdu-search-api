// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

type Order string

const (
	OrderAsc  Order = "ASC"
	OrderDesc Order = "DESC"
)

func (e Order) ToPointer() *Order {
	return &e
}
func (e *Order) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ASC":
		fallthrough
	case "DESC":
		*e = Order(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Order: %v", v)
	}
}

type SortQuery struct {
	Field []string `json:"field,omitempty"`
	Order []Order  `json:"order,omitempty"`
}

func (o *SortQuery) GetField() []string {
	if o == nil {
		return nil
	}
	return o.Field
}

func (o *SortQuery) GetOrder() []Order {
	if o == nil {
		return nil
	}
	return o.Order
}
