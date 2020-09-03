// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// AvailableServiceAliasesOperations contains the methods for the AvailableServiceAliases group.
type AvailableServiceAliasesOperations interface {
	// List - Gets all available service aliases for this subscription in this region.
	List(location string) (AvailableServiceAliasesResultPager, error)
	// ListByResourceGroup - Gets all available service aliases for this resource group in this region.
	ListByResourceGroup(resourceGroupName string, location string) (AvailableServiceAliasesResultPager, error)
}

// AvailableServiceAliasesClient implements the AvailableServiceAliasesOperations interface.
// Don't use this type directly, use NewAvailableServiceAliasesClient() instead.
type AvailableServiceAliasesClient struct {
	*Client
	subscriptionID string
}

// NewAvailableServiceAliasesClient creates a new instance of AvailableServiceAliasesClient with the specified values.
func NewAvailableServiceAliasesClient(c *Client, subscriptionID string) AvailableServiceAliasesOperations {
	return &AvailableServiceAliasesClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *AvailableServiceAliasesClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// List - Gets all available service aliases for this subscription in this region.
func (client *AvailableServiceAliasesClient) List(location string) (AvailableServiceAliasesResultPager, error) {
	req, err := client.ListCreateRequest(location)
	if err != nil {
		return nil, err
	}
	return &availableServiceAliasesResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.ListHandleResponse,
		advancer: func(resp *AvailableServiceAliasesResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.AvailableServiceAliasesResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.AvailableServiceAliasesResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// ListCreateRequest creates the List request.
func (client *AvailableServiceAliasesClient) ListCreateRequest(location string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/availableServiceAliases"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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
func (client *AvailableServiceAliasesClient) ListHandleResponse(resp *azcore.Response) (*AvailableServiceAliasesResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := AvailableServiceAliasesResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailableServiceAliasesResult)
}

// ListHandleError handles the List error response.
func (client *AvailableServiceAliasesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListByResourceGroup - Gets all available service aliases for this resource group in this region.
func (client *AvailableServiceAliasesClient) ListByResourceGroup(resourceGroupName string, location string) (AvailableServiceAliasesResultPager, error) {
	req, err := client.ListByResourceGroupCreateRequest(resourceGroupName, location)
	if err != nil {
		return nil, err
	}
	return &availableServiceAliasesResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.ListByResourceGroupHandleResponse,
		advancer: func(resp *AvailableServiceAliasesResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.AvailableServiceAliasesResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.AvailableServiceAliasesResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// ListByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AvailableServiceAliasesClient) ListByResourceGroupCreateRequest(resourceGroupName string, location string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableServiceAliases"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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

// ListByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AvailableServiceAliasesClient) ListByResourceGroupHandleResponse(resp *azcore.Response) (*AvailableServiceAliasesResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListByResourceGroupHandleError(resp)
	}
	result := AvailableServiceAliasesResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailableServiceAliasesResult)
}

// ListByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *AvailableServiceAliasesClient) ListByResourceGroupHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
