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

// NewGetSplitTestParams creates a new GetSplitTestParams object
// with the default values initialized.
func NewGetSplitTestParams() *GetSplitTestParams {
	var ()
	return &GetSplitTestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSplitTestParamsWithTimeout creates a new GetSplitTestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSplitTestParamsWithTimeout(timeout time.Duration) *GetSplitTestParams {
	var ()
	return &GetSplitTestParams{

		timeout: timeout,
	}
}

// NewGetSplitTestParamsWithContext creates a new GetSplitTestParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSplitTestParamsWithContext(ctx context.Context) *GetSplitTestParams {
	var ()
	return &GetSplitTestParams{

		Context: ctx,
	}
}

// NewGetSplitTestParamsWithHTTPClient creates a new GetSplitTestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSplitTestParamsWithHTTPClient(client *http.Client) *GetSplitTestParams {
	var ()
	return &GetSplitTestParams{
		HTTPClient: client,
	}
}

/*
GetSplitTestParams contains all the parameters to send to the API endpoint
for the get split test operation typically these are written to a http.Request
*/
type GetSplitTestParams struct {

	/*SiteID*/
	SiteID string
	/*SplitTestID*/
	SplitTestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get split test params
func (o *GetSplitTestParams) WithTimeout(timeout time.Duration) *GetSplitTestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get split test params
func (o *GetSplitTestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get split test params
func (o *GetSplitTestParams) WithContext(ctx context.Context) *GetSplitTestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get split test params
func (o *GetSplitTestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get split test params
func (o *GetSplitTestParams) WithHTTPClient(client *http.Client) *GetSplitTestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get split test params
func (o *GetSplitTestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the get split test params
func (o *GetSplitTestParams) WithSiteID(siteID string) *GetSplitTestParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get split test params
func (o *GetSplitTestParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithSplitTestID adds the splitTestID to the get split test params
func (o *GetSplitTestParams) WithSplitTestID(splitTestID string) *GetSplitTestParams {
	o.SetSplitTestID(splitTestID)
	return o
}

// SetSplitTestID adds the splitTestId to the get split test params
func (o *GetSplitTestParams) SetSplitTestID(splitTestID string) {
	o.SplitTestID = splitTestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSplitTestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	// path param split_test_id
	if err := r.SetPathParam("split_test_id", o.SplitTestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
