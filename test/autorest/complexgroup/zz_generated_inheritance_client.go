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

// InheritanceClient contains the methods for the Inheritance group.
// Don't use this type directly, use NewInheritanceClient() instead.
type InheritanceClient struct {
	con *Connection
}

// NewInheritanceClient creates a new instance of InheritanceClient with the specified values.
func NewInheritanceClient(con *Connection) *InheritanceClient {
	return &InheritanceClient{con: con}
}

// GetValid - Get complex types that extend others
func (client *InheritanceClient) GetValid(ctx context.Context, options *InheritanceGetValidOptions) (SiameseResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return SiameseResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SiameseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SiameseResponse{}, client.getValidHandleError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *InheritanceClient) getValidCreateRequest(ctx context.Context, options *InheritanceGetValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/inheritance/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *InheritanceClient) getValidHandleResponse(resp *azcore.Response) (SiameseResponse, error) {
	var val *Siamese
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return SiameseResponse{}, err
	}
	return SiameseResponse{RawResponse: resp.Response, Siamese: val}, nil
}

// getValidHandleError handles the GetValid error response.
func (client *InheritanceClient) getValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutValid - Put complex types that extend others
func (client *InheritanceClient) PutValid(ctx context.Context, complexBody Siamese, options *InheritancePutValidOptions) (*http.Response, error) {
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
func (client *InheritanceClient) putValidCreateRequest(ctx context.Context, complexBody Siamese, options *InheritancePutValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/inheritance/valid"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putValidHandleError handles the PutValid error response.
func (client *InheritanceClient) putValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}
