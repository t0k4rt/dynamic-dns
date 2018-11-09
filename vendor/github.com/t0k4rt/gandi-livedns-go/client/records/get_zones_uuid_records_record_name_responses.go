// Code generated by go-swagger; DO NOT EDIT.

package records

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	models "github.com/t0k4rt/gandi-livedns-go/models"
)

// GetZonesUUIDRecordsRecordNameReader is a Reader for the GetZonesUUIDRecordsRecordName structure.
type GetZonesUUIDRecordsRecordNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZonesUUIDRecordsRecordNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetZonesUUIDRecordsRecordNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetZonesUUIDRecordsRecordNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetZonesUUIDRecordsRecordNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetZonesUUIDRecordsRecordNameOK creates a GetZonesUUIDRecordsRecordNameOK with default headers values
func NewGetZonesUUIDRecordsRecordNameOK() *GetZonesUUIDRecordsRecordNameOK {
	return &GetZonesUUIDRecordsRecordNameOK{}
}

/*GetZonesUUIDRecordsRecordNameOK handles this case with default header values.

OK
*/
type GetZonesUUIDRecordsRecordNameOK struct {
	Payload []*models.Record
}

func (o *GetZonesUUIDRecordsRecordNameOK) Error() string {
	return fmt.Sprintf("[GET /zones/{uuid}/records/{record_name}][%d] getZonesUuidRecordsRecordNameOK  %+v", 200, o.Payload)
}

func (o *GetZonesUUIDRecordsRecordNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZonesUUIDRecordsRecordNameBadRequest creates a GetZonesUUIDRecordsRecordNameBadRequest with default headers values
func NewGetZonesUUIDRecordsRecordNameBadRequest() *GetZonesUUIDRecordsRecordNameBadRequest {
	return &GetZonesUUIDRecordsRecordNameBadRequest{}
}

/*GetZonesUUIDRecordsRecordNameBadRequest handles this case with default header values.

Not OK
*/
type GetZonesUUIDRecordsRecordNameBadRequest struct {
	Payload *models.Return400
}

func (o *GetZonesUUIDRecordsRecordNameBadRequest) Error() string {
	return fmt.Sprintf("[GET /zones/{uuid}/records/{record_name}][%d] getZonesUuidRecordsRecordNameBadRequest  %+v", 400, o.Payload)
}

func (o *GetZonesUUIDRecordsRecordNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return400)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetZonesUUIDRecordsRecordNameDefault creates a GetZonesUUIDRecordsRecordNameDefault with default headers values
func NewGetZonesUUIDRecordsRecordNameDefault(code int) *GetZonesUUIDRecordsRecordNameDefault {
	return &GetZonesUUIDRecordsRecordNameDefault{
		_statusCode: code,
	}
}

/*GetZonesUUIDRecordsRecordNameDefault handles this case with default header values.

Unexpected error
*/
type GetZonesUUIDRecordsRecordNameDefault struct {
	_statusCode int

	Payload *models.Return40x
}

// Code gets the status code for the get zones UUID records record name default response
func (o *GetZonesUUIDRecordsRecordNameDefault) Code() int {
	return o._statusCode
}

func (o *GetZonesUUIDRecordsRecordNameDefault) Error() string {
	return fmt.Sprintf("[GET /zones/{uuid}/records/{record_name}][%d] GetZonesUUIDRecordsRecordName default  %+v", o._statusCode, o.Payload)
}

func (o *GetZonesUUIDRecordsRecordNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Return40x)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
