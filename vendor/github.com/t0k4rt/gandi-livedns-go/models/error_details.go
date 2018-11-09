// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ErrorDetails error details
// swagger:model ErrorDetails
type ErrorDetails struct {

	// description
	Description string `json:"description,omitempty"`

	// location
	Location string `json:"location,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this error details
func (m *ErrorDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorDetails) UnmarshalBinary(b []byte) error {
	var res ErrorDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}