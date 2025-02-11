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

// MLASettings m l a settings
//
// swagger:model MLASettings
type MLASettings struct {

	// LoggingEnabled is the flag for enabling logging in user cluster.
	LoggingEnabled bool `json:"loggingEnabled,omitempty"`

	// MonitoringEnabled is the flag for enabling monitoring in user cluster.
	MonitoringEnabled bool `json:"monitoringEnabled,omitempty"`

	// logging resources
	LoggingResources *ResourceRequirements `json:"loggingResources,omitempty"`

	// monitoring resources
	MonitoringResources *ResourceRequirements `json:"monitoringResources,omitempty"`
}

// Validate validates this m l a settings
func (m *MLASettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoggingResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitoringResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MLASettings) validateLoggingResources(formats strfmt.Registry) error {
	if swag.IsZero(m.LoggingResources) { // not required
		return nil
	}

	if m.LoggingResources != nil {
		if err := m.LoggingResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loggingResources")
			}
			return err
		}
	}

	return nil
}

func (m *MLASettings) validateMonitoringResources(formats strfmt.Registry) error {
	if swag.IsZero(m.MonitoringResources) { // not required
		return nil
	}

	if m.MonitoringResources != nil {
		if err := m.MonitoringResources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitoringResources")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this m l a settings based on the context it is used
func (m *MLASettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLoggingResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMonitoringResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MLASettings) contextValidateLoggingResources(ctx context.Context, formats strfmt.Registry) error {

	if m.LoggingResources != nil {
		if err := m.LoggingResources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("loggingResources")
			}
			return err
		}
	}

	return nil
}

func (m *MLASettings) contextValidateMonitoringResources(ctx context.Context, formats strfmt.Registry) error {

	if m.MonitoringResources != nil {
		if err := m.MonitoringResources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("monitoringResources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MLASettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MLASettings) UnmarshalBinary(b []byte) error {
	var res MLASettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
