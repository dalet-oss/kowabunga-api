// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ZoneSubnet A zone/subnet map.
//
// swagger:model ZoneSubnet
type ZoneSubnet struct {

	// The zone key.
	Key string `json:"key,omitempty"`

	// The subnet ID.
	Value string `json:"value,omitempty"`
}

// Validate validates this zone subnet
func (m *ZoneSubnet) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this zone subnet based on context it is used
func (m *ZoneSubnet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ZoneSubnet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZoneSubnet) UnmarshalBinary(b []byte) error {
	var res ZoneSubnet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
