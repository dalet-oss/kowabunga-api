// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InstanceState instance state
//
// swagger:model InstanceState
type InstanceState struct {

	// the reason of the state of the VM
	Reason string `json:"reason,omitempty"`

	// the state of the VM
	State string `json:"state,omitempty"`
}

// Validate validates this instance state
func (m *InstanceState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this instance state based on context it is used
func (m *InstanceState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstanceState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceState) UnmarshalBinary(b []byte) error {
	var res InstanceState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}