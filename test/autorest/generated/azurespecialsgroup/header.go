// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// HeaderOperations contains the methods for the Header group.
type HeaderOperations interface {
	// CustomNamedRequestID - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
	CustomNamedRequestID(ctx context.Context, fooClientRequestId string) (*HeaderCustomNamedRequestIDResponse, error)
	// CustomNamedRequestIDHead - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
	CustomNamedRequestIDHead(ctx context.Context, fooClientRequestId string) (*HeaderCustomNamedRequestIDHeadResponse, error)
	// CustomNamedRequestIDParamGrouping - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request, via a parameter group
	CustomNamedRequestIDParamGrouping(ctx context.Context, headerCustomNamedRequestIdParamGroupingParameters HeaderCustomNamedRequestIDParamGroupingParameters) (*HeaderCustomNamedRequestIDParamGroupingResponse, error)
}

// headerOperations implements the HeaderOperations interface.
type headerOperations struct {
	*Client
}

// CustomNamedRequestID - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
func (client *headerOperations) CustomNamedRequestID(ctx context.Context, fooClientRequestId string) (*HeaderCustomNamedRequestIDResponse, error) {
	req, err := client.customNamedRequestIdCreateRequest(fooClientRequestId)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.customNamedRequestIdHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// customNamedRequestIdCreateRequest creates the CustomNamedRequestID request.
func (client *headerOperations) customNamedRequestIdCreateRequest(fooClientRequestId string) (*azcore.Request, error) {
	urlPath := "/azurespecials/customNamedRequestId"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	req.Header.Set("foo-client-request-id", fooClientRequestId)
	return req, nil
}

// customNamedRequestIdHandleResponse handles the CustomNamedRequestID response.
func (client *headerOperations) customNamedRequestIdHandleResponse(resp *azcore.Response) (*HeaderCustomNamedRequestIDResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.customNamedRequestIdHandleError(resp)
	}
	result := HeaderCustomNamedRequestIDResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("foo-request-id"); val != "" {
		result.FooRequestId = &val
	}
	return &result, nil
}

// customNamedRequestIdHandleError handles the CustomNamedRequestID error response.
func (client *headerOperations) customNamedRequestIdHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// CustomNamedRequestIDHead - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request
func (client *headerOperations) CustomNamedRequestIDHead(ctx context.Context, fooClientRequestId string) (*HeaderCustomNamedRequestIDHeadResponse, error) {
	req, err := client.customNamedRequestIdHeadCreateRequest(fooClientRequestId)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.customNamedRequestIdHeadHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// customNamedRequestIdHeadCreateRequest creates the CustomNamedRequestIDHead request.
func (client *headerOperations) customNamedRequestIdHeadCreateRequest(fooClientRequestId string) (*azcore.Request, error) {
	urlPath := "/azurespecials/customNamedRequestIdHead"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodHead, *u)
	req.Header.Set("foo-client-request-id", fooClientRequestId)
	return req, nil
}

// customNamedRequestIdHeadHandleResponse handles the CustomNamedRequestIDHead response.
func (client *headerOperations) customNamedRequestIdHeadHandleResponse(resp *azcore.Response) (*HeaderCustomNamedRequestIDHeadResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.customNamedRequestIdHeadHandleError(resp)
	}
	result := HeaderCustomNamedRequestIDHeadResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("foo-request-id"); val != "" {
		result.FooRequestId = &val
	}
	return &result, nil
}

// customNamedRequestIdHeadHandleError handles the CustomNamedRequestIDHead error response.
func (client *headerOperations) customNamedRequestIdHeadHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// CustomNamedRequestIDParamGrouping - Send foo-client-request-id = 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0 in the header of the request, via a parameter group
func (client *headerOperations) CustomNamedRequestIDParamGrouping(ctx context.Context, headerCustomNamedRequestIdParamGroupingParameters HeaderCustomNamedRequestIDParamGroupingParameters) (*HeaderCustomNamedRequestIDParamGroupingResponse, error) {
	req, err := client.customNamedRequestIdParamGroupingCreateRequest(headerCustomNamedRequestIdParamGroupingParameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.customNamedRequestIdParamGroupingHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// customNamedRequestIdParamGroupingCreateRequest creates the CustomNamedRequestIDParamGrouping request.
func (client *headerOperations) customNamedRequestIdParamGroupingCreateRequest(headerCustomNamedRequestIdParamGroupingParameters HeaderCustomNamedRequestIDParamGroupingParameters) (*azcore.Request, error) {
	urlPath := "/azurespecials/customNamedRequestIdParamGrouping"
	u, err := client.u.Parse(urlPath)
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	req.Header.Set("foo-client-request-id", headerCustomNamedRequestIdParamGroupingParameters.FooClientRequestId)
	return req, nil
}

// customNamedRequestIdParamGroupingHandleResponse handles the CustomNamedRequestIDParamGrouping response.
func (client *headerOperations) customNamedRequestIdParamGroupingHandleResponse(resp *azcore.Response) (*HeaderCustomNamedRequestIDParamGroupingResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.customNamedRequestIdParamGroupingHandleError(resp)
	}
	result := HeaderCustomNamedRequestIDParamGroupingResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("foo-request-id"); val != "" {
		result.FooRequestId = &val
	}
	return &result, nil
}

// customNamedRequestIdParamGroupingHandleError handles the CustomNamedRequestIDParamGrouping error response.
func (client *headerOperations) customNamedRequestIdParamGroupingHandleError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}