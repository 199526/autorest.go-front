package modelflatteninggroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ProvisioningStateValues enumerates the values for provisioning state values.
type ProvisioningStateValues string

const (
	// Accepted ...
	Accepted ProvisioningStateValues = "Accepted"
	// Canceled ...
	Canceled ProvisioningStateValues = "canceled"
	// Created ...
	Created ProvisioningStateValues = "Created"
	// Creating ...
	Creating ProvisioningStateValues = "Creating"
	// Deleted ...
	Deleted ProvisioningStateValues = "Deleted"
	// Deleting ...
	Deleting ProvisioningStateValues = "Deleting"
	// Failed ...
	Failed ProvisioningStateValues = "Failed"
	// OK ...
	OK ProvisioningStateValues = "OK"
	// Succeeded ...
	Succeeded ProvisioningStateValues = "Succeeded"
	// Updated ...
	Updated ProvisioningStateValues = "Updated"
	// Updating ...
	Updating ProvisioningStateValues = "Updating"
)

// PossibleProvisioningStateValuesValues returns an array of possible values for the ProvisioningStateValues const type.
func PossibleProvisioningStateValuesValues() []ProvisioningStateValues {
	return []ProvisioningStateValues{Accepted, Canceled, Created, Creating, Deleted, Deleting, Failed, OK, Succeeded, Updated, Updating}
}