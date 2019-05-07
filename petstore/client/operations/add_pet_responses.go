// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tojen/openapi-examples/petstore/models"
)

// AddPetReader is a Reader for the AddPet structure.
type AddPetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddPetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddPetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPetOK creates a AddPetOK with default headers values
func NewAddPetOK() *AddPetOK {
	return &AddPetOK{}
}

/*AddPetOK handles this case with default header values.

pet response
*/
type AddPetOK struct {
	Payload *models.Pet
}

func (o *AddPetOK) Error() string {
	return fmt.Sprintf("[POST /pets][%d] addPetOK  %+v", 200, o.Payload)
}

func (o *AddPetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Pet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPetDefault creates a AddPetDefault with default headers values
func NewAddPetDefault(code int) *AddPetDefault {
	return &AddPetDefault{
		_statusCode: code,
	}
}

/*AddPetDefault handles this case with default header values.

unexpected error
*/
type AddPetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the add pet default response
func (o *AddPetDefault) Code() int {
	return o._statusCode
}

func (o *AddPetDefault) Error() string {
	return fmt.Sprintf("[POST /pets][%d] addPet default  %+v", o._statusCode, o.Payload)
}

func (o *AddPetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
