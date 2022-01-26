//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// dataFlowClientGetDataFlowsByWorkspacePager provides operations for iterating over paged responses.
type dataFlowClientGetDataFlowsByWorkspacePager struct {
	client    *dataFlowClient
	current   dataFlowClientGetDataFlowsByWorkspaceResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, dataFlowClientGetDataFlowsByWorkspaceResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *dataFlowClientGetDataFlowsByWorkspacePager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DataFlowListResponse.NextLink == nil || len(*p.current.DataFlowListResponse.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *dataFlowClientGetDataFlowsByWorkspacePager) NextPage(ctx context.Context) (dataFlowClientGetDataFlowsByWorkspaceResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return dataFlowClientGetDataFlowsByWorkspaceResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return dataFlowClientGetDataFlowsByWorkspaceResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return dataFlowClientGetDataFlowsByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return dataFlowClientGetDataFlowsByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getDataFlowsByWorkspaceHandleResponse(resp)
	if err != nil {
		return dataFlowClientGetDataFlowsByWorkspaceResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// dataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspacePager provides operations for iterating over paged responses.
type dataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspacePager struct {
	client    *dataFlowDebugSessionClient
	current   dataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, dataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *dataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspacePager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.QueryDataFlowDebugSessionsResponse.NextLink == nil || len(*p.current.QueryDataFlowDebugSessionsResponse.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *dataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspacePager) NextPage(ctx context.Context) (dataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return dataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return dataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return dataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return dataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.queryDataFlowDebugSessionsByWorkspaceHandleResponse(resp)
	if err != nil {
		return dataFlowDebugSessionClientQueryDataFlowDebugSessionsByWorkspaceResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// datasetClientGetDatasetsByWorkspacePager provides operations for iterating over paged responses.
type datasetClientGetDatasetsByWorkspacePager struct {
	client    *datasetClient
	current   datasetClientGetDatasetsByWorkspaceResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, datasetClientGetDatasetsByWorkspaceResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *datasetClientGetDatasetsByWorkspacePager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DatasetListResponse.NextLink == nil || len(*p.current.DatasetListResponse.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *datasetClientGetDatasetsByWorkspacePager) NextPage(ctx context.Context) (datasetClientGetDatasetsByWorkspaceResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return datasetClientGetDatasetsByWorkspaceResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return datasetClientGetDatasetsByWorkspaceResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return datasetClientGetDatasetsByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return datasetClientGetDatasetsByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getDatasetsByWorkspaceHandleResponse(resp)
	if err != nil {
		return datasetClientGetDatasetsByWorkspaceResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// libraryClientListPager provides operations for iterating over paged responses.
type libraryClientListPager struct {
	client    *libraryClient
	current   libraryClientListResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, libraryClientListResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *libraryClientListPager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LibraryListResponse.NextLink == nil || len(*p.current.LibraryListResponse.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *libraryClientListPager) NextPage(ctx context.Context) (libraryClientListResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return libraryClientListResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return libraryClientListResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return libraryClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return libraryClientListResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		return libraryClientListResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// linkedServiceClientGetLinkedServicesByWorkspacePager provides operations for iterating over paged responses.
type linkedServiceClientGetLinkedServicesByWorkspacePager struct {
	client    *linkedServiceClient
	current   linkedServiceClientGetLinkedServicesByWorkspaceResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, linkedServiceClientGetLinkedServicesByWorkspaceResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *linkedServiceClientGetLinkedServicesByWorkspacePager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LinkedServiceListResponse.NextLink == nil || len(*p.current.LinkedServiceListResponse.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *linkedServiceClientGetLinkedServicesByWorkspacePager) NextPage(ctx context.Context) (linkedServiceClientGetLinkedServicesByWorkspaceResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return linkedServiceClientGetLinkedServicesByWorkspaceResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return linkedServiceClientGetLinkedServicesByWorkspaceResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return linkedServiceClientGetLinkedServicesByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return linkedServiceClientGetLinkedServicesByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getLinkedServicesByWorkspaceHandleResponse(resp)
	if err != nil {
		return linkedServiceClientGetLinkedServicesByWorkspaceResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// notebookClientGetNotebookSummaryByWorkSpacePager provides operations for iterating over paged responses.
type notebookClientGetNotebookSummaryByWorkSpacePager struct {
	client    *notebookClient
	current   notebookClientGetNotebookSummaryByWorkSpaceResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, notebookClientGetNotebookSummaryByWorkSpaceResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *notebookClientGetNotebookSummaryByWorkSpacePager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.NotebookListResponse.NextLink == nil || len(*p.current.NotebookListResponse.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *notebookClientGetNotebookSummaryByWorkSpacePager) NextPage(ctx context.Context) (notebookClientGetNotebookSummaryByWorkSpaceResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return notebookClientGetNotebookSummaryByWorkSpaceResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return notebookClientGetNotebookSummaryByWorkSpaceResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return notebookClientGetNotebookSummaryByWorkSpaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return notebookClientGetNotebookSummaryByWorkSpaceResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getNotebookSummaryByWorkSpaceHandleResponse(resp)
	if err != nil {
		return notebookClientGetNotebookSummaryByWorkSpaceResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// notebookClientGetNotebooksByWorkspacePager provides operations for iterating over paged responses.
type notebookClientGetNotebooksByWorkspacePager struct {
	client    *notebookClient
	current   notebookClientGetNotebooksByWorkspaceResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, notebookClientGetNotebooksByWorkspaceResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *notebookClientGetNotebooksByWorkspacePager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.NotebookListResponse.NextLink == nil || len(*p.current.NotebookListResponse.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *notebookClientGetNotebooksByWorkspacePager) NextPage(ctx context.Context) (notebookClientGetNotebooksByWorkspaceResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return notebookClientGetNotebooksByWorkspaceResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return notebookClientGetNotebooksByWorkspaceResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return notebookClientGetNotebooksByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return notebookClientGetNotebooksByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getNotebooksByWorkspaceHandleResponse(resp)
	if err != nil {
		return notebookClientGetNotebooksByWorkspaceResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// pipelineClientGetPipelinesByWorkspacePager provides operations for iterating over paged responses.
type pipelineClientGetPipelinesByWorkspacePager struct {
	client    *pipelineClient
	current   pipelineClientGetPipelinesByWorkspaceResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, pipelineClientGetPipelinesByWorkspaceResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *pipelineClientGetPipelinesByWorkspacePager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PipelineListResponse.NextLink == nil || len(*p.current.PipelineListResponse.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *pipelineClientGetPipelinesByWorkspacePager) NextPage(ctx context.Context) (pipelineClientGetPipelinesByWorkspaceResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return pipelineClientGetPipelinesByWorkspaceResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return pipelineClientGetPipelinesByWorkspaceResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return pipelineClientGetPipelinesByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return pipelineClientGetPipelinesByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getPipelinesByWorkspaceHandleResponse(resp)
	if err != nil {
		return pipelineClientGetPipelinesByWorkspaceResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// sparkJobDefinitionClientGetSparkJobDefinitionsByWorkspacePager provides operations for iterating over paged responses.
type sparkJobDefinitionClientGetSparkJobDefinitionsByWorkspacePager struct {
	client    *sparkJobDefinitionClient
	current   sparkJobDefinitionClientGetSparkJobDefinitionsByWorkspaceResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, sparkJobDefinitionClientGetSparkJobDefinitionsByWorkspaceResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *sparkJobDefinitionClientGetSparkJobDefinitionsByWorkspacePager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SparkJobDefinitionsListResponse.NextLink == nil || len(*p.current.SparkJobDefinitionsListResponse.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *sparkJobDefinitionClientGetSparkJobDefinitionsByWorkspacePager) NextPage(ctx context.Context) (sparkJobDefinitionClientGetSparkJobDefinitionsByWorkspaceResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return sparkJobDefinitionClientGetSparkJobDefinitionsByWorkspaceResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return sparkJobDefinitionClientGetSparkJobDefinitionsByWorkspaceResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return sparkJobDefinitionClientGetSparkJobDefinitionsByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return sparkJobDefinitionClientGetSparkJobDefinitionsByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getSparkJobDefinitionsByWorkspaceHandleResponse(resp)
	if err != nil {
		return sparkJobDefinitionClientGetSparkJobDefinitionsByWorkspaceResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// sqlScriptClientGetSQLScriptsByWorkspacePager provides operations for iterating over paged responses.
type sqlScriptClientGetSQLScriptsByWorkspacePager struct {
	client    *sqlScriptClient
	current   sqlScriptClientGetSQLScriptsByWorkspaceResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, sqlScriptClientGetSQLScriptsByWorkspaceResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *sqlScriptClientGetSQLScriptsByWorkspacePager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SQLScriptsListResponse.NextLink == nil || len(*p.current.SQLScriptsListResponse.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *sqlScriptClientGetSQLScriptsByWorkspacePager) NextPage(ctx context.Context) (sqlScriptClientGetSQLScriptsByWorkspaceResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return sqlScriptClientGetSQLScriptsByWorkspaceResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return sqlScriptClientGetSQLScriptsByWorkspaceResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return sqlScriptClientGetSQLScriptsByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return sqlScriptClientGetSQLScriptsByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getSQLScriptsByWorkspaceHandleResponse(resp)
	if err != nil {
		return sqlScriptClientGetSQLScriptsByWorkspaceResponse{}, err
	}
	p.current = result
	return p.current, nil
}

// triggerClientGetTriggersByWorkspacePager provides operations for iterating over paged responses.
type triggerClientGetTriggersByWorkspacePager struct {
	client    *triggerClient
	current   triggerClientGetTriggersByWorkspaceResponse
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, triggerClientGetTriggersByWorkspaceResponse) (*policy.Request, error)
}

// More returns true if there are more pages to retrieve.
func (p *triggerClientGetTriggersByWorkspacePager) More() bool {
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.TriggerListResponse.NextLink == nil || len(*p.current.TriggerListResponse.NextLink) == 0 {
			return false
		}
	}
	return true
}

// NextPage advances the pager to the next page.
func (p *triggerClientGetTriggersByWorkspacePager) NextPage(ctx context.Context) (triggerClientGetTriggersByWorkspaceResponse, error) {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if !p.More() {
			return triggerClientGetTriggersByWorkspaceResponse{}, errors.New("no more pages")
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		return triggerClientGetTriggersByWorkspaceResponse{}, err
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		return triggerClientGetTriggersByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {

		return triggerClientGetTriggersByWorkspaceResponse{}, runtime.NewResponseError(resp)
	}
	result, err := p.client.getTriggersByWorkspaceHandleResponse(resp)
	if err != nil {
		return triggerClientGetTriggersByWorkspaceResponse{}, err
	}
	p.current = result
	return p.current, nil
}
