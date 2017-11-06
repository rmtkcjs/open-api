package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*AssetForm asset form

swagger:model assetForm
*/
type AssetForm struct {

	/* fields
	 */
	Fields map[string]string `json:"fields,omitempty"`

	/* url
	 */
	URL string `json:"url,omitempty"`
}

// Validate validates this asset form
func (m *AssetForm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFields(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetForm) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	if err := validate.Required("fields", "body", m.Fields); err != nil {
		return err
	}

	return nil
}