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

// AvailableResourceGroupDelegationsOperations contains the methods for the AvailableResourceGroupDelegations group.
type AvailableResourceGroupDelegationsOperations interface {
	// List - Gets all of the available subnet delegations for this resource group in this region.
	List(location string, resourceGroupName string) (AvailableDelegationsResultPager, error)
}

// AvailableResourceGroupDelegationsClient implements the AvailableResourceGroupDelegationsOperations interface.
// Don't use this type directly, use NewAvailableResourceGroupDelegationsClient() instead.
type AvailableResourceGroupDelegationsClient struct {
	*Client
	subscriptionID string
}

// NewAvailableResourceGroupDelegationsClient creates a new instance of AvailableResourceGroupDelegationsClient with the specified values.
func NewAvailableResourceGroupDelegationsClient(c *Client, subscriptionID string) AvailableResourceGroupDelegationsOperations {
	return &AvailableResourceGroupDelegationsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *AvailableResourceGroupDelegationsClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// List - Gets all of the available subnet delegations for this resource group in this region.
func (client *AvailableResourceGroupDelegationsClient) List(location string, resourceGroupName string) (AvailableDelegationsResultPager, error) {
	req, err := client.ListCreateRequest(location, resourceGroupName)
	if err != nil {
		return nil, err
	}
	return &availableDelegationsResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.ListHandleResponse,
		advancer: func(resp *AvailableDelegationsResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.AvailableDelegationsResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.AvailableDelegationsResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// ListCreateRequest creates the List request.
func (client *AvailableResourceGroupDelegationsClient) ListCreateRequest(location string, resourceGroupName string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/locations/{location}/availableDelegations"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *AvailableResourceGroupDelegationsClient) ListHandleResponse(resp *azcore.Response) (*AvailableDelegationsResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := AvailableDelegationsResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.AvailableDelegationsResult)
}

// ListHandleError handles the List error response.
func (client *AvailableResourceGroupDelegationsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
