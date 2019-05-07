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

// DeletePetReader is a Reader for the DeletePet structure.
type DeletePetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeletePetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeletePetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeletePetNoContent creates a DeletePetNoContent with default headers values
func NewDeletePetNoContent() *DeletePetNoContent {
	return &DeletePetNoContent{}
}

/*DeletePetNoContent handles this case with default header values.

pet deleted
*/
type DeletePetNoContent struct {
}

func (o *DeletePetNoContent) Error() string {
	return fmt.Sprintf("[DELETE /pets/{id}][%d] deletePetNoContent ", 204)
}

func (o *DeletePetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePetDefault creates a DeletePetDefault with default headers values
func NewDeletePetDefault(code int) *DeletePetDefault {
	return &DeletePetDefault{
		_statusCode: code,
	}
}

/*DeletePetDefault handles this case with default header values.

unexpected error
*/
type DeletePetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete pet default response
func (o *DeletePetDefault) Code() int {
	return o._statusCode
}

func (o *DeletePetDefault) Error() string {
	return fmt.Sprintf("[DELETE /pets/{id}][%d] deletePet default  %+v", o._statusCode, o.Payload)
}

func (o *DeletePetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
