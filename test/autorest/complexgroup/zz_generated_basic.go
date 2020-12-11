// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// BasicClient contains the methods for the Basic group.
// Don't use this type directly, use NewBasicClient() instead.
type BasicClient struct {
	con *Connection
}

// NewBasicClient creates a new instance of BasicClient with the specified values.
func NewBasicClient(con *Connection) *BasicClient {
	return &BasicClient{con: con}
}

// GetEmpty - Get a basic complex type that is empty
func (client *BasicClient) GetEmpty(ctx context.Context, options *BasicGetEmptyOptions) (BasicResponse, error) {
	req, err := client.getEmptyCreateRequest(ctx, options)
	if err != nil {
		return BasicResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BasicResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BasicResponse{}, client.getEmptyHandleError(resp)
	}
	return client.getEmptyHandleResponse(resp)
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *BasicClient) getEmptyCreateRequest(ctx context.Context, options *BasicGetEmptyOptions) (*azcore.Request, error) {
	urlPath := "/complex/basic/empty"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *BasicClient) getEmptyHandleResponse(resp *azcore.Response) (BasicResponse, error) {
	var val *Basic
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return BasicResponse{}, err
	}
	return BasicResponse{RawResponse: resp.Response, Basic: val}, nil
}

// getEmptyHandleError handles the GetEmpty error response.
func (client *BasicClient) getEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetInvalid - Get a basic complex type that is invalid for the local strong type
func (client *BasicClient) GetInvalid(ctx context.Context, options *BasicGetInvalidOptions) (BasicResponse, error) {
	req, err := client.getInvalidCreateRequest(ctx, options)
	if err != nil {
		return BasicResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BasicResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BasicResponse{}, client.getInvalidHandleError(resp)
	}
	return client.getInvalidHandleResponse(resp)
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *BasicClient) getInvalidCreateRequest(ctx context.Context, options *BasicGetInvalidOptions) (*azcore.Request, error) {
	urlPath := "/complex/basic/invalid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *BasicClient) getInvalidHandleResponse(resp *azcore.Response) (BasicResponse, error) {
	var val *Basic
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return BasicResponse{}, err
	}
	return BasicResponse{RawResponse: resp.Response, Basic: val}, nil
}

// getInvalidHandleError handles the GetInvalid error response.
func (client *BasicClient) getInvalidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNotProvided - Get a basic complex type while the server doesn't provide a response payload
func (client *BasicClient) GetNotProvided(ctx context.Context, options *BasicGetNotProvidedOptions) (BasicResponse, error) {
	req, err := client.getNotProvidedCreateRequest(ctx, options)
	if err != nil {
		return BasicResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BasicResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BasicResponse{}, client.getNotProvidedHandleError(resp)
	}
	return client.getNotProvidedHandleResponse(resp)
}

// getNotProvidedCreateRequest creates the GetNotProvided request.
func (client *BasicClient) getNotProvidedCreateRequest(ctx context.Context, options *BasicGetNotProvidedOptions) (*azcore.Request, error) {
	urlPath := "/complex/basic/notprovided"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNotProvidedHandleResponse handles the GetNotProvided response.
func (client *BasicClient) getNotProvidedHandleResponse(resp *azcore.Response) (BasicResponse, error) {
	var val *Basic
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return BasicResponse{}, err
	}
	return BasicResponse{RawResponse: resp.Response, Basic: val}, nil
}

// getNotProvidedHandleError handles the GetNotProvided error response.
func (client *BasicClient) getNotProvidedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNull - Get a basic complex type whose properties are null
func (client *BasicClient) GetNull(ctx context.Context, options *BasicGetNullOptions) (BasicResponse, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return BasicResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BasicResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BasicResponse{}, client.getNullHandleError(resp)
	}
	return client.getNullHandleResponse(resp)
}

// getNullCreateRequest creates the GetNull request.
func (client *BasicClient) getNullCreateRequest(ctx context.Context, options *BasicGetNullOptions) (*azcore.Request, error) {
	urlPath := "/complex/basic/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *BasicClient) getNullHandleResponse(resp *azcore.Response) (BasicResponse, error) {
	var val *Basic
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return BasicResponse{}, err
	}
	return BasicResponse{RawResponse: resp.Response, Basic: val}, nil
}

// getNullHandleError handles the GetNull error response.
func (client *BasicClient) getNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetValid - Get complex type {id: 2, name: 'abc', color: 'YELLOW'}
func (client *BasicClient) GetValid(ctx context.Context, options *BasicGetValidOptions) (BasicResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return BasicResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BasicResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BasicResponse{}, client.getValidHandleError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *BasicClient) getValidCreateRequest(ctx context.Context, options *BasicGetValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/basic/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *BasicClient) getValidHandleResponse(resp *azcore.Response) (BasicResponse, error) {
	var val *Basic
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return BasicResponse{}, err
	}
	return BasicResponse{RawResponse: resp.Response, Basic: val}, nil
}

// getValidHandleError handles the GetValid error response.
func (client *BasicClient) getValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutValid - Please put {id: 2, name: 'abc', color: 'Magenta'}
func (client *BasicClient) PutValid(ctx context.Context, complexBody Basic, options *BasicPutValidOptions) (*http.Response, error) {
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putValidHandleError(resp)
	}
	return resp.Response, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *BasicClient) putValidCreateRequest(ctx context.Context, complexBody Basic, options *BasicPutValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/basic/valid"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2016-02-29")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putValidHandleError handles the PutValid error response.
func (client *BasicClient) putValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}