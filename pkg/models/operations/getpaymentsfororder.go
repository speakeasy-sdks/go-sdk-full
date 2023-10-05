// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk-full/pkg/utils"
	"net/http"
)

type GetPaymentsfororderRequest struct {
	OrderID       string  `pathParam:"style=simple,explode=false,name=order_id"`
	XAPIVersion   *string `default:"2022-09-01" header:"style=simple,explode=false,name=x-api-version"`
	XClientID     string  `header:"style=simple,explode=false,name=x-client-id"`
	XClientSecret string  `header:"style=simple,explode=false,name=x-client-secret"`
}

func (g GetPaymentsfororderRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetPaymentsfororderRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetPaymentsfororderRequest) GetOrderID() string {
	if o == nil {
		return ""
	}
	return o.OrderID
}

func (o *GetPaymentsfororderRequest) GetXAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.XAPIVersion
}

func (o *GetPaymentsfororderRequest) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *GetPaymentsfororderRequest) GetXClientSecret() string {
	if o == nil {
		return ""
	}
	return o.XClientSecret
}

type GetPaymentsfororder200ApplicationJSONType string

const (
	GetPaymentsfororder200ApplicationJSONTypePaymentsEntity GetPaymentsfororder200ApplicationJSONType = "PaymentsEntity"
)

type GetPaymentsfororder200ApplicationJSON struct {
	PaymentsEntity *shared.PaymentsEntity

	Type GetPaymentsfororder200ApplicationJSONType
}

func CreateGetPaymentsfororder200ApplicationJSONPaymentsEntity(paymentsEntity shared.PaymentsEntity) GetPaymentsfororder200ApplicationJSON {
	typ := GetPaymentsfororder200ApplicationJSONTypePaymentsEntity

	return GetPaymentsfororder200ApplicationJSON{
		PaymentsEntity: &paymentsEntity,
		Type:           typ,
	}
}

func (u *GetPaymentsfororder200ApplicationJSON) UnmarshalJSON(data []byte) error {

	paymentsEntity := new(shared.PaymentsEntity)
	if err := utils.UnmarshalJSON(data, &paymentsEntity, "", true, true); err == nil {
		u.PaymentsEntity = paymentsEntity
		u.Type = GetPaymentsfororder200ApplicationJSONTypePaymentsEntity
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u GetPaymentsfororder200ApplicationJSON) MarshalJSON() ([]byte, error) {
	if u.PaymentsEntity != nil {
		return utils.MarshalJSON(u.PaymentsEntity, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type GetPaymentsfororderResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Any bad or invalid request will lead to following error object
	ErrorResponse *shared.ErrorResponse
	// OK
	GetPaymentsfororder200ApplicationJSONOneOf *GetPaymentsfororder200ApplicationJSON
	Headers                                    map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetPaymentsfororderResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPaymentsfororderResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *GetPaymentsfororderResponse) GetGetPaymentsfororder200ApplicationJSONOneOf() *GetPaymentsfororder200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetPaymentsfororder200ApplicationJSONOneOf
}

func (o *GetPaymentsfororderResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *GetPaymentsfororderResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPaymentsfororderResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
