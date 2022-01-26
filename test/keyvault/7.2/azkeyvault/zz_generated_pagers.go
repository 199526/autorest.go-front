//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azkeyvault

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// ClientGetCertificateIssuersPager provides operations for iterating over paged responses.
type ClientGetCertificateIssuersPager struct {
	client    *Client
	current   ClientGetCertificateIssuersResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetCertificateIssuersResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClientGetCertificateIssuersPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CertificateIssuerListResult.NextLink == nil || len(*p.current.CertificateIssuerListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClientGetCertificateIssuersPager) NextPage(ctx context.Context) (ClientGetCertificateIssuersResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClientGetCertificateIssuersResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClientGetCertificateIssuersResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClientGetCertificateIssuersResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClientGetCertificateIssuersResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getCertificateIssuersHandleResponse(resp)
	if err != nil {
		return ClientGetCertificateIssuersResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ClientGetCertificateVersionsPager provides operations for iterating over paged responses.
type ClientGetCertificateVersionsPager struct {
	client    *Client
	current   ClientGetCertificateVersionsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetCertificateVersionsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClientGetCertificateVersionsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CertificateListResult.NextLink == nil || len(*p.current.CertificateListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClientGetCertificateVersionsPager) NextPage(ctx context.Context) (ClientGetCertificateVersionsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClientGetCertificateVersionsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClientGetCertificateVersionsResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClientGetCertificateVersionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClientGetCertificateVersionsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getCertificateVersionsHandleResponse(resp)
	if err != nil {
		return ClientGetCertificateVersionsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ClientGetCertificatesPager provides operations for iterating over paged responses.
type ClientGetCertificatesPager struct {
	client    *Client
	current   ClientGetCertificatesResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetCertificatesResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClientGetCertificatesPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CertificateListResult.NextLink == nil || len(*p.current.CertificateListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClientGetCertificatesPager) NextPage(ctx context.Context) (ClientGetCertificatesResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClientGetCertificatesResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClientGetCertificatesResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClientGetCertificatesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClientGetCertificatesResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getCertificatesHandleResponse(resp)
	if err != nil {
		return ClientGetCertificatesResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ClientGetDeletedCertificatesPager provides operations for iterating over paged responses.
type ClientGetDeletedCertificatesPager struct {
	client    *Client
	current   ClientGetDeletedCertificatesResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetDeletedCertificatesResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClientGetDeletedCertificatesPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedCertificateListResult.NextLink == nil || len(*p.current.DeletedCertificateListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClientGetDeletedCertificatesPager) NextPage(ctx context.Context) (ClientGetDeletedCertificatesResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClientGetDeletedCertificatesResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClientGetDeletedCertificatesResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClientGetDeletedCertificatesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClientGetDeletedCertificatesResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getDeletedCertificatesHandleResponse(resp)
	if err != nil {
		return ClientGetDeletedCertificatesResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ClientGetDeletedKeysPager provides operations for iterating over paged responses.
type ClientGetDeletedKeysPager struct {
	client    *Client
	current   ClientGetDeletedKeysResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetDeletedKeysResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClientGetDeletedKeysPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedKeyListResult.NextLink == nil || len(*p.current.DeletedKeyListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClientGetDeletedKeysPager) NextPage(ctx context.Context) (ClientGetDeletedKeysResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClientGetDeletedKeysResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClientGetDeletedKeysResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClientGetDeletedKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClientGetDeletedKeysResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getDeletedKeysHandleResponse(resp)
	if err != nil {
		return ClientGetDeletedKeysResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ClientGetDeletedSasDefinitionsPager provides operations for iterating over paged responses.
type ClientGetDeletedSasDefinitionsPager struct {
	client    *Client
	current   ClientGetDeletedSasDefinitionsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetDeletedSasDefinitionsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClientGetDeletedSasDefinitionsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedSasDefinitionListResult.NextLink == nil || len(*p.current.DeletedSasDefinitionListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClientGetDeletedSasDefinitionsPager) NextPage(ctx context.Context) (ClientGetDeletedSasDefinitionsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClientGetDeletedSasDefinitionsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClientGetDeletedSasDefinitionsResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClientGetDeletedSasDefinitionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClientGetDeletedSasDefinitionsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getDeletedSasDefinitionsHandleResponse(resp)
	if err != nil {
		return ClientGetDeletedSasDefinitionsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ClientGetDeletedSecretsPager provides operations for iterating over paged responses.
type ClientGetDeletedSecretsPager struct {
	client    *Client
	current   ClientGetDeletedSecretsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetDeletedSecretsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClientGetDeletedSecretsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedSecretListResult.NextLink == nil || len(*p.current.DeletedSecretListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClientGetDeletedSecretsPager) NextPage(ctx context.Context) (ClientGetDeletedSecretsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClientGetDeletedSecretsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClientGetDeletedSecretsResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClientGetDeletedSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClientGetDeletedSecretsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getDeletedSecretsHandleResponse(resp)
	if err != nil {
		return ClientGetDeletedSecretsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ClientGetDeletedStorageAccountsPager provides operations for iterating over paged responses.
type ClientGetDeletedStorageAccountsPager struct {
	client    *Client
	current   ClientGetDeletedStorageAccountsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetDeletedStorageAccountsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClientGetDeletedStorageAccountsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedStorageListResult.NextLink == nil || len(*p.current.DeletedStorageListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClientGetDeletedStorageAccountsPager) NextPage(ctx context.Context) (ClientGetDeletedStorageAccountsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClientGetDeletedStorageAccountsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClientGetDeletedStorageAccountsResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClientGetDeletedStorageAccountsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClientGetDeletedStorageAccountsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getDeletedStorageAccountsHandleResponse(resp)
	if err != nil {
		return ClientGetDeletedStorageAccountsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ClientGetKeyVersionsPager provides operations for iterating over paged responses.
type ClientGetKeyVersionsPager struct {
	client    *Client
	current   ClientGetKeyVersionsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetKeyVersionsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClientGetKeyVersionsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.KeyListResult.NextLink == nil || len(*p.current.KeyListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClientGetKeyVersionsPager) NextPage(ctx context.Context) (ClientGetKeyVersionsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClientGetKeyVersionsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClientGetKeyVersionsResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClientGetKeyVersionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClientGetKeyVersionsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getKeyVersionsHandleResponse(resp)
	if err != nil {
		return ClientGetKeyVersionsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ClientGetKeysPager provides operations for iterating over paged responses.
type ClientGetKeysPager struct {
	client    *Client
	current   ClientGetKeysResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetKeysResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClientGetKeysPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.KeyListResult.NextLink == nil || len(*p.current.KeyListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClientGetKeysPager) NextPage(ctx context.Context) (ClientGetKeysResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClientGetKeysResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClientGetKeysResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClientGetKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClientGetKeysResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getKeysHandleResponse(resp)
	if err != nil {
		return ClientGetKeysResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ClientGetSasDefinitionsPager provides operations for iterating over paged responses.
type ClientGetSasDefinitionsPager struct {
	client    *Client
	current   ClientGetSasDefinitionsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetSasDefinitionsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClientGetSasDefinitionsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SasDefinitionListResult.NextLink == nil || len(*p.current.SasDefinitionListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClientGetSasDefinitionsPager) NextPage(ctx context.Context) (ClientGetSasDefinitionsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClientGetSasDefinitionsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClientGetSasDefinitionsResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClientGetSasDefinitionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClientGetSasDefinitionsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getSasDefinitionsHandleResponse(resp)
	if err != nil {
		return ClientGetSasDefinitionsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ClientGetSecretVersionsPager provides operations for iterating over paged responses.
type ClientGetSecretVersionsPager struct {
	client    *Client
	current   ClientGetSecretVersionsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetSecretVersionsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClientGetSecretVersionsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SecretListResult.NextLink == nil || len(*p.current.SecretListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClientGetSecretVersionsPager) NextPage(ctx context.Context) (ClientGetSecretVersionsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClientGetSecretVersionsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClientGetSecretVersionsResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClientGetSecretVersionsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClientGetSecretVersionsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getSecretVersionsHandleResponse(resp)
	if err != nil {
		return ClientGetSecretVersionsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ClientGetSecretsPager provides operations for iterating over paged responses.
type ClientGetSecretsPager struct {
	client    *Client
	current   ClientGetSecretsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetSecretsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClientGetSecretsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SecretListResult.NextLink == nil || len(*p.current.SecretListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClientGetSecretsPager) NextPage(ctx context.Context) (ClientGetSecretsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClientGetSecretsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClientGetSecretsResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClientGetSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClientGetSecretsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getSecretsHandleResponse(resp)
	if err != nil {
		return ClientGetSecretsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// ClientGetStorageAccountsPager provides operations for iterating over paged responses.
type ClientGetStorageAccountsPager struct {
	client    *Client
	current   ClientGetStorageAccountsResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ClientGetStorageAccountsResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *ClientGetStorageAccountsPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.StorageListResult.NextLink == nil || len(*p.current.StorageListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *ClientGetStorageAccountsPager) NextPage(ctx context.Context) (ClientGetStorageAccountsResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return ClientGetStorageAccountsResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return ClientGetStorageAccountsResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return ClientGetStorageAccountsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return ClientGetStorageAccountsResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getStorageAccountsHandleResponse(resp)
	if err != nil {
		return ClientGetStorageAccountsResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// RoleAssignmentsClientListForScopePager provides operations for iterating over paged responses.
type RoleAssignmentsClientListForScopePager struct {
	client    *RoleAssignmentsClient
	current   RoleAssignmentsClientListForScopeResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RoleAssignmentsClientListForScopeResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *RoleAssignmentsClientListForScopePager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleAssignmentListResult.NextLink == nil || len(*p.current.RoleAssignmentListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *RoleAssignmentsClientListForScopePager) NextPage(ctx context.Context) (RoleAssignmentsClientListForScopeResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return RoleAssignmentsClientListForScopeResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return RoleAssignmentsClientListForScopeResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return RoleAssignmentsClientListForScopeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return RoleAssignmentsClientListForScopeResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listForScopeHandleResponse(resp)
	if err != nil {
		return RoleAssignmentsClientListForScopeResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// RoleDefinitionsClientListPager provides operations for iterating over paged responses.
type RoleDefinitionsClientListPager struct {
	client    *RoleDefinitionsClient
	current   RoleDefinitionsClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RoleDefinitionsClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *RoleDefinitionsClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleDefinitionListResult.NextLink == nil || len(*p.current.RoleDefinitionListResult.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *RoleDefinitionsClientListPager) NextPage(ctx context.Context) (RoleDefinitionsClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return RoleDefinitionsClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return RoleDefinitionsClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return RoleDefinitionsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return RoleDefinitionsClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return RoleDefinitionsClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}
