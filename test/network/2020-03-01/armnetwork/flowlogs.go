// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

// FlowLogsOperations contains the methods for the FlowLogs group.
type FlowLogsOperations interface {
	// BeginCreateOrUpdate - Create or update a flow log for the specified network security group.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog) (*FlowLogPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (FlowLogPoller, error)
	// BeginDelete - Deletes the specified flow log resource.
	BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets a flow log resource by name.
	Get(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string) (*FlowLogResponse, error)
	// List - Lists all flow log resources for the specified Network Watcher.
	List(resourceGroupName string, networkWatcherName string) (FlowLogListResultPager, error)
}

// FlowLogsClient implements the FlowLogsOperations interface.
// Don't use this type directly, use NewFlowLogsClient() instead.
type FlowLogsClient struct {
	*Client
	subscriptionID string
}

// NewFlowLogsClient creates a new instance of FlowLogsClient with the specified values.
func NewFlowLogsClient(c *Client, subscriptionID string) FlowLogsOperations {
	return &FlowLogsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *FlowLogsClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// CreateOrUpdate - Create or update a flow log for the specified network security group.
func (client *FlowLogsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog) (*FlowLogPollerResponse, error) {
	req, err := client.CreateOrUpdateCreateRequest(resourceGroupName, networkWatcherName, flowLogName, parameters)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.CreateOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("FlowLogsClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &flowLogPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*FlowLogResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *FlowLogsClient) ResumeCreateOrUpdate(token string) (FlowLogPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("FlowLogsClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &flowLogPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FlowLogsClient) CreateOrUpdateCreateRequest(resourceGroupName string, networkWatcherName string, flowLogName string, parameters FlowLog) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPut, *u)
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *FlowLogsClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*FlowLogPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return &FlowLogPollerResponse{RawResponse: resp.Response}, nil
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *FlowLogsClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified flow log resource.
func (client *FlowLogsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string) (*HTTPPollerResponse, error) {
	req, err := client.DeleteCreateRequest(resourceGroupName, networkWatcherName, flowLogName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.DeleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("FlowLogsClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *FlowLogsClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("FlowLogsClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *FlowLogsClient) DeleteCreateRequest(resourceGroupName string, networkWatcherName string, flowLogName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodDelete, *u)
	return req, nil
}

// DeleteHandleResponse handles the Delete response.
func (client *FlowLogsClient) DeleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// DeleteHandleError handles the Delete error response.
func (client *FlowLogsClient) DeleteHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets a flow log resource by name.
func (client *FlowLogsClient) Get(ctx context.Context, resourceGroupName string, networkWatcherName string, flowLogName string) (*FlowLogResponse, error) {
	req, err := client.GetCreateRequest(resourceGroupName, networkWatcherName, flowLogName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *FlowLogsClient) GetCreateRequest(resourceGroupName string, networkWatcherName string, flowLogName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs/{flowLogName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{flowLogName}", url.PathEscape(flowLogName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *FlowLogsClient) GetHandleResponse(resp *azcore.Response) (*FlowLogResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result := FlowLogResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FlowLog)
}

// GetHandleError handles the Get error response.
func (client *FlowLogsClient) GetHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Lists all flow log resources for the specified Network Watcher.
func (client *FlowLogsClient) List(resourceGroupName string, networkWatcherName string) (FlowLogListResultPager, error) {
	req, err := client.ListCreateRequest(resourceGroupName, networkWatcherName)
	if err != nil {
		return nil, err
	}
	return &flowLogListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.ListHandleResponse,
		advancer: func(resp *FlowLogListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.FlowLogListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.FlowLogListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// ListCreateRequest creates the List request.
func (client *FlowLogsClient) ListCreateRequest(resourceGroupName string, networkWatcherName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/flowLogs"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *FlowLogsClient) ListHandleResponse(resp *azcore.Response) (*FlowLogListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := FlowLogListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.FlowLogListResult)
}

// ListHandleError handles the List error response.
func (client *FlowLogsClient) ListHandleError(resp *azcore.Response) error {
	var err ErrorResponse
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
