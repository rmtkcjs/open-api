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

// NewListSiteDevServerHooksParams creates a new ListSiteDevServerHooksParams object
// with the default values initialized.
func NewListSiteDevServerHooksParams() *ListSiteDevServerHooksParams {
	var ()
	return &ListSiteDevServerHooksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSiteDevServerHooksParamsWithTimeout creates a new ListSiteDevServerHooksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSiteDevServerHooksParamsWithTimeout(timeout time.Duration) *ListSiteDevServerHooksParams {
	var ()
	return &ListSiteDevServerHooksParams{

		timeout: timeout,
	}
}

// NewListSiteDevServerHooksParamsWithContext creates a new ListSiteDevServerHooksParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSiteDevServerHooksParamsWithContext(ctx context.Context) *ListSiteDevServerHooksParams {
	var ()
	return &ListSiteDevServerHooksParams{

		Context: ctx,
	}
}

// NewListSiteDevServerHooksParamsWithHTTPClient creates a new ListSiteDevServerHooksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSiteDevServerHooksParamsWithHTTPClient(client *http.Client) *ListSiteDevServerHooksParams {
	var ()
	return &ListSiteDevServerHooksParams{
		HTTPClient: client,
	}
}

/*
ListSiteDevServerHooksParams contains all the parameters to send to the API endpoint
for the list site dev server hooks operation typically these are written to a http.Request
*/
type ListSiteDevServerHooksParams struct {

	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list site dev server hooks params
func (o *ListSiteDevServerHooksParams) WithTimeout(timeout time.Duration) *ListSiteDevServerHooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list site dev server hooks params
func (o *ListSiteDevServerHooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list site dev server hooks params
func (o *ListSiteDevServerHooksParams) WithContext(ctx context.Context) *ListSiteDevServerHooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list site dev server hooks params
func (o *ListSiteDevServerHooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list site dev server hooks params
func (o *ListSiteDevServerHooksParams) WithHTTPClient(client *http.Client) *ListSiteDevServerHooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list site dev server hooks params
func (o *ListSiteDevServerHooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the list site dev server hooks params
func (o *ListSiteDevServerHooksParams) WithSiteID(siteID string) *ListSiteDevServerHooksParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the list site dev server hooks params
func (o *ListSiteDevServerHooksParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *ListSiteDevServerHooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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