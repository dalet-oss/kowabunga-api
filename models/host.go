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

// Host host
//
// swagger:model Host
type Host struct {

	// The host libvirt's IPv4 address.
	// Required: true
	Address *string `json:"address"`

	// Cost associated to the host.
	Cost *Cost `json:"cost,omitempty"`

	// The host description.
	Description string `json:"description,omitempty"`

	// The host ID (auto-generated).
	ID string `json:"id,omitempty"`

	// The host name.
	// Required: true
	Name *string `json:"name"`

	// The host libvirt's port.
	Port int64 `json:"port,omitempty"`

	// The protocol to use to issue libvirt connection.
	// Required: true
	// Enum: [tcp tls]
	Protocol *string `json:"protocol"`

	// tls
	TLS *HostTLS `json:"tls,omitempty"`
}

// Validate validates this host
func (m *Host) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCost(formats); err != nil {
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

func (m *Host) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateCost(formats strfmt.Registry) error {
	if swag.IsZero(m.Cost) { // not required
		return nil
	}

	if m.Cost != nil {
		if err := m.Cost.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cost")
			}
			return err
		}
	}

	return nil
}

func (m *Host) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var hostTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcp","tls"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hostTypeProtocolPropEnum = append(hostTypeProtocolPropEnum, v)
	}
}

const (

	// HostProtocolTCP captures enum value "tcp"
	HostProtocolTCP string = "tcp"

	// HostProtocolTLS captures enum value "tls"
	HostProtocolTLS string = "tls"
)

// prop value enum
func (m *Host) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, hostTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Host) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", m.Protocol); err != nil {
		return err
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", *m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *Host) validateTLS(formats strfmt.Registry) error {
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

// ContextValidate validate this host based on the context it is used
func (m *Host) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTLS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Host) contextValidateCost(ctx context.Context, formats strfmt.Registry) error {

	if m.Cost != nil {
		if err := m.Cost.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cost")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cost")
			}
			return err
		}
	}

	return nil
}

func (m *Host) contextValidateTLS(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Host) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Host) UnmarshalBinary(b []byte) error {
	var res Host
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// HostTLS The host libvirt's TLS configuration.
//
// swagger:model HostTLS
type HostTLS struct {

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

// Validate validates this host TLS
func (m *HostTLS) Validate(formats strfmt.Registry) error {
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

func (m *HostTLS) validateCa(formats strfmt.Registry) error {

	if err := validate.Required("tls"+"."+"ca", "body", m.Ca); err != nil {
		return err
	}

	return nil
}

func (m *HostTLS) validateCert(formats strfmt.Registry) error {

	if err := validate.Required("tls"+"."+"cert", "body", m.Cert); err != nil {
		return err
	}

	return nil
}

func (m *HostTLS) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("tls"+"."+"key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this host TLS based on context it is used
func (m *HostTLS) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostTLS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostTLS) UnmarshalBinary(b []byte) error {
	var res HostTLS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
