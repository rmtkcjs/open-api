package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListSiteFormsParams creates a new ListSiteFormsParams object
// with the default values initialized.
func NewListSiteFormsParams() *ListSiteFormsParams {
	var ()
	return &ListSiteFormsParams{}
}

/*ListSiteFormsParams contains all the parameters to send to the API endpoint
for the list site forms operation typically these are written to a http.Request
*/
type ListSiteFormsParams struct {

	/*SiteID*/
	SiteID string
}

// WithSiteID adds the siteId to the list site forms params
func (o *ListSiteFormsParams) WithSiteID(SiteID string) *ListSiteFormsParams {
	o.SiteID = SiteID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ListSiteFormsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}