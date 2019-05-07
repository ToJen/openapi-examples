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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/tojen/openapi-examples/petstore/models"
)

// NewAddPetParams creates a new AddPetParams object
// with the default values initialized.
func NewAddPetParams() *AddPetParams {
	var ()
	return &AddPetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddPetParamsWithTimeout creates a new AddPetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddPetParamsWithTimeout(timeout time.Duration) *AddPetParams {
	var ()
	return &AddPetParams{

		timeout: timeout,
	}
}

// NewAddPetParamsWithContext creates a new AddPetParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddPetParamsWithContext(ctx context.Context) *AddPetParams {
	var ()
	return &AddPetParams{

		Context: ctx,
	}
}

// NewAddPetParamsWithHTTPClient creates a new AddPetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddPetParamsWithHTTPClient(client *http.Client) *AddPetParams {
	var ()
	return &AddPetParams{
		HTTPClient: client,
	}
}

/*AddPetParams contains all the parameters to send to the API endpoint
for the add pet operation typically these are written to a http.Request
*/
type AddPetParams struct {

	/*Pet
	  Pet to add to the store

	*/
	Pet *models.NewPet

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add pet params
func (o *AddPetParams) WithTimeout(timeout time.Duration) *AddPetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add pet params
func (o *AddPetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add pet params
func (o *AddPetParams) WithContext(ctx context.Context) *AddPetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add pet params
func (o *AddPetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add pet params
func (o *AddPetParams) WithHTTPClient(client *http.Client) *AddPetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add pet params
func (o *AddPetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPet adds the pet to the add pet params
func (o *AddPetParams) WithPet(pet *models.NewPet) *AddPetParams {
	o.SetPet(pet)
	return o
}

// SetPet adds the pet to the add pet params
func (o *AddPetParams) SetPet(pet *models.NewPet) {
	o.Pet = pet
}

// WriteToRequest writes these params to a swagger request
func (o *AddPetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Pet != nil {
		if err := r.SetBodyParam(o.Pet); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
