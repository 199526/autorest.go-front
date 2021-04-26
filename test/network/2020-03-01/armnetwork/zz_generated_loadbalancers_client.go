// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// LoadBalancersClient contains the methods for the LoadBalancers group.
// Don't use this type directly, use NewLoadBalancersClient() instead.
type LoadBalancersClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewLoadBalancersClient creates a new instance of LoadBalancersClient with the specified values.
func NewLoadBalancersClient(con *armcore.Connection, subscriptionID string) *LoadBalancersClient {
	return &LoadBalancersClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a load balancer.
func (client *LoadBalancersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersBeginCreateOrUpdateOptions) (LoadBalancerPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, loadBalancerName, parameters, options)
	if err != nil {
		return LoadBalancerPollerResponse{}, err
	}
	result := LoadBalancerPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LoadBalancersClient.CreateOrUpdate", "azure-async-operation", resp, client.createOrUpdateHandleError)
	if err != nil {
		return LoadBalancerPollerResponse{}, err
	}
	poller := &loadBalancerPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (LoadBalancerResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new LoadBalancerPoller from the specified resume token.
// token - The value must come from a previous call to LoadBalancerPoller.ResumeToken().
func (client *LoadBalancersClient) ResumeCreateOrUpdate(ctx context.Context, token string) (LoadBalancerPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("LoadBalancersClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return LoadBalancerPollerResponse{}, err
	}
	poller := &loadBalancerPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return LoadBalancerPollerResponse{}, err
	}
	result := LoadBalancerPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (LoadBalancerResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a load balancer.
func (client *LoadBalancersClient) createOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, loadBalancerName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LoadBalancersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *LoadBalancersClient) createOrUpdateHandleResponse(resp *azcore.Response) (LoadBalancerResponse, error) {
	var val *LoadBalancer
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LoadBalancerResponse{}, err
	}
	return LoadBalancerResponse{RawResponse: resp.Response, LoadBalancer: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *LoadBalancersClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Deletes the specified load balancer.
func (client *LoadBalancersClient) BeginDelete(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, loadBalancerName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("LoadBalancersClient.Delete", "location", resp, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *LoadBalancersClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("LoadBalancersClient.Delete", token, client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes the specified load balancer.
func (client *LoadBalancersClient) deleteOperation(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *LoadBalancersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *LoadBalancersClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Gets the specified load balancer.
func (client *LoadBalancersClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersGetOptions) (LoadBalancerResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
	if err != nil {
		return LoadBalancerResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LoadBalancerResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return LoadBalancerResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LoadBalancersClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *LoadBalancersGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoadBalancersClient) getHandleResponse(resp *azcore.Response) (LoadBalancerResponse, error) {
	var val *LoadBalancer
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LoadBalancerResponse{}, err
	}
	return LoadBalancerResponse{RawResponse: resp.Response, LoadBalancer: val}, nil
}

// getHandleError handles the Get error response.
func (client *LoadBalancersClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - Gets all the load balancers in a resource group.
func (client *LoadBalancersClient) List(resourceGroupName string, options *LoadBalancersListOptions) LoadBalancerListResultPager {
	return &loadBalancerListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp LoadBalancerListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *LoadBalancersClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *LoadBalancersListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LoadBalancersClient) listHandleResponse(resp *azcore.Response) (LoadBalancerListResultResponse, error) {
	var val *LoadBalancerListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LoadBalancerListResultResponse{}, err
	}
	return LoadBalancerListResultResponse{RawResponse: resp.Response, LoadBalancerListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *LoadBalancersClient) listHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListAll - Gets all the load balancers in a subscription.
func (client *LoadBalancersClient) ListAll(options *LoadBalancersListAllOptions) LoadBalancerListResultPager {
	return &loadBalancerListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp LoadBalancerListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LoadBalancerListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client *LoadBalancersClient) listAllCreateRequest(ctx context.Context, options *LoadBalancersListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/loadBalancers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *LoadBalancersClient) listAllHandleResponse(resp *azcore.Response) (LoadBalancerListResultResponse, error) {
	var val *LoadBalancerListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LoadBalancerListResultResponse{}, err
	}
	return LoadBalancerListResultResponse{RawResponse: resp.Response, LoadBalancerListResult: val}, nil
}

// listAllHandleError handles the ListAll error response.
func (client *LoadBalancersClient) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// UpdateTags - Updates a load balancer tags.
func (client *LoadBalancersClient) UpdateTags(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters TagsObject, options *LoadBalancersUpdateTagsOptions) (LoadBalancerResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, loadBalancerName, parameters, options)
	if err != nil {
		return LoadBalancerResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LoadBalancerResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return LoadBalancerResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *LoadBalancersClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, parameters TagsObject, options *LoadBalancersUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *LoadBalancersClient) updateTagsHandleResponse(resp *azcore.Response) (LoadBalancerResponse, error) {
	var val *LoadBalancer
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LoadBalancerResponse{}, err
	}
	return LoadBalancerResponse{RawResponse: resp.Response, LoadBalancer: val}, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *LoadBalancersClient) updateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}
