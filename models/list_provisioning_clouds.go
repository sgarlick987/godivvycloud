// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ListProvisioningClouds ListProvisioningClouds
// swagger:model ListProvisioningClouds
type ListProvisioningClouds struct {

	// clouds
	// Required: true
	Clouds []*Cloud2 `json:"clouds"`
}

// Validate validates this list provisioning clouds
func (m *ListProvisioningClouds) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClouds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListProvisioningClouds) validateClouds(formats strfmt.Registry) error {

	if err := validate.Required("clouds", "body", m.Clouds); err != nil {
		return err
	}

	for i := 0; i < len(m.Clouds); i++ {
		if swag.IsZero(m.Clouds[i]) { // not required
			continue
		}

		if m.Clouds[i] != nil {
			if err := m.Clouds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clouds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListProvisioningClouds) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListProvisioningClouds) UnmarshalBinary(b []byte) error {
	var res ListProvisioningClouds
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
