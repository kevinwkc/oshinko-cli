package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/radanalyticsio/oshinko-cli/rest/models"
)

// FindClustersReader is a Reader for the FindClusters structure.
type FindClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFindClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewFindClustersOK creates a FindClustersOK with default headers values
func NewFindClustersOK() *FindClustersOK {
	return &FindClustersOK{}
}

/*FindClustersOK handles this case with default header values.

Clusters response
*/
type FindClustersOK struct {
	Payload FindClustersOKBodyBody
}

func (o *FindClustersOK) Error() string {
	return fmt.Sprintf("[GET /clusters][%d] findClustersOK  %+v", 200, o.Payload)
}

func (o *FindClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindClustersDefault creates a FindClustersDefault with default headers values
func NewFindClustersDefault(code int) *FindClustersDefault {
	return &FindClustersDefault{
		_statusCode: code,
	}
}

/*FindClustersDefault handles this case with default header values.

Unexpected error
*/
type FindClustersDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the find clusters default response
func (o *FindClustersDefault) Code() int {
	return o._statusCode
}

func (o *FindClustersDefault) Error() string {
	return fmt.Sprintf("[GET /clusters][%d] findClusters default  %+v", o._statusCode, o.Payload)
}

func (o *FindClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*FindClustersOKBodyBody find clusters o k body body

swagger:model FindClustersOKBodyBody
*/
type FindClustersOKBodyBody struct {

	/* clusters

	Required: true
	*/
	Clusters []*ClustersItems0 `json:"clusters"`
}

// Validate validates this find clusters o k body body
func (o *FindClustersOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClusters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindClustersOKBodyBody) validateClusters(formats strfmt.Registry) error {

	if err := validate.Required("findClustersOK"+"."+"clusters", "body", o.Clusters); err != nil {
		return err
	}

	for i := 0; i < len(o.Clusters); i++ {

		if swag.IsZero(o.Clusters[i]) { // not required
			continue
		}

		if o.Clusters[i] != nil {

			if err := o.Clusters[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

/*ClustersItems0 clusters items0

swagger:model ClustersItems0
*/
type ClustersItems0 struct {

	/* URL for the cluster information page in oshinko-rest

	Required: true
	*/
	Href *string `json:"href"`

	/* URL to the spark master

	Required: true
	*/
	MasterURL *string `json:"masterUrl"`

	/* URL to the spark master web UI

	Required: true
	*/
	MasterWebURL *string `json:"masterWebUrl"`

	/* Name of the cluster

	Required: true
	*/
	Name *string `json:"name"`

	/* Current state of the cluster

	Required: true
	*/
	Status *string `json:"status"`

	/* Number of worker nodes in the cluster

	Required: true
	*/
	WorkerCount *int64 `json:"workerCount"`
}

// Validate validates this clusters items0
func (o *ClustersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateHref(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateMasterURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateMasterWebURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := o.validateWorkerCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClustersItems0) validateHref(formats strfmt.Registry) error {

	if err := validate.Required("href", "body", o.Href); err != nil {
		return err
	}

	return nil
}

func (o *ClustersItems0) validateMasterURL(formats strfmt.Registry) error {

	if err := validate.Required("masterUrl", "body", o.MasterURL); err != nil {
		return err
	}

	return nil
}

func (o *ClustersItems0) validateMasterWebURL(formats strfmt.Registry) error {

	if err := validate.Required("masterWebUrl", "body", o.MasterWebURL); err != nil {
		return err
	}

	return nil
}

func (o *ClustersItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *ClustersItems0) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

func (o *ClustersItems0) validateWorkerCount(formats strfmt.Registry) error {

	if err := validate.Required("workerCount", "body", o.WorkerCount); err != nil {
		return err
	}

	return nil
}
