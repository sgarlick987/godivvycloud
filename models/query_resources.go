// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// QueryResources QueryResources
// swagger:model QueryResources
type QueryResources struct {

	// counts
	// Required: true
	Counts *Counts1 `json:"counts"`

	// filters
	// Required: true
	Filters []string `json:"filters"`

	// limit
	// Required: true
	Limit *int32 `json:"limit"`

	// offset
	// Required: true
	Offset *int32 `json:"offset"`

	// order by
	// Required: true
	OrderBy *string `json:"order_by"`

	// resources
	// Required: true
	Resources []string `json:"resources"`

	// scopes
	// Required: true
	Scopes []string `json:"scopes"`

	// selected resource type
	SelectedResourceType string `json:"selected_resource_type,omitempty"`

	// supported types
	// Required: true
	SupportedTypes []string `json:"supported_types"`
}

// Validate validates this query resources
func (m *QueryResources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffset(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScopes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportedTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueryResources) validateCounts(formats strfmt.Registry) error {

	if err := validate.Required("counts", "body", m.Counts); err != nil {
		return err
	}

	if m.Counts != nil {
		if err := m.Counts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counts")
			}
			return err
		}
	}

	return nil
}

func (m *QueryResources) validateFilters(formats strfmt.Registry) error {

	if err := validate.Required("filters", "body", m.Filters); err != nil {
		return err
	}

	return nil
}

func (m *QueryResources) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	return nil
}

func (m *QueryResources) validateOffset(formats strfmt.Registry) error {

	if err := validate.Required("offset", "body", m.Offset); err != nil {
		return err
	}

	return nil
}

func (m *QueryResources) validateOrderBy(formats strfmt.Registry) error {

	if err := validate.Required("order_by", "body", m.OrderBy); err != nil {
		return err
	}

	return nil
}

func (m *QueryResources) validateResources(formats strfmt.Registry) error {

	if err := validate.Required("resources", "body", m.Resources); err != nil {
		return err
	}

	return nil
}

func (m *QueryResources) validateScopes(formats strfmt.Registry) error {

	if err := validate.Required("scopes", "body", m.Scopes); err != nil {
		return err
	}

	return nil
}

func (m *QueryResources) validateSupportedTypes(formats strfmt.Registry) error {

	if err := validate.Required("supported_types", "body", m.SupportedTypes); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueryResources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueryResources) UnmarshalBinary(b []byte) error {
	var res QueryResources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}