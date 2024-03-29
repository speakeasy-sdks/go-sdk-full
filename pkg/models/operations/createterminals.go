// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/go-sdk-full/v3/pkg/models/shared"
	"github.com/speakeasy-sdks/go-sdk-full/v3/pkg/utils"
	"net/http"
)

type CreateTerminalsRequest struct {
	CreateTerminalRequest *shared.CreateTerminalRequest `request:"mediaType=application/json"`
	XAPIVersion           *string                       `default:"2022-09-01" header:"style=simple,explode=false,name=x-api-version"`
	XClientID             string                        `header:"style=simple,explode=false,name=x-client-id"`
	XClientSecret         string                        `header:"style=simple,explode=false,name=x-client-secret"`
}

func (c CreateTerminalsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateTerminalsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateTerminalsRequest) GetCreateTerminalRequest() *shared.CreateTerminalRequest {
	if o == nil {
		return nil
	}
	return o.CreateTerminalRequest
}

func (o *CreateTerminalsRequest) GetXAPIVersion() *string {
	if o == nil {
		return nil
	}
	return o.XAPIVersion
}

func (o *CreateTerminalsRequest) GetXClientID() string {
	if o == nil {
		return ""
	}
	return o.XClientID
}

func (o *CreateTerminalsRequest) GetXClientSecret() string {
	if o == nil {
		return ""
	}
	return o.XClientSecret
}

type CreateTerminalsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Terminal created
	TerminalResponse *shared.TerminalResponse
}

func (o *CreateTerminalsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTerminalsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return map[string][]string{}
	}
	return o.Headers
}

func (o *CreateTerminalsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTerminalsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTerminalsResponse) GetTerminalResponse() *shared.TerminalResponse {
	if o == nil {
		return nil
	}
	return o.TerminalResponse
}
