// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Instance instance
//
// swagger:model Instance
type Instance struct {

	// The instance ID  (auto-generated).
	ID string `json:"id,omitempty"`

	// the name of the Virtual Machine
	Name string `json:"name,omitempty"`

	// is the VM a template ?
	Template bool `json:"template,omitempty"`

	// topology
	Topology *InstanceTopology `json:"topology,omitempty"`

	// the ID of the Virtual Machine
	VMID string `json:"vm_id,omitempty"`

	// the UUID of the Virtual Machine
	VMUUID string `json:"vm_uuid,omitempty"`
}

// Validate validates this instance
func (m *Instance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTopology(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Instance) validateTopology(formats strfmt.Registry) error {
	if swag.IsZero(m.Topology) { // not required
		return nil
	}

	if m.Topology != nil {
		if err := m.Topology.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("topology")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("topology")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this instance based on the context it is used
func (m *Instance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTopology(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Instance) contextValidateTopology(ctx context.Context, formats strfmt.Registry) error {

	if m.Topology != nil {
		if err := m.Topology.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("topology")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("topology")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Instance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Instance) UnmarshalBinary(b []byte) error {
	var res Instance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}