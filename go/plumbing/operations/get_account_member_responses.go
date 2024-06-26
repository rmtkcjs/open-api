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

// GetAccountMemberReader is a Reader for the GetAccountMember structure.
type GetAccountMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetAccountMemberDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAccountMemberOK creates a GetAccountMemberOK with default headers values
func NewGetAccountMemberOK() *GetAccountMemberOK {
	return &GetAccountMemberOK{}
}

/*
GetAccountMemberOK handles this case with default header values.

OK
*/
type GetAccountMemberOK struct {
	Payload *models.Member
}

func (o *GetAccountMemberOK) Error() string {
	return fmt.Sprintf("[GET /{account_slug}/members/{member_id}][%d] getAccountMemberOK  %+v", 200, o.Payload)
}

func (o *GetAccountMemberOK) GetPayload() *models.Member {
	return o.Payload
}

func (o *GetAccountMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Member)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountMemberDefault creates a GetAccountMemberDefault with default headers values
func NewGetAccountMemberDefault(code int) *GetAccountMemberDefault {
	return &GetAccountMemberDefault{
		_statusCode: code,
	}
}

/*
GetAccountMemberDefault handles this case with default header values.

error
*/
type GetAccountMemberDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get account member default response
func (o *GetAccountMemberDefault) Code() int {
	return o._statusCode
}

func (o *GetAccountMemberDefault) Error() string {
	return fmt.Sprintf("[GET /{account_slug}/members/{member_id}][%d] getAccountMember default  %+v", o._statusCode, o.Payload)
}

func (o *GetAccountMemberDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetAccountMemberDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
