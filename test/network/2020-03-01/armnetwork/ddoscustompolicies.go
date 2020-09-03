// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

// DdosCustomPoliciesOperations contains the methods for the DdosCustomPolicies group.
type DdosCustomPoliciesOperations interface {
	// BeginCreateOrUpdate - Creates or updates a DDoS custom policy.
	BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters DdosCustomPolicy) (*DdosCustomPolicyPollerResponse, error)
	// ResumeCreateOrUpdate - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeCreateOrUpdate(token string) (DdosCustomPolicyPoller, error)
	// BeginDelete - Deletes the specified DDoS custom policy.
	BeginDelete(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Gets information about the specified DDoS custom policy.
	Get(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string) (*DdosCustomPolicyResponse, error)
	// UpdateTags - Update a DDoS custom policy tags.
	UpdateTags(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters TagsObject) (*DdosCustomPolicyResponse, error)
}

// DdosCustomPoliciesClient implements the DdosCustomPoliciesOperations interface.
// Don't use this type directly, use NewDdosCustomPoliciesClient() instead.
type DdosCustomPoliciesClient struct {
	*Client
	subscriptionID string
}

// NewDdosCustomPoliciesClient creates a new instance of DdosCustomPoliciesClient with the specified values.
func NewDdosCustomPoliciesClient(c *Client, subscriptionID string) DdosCustomPoliciesOperations {
	return &DdosCustomPoliciesClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *DdosCustomPoliciesClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// CreateOrUpdate - Creates or updates a DDoS custom policy.
func (client *DdosCustomPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters DdosCustomPolicy) (*DdosCustomPolicyPollerResponse, error) {
	req, err := client.CreateOrUpdateCreateRequest(resourceGroupName, ddosCustomPolicyName, parameters)
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
	pt, err := armcore.NewPoller("DdosCustomPoliciesClient.CreateOrUpdate", "azure-async-operation", resp, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	poller := &ddosCustomPolicyPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*DdosCustomPolicyResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *DdosCustomPoliciesClient) ResumeCreateOrUpdate(token string) (DdosCustomPolicyPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("DdosCustomPoliciesClient.CreateOrUpdate", token, client.CreateOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &ddosCustomPolicyPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DdosCustomPoliciesClient) CreateOrUpdateCreateRequest(resourceGroupName string, ddosCustomPolicyName string, parameters DdosCustomPolicy) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ddosCustomPolicyName}", url.PathEscape(ddosCustomPolicyName))
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
func (client *DdosCustomPoliciesClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*DdosCustomPolicyPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusNoContent) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	return &DdosCustomPolicyPollerResponse{RawResponse: resp.Response}, nil
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DdosCustomPoliciesClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes the specified DDoS custom policy.
func (client *DdosCustomPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string) (*HTTPPollerResponse, error) {
	req, err := client.DeleteCreateRequest(resourceGroupName, ddosCustomPolicyName)
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
	pt, err := armcore.NewPoller("DdosCustomPoliciesClient.Delete", "location", resp, client.DeleteHandleError)
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

func (client *DdosCustomPoliciesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("DdosCustomPoliciesClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *DdosCustomPoliciesClient) DeleteCreateRequest(resourceGroupName string, ddosCustomPolicyName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ddosCustomPolicyName}", url.PathEscape(ddosCustomPolicyName))
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
func (client *DdosCustomPoliciesClient) DeleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// DeleteHandleError handles the Delete error response.
func (client *DdosCustomPoliciesClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Gets information about the specified DDoS custom policy.
func (client *DdosCustomPoliciesClient) Get(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string) (*DdosCustomPolicyResponse, error) {
	req, err := client.GetCreateRequest(resourceGroupName, ddosCustomPolicyName)
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
func (client *DdosCustomPoliciesClient) GetCreateRequest(resourceGroupName string, ddosCustomPolicyName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ddosCustomPolicyName}", url.PathEscape(ddosCustomPolicyName))
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
func (client *DdosCustomPoliciesClient) GetHandleResponse(resp *azcore.Response) (*DdosCustomPolicyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result := DdosCustomPolicyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DdosCustomPolicy)
}

// GetHandleError handles the Get error response.
func (client *DdosCustomPoliciesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// UpdateTags - Update a DDoS custom policy tags.
func (client *DdosCustomPoliciesClient) UpdateTags(ctx context.Context, resourceGroupName string, ddosCustomPolicyName string, parameters TagsObject) (*DdosCustomPolicyResponse, error) {
	req, err := client.UpdateTagsCreateRequest(resourceGroupName, ddosCustomPolicyName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.UpdateTagsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateTagsCreateRequest creates the UpdateTags request.
func (client *DdosCustomPoliciesClient) UpdateTagsCreateRequest(resourceGroupName string, ddosCustomPolicyName string, parameters TagsObject) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ddosCustomPolicies/{ddosCustomPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{ddosCustomPolicyName}", url.PathEscape(ddosCustomPolicyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2020-03-01")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodPatch, *u)
	return req, req.MarshalAsJSON(parameters)
}

// UpdateTagsHandleResponse handles the UpdateTags response.
func (client *DdosCustomPoliciesClient) UpdateTagsHandleResponse(resp *azcore.Response) (*DdosCustomPolicyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.UpdateTagsHandleError(resp)
	}
	result := DdosCustomPolicyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.DdosCustomPolicy)
}

// UpdateTagsHandleError handles the UpdateTags error response.
func (client *DdosCustomPoliciesClient) UpdateTagsHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
