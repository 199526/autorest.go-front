// +build go1.13

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
	"strings"
)

// VpnSiteLinksClient contains the methods for the VpnSiteLinks group.
// Don't use this type directly, use NewVpnSiteLinksClient() instead.
type VpnSiteLinksClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVpnSiteLinksClient creates a new instance of VpnSiteLinksClient with the specified values.
func NewVpnSiteLinksClient(con *armcore.Connection, subscriptionID string) *VpnSiteLinksClient {
	return &VpnSiteLinksClient{con: con, subscriptionID: subscriptionID}
}

// Get - Retrieves the details of a VPN site link.
func (client *VpnSiteLinksClient) Get(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteLinkName string, options *VpnSiteLinksGetOptions) (VpnSiteLinkResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vpnSiteName, vpnSiteLinkName, options)
	if err != nil {
		return VpnSiteLinkResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VpnSiteLinkResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VpnSiteLinkResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VpnSiteLinksClient) getCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, vpnSiteLinkName string, options *VpnSiteLinksGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks/{vpnSiteLinkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteLinkName}", url.PathEscape(vpnSiteLinkName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VpnSiteLinksClient) getHandleResponse(resp *azcore.Response) (VpnSiteLinkResponse, error) {
	var val *VpnSiteLink
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VpnSiteLinkResponse{}, err
	}
	return VpnSiteLinkResponse{RawResponse: resp.Response, VpnSiteLink: val}, nil
}

// getHandleError handles the Get error response.
func (client *VpnSiteLinksClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListByVpnSite - Lists all the vpnSiteLinks in a resource group for a vpn site.
func (client *VpnSiteLinksClient) ListByVpnSite(resourceGroupName string, vpnSiteName string, options *VpnSiteLinksListByVpnSiteOptions) ListVpnSiteLinksResultPager {
	return &listVpnSiteLinksResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByVpnSiteCreateRequest(ctx, resourceGroupName, vpnSiteName, options)
		},
		responder: client.listByVpnSiteHandleResponse,
		errorer:   client.listByVpnSiteHandleError,
		advancer: func(ctx context.Context, resp ListVpnSiteLinksResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVpnSiteLinksResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByVpnSiteCreateRequest creates the ListByVpnSite request.
func (client *VpnSiteLinksClient) listByVpnSiteCreateRequest(ctx context.Context, resourceGroupName string, vpnSiteName string, options *VpnSiteLinksListByVpnSiteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnSites/{vpnSiteName}/vpnSiteLinks"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vpnSiteName}", url.PathEscape(vpnSiteName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByVpnSiteHandleResponse handles the ListByVpnSite response.
func (client *VpnSiteLinksClient) listByVpnSiteHandleResponse(resp *azcore.Response) (ListVpnSiteLinksResultResponse, error) {
	var val *ListVpnSiteLinksResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ListVpnSiteLinksResultResponse{}, err
	}
	return ListVpnSiteLinksResultResponse{RawResponse: resp.Response, ListVpnSiteLinksResult: val}, nil
}

// listByVpnSiteHandleError handles the ListByVpnSite error response.
func (client *VpnSiteLinksClient) listByVpnSiteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}