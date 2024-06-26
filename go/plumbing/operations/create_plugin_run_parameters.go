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

	"github.com/netlify/open-api/v2/go/models"
)

// NewCreatePluginRunParams creates a new CreatePluginRunParams object
// with the default values initialized.
func NewCreatePluginRunParams() *CreatePluginRunParams {
	var ()
	return &CreatePluginRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePluginRunParamsWithTimeout creates a new CreatePluginRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePluginRunParamsWithTimeout(timeout time.Duration) *CreatePluginRunParams {
	var ()
	return &CreatePluginRunParams{

		timeout: timeout,
	}
}

// NewCreatePluginRunParamsWithContext creates a new CreatePluginRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePluginRunParamsWithContext(ctx context.Context) *CreatePluginRunParams {
	var ()
	return &CreatePluginRunParams{

		Context: ctx,
	}
}

// NewCreatePluginRunParamsWithHTTPClient creates a new CreatePluginRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePluginRunParamsWithHTTPClient(client *http.Client) *CreatePluginRunParams {
	var ()
	return &CreatePluginRunParams{
		HTTPClient: client,
	}
}

/*
CreatePluginRunParams contains all the parameters to send to the API endpoint
for the create plugin run operation typically these are written to a http.Request
*/
type CreatePluginRunParams struct {

	/*DeployID*/
	DeployID string
	/*PluginRun*/
	PluginRun *models.PluginRunData

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create plugin run params
func (o *CreatePluginRunParams) WithTimeout(timeout time.Duration) *CreatePluginRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create plugin run params
func (o *CreatePluginRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create plugin run params
func (o *CreatePluginRunParams) WithContext(ctx context.Context) *CreatePluginRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create plugin run params
func (o *CreatePluginRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create plugin run params
func (o *CreatePluginRunParams) WithHTTPClient(client *http.Client) *CreatePluginRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create plugin run params
func (o *CreatePluginRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeployID adds the deployID to the create plugin run params
func (o *CreatePluginRunParams) WithDeployID(deployID string) *CreatePluginRunParams {
	o.SetDeployID(deployID)
	return o
}

// SetDeployID adds the deployId to the create plugin run params
func (o *CreatePluginRunParams) SetDeployID(deployID string) {
	o.DeployID = deployID
}

// WithPluginRun adds the pluginRun to the create plugin run params
func (o *CreatePluginRunParams) WithPluginRun(pluginRun *models.PluginRunData) *CreatePluginRunParams {
	o.SetPluginRun(pluginRun)
	return o
}

// SetPluginRun adds the pluginRun to the create plugin run params
func (o *CreatePluginRunParams) SetPluginRun(pluginRun *models.PluginRunData) {
	o.PluginRun = pluginRun
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePluginRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deploy_id
	if err := r.SetPathParam("deploy_id", o.DeployID); err != nil {
		return err
	}

	if o.PluginRun != nil {
		if err := r.SetBodyParam(o.PluginRun); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
