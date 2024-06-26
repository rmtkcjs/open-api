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

// NewGetDNSRecordsParams creates a new GetDNSRecordsParams object
// with the default values initialized.
func NewGetDNSRecordsParams() *GetDNSRecordsParams {
	var ()
	return &GetDNSRecordsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDNSRecordsParamsWithTimeout creates a new GetDNSRecordsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDNSRecordsParamsWithTimeout(timeout time.Duration) *GetDNSRecordsParams {
	var ()
	return &GetDNSRecordsParams{

		timeout: timeout,
	}
}

// NewGetDNSRecordsParamsWithContext creates a new GetDNSRecordsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDNSRecordsParamsWithContext(ctx context.Context) *GetDNSRecordsParams {
	var ()
	return &GetDNSRecordsParams{

		Context: ctx,
	}
}

// NewGetDNSRecordsParamsWithHTTPClient creates a new GetDNSRecordsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDNSRecordsParamsWithHTTPClient(client *http.Client) *GetDNSRecordsParams {
	var ()
	return &GetDNSRecordsParams{
		HTTPClient: client,
	}
}

/*
GetDNSRecordsParams contains all the parameters to send to the API endpoint
for the get Dns records operation typically these are written to a http.Request
*/
type GetDNSRecordsParams struct {

	/*ZoneID*/
	ZoneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Dns records params
func (o *GetDNSRecordsParams) WithTimeout(timeout time.Duration) *GetDNSRecordsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Dns records params
func (o *GetDNSRecordsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Dns records params
func (o *GetDNSRecordsParams) WithContext(ctx context.Context) *GetDNSRecordsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Dns records params
func (o *GetDNSRecordsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Dns records params
func (o *GetDNSRecordsParams) WithHTTPClient(client *http.Client) *GetDNSRecordsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Dns records params
func (o *GetDNSRecordsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithZoneID adds the zoneID to the get Dns records params
func (o *GetDNSRecordsParams) WithZoneID(zoneID string) *GetDNSRecordsParams {
	o.SetZoneID(zoneID)
	return o
}

// SetZoneID adds the zoneId to the get Dns records params
func (o *GetDNSRecordsParams) SetZoneID(zoneID string) {
	o.ZoneID = zoneID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDNSRecordsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param zone_id
	if err := r.SetPathParam("zone_id", o.ZoneID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
