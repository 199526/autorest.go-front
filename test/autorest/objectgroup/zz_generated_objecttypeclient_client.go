// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package objectgroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ObjectTypeClient contains the methods for the ObjectTypeClient group.
// Don't use this type directly, use NewObjectTypeClient() instead.
type ObjectTypeClient struct {
	con *Connection
}

// NewObjectTypeClient creates a new instance of ObjectTypeClient with the specified values.
func NewObjectTypeClient(con *Connection) *ObjectTypeClient {
	return &ObjectTypeClient{con: con}
}

// Get - Basic get that returns an object. Returns object { 'message': 'An object was successfully returned' }
func (client *ObjectTypeClient) Get(ctx context.Context, options *ObjectTypeClientGetOptions) (InterfaceResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return InterfaceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return InterfaceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return InterfaceResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ObjectTypeClient) getCreateRequest(ctx context.Context, options *ObjectTypeClientGetOptions) (*azcore.Request, error) {
	urlPath := "/objectType/get"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ObjectTypeClient) getHandleResponse(resp *azcore.Response) (InterfaceResponse, error) {
	var val interface{}
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return InterfaceResponse{}, err
	}
	return InterfaceResponse{RawResponse: resp.Response, Interface: val}, nil
}

// getHandleError handles the Get error response.
func (client *ObjectTypeClient) getHandleError(resp *azcore.Response) error {
	var err interface{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(fmt.Errorf("%v", err), resp.Response)
}

// Put - Basic put that puts an object. Pass in {'foo': 'bar'} to get a 200 and anything else to get an object error.
func (client *ObjectTypeClient) Put(ctx context.Context, putObject interface{}, options *ObjectTypeClientPutOptions) (*http.Response, error) {
	req, err := client.putCreateRequest(ctx, putObject, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putHandleError(resp)
	}
	return resp.Response, nil
}

// putCreateRequest creates the Put request.
func (client *ObjectTypeClient) putCreateRequest(ctx context.Context, putObject interface{}, options *ObjectTypeClientPutOptions) (*azcore.Request, error) {
	urlPath := "/objectType/put"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(putObject)
}

// putHandleError handles the Put error response.
func (client *ObjectTypeClient) putHandleError(resp *azcore.Response) error {
	var err interface{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(fmt.Errorf("%v", err), resp.Response)
}
