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
	"github.com/go-openapi/swag"
)

// NewListAccountAuditEventsParams creates a new ListAccountAuditEventsParams object
// with the default values initialized.
func NewListAccountAuditEventsParams() *ListAccountAuditEventsParams {
	var ()
	return &ListAccountAuditEventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAccountAuditEventsParamsWithTimeout creates a new ListAccountAuditEventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAccountAuditEventsParamsWithTimeout(timeout time.Duration) *ListAccountAuditEventsParams {
	var ()
	return &ListAccountAuditEventsParams{

		timeout: timeout,
	}
}

// NewListAccountAuditEventsParamsWithContext creates a new ListAccountAuditEventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAccountAuditEventsParamsWithContext(ctx context.Context) *ListAccountAuditEventsParams {
	var ()
	return &ListAccountAuditEventsParams{

		Context: ctx,
	}
}

// NewListAccountAuditEventsParamsWithHTTPClient creates a new ListAccountAuditEventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAccountAuditEventsParamsWithHTTPClient(client *http.Client) *ListAccountAuditEventsParams {
	var ()
	return &ListAccountAuditEventsParams{
		HTTPClient: client,
	}
}

/*
ListAccountAuditEventsParams contains all the parameters to send to the API endpoint
for the list account audit events operation typically these are written to a http.Request
*/
type ListAccountAuditEventsParams struct {

	/*AccountID*/
	AccountID string
	/*LogType*/
	LogType *string
	/*Page*/
	Page *int32
	/*PerPage*/
	PerPage *int32
	/*Query*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list account audit events params
func (o *ListAccountAuditEventsParams) WithTimeout(timeout time.Duration) *ListAccountAuditEventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list account audit events params
func (o *ListAccountAuditEventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list account audit events params
func (o *ListAccountAuditEventsParams) WithContext(ctx context.Context) *ListAccountAuditEventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list account audit events params
func (o *ListAccountAuditEventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list account audit events params
func (o *ListAccountAuditEventsParams) WithHTTPClient(client *http.Client) *ListAccountAuditEventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list account audit events params
func (o *ListAccountAuditEventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the list account audit events params
func (o *ListAccountAuditEventsParams) WithAccountID(accountID string) *ListAccountAuditEventsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the list account audit events params
func (o *ListAccountAuditEventsParams) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WithLogType adds the logType to the list account audit events params
func (o *ListAccountAuditEventsParams) WithLogType(logType *string) *ListAccountAuditEventsParams {
	o.SetLogType(logType)
	return o
}

// SetLogType adds the logType to the list account audit events params
func (o *ListAccountAuditEventsParams) SetLogType(logType *string) {
	o.LogType = logType
}

// WithPage adds the page to the list account audit events params
func (o *ListAccountAuditEventsParams) WithPage(page *int32) *ListAccountAuditEventsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list account audit events params
func (o *ListAccountAuditEventsParams) SetPage(page *int32) {
	o.Page = page
}

// WithPerPage adds the perPage to the list account audit events params
func (o *ListAccountAuditEventsParams) WithPerPage(perPage *int32) *ListAccountAuditEventsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the list account audit events params
func (o *ListAccountAuditEventsParams) SetPerPage(perPage *int32) {
	o.PerPage = perPage
}

// WithQuery adds the query to the list account audit events params
func (o *ListAccountAuditEventsParams) WithQuery(query *string) *ListAccountAuditEventsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the list account audit events params
func (o *ListAccountAuditEventsParams) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *ListAccountAuditEventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param account_id
	if err := r.SetPathParam("account_id", o.AccountID); err != nil {
		return err
	}

	if o.LogType != nil {

		// query param log_type
		var qrLogType string
		if o.LogType != nil {
			qrLogType = *o.LogType
		}
		qLogType := qrLogType
		if qLogType != "" {
			if err := r.SetQueryParam("log_type", qLogType); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int32
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt32(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}

	}

	if o.Query != nil {

		// query param query
		var qrQuery string
		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {
			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
