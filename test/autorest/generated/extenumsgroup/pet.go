// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package extenumsgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// PetOperations contains the methods for the Pet group.
type PetOperations interface {
	// AddPet - add pet
	AddPet(ctx context.Context, petAddPetOptions *PetAddPetOptions) (*PetResponse, error)
	// GetByPetID - get pet by id
	GetByPetID(ctx context.Context, petId string) (*PetResponse, error)
}

// PetClient implements the PetOperations interface.
// Don't use this type directly, use NewPetClient() instead.
type PetClient struct {
	*Client
}

// NewPetClient creates a new instance of PetClient with the specified values.
func NewPetClient(c *Client) PetOperations {
	return &PetClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *PetClient) Do(ctx context.Context, req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(ctx, req)
}

// AddPet - add pet
func (client *PetClient) AddPet(ctx context.Context, petAddPetOptions *PetAddPetOptions) (*PetResponse, error) {
	req, err := client.AddPetCreateRequest(petAddPetOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.AddPetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AddPetCreateRequest creates the AddPet request.
func (client *PetClient) AddPetCreateRequest(petAddPetOptions *PetAddPetOptions) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/extensibleenums/pet/addPet"
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodPost, *u)
	if petAddPetOptions != nil {
		return req, req.MarshalAsJSON(petAddPetOptions.PetParam)
	}
	return req, nil
}

// AddPetHandleResponse handles the AddPet response.
func (client *PetClient) AddPetHandleResponse(resp *azcore.Response) (*PetResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.AddPetHandleError(resp)
	}
	result := PetResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Pet)
}

// AddPetHandleError handles the AddPet error response.
func (client *PetClient) AddPetHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}

// GetByPetID - get pet by id
func (client *PetClient) GetByPetID(ctx context.Context, petId string) (*PetResponse, error) {
	req, err := client.GetByPetIDCreateRequest(petId)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetByPetIDHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByPetIDCreateRequest creates the GetByPetID request.
func (client *PetClient) GetByPetIDCreateRequest(petId string) (*azcore.Request, error) {
	u, err := url.Parse(client.u)
	if err != nil {
		return nil, err
	}
	urlPath := "/extensibleenums/pet/{petId}"
	urlPath = strings.ReplaceAll(urlPath, "{petId}", url.PathEscape(petId))
	u, err = u.Parse(path.Join(u.Path, urlPath))
	if err != nil {
		return nil, err
	}
	req := azcore.NewRequest(http.MethodGet, *u)
	return req, nil
}

// GetByPetIDHandleResponse handles the GetByPetID response.
func (client *PetClient) GetByPetIDHandleResponse(resp *azcore.Response) (*PetResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetByPetIDHandleError(resp)
	}
	result := PetResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Pet)
}

// GetByPetIDHandleError handles the GetByPetID error response.
func (client *PetClient) GetByPetIDHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return errors.New(resp.Status)
	}
	return errors.New(string(body))
}
