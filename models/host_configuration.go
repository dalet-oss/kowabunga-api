// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HostConfiguration host configuration
//
// swagger:model HostConfiguration
type HostConfiguration struct {

	// The host libvirt's IPv4 address.
	// Required: true
	Address *string `json:"address"`

	// The host description.
	Description string `json:"description,omitempty"`

	// The host name.
	// Required: true
	Name *string `json:"name"`

	// The host libvirt's port.
	Port int64 `json:"port,omitempty"`

	// The protocol to use to issue libvirt connection.
	// Required: true
	// Enum: [tcp tls]
	Protocol *string `json:"protocol"`

	// The host libvirt's TLS configuration.
	TLS *HostConfigurationTLS `json:"tls,omitempty"`
}

// Validate validates this host configuration
func (m *HostConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTLS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostConfiguration) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *HostConfiguration) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var hostConfigurationTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcp","tls"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hostConfigurationTypeProtocolPropEnum = append(hostConfigurationTypeProtocolPropEnum, v)
	}
}

const (

	// HostConfigurationProtocolTCP captures enum value "tcp"
	HostConfigurationProtocolTCP string = "tcp"

	// HostConfigurationProtocolTLS captures enum value "tls"
	HostConfigurationProtocolTLS string = "tls"
)

// prop value enum
func (m *HostConfiguration) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hostConfigurationTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HostConfiguration) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *HostConfiguration) validateTLS(formats strfmt.Registry) error {
	if swag.IsZero(m.TLS) { // not required
		return nil
	}

	if m.TLS != nil {
		if err := m.TLS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this host configuration based on the context it is used
func (m *HostConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTLS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostConfiguration) contextValidateTLS(ctx context.Context, formats strfmt.Registry) error {

	if m.TLS != nil {
		if err := m.TLS.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tls")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tls")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostConfiguration) UnmarshalBinary(b []byte) error {
	var res HostConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}