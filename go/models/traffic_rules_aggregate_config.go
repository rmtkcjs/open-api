// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TrafficRulesAggregateConfig traffic rules aggregate config
//
// swagger:model trafficRulesAggregateConfig
type TrafficRulesAggregateConfig struct {

	// keys
	Keys []*TrafficRulesAggregateConfigKeysItems `json:"keys"`
}

// Validate validates this traffic rules aggregate config
func (m *TrafficRulesAggregateConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrafficRulesAggregateConfig) validateKeys(formats strfmt.Registry) error {

	if swag.IsZero(m.Keys) { // not required
		return nil
	}

	for i := 0; i < len(m.Keys); i++ {
		if swag.IsZero(m.Keys[i]) { // not required
			continue
		}

		if m.Keys[i] != nil {
			if err := m.Keys[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keys" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrafficRulesAggregateConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrafficRulesAggregateConfig) UnmarshalBinary(b []byte) error {
	var res TrafficRulesAggregateConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
