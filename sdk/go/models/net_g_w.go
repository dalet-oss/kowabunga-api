// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NetGW net g w
//
// swagger:model NetGW
type NetGW struct {

	// The network gateway IPv4 address.
	// Required: true
	Address *string `json:"address"`

	// The network gateway description.
	Description string `json:"description,omitempty"`

	// The network gateway ID (auto-generated).
	ID string `json:"id,omitempty"`

	// The network gateway name.
	// Required: true
	Name *string `json:"name"`

	// The network gateway service port (default to 8080).
	Port *int64 `json:"port,omitempty"`

	// The network gateway admin API token.
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this net g w
func (m *NetGW) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetGW) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *NetGW) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NetGW) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this net g w based on context it is used
func (m *NetGW) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetGW) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetGW) UnmarshalBinary(b []byte) error {
	var res NetGW
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}