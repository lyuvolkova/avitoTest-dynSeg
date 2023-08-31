// Code generated by go-swagger; DO NOT EDIT.

package segments

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

	"github.com/lyuvolkova/avitoTest-dynSeg/tests/client/models"
)

// NewCreateSegmentParams creates a new CreateSegmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSegmentParams() *CreateSegmentParams {
	return &CreateSegmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSegmentParamsWithTimeout creates a new CreateSegmentParams object
// with the ability to set a timeout on a request.
func NewCreateSegmentParamsWithTimeout(timeout time.Duration) *CreateSegmentParams {
	return &CreateSegmentParams{
		timeout: timeout,
	}
}

// NewCreateSegmentParamsWithContext creates a new CreateSegmentParams object
// with the ability to set a context for a request.
func NewCreateSegmentParamsWithContext(ctx context.Context) *CreateSegmentParams {
	return &CreateSegmentParams{
		Context: ctx,
	}
}

// NewCreateSegmentParamsWithHTTPClient creates a new CreateSegmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSegmentParamsWithHTTPClient(client *http.Client) *CreateSegmentParams {
	return &CreateSegmentParams{
		HTTPClient: client,
	}
}

/*
CreateSegmentParams contains all the parameters to send to the API endpoint

	for the create segment operation.

	Typically these are written to a http.Request.
*/
type CreateSegmentParams struct {

	// Body.
	Body *models.CreateSegmentRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create segment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSegmentParams) WithDefaults() *CreateSegmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create segment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSegmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create segment params
func (o *CreateSegmentParams) WithTimeout(timeout time.Duration) *CreateSegmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create segment params
func (o *CreateSegmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create segment params
func (o *CreateSegmentParams) WithContext(ctx context.Context) *CreateSegmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create segment params
func (o *CreateSegmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create segment params
func (o *CreateSegmentParams) WithHTTPClient(client *http.Client) *CreateSegmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create segment params
func (o *CreateSegmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create segment params
func (o *CreateSegmentParams) WithBody(body *models.CreateSegmentRequest) *CreateSegmentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create segment params
func (o *CreateSegmentParams) SetBody(body *models.CreateSegmentRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSegmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
