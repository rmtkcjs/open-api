// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewRollbackSiteDeployParams creates a new RollbackSiteDeployParams object
// with the default values initialized.
func NewRollbackSiteDeployParams() *RollbackSiteDeployParams {
	var ()
	return &RollbackSiteDeployParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRollbackSiteDeployParamsWithTimeout creates a new RollbackSiteDeployParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRollbackSiteDeployParamsWithTimeout(timeout time.Duration) *RollbackSiteDeployParams {
	var ()
	return &RollbackSiteDeployParams{

		timeout: timeout,
	}
}

// NewRollbackSiteDeployParamsWithContext creates a new RollbackSiteDeployParams object
// with the default values initialized, and the ability to set a context for a request
func NewRollbackSiteDeployParamsWithContext(ctx context.Context) *RollbackSiteDeployParams {
	var ()
	return &RollbackSiteDeployParams{

		Context: ctx,
	}
}

// NewRollbackSiteDeployParamsWithHTTPClient creates a new RollbackSiteDeployParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRollbackSiteDeployParamsWithHTTPClient(client *http.Client) *RollbackSiteDeployParams {
	var ()
	return &RollbackSiteDeployParams{
		HTTPClient: client,
	}
}

/*
RollbackSiteDeployParams contains all the parameters to send to the API endpoint
for the rollback site deploy operation typically these are written to a http.Request
*/
type RollbackSiteDeployParams struct {

	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rollback site deploy params
func (o *RollbackSiteDeployParams) WithTimeout(timeout time.Duration) *RollbackSiteDeployParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rollback site deploy params
func (o *RollbackSiteDeployParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rollback site deploy params
func (o *RollbackSiteDeployParams) WithContext(ctx context.Context) *RollbackSiteDeployParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rollback site deploy params
func (o *RollbackSiteDeployParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rollback site deploy params
func (o *RollbackSiteDeployParams) WithHTTPClient(client *http.Client) *RollbackSiteDeployParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rollback site deploy params
func (o *RollbackSiteDeployParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the rollback site deploy params
func (o *RollbackSiteDeployParams) WithSiteID(siteID string) *RollbackSiteDeployParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the rollback site deploy params
func (o *RollbackSiteDeployParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *RollbackSiteDeployParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
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
