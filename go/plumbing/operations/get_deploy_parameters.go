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

// NewGetDeployParams creates a new GetDeployParams object
// with the default values initialized.
func NewGetDeployParams() *GetDeployParams {
	var ()
	return &GetDeployParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeployParamsWithTimeout creates a new GetDeployParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeployParamsWithTimeout(timeout time.Duration) *GetDeployParams {
	var ()
	return &GetDeployParams{

		timeout: timeout,
	}
}

// NewGetDeployParamsWithContext creates a new GetDeployParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeployParamsWithContext(ctx context.Context) *GetDeployParams {
	var ()
	return &GetDeployParams{

		Context: ctx,
	}
}

// NewGetDeployParamsWithHTTPClient creates a new GetDeployParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeployParamsWithHTTPClient(client *http.Client) *GetDeployParams {
	var ()
	return &GetDeployParams{
		HTTPClient: client,
	}
}

/*
GetDeployParams contains all the parameters to send to the API endpoint
for the get deploy operation typically these are written to a http.Request
*/
type GetDeployParams struct {

	/*DeployID*/
	DeployID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deploy params
func (o *GetDeployParams) WithTimeout(timeout time.Duration) *GetDeployParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deploy params
func (o *GetDeployParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deploy params
func (o *GetDeployParams) WithContext(ctx context.Context) *GetDeployParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deploy params
func (o *GetDeployParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deploy params
func (o *GetDeployParams) WithHTTPClient(client *http.Client) *GetDeployParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deploy params
func (o *GetDeployParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeployID adds the deployID to the get deploy params
func (o *GetDeployParams) WithDeployID(deployID string) *GetDeployParams {
	o.SetDeployID(deployID)
	return o
}

// SetDeployID adds the deployId to the get deploy params
func (o *GetDeployParams) SetDeployID(deployID string) {
	o.DeployID = deployID
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeployParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deploy_id
	if err := r.SetPathParam("deploy_id", o.DeployID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
