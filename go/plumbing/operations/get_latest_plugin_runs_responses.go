// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/v2/go/models"
)

// GetLatestPluginRunsReader is a Reader for the GetLatestPluginRuns structure.
type GetLatestPluginRunsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLatestPluginRunsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLatestPluginRunsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLatestPluginRunsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLatestPluginRunsOK creates a GetLatestPluginRunsOK with default headers values
func NewGetLatestPluginRunsOK() *GetLatestPluginRunsOK {
	return &GetLatestPluginRunsOK{}
}

/*
GetLatestPluginRunsOK handles this case with default header values.

OK
*/
type GetLatestPluginRunsOK struct {
	Payload []*models.PluginRun
}

func (o *GetLatestPluginRunsOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/plugin_runs/latest][%d] getLatestPluginRunsOK  %+v", 200, o.Payload)
}

func (o *GetLatestPluginRunsOK) GetPayload() []*models.PluginRun {
	return o.Payload
}

func (o *GetLatestPluginRunsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestPluginRunsDefault creates a GetLatestPluginRunsDefault with default headers values
func NewGetLatestPluginRunsDefault(code int) *GetLatestPluginRunsDefault {
	return &GetLatestPluginRunsDefault{
		_statusCode: code,
	}
}

/*
GetLatestPluginRunsDefault handles this case with default header values.

error
*/
type GetLatestPluginRunsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get latest plugin runs default response
func (o *GetLatestPluginRunsDefault) Code() int {
	return o._statusCode
}

func (o *GetLatestPluginRunsDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/plugin_runs/latest][%d] getLatestPluginRuns default  %+v", o._statusCode, o.Payload)
}

func (o *GetLatestPluginRunsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLatestPluginRunsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
