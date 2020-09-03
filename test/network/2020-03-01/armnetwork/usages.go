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

// UsagesOperations contains the methods for the Usages group.
type UsagesOperations interface {
	// List - List network usages for a subscription.
	List(location string) (UsagesListResultPager, error)
}

// UsagesClient implements the UsagesOperations interface.
// Don't use this type directly, use NewUsagesClient() instead.
type UsagesClient struct {
	*Client
	subscriptionID string
}

// NewUsagesClient creates a new instance of UsagesClient with the specified values.
func NewUsagesClient(c *Client, subscriptionID string) UsagesOperations {
	return &UsagesClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *UsagesClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// List - List network usages for a subscription.
func (client *UsagesClient) List(location string) (UsagesListResultPager, error) {
	req, err := client.ListCreateRequest(location)
	if err != nil {
		return nil, err
	}
	return &usagesListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.ListHandleResponse,
		advancer: func(resp *UsagesListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.UsagesListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.UsagesListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// ListCreateRequest creates the List request.
func (client *UsagesClient) ListCreateRequest(location string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/usages"
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
func (client *UsagesClient) ListHandleResponse(resp *azcore.Response) (*UsagesListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := UsagesListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.UsagesListResult)
}

// ListHandleError handles the List error response.
func (client *UsagesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
