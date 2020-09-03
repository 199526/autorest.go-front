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
)

// APIVersionDefaultOperations contains the methods for the APIVersionDefault group.
type APIVersionDefaultOperations interface {
	// GetMethodGlobalNotProvidedValid - GET method with api-version modeled in global settings.
	GetMethodGlobalNotProvidedValid(ctx context.Context) (*http.Response, error)
	// GetMethodGlobalValid - GET method with api-version modeled in global settings.
	GetMethodGlobalValid(ctx context.Context) (*http.Response, error)
	// GetPathGlobalValid - GET method with api-version modeled in global settings.
	GetPathGlobalValid(ctx context.Context) (*http.Response, error)
	// GetSwaggerGlobalValid - GET method with api-version modeled in global settings.
	GetSwaggerGlobalValid(ctx context.Context) (*http.Response, error)
}

// APIVersionDefaultClient implements the APIVersionDefaultOperations interface.
// Don't use this type directly, use NewAPIVersionDefaultClient() instead.
type APIVersionDefaultClient struct {
	*Client
}

// NewAPIVersionDefaultClient creates a new instance of APIVersionDefaultClient with the specified values.
func NewAPIVersionDefaultClient(c *Client) APIVersionDefaultOperations {
	return &APIVersionDefaultClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *APIVersionDefaultClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// GetMethodGlobalNotProvidedValid - GET method with api-version modeled in global settings.
func (client *APIVersionDefaultClient) GetMethodGlobalNotProvidedValid(ctx context.Context) (*http.Response, error) {
	req, err := client.GetMethodGlobalNotProvidedValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetMethodGlobalNotProvidedValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMethodGlobalNotProvidedValidCreateRequest creates the GetMethodGlobalNotProvidedValid request.
func (client *APIVersionDefaultClient) GetMethodGlobalNotProvidedValidCreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/azurespecials/apiVersion/method/string/none/query/globalNotProvided/2015-07-01-preview"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2015-07-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// GetMethodGlobalNotProvidedValidHandleResponse handles the GetMethodGlobalNotProvidedValid response.
func (client *APIVersionDefaultClient) GetMethodGlobalNotProvidedValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetMethodGlobalNotProvidedValidHandleError(resp)
	}
	return resp.Response, nil
}

// GetMethodGlobalNotProvidedValidHandleError handles the GetMethodGlobalNotProvidedValid error response.
func (client *APIVersionDefaultClient) GetMethodGlobalNotProvidedValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetMethodGlobalValid - GET method with api-version modeled in global settings.
func (client *APIVersionDefaultClient) GetMethodGlobalValid(ctx context.Context) (*http.Response, error) {
	req, err := client.GetMethodGlobalValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetMethodGlobalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMethodGlobalValidCreateRequest creates the GetMethodGlobalValid request.
func (client *APIVersionDefaultClient) GetMethodGlobalValidCreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/azurespecials/apiVersion/method/string/none/query/global/2015-07-01-preview"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2015-07-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// GetMethodGlobalValidHandleResponse handles the GetMethodGlobalValid response.
func (client *APIVersionDefaultClient) GetMethodGlobalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetMethodGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// GetMethodGlobalValidHandleError handles the GetMethodGlobalValid error response.
func (client *APIVersionDefaultClient) GetMethodGlobalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetPathGlobalValid - GET method with api-version modeled in global settings.
func (client *APIVersionDefaultClient) GetPathGlobalValid(ctx context.Context) (*http.Response, error) {
	req, err := client.GetPathGlobalValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetPathGlobalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPathGlobalValidCreateRequest creates the GetPathGlobalValid request.
func (client *APIVersionDefaultClient) GetPathGlobalValidCreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/azurespecials/apiVersion/path/string/none/query/global/2015-07-01-preview"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2015-07-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// GetPathGlobalValidHandleResponse handles the GetPathGlobalValid response.
func (client *APIVersionDefaultClient) GetPathGlobalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetPathGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// GetPathGlobalValidHandleError handles the GetPathGlobalValid error response.
func (client *APIVersionDefaultClient) GetPathGlobalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetSwaggerGlobalValid - GET method with api-version modeled in global settings.
func (client *APIVersionDefaultClient) GetSwaggerGlobalValid(ctx context.Context) (*http.Response, error) {
	req, err := client.GetSwaggerGlobalValidCreateRequest()
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetSwaggerGlobalValidHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetSwaggerGlobalValidCreateRequest creates the GetSwaggerGlobalValid request.
func (client *APIVersionDefaultClient) GetSwaggerGlobalValidCreateRequest() (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/azurespecials/apiVersion/swagger/string/none/query/global/2015-07-01-preview"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Set("api-version", "2015-07-01-preview")
	u.RawQuery = query.Encode()
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// GetSwaggerGlobalValidHandleResponse handles the GetSwaggerGlobalValid response.
func (client *APIVersionDefaultClient) GetSwaggerGlobalValidHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetSwaggerGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// GetSwaggerGlobalValidHandleError handles the GetSwaggerGlobalValid error response.
func (client *APIVersionDefaultClient) GetSwaggerGlobalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
