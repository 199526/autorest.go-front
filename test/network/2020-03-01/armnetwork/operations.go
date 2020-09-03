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
)

// Operations contains the methods for the Operations group.
type Operations interface {
	// List - Lists all of the available Network Rest API operations.
	List() (OperationListResultPager, error)
}

// OperationsClient implements the Operations interface.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	*Client
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
func NewOperationsClient(c *Client) Operations {
	return &OperationsClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *OperationsClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// List - Lists all of the available Network Rest API operations.
func (client *OperationsClient) List() (OperationListResultPager, error) {
	req, err := client.ListCreateRequest()
	if err != nil {
		return nil, err
	}
	return &operationListResultPager{
		pipeline:  client.p,
		request:   req,
		responder: client.ListHandleResponse,
		advancer: func(resp *OperationListResultResponse) (*azcore.Request, error) {
			u, err := url.Parse(*resp.OperationListResult.NextLink)
			if err != nil {
				return nil, fmt.Errorf("invalid NextLink: %w", err)
			}
			if u.Scheme == "" {
				return nil, fmt.Errorf("no scheme detected in NextLink %s", *resp.OperationListResult.NextLink)
			}
			return azcore.NewRequest(http.MethodGet, *u), nil
		},
	}, nil
}

// ListCreateRequest creates the List request.
func (client *OperationsClient) ListCreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/providers/Microsoft.Network/operations"
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
func (client *OperationsClient) ListHandleResponse(resp *azcore.Response) (*OperationListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := OperationListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.OperationListResult)
}

// ListHandleError handles the List error response.
func (client *OperationsClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
