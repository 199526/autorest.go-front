// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// GalleryApplicationVersionsClient contains the methods for the GalleryApplicationVersions group.
// Don't use this type directly, use NewGalleryApplicationVersionsClient() instead.
type GalleryApplicationVersionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewGalleryApplicationVersionsClient creates a new instance of GalleryApplicationVersionsClient with the specified values.
func NewGalleryApplicationVersionsClient(con *armcore.Connection, subscriptionID string) *GalleryApplicationVersionsClient {
	return &GalleryApplicationVersionsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Create or update a gallery Application Version.
func (client *GalleryApplicationVersionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersion, options *GalleryApplicationVersionsBeginCreateOrUpdateOptions) (GalleryApplicationVersionPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
	if err != nil {
		return GalleryApplicationVersionPollerResponse{}, err
	}
	result := GalleryApplicationVersionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GalleryApplicationVersionsClient.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return GalleryApplicationVersionPollerResponse{}, err
	}
	poller := &galleryApplicationVersionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (GalleryApplicationVersionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new GalleryApplicationVersionPoller from the specified resume token.
// token - The value must come from a previous call to GalleryApplicationVersionPoller.ResumeToken().
func (client *GalleryApplicationVersionsClient) ResumeCreateOrUpdate(ctx context.Context, token string) (GalleryApplicationVersionPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("GalleryApplicationVersionsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return GalleryApplicationVersionPollerResponse{}, err
	}
	poller := &galleryApplicationVersionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return GalleryApplicationVersionPollerResponse{}, err
	}
	result := GalleryApplicationVersionPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (GalleryApplicationVersionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateOrUpdate - Create or update a gallery Application Version.
func (client *GalleryApplicationVersionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersion, options *GalleryApplicationVersionsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GalleryApplicationVersionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersion, options *GalleryApplicationVersionsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	if galleryApplicationVersionName == "" {
		return nil, errors.New("parameter galleryApplicationVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(galleryApplicationVersion)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *GalleryApplicationVersionsClient) createOrUpdateHandleResponse(resp *azcore.Response) (GalleryApplicationVersionResponse, error) {
	var val *GalleryApplicationVersion
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return GalleryApplicationVersionResponse{}, err
	}
	return GalleryApplicationVersionResponse{RawResponse: resp.Response, GalleryApplicationVersion: val}, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *GalleryApplicationVersionsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginDelete - Delete a gallery Application Version.
func (client *GalleryApplicationVersionsClient) BeginDelete(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GalleryApplicationVersionsClient.Delete", "", resp, client.deleteHandleError)
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
func (client *GalleryApplicationVersionsClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("GalleryApplicationVersionsClient.Delete", token, client.deleteHandleError)
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

// Delete - Delete a gallery Application Version.
func (client *GalleryApplicationVersionsClient) deleteOperation(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, options)
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
func (client *GalleryApplicationVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	if galleryApplicationVersionName == "" {
		return nil, errors.New("parameter galleryApplicationVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *GalleryApplicationVersionsClient) deleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get - Retrieves information about a gallery Application Version.
func (client *GalleryApplicationVersionsClient) Get(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsGetOptions) (GalleryApplicationVersionResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, options)
	if err != nil {
		return GalleryApplicationVersionResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return GalleryApplicationVersionResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return GalleryApplicationVersionResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *GalleryApplicationVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, options *GalleryApplicationVersionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	if galleryApplicationVersionName == "" {
		return nil, errors.New("parameter galleryApplicationVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", string(*options.Expand))
	}
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GalleryApplicationVersionsClient) getHandleResponse(resp *azcore.Response) (GalleryApplicationVersionResponse, error) {
	var val *GalleryApplicationVersion
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return GalleryApplicationVersionResponse{}, err
	}
	return GalleryApplicationVersionResponse{RawResponse: resp.Response, GalleryApplicationVersion: val}, nil
}

// getHandleError handles the Get error response.
func (client *GalleryApplicationVersionsClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByGalleryApplication - List gallery Application Versions in a gallery Application Definition.
func (client *GalleryApplicationVersionsClient) ListByGalleryApplication(resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationVersionsListByGalleryApplicationOptions) GalleryApplicationVersionListPager {
	return &galleryApplicationVersionListPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByGalleryApplicationCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, options)
		},
		responder: client.listByGalleryApplicationHandleResponse,
		errorer:   client.listByGalleryApplicationHandleError,
		advancer: func(ctx context.Context, resp GalleryApplicationVersionListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.GalleryApplicationVersionList.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByGalleryApplicationCreateRequest creates the ListByGalleryApplication request.
func (client *GalleryApplicationVersionsClient) listByGalleryApplicationCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, options *GalleryApplicationVersionsListByGalleryApplicationOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByGalleryApplicationHandleResponse handles the ListByGalleryApplication response.
func (client *GalleryApplicationVersionsClient) listByGalleryApplicationHandleResponse(resp *azcore.Response) (GalleryApplicationVersionListResponse, error) {
	var val *GalleryApplicationVersionList
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return GalleryApplicationVersionListResponse{}, err
	}
	return GalleryApplicationVersionListResponse{RawResponse: resp.Response, GalleryApplicationVersionList: val}, nil
}

// listByGalleryApplicationHandleError handles the ListByGalleryApplication error response.
func (client *GalleryApplicationVersionsClient) listByGalleryApplicationHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// BeginUpdate - Update a gallery Application Version.
func (client *GalleryApplicationVersionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersionUpdate, options *GalleryApplicationVersionsBeginUpdateOptions) (GalleryApplicationVersionPollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
	if err != nil {
		return GalleryApplicationVersionPollerResponse{}, err
	}
	result := GalleryApplicationVersionPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("GalleryApplicationVersionsClient.Update", "", resp, client.updateHandleError)
	if err != nil {
		return GalleryApplicationVersionPollerResponse{}, err
	}
	poller := &galleryApplicationVersionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (GalleryApplicationVersionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new GalleryApplicationVersionPoller from the specified resume token.
// token - The value must come from a previous call to GalleryApplicationVersionPoller.ResumeToken().
func (client *GalleryApplicationVersionsClient) ResumeUpdate(ctx context.Context, token string) (GalleryApplicationVersionPollerResponse, error) {
	pt, err := armcore.NewPollerFromResumeToken("GalleryApplicationVersionsClient.Update", token, client.updateHandleError)
	if err != nil {
		return GalleryApplicationVersionPollerResponse{}, err
	}
	poller := &galleryApplicationVersionPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return GalleryApplicationVersionPollerResponse{}, err
	}
	result := GalleryApplicationVersionPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (GalleryApplicationVersionResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Update a gallery Application Version.
func (client *GalleryApplicationVersionsClient) update(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersionUpdate, options *GalleryApplicationVersionsBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, galleryName, galleryApplicationName, galleryApplicationVersionName, galleryApplicationVersion, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *GalleryApplicationVersionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, galleryName string, galleryApplicationName string, galleryApplicationVersionName string, galleryApplicationVersion GalleryApplicationVersionUpdate, options *GalleryApplicationVersionsBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/galleries/{galleryName}/applications/{galleryApplicationName}/versions/{galleryApplicationVersionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if galleryName == "" {
		return nil, errors.New("parameter galleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryName}", url.PathEscape(galleryName))
	if galleryApplicationName == "" {
		return nil, errors.New("parameter galleryApplicationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationName}", url.PathEscape(galleryApplicationName))
	if galleryApplicationVersionName == "" {
		return nil, errors.New("parameter galleryApplicationVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryApplicationVersionName}", url.PathEscape(galleryApplicationVersionName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(galleryApplicationVersion)
}

// updateHandleResponse handles the Update response.
func (client *GalleryApplicationVersionsClient) updateHandleResponse(resp *azcore.Response) (GalleryApplicationVersionResponse, error) {
	var val *GalleryApplicationVersion
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return GalleryApplicationVersionResponse{}, err
	}
	return GalleryApplicationVersionResponse{RawResponse: resp.Response, GalleryApplicationVersion: val}, nil
}

// updateHandleError handles the Update error response.
func (client *GalleryApplicationVersionsClient) updateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return azcore.NewResponseError(resp.UnmarshalError(err), resp.Response)
	}
	return azcore.NewResponseError(&err, resp.Response)
}
