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

// HostConfigurationTLS host configuration TLS
//
// swagger:model HostConfigurationTLS
type HostConfigurationTLS struct {

	// The TLS certificate of authority.
	// Required: true
	Ca *string `json:"ca"`

	// The TLS client public cert.
	// Required: true
	Cert *string `json:"cert"`

	// The TLS client private key.
	// Required: true
	Key *string `json:"key"`
}

// Validate validates this host configuration TLS
func (m *HostConfigurationTLS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCert(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostConfigurationTLS) validateCa(formats strfmt.Registry) error {

	if err := validate.Required("ca", "body", m.Ca); err != nil {
		return err
	}

	return nil
}

func (m *HostConfigurationTLS) validateCert(formats strfmt.Registry) error {

	if err := validate.Required("cert", "body", m.Cert); err != nil {
		return err
	}

	return nil
}

func (m *HostConfigurationTLS) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this host configuration TLS based on context it is used
func (m *HostConfigurationTLS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostConfigurationTLS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostConfigurationTLS) UnmarshalBinary(b []byte) error {
	var res HostConfigurationTLS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}