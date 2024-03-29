// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-opportunity/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteOpportunityRequest struct {
	// Opportunity ID
	OpportunityID string `pathParam:"style=simple,explode=false,name=opportunityId"`
}

func (o *DeleteOpportunityRequest) GetOpportunityID() string {
	if o == nil {
		return ""
	}
	return o.OpportunityID
}

type DeleteOpportunityResponse struct {
	// Any error based on client data errors
	ClientError *shared.ClientError
	// HTTP response content type for this operation
	ContentType string
	// Opportunity response
	Opportunity *shared.Opportunity
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Any error based on the server-side
	ServerError *shared.ServerError
}

func (o *DeleteOpportunityResponse) GetClientError() *shared.ClientError {
	if o == nil {
		return nil
	}
	return o.ClientError
}

func (o *DeleteOpportunityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteOpportunityResponse) GetOpportunity() *shared.Opportunity {
	if o == nil {
		return nil
	}
	return o.Opportunity
}

func (o *DeleteOpportunityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteOpportunityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteOpportunityResponse) GetServerError() *shared.ServerError {
	if o == nil {
		return nil
	}
	return o.ServerError
}
