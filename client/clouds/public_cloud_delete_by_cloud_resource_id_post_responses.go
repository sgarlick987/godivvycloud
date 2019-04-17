// Code generated by go-swagger; DO NOT EDIT.

package clouds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PublicCloudDeleteByCloudResourceIDPostReader is a Reader for the PublicCloudDeleteByCloudResourceIDPost structure.
type PublicCloudDeleteByCloudResourceIDPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicCloudDeleteByCloudResourceIDPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPublicCloudDeleteByCloudResourceIDPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPublicCloudDeleteByCloudResourceIDPostOK creates a PublicCloudDeleteByCloudResourceIDPostOK with default headers values
func NewPublicCloudDeleteByCloudResourceIDPostOK() *PublicCloudDeleteByCloudResourceIDPostOK {
	return &PublicCloudDeleteByCloudResourceIDPostOK{}
}

/*PublicCloudDeleteByCloudResourceIDPostOK handles this case with default header values.

PublicCloudDeleteByCloudResourceIDPostOK public cloud delete by cloud resource Id post o k
*/
type PublicCloudDeleteByCloudResourceIDPostOK struct {
}

func (o *PublicCloudDeleteByCloudResourceIDPostOK) Error() string {
	return fmt.Sprintf("[POST /public/cloud/{cloud_resource_id}/delete][%d] publicCloudDeleteByCloudResourceIdPostOK ", 200)
}

func (o *PublicCloudDeleteByCloudResourceIDPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}