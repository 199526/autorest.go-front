// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type triggerClient struct {
	con *connection
}

// BeginCreateOrUpdateTrigger - Creates or updates a trigger.
func (client *triggerClient) BeginCreateOrUpdateTrigger(ctx context.Context, triggerName string, trigger TriggerResource, options *TriggerBeginCreateOrUpdateTriggerOptions) (TriggerResourcePollerResponse, error) {
	resp, err := client.createOrUpdateTrigger(ctx, triggerName, trigger, options)
	if err != nil {
		return TriggerResourcePollerResponse{}, err
	}
	result := TriggerResourcePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("triggerClient.CreateOrUpdateTrigger", resp, client.con.Pipeline(), client.createOrUpdateTriggerHandleError)
	if err != nil {
		return TriggerResourcePollerResponse{}, err
	}
	poller := &triggerResourcePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TriggerResourceResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdateTrigger creates a new TriggerResourcePoller from the specified resume token.
// token - The value must come from a previous call to TriggerResourcePoller.ResumeToken().
func (client *triggerClient) ResumeCreateOrUpdateTrigger(ctx context.Context, token string) (TriggerResourcePollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("triggerClient.CreateOrUpdateTrigger", token, client.con.Pipeline(), client.createOrUpdateTriggerHandleError)
	if err != nil {
		return TriggerResourcePollerResponse{}, err
	}
	poller := &triggerResourcePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return TriggerResourcePollerResponse{}, err
	}
	result := TriggerResourcePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TriggerResourceResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdateTrigger - Creates or updates a trigger.
func (client *triggerClient) createOrUpdateTrigger(ctx context.Context, triggerName string, trigger TriggerResource, options *TriggerBeginCreateOrUpdateTriggerOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateTriggerCreateRequest(ctx, triggerName, trigger, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateTriggerHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateTriggerCreateRequest creates the CreateOrUpdateTrigger request.
func (client *triggerClient) createOrUpdateTriggerCreateRequest(ctx context.Context, triggerName string, trigger TriggerResource, options *TriggerBeginCreateOrUpdateTriggerOptions) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(trigger)
}

// createOrUpdateTriggerHandleResponse handles the CreateOrUpdateTrigger response.
func (client *triggerClient) createOrUpdateTriggerHandleResponse(resp *azcore.Response) (TriggerResourceResponse, error) {
	var val *TriggerResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TriggerResourceResponse{}, err
	}
	return TriggerResourceResponse{RawResponse: resp.Response, TriggerResource: val}, nil
}

// createOrUpdateTriggerHandleError handles the CreateOrUpdateTrigger error response.
func (client *triggerClient) createOrUpdateTriggerHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDeleteTrigger - Deletes a trigger.
func (client *triggerClient) BeginDeleteTrigger(ctx context.Context, triggerName string, options *TriggerBeginDeleteTriggerOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteTrigger(ctx, triggerName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("triggerClient.DeleteTrigger", resp, client.con.Pipeline(), client.deleteTriggerHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDeleteTrigger creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *triggerClient) ResumeDeleteTrigger(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("triggerClient.DeleteTrigger", token, client.con.Pipeline(), client.deleteTriggerHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
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

// DeleteTrigger - Deletes a trigger.
func (client *triggerClient) deleteTrigger(ctx context.Context, triggerName string, options *TriggerBeginDeleteTriggerOptions) (*azcore.Response, error) {
	req, err := client.deleteTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteTriggerHandleError(resp)
	}
	return resp, nil
}

// deleteTriggerCreateRequest creates the DeleteTrigger request.
func (client *triggerClient) deleteTriggerCreateRequest(ctx context.Context, triggerName string, options *TriggerBeginDeleteTriggerOptions) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteTriggerHandleError handles the DeleteTrigger error response.
func (client *triggerClient) deleteTriggerHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetEventSubscriptionStatus - Get a trigger's event subscription status.
func (client *triggerClient) GetEventSubscriptionStatus(ctx context.Context, triggerName string, options *TriggerGetEventSubscriptionStatusOptions) (TriggerSubscriptionOperationStatusResponse, error) {
	req, err := client.getEventSubscriptionStatusCreateRequest(ctx, triggerName, options)
	if err != nil {
		return TriggerSubscriptionOperationStatusResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TriggerSubscriptionOperationStatusResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TriggerSubscriptionOperationStatusResponse{}, client.getEventSubscriptionStatusHandleError(resp)
	}
	return client.getEventSubscriptionStatusHandleResponse(resp)
}

// getEventSubscriptionStatusCreateRequest creates the GetEventSubscriptionStatus request.
func (client *triggerClient) getEventSubscriptionStatusCreateRequest(ctx context.Context, triggerName string, options *TriggerGetEventSubscriptionStatusOptions) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}/getEventSubscriptionStatus"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getEventSubscriptionStatusHandleResponse handles the GetEventSubscriptionStatus response.
func (client *triggerClient) getEventSubscriptionStatusHandleResponse(resp *azcore.Response) (TriggerSubscriptionOperationStatusResponse, error) {
	var val *TriggerSubscriptionOperationStatus
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TriggerSubscriptionOperationStatusResponse{}, err
	}
	return TriggerSubscriptionOperationStatusResponse{RawResponse: resp.Response, TriggerSubscriptionOperationStatus: val}, nil
}

// getEventSubscriptionStatusHandleError handles the GetEventSubscriptionStatus error response.
func (client *triggerClient) getEventSubscriptionStatusHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetTrigger - Gets a trigger.
func (client *triggerClient) GetTrigger(ctx context.Context, triggerName string, options *TriggerGetTriggerOptions) (TriggerResourceResponse, error) {
	req, err := client.getTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return TriggerResourceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TriggerResourceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return TriggerResourceResponse{}, client.getTriggerHandleError(resp)
	}
	return client.getTriggerHandleResponse(resp)
}

// getTriggerCreateRequest creates the GetTrigger request.
func (client *triggerClient) getTriggerCreateRequest(ctx context.Context, triggerName string, options *TriggerGetTriggerOptions) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getTriggerHandleResponse handles the GetTrigger response.
func (client *triggerClient) getTriggerHandleResponse(resp *azcore.Response) (TriggerResourceResponse, error) {
	var val *TriggerResource
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TriggerResourceResponse{}, err
	}
	return TriggerResourceResponse{RawResponse: resp.Response, TriggerResource: val}, nil
}

// getTriggerHandleError handles the GetTrigger error response.
func (client *triggerClient) getTriggerHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetTriggersByWorkspace - Lists triggers.
func (client *triggerClient) GetTriggersByWorkspace(options *TriggerGetTriggersByWorkspaceOptions) TriggerListResponsePager {
	return &triggerListResponsePager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getTriggersByWorkspaceCreateRequest(ctx, options)
		},
		responder: client.getTriggersByWorkspaceHandleResponse,
		errorer:   client.getTriggersByWorkspaceHandleError,
		advancer: func(ctx context.Context, resp TriggerListResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.TriggerListResponse.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// getTriggersByWorkspaceCreateRequest creates the GetTriggersByWorkspace request.
func (client *triggerClient) getTriggersByWorkspaceCreateRequest(ctx context.Context, options *TriggerGetTriggersByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/triggers"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getTriggersByWorkspaceHandleResponse handles the GetTriggersByWorkspace response.
func (client *triggerClient) getTriggersByWorkspaceHandleResponse(resp *azcore.Response) (TriggerListResponseResponse, error) {
	var val *TriggerListResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TriggerListResponseResponse{}, err
	}
	return TriggerListResponseResponse{RawResponse: resp.Response, TriggerListResponse: val}, nil
}

// getTriggersByWorkspaceHandleError handles the GetTriggersByWorkspace error response.
func (client *triggerClient) getTriggersByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginStartTrigger - Starts a trigger.
func (client *triggerClient) BeginStartTrigger(ctx context.Context, triggerName string, options *TriggerBeginStartTriggerOptions) (HTTPPollerResponse, error) {
	resp, err := client.startTrigger(ctx, triggerName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("triggerClient.StartTrigger", resp, client.con.Pipeline(), client.startTriggerHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeStartTrigger creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *triggerClient) ResumeStartTrigger(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("triggerClient.StartTrigger", token, client.con.Pipeline(), client.startTriggerHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
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

// StartTrigger - Starts a trigger.
func (client *triggerClient) startTrigger(ctx context.Context, triggerName string, options *TriggerBeginStartTriggerOptions) (*azcore.Response, error) {
	req, err := client.startTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.startTriggerHandleError(resp)
	}
	return resp, nil
}

// startTriggerCreateRequest creates the StartTrigger request.
func (client *triggerClient) startTriggerCreateRequest(ctx context.Context, triggerName string, options *TriggerBeginStartTriggerOptions) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}/start"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// startTriggerHandleError handles the StartTrigger error response.
func (client *triggerClient) startTriggerHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginStopTrigger - Stops a trigger.
func (client *triggerClient) BeginStopTrigger(ctx context.Context, triggerName string, options *TriggerBeginStopTriggerOptions) (HTTPPollerResponse, error) {
	resp, err := client.stopTrigger(ctx, triggerName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("triggerClient.StopTrigger", resp, client.con.Pipeline(), client.stopTriggerHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeStopTrigger creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *triggerClient) ResumeStopTrigger(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("triggerClient.StopTrigger", token, client.con.Pipeline(), client.stopTriggerHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
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

// StopTrigger - Stops a trigger.
func (client *triggerClient) stopTrigger(ctx context.Context, triggerName string, options *TriggerBeginStopTriggerOptions) (*azcore.Response, error) {
	req, err := client.stopTriggerCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.stopTriggerHandleError(resp)
	}
	return resp, nil
}

// stopTriggerCreateRequest creates the StopTrigger request.
func (client *triggerClient) stopTriggerCreateRequest(ctx context.Context, triggerName string, options *TriggerBeginStopTriggerOptions) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}/stop"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// stopTriggerHandleError handles the StopTrigger error response.
func (client *triggerClient) stopTriggerHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginSubscribeTriggerToEvents - Subscribe event trigger to events.
func (client *triggerClient) BeginSubscribeTriggerToEvents(ctx context.Context, triggerName string, options *TriggerBeginSubscribeTriggerToEventsOptions) (TriggerSubscriptionOperationStatusPollerResponse, error) {
	resp, err := client.subscribeTriggerToEvents(ctx, triggerName, options)
	if err != nil {
		return TriggerSubscriptionOperationStatusPollerResponse{}, err
	}
	result := TriggerSubscriptionOperationStatusPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("triggerClient.SubscribeTriggerToEvents", resp, client.con.Pipeline(), client.subscribeTriggerToEventsHandleError)
	if err != nil {
		return TriggerSubscriptionOperationStatusPollerResponse{}, err
	}
	poller := &triggerSubscriptionOperationStatusPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TriggerSubscriptionOperationStatusResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeSubscribeTriggerToEvents creates a new TriggerSubscriptionOperationStatusPoller from the specified resume token.
// token - The value must come from a previous call to TriggerSubscriptionOperationStatusPoller.ResumeToken().
func (client *triggerClient) ResumeSubscribeTriggerToEvents(ctx context.Context, token string) (TriggerSubscriptionOperationStatusPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("triggerClient.SubscribeTriggerToEvents", token, client.con.Pipeline(), client.subscribeTriggerToEventsHandleError)
	if err != nil {
		return TriggerSubscriptionOperationStatusPollerResponse{}, err
	}
	poller := &triggerSubscriptionOperationStatusPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return TriggerSubscriptionOperationStatusPollerResponse{}, err
	}
	result := TriggerSubscriptionOperationStatusPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TriggerSubscriptionOperationStatusResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// SubscribeTriggerToEvents - Subscribe event trigger to events.
func (client *triggerClient) subscribeTriggerToEvents(ctx context.Context, triggerName string, options *TriggerBeginSubscribeTriggerToEventsOptions) (*azcore.Response, error) {
	req, err := client.subscribeTriggerToEventsCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.subscribeTriggerToEventsHandleError(resp)
	}
	return resp, nil
}

// subscribeTriggerToEventsCreateRequest creates the SubscribeTriggerToEvents request.
func (client *triggerClient) subscribeTriggerToEventsCreateRequest(ctx context.Context, triggerName string, options *TriggerBeginSubscribeTriggerToEventsOptions) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}/subscribeToEvents"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// subscribeTriggerToEventsHandleResponse handles the SubscribeTriggerToEvents response.
func (client *triggerClient) subscribeTriggerToEventsHandleResponse(resp *azcore.Response) (TriggerSubscriptionOperationStatusResponse, error) {
	var val *TriggerSubscriptionOperationStatus
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TriggerSubscriptionOperationStatusResponse{}, err
	}
	return TriggerSubscriptionOperationStatusResponse{RawResponse: resp.Response, TriggerSubscriptionOperationStatus: val}, nil
}

// subscribeTriggerToEventsHandleError handles the SubscribeTriggerToEvents error response.
func (client *triggerClient) subscribeTriggerToEventsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginUnsubscribeTriggerFromEvents - Unsubscribe event trigger from events.
func (client *triggerClient) BeginUnsubscribeTriggerFromEvents(ctx context.Context, triggerName string, options *TriggerBeginUnsubscribeTriggerFromEventsOptions) (TriggerSubscriptionOperationStatusPollerResponse, error) {
	resp, err := client.unsubscribeTriggerFromEvents(ctx, triggerName, options)
	if err != nil {
		return TriggerSubscriptionOperationStatusPollerResponse{}, err
	}
	result := TriggerSubscriptionOperationStatusPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("triggerClient.UnsubscribeTriggerFromEvents", resp, client.con.Pipeline(), client.unsubscribeTriggerFromEventsHandleError)
	if err != nil {
		return TriggerSubscriptionOperationStatusPollerResponse{}, err
	}
	poller := &triggerSubscriptionOperationStatusPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TriggerSubscriptionOperationStatusResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUnsubscribeTriggerFromEvents creates a new TriggerSubscriptionOperationStatusPoller from the specified resume token.
// token - The value must come from a previous call to TriggerSubscriptionOperationStatusPoller.ResumeToken().
func (client *triggerClient) ResumeUnsubscribeTriggerFromEvents(ctx context.Context, token string) (TriggerSubscriptionOperationStatusPollerResponse, error) {
	pt, err := azcore.NewLROPollerFromResumeToken("triggerClient.UnsubscribeTriggerFromEvents", token, client.con.Pipeline(), client.unsubscribeTriggerFromEventsHandleError)
	if err != nil {
		return TriggerSubscriptionOperationStatusPollerResponse{}, err
	}
	poller := &triggerSubscriptionOperationStatusPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return TriggerSubscriptionOperationStatusPollerResponse{}, err
	}
	result := TriggerSubscriptionOperationStatusPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TriggerSubscriptionOperationStatusResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// UnsubscribeTriggerFromEvents - Unsubscribe event trigger from events.
func (client *triggerClient) unsubscribeTriggerFromEvents(ctx context.Context, triggerName string, options *TriggerBeginUnsubscribeTriggerFromEventsOptions) (*azcore.Response, error) {
	req, err := client.unsubscribeTriggerFromEventsCreateRequest(ctx, triggerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.unsubscribeTriggerFromEventsHandleError(resp)
	}
	return resp, nil
}

// unsubscribeTriggerFromEventsCreateRequest creates the UnsubscribeTriggerFromEvents request.
func (client *triggerClient) unsubscribeTriggerFromEventsCreateRequest(ctx context.Context, triggerName string, options *TriggerBeginUnsubscribeTriggerFromEventsOptions) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}/unsubscribeFromEvents"
	if triggerName == "" {
		return nil, errors.New("parameter triggerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// unsubscribeTriggerFromEventsHandleResponse handles the UnsubscribeTriggerFromEvents response.
func (client *triggerClient) unsubscribeTriggerFromEventsHandleResponse(resp *azcore.Response) (TriggerSubscriptionOperationStatusResponse, error) {
	var val *TriggerSubscriptionOperationStatus
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TriggerSubscriptionOperationStatusResponse{}, err
	}
	return TriggerSubscriptionOperationStatusResponse{RawResponse: resp.Response, TriggerSubscriptionOperationStatus: val}, nil
}

// unsubscribeTriggerFromEventsHandleError handles the UnsubscribeTriggerFromEvents error response.
func (client *triggerClient) unsubscribeTriggerFromEventsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err.InnerError); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}
