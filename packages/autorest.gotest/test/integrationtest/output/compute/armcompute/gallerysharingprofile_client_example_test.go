//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/compute/armcompute"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2020-09-30/examples/AddToSharingProfileInAGallery.json
func ExampleGallerySharingProfileClient_BeginUpdate_addToSharingProfileInAGallery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewGallerySharingProfileClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "myResourceGroup", "myGalleryName", armcompute.SharingUpdate{
		Groups: []*armcompute.SharingProfileGroup{
			{
				Type: to.Ptr(armcompute.SharingProfileGroupTypesSubscriptions),
				IDs: []*string{
					to.Ptr("34a4ab42-0d72-47d9-bd1a-aed207386dac"),
					to.Ptr("380fd389-260b-41aa-bad9-0a83108c370b")},
			},
			{
				Type: to.Ptr(armcompute.SharingProfileGroupTypesAADTenants),
				IDs: []*string{
					to.Ptr("c24c76aa-8897-4027-9b03-8f7928b54ff6")},
			}},
		OperationType: to.Ptr(armcompute.SharingUpdateOperationTypesAdd),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/compute/resource-manager/Microsoft.Compute/stable/2020-09-30/examples/ResetSharingProfileInAGallery.json
func ExampleGallerySharingProfileClient_BeginUpdate_resetSharingProfileInAGallery() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armcompute.NewGallerySharingProfileClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx, "myResourceGroup", "myGalleryName", armcompute.SharingUpdate{
		OperationType: to.Ptr(armcompute.SharingUpdateOperationTypesReset),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}
