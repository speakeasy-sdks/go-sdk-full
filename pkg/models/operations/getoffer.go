// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
	"net/http"
)

type GetOfferRequest struct {
	OfferID       string  `pathParam:"style=simple,explode=false,name=offer_id"`
	XAPIVersion   *string `header:"style=simple,explode=false,name=x-api-version"`
	XClientID     string  `header:"style=simple,explode=false,name=x-client-id"`
	XClientSecret string  `header:"style=simple,explode=false,name=x-client-secret"`
}

func (o *GetOfferRequest) GetOfferID() string {
	if o == nil {
		return ""
	}
	return o.OfferID
}

func (o *GetOfferRequest) GetXAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.XAPIVersion
}

func (o *GetOfferRequest) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *GetOfferRequest) GetXClientSecret() string {
	if o == nil {
		return ""
	}
	return o.XClientSecret
}

type GetOfferResponse struct {
	ContentType string
	Headers     map[string][]string
	// OK
	OfferEntity *shared.OfferEntity
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetOfferResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetOfferResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetOfferResponse) GetOfferEntity() *shared.OfferEntity {
	if o == nil {
		return nil
	}
	return o.OfferEntity
}

func (o *GetOfferResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetOfferResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
