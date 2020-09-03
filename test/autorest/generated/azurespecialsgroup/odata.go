// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

// OdataOperations contains the methods for the Odata group.
type OdataOperations interface {
	// GetWithFilter - Specify filter parameter with value '$filter=id gt 5 and name eq 'foo'&$orderby=id&$top=10'
	GetWithFilter(ctx context.Context, odataGetWithFilterOptions *OdataGetWithFilterOptions) (*http.Response, error)
}

// OdataClient implements the OdataOperations interface.
// Don't use this type directly, use NewOdataClient() instead.
type OdataClient struct {
	*Client
}

// NewOdataClient creates a new instance of OdataClient with the specified values.
func NewOdataClient(c *Client) OdataOperations {
	return &OdataClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *OdataClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// GetWithFilter - Specify filter parameter with value '$filter=id gt 5 and name eq 'foo'&$orderby=id&$top=10'
func (client *OdataClient) GetWithFilter(ctx context.Context, odataGetWithFilterOptions *OdataGetWithFilterOptions) (*http.Response, error) {
	req, err := client.GetWithFilterCreateRequest(odataGetWithFilterOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetWithFilterHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetWithFilterCreateRequest creates the GetWithFilter request.
func (client *OdataClient) GetWithFilterCreateRequest(odataGetWithFilterOptions *OdataGetWithFilterOptions) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/azurespecials/odata/filter"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	if odataGetWithFilterOptions != nil && odataGetWithFilterOptions.Filter != nil {
		query.Set("$filter", *odataGetWithFilterOptions.Filter)
	}
	if odataGetWithFilterOptions != nil && odataGetWithFilterOptions.Top != nil {
		query.Set("$top", strconv.FormatInt(int64(*odataGetWithFilterOptions.Top), 10))
	}
	if odataGetWithFilterOptions != nil && odataGetWithFilterOptions.Orderby != nil {
		query.Set("$orderby", *odataGetWithFilterOptions.Orderby)
	}
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// GetWithFilterHandleResponse handles the GetWithFilter response.
func (client *OdataClient) GetWithFilterHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetWithFilterHandleError(resp)
	}
	return resp.Response, nil
}

// GetWithFilterHandleError handles the GetWithFilter error response.
func (client *OdataClient) GetWithFilterHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
