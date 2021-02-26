// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

type dataFlowClient struct {
	con *connection
}

// CreateOrUpdateDataFlow - Creates or updates a data flow.
func (client *dataFlowClient) createOrUpdateDataFlow(ctx context.Context, dataFlowName string, dataFlow DataFlowResource, options *DataFlowBeginCreateOrUpdateDataFlowOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateDataFlowCreateRequest(ctx, dataFlowName, dataFlow, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateDataFlowHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateDataFlowCreateRequest creates the CreateOrUpdateDataFlow request.
func (client *dataFlowClient) createOrUpdateDataFlowCreateRequest(ctx context.Context, dataFlowName string, dataFlow DataFlowResource, options *DataFlowBeginCreateOrUpdateDataFlowOptions) (*azcore.Request, error) {
	urlPath := "/dataflows/{dataFlowName}"
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(dataFlow)
}

// createOrUpdateDataFlowHandleResponse handles the CreateOrUpdateDataFlow response.
func (client *dataFlowClient) createOrUpdateDataFlowHandleResponse(resp *azcore.Response) (DataFlowResourceResponse, error) {
	var val *DataFlowResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DataFlowResourceResponse{}, err
	}
	return DataFlowResourceResponse{RawResponse: resp.Response, DataFlowResource: val}, nil
}

// createOrUpdateDataFlowHandleError handles the CreateOrUpdateDataFlow error response.
func (client *dataFlowClient) createOrUpdateDataFlowHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// DeleteDataFlow - Deletes a data flow.
func (client *dataFlowClient) deleteDataFlow(ctx context.Context, dataFlowName string, options *DataFlowBeginDeleteDataFlowOptions) (*azcore.Response, error) {
	req, err := client.deleteDataFlowCreateRequest(ctx, dataFlowName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteDataFlowHandleError(resp)
	}
	return resp, nil
}

// deleteDataFlowCreateRequest creates the DeleteDataFlow request.
func (client *dataFlowClient) deleteDataFlowCreateRequest(ctx context.Context, dataFlowName string, options *DataFlowBeginDeleteDataFlowOptions) (*azcore.Request, error) {
	urlPath := "/dataflows/{dataFlowName}"
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteDataFlowHandleError handles the DeleteDataFlow error response.
func (client *dataFlowClient) deleteDataFlowHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDataFlow - Gets a data flow.
func (client *dataFlowClient) GetDataFlow(ctx context.Context, dataFlowName string, options *DataFlowGetDataFlowOptions) (DataFlowResourceResponse, error) {
	req, err := client.getDataFlowCreateRequest(ctx, dataFlowName, options)
	if err != nil {
		return DataFlowResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return DataFlowResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DataFlowResourceResponse{}, client.getDataFlowHandleError(resp)
	}
	return client.getDataFlowHandleResponse(resp)
}

// getDataFlowCreateRequest creates the GetDataFlow request.
func (client *dataFlowClient) getDataFlowCreateRequest(ctx context.Context, dataFlowName string, options *DataFlowGetDataFlowOptions) (*azcore.Request, error) {
	urlPath := "/dataflows/{dataFlowName}"
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getDataFlowHandleResponse handles the GetDataFlow response.
func (client *dataFlowClient) getDataFlowHandleResponse(resp *azcore.Response) (DataFlowResourceResponse, error) {
	var val *DataFlowResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DataFlowResourceResponse{}, err
	}
	return DataFlowResourceResponse{RawResponse: resp.Response, DataFlowResource: val}, nil
}

// getDataFlowHandleError handles the GetDataFlow error response.
func (client *dataFlowClient) getDataFlowHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDataFlowsByWorkspace - Lists data flows.
func (client *dataFlowClient) GetDataFlowsByWorkspace(options *DataFlowGetDataFlowsByWorkspaceOptions) DataFlowListResponsePager {
	return &dataFlowListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getDataFlowsByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.getDataFlowsByWorkspaceHandleResponse,
		errorer:   client.getDataFlowsByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp DataFlowListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DataFlowListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// getDataFlowsByWorkspaceCreateRequest creates the GetDataFlowsByWorkspace request.
func (client *dataFlowClient) getDataFlowsByWorkspaceCreateRequest(ctx context.Context, options *DataFlowGetDataFlowsByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/dataflows"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getDataFlowsByWorkspaceHandleResponse handles the GetDataFlowsByWorkspace response.
func (client *dataFlowClient) getDataFlowsByWorkspaceHandleResponse(resp *azcore.Response) (DataFlowListResponseResponse, error) {
	var val *DataFlowListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DataFlowListResponseResponse{}, err
	}
	return DataFlowListResponseResponse{RawResponse: resp.Response, DataFlowListResponse: val}, nil
}

// getDataFlowsByWorkspaceHandleError handles the GetDataFlowsByWorkspace error response.
func (client *dataFlowClient) getDataFlowsByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// RenameDataFlow - Renames a dataflow.
func (client *dataFlowClient) renameDataFlow(ctx context.Context, dataFlowName string, request ArtifactRenameRequest, options *DataFlowBeginRenameDataFlowOptions) (*azcore.Response, error) {
	req, err := client.renameDataFlowCreateRequest(ctx, dataFlowName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.renameDataFlowHandleError(resp)
	}
	return resp, nil
}

// renameDataFlowCreateRequest creates the RenameDataFlow request.
func (client *dataFlowClient) renameDataFlowCreateRequest(ctx context.Context, dataFlowName string, request ArtifactRenameRequest, options *DataFlowBeginRenameDataFlowOptions) (*azcore.Request, error) {
	urlPath := "/dataflows/{dataFlowName}/rename"
	urlPath = strings.ReplaceAll(urlPath, "{dataFlowName}", url.PathEscape(dataFlowName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// renameDataFlowHandleError handles the RenameDataFlow error response.
func (client *dataFlowClient) renameDataFlowHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}