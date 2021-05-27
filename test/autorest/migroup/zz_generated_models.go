// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package migroup

type Cat struct {
	Feline
	Pet
	LikesMilk *bool `json:"likesMilk,omitempty"`
}

// Implements the error and azcore.HTTPResponse interfaces.
type Error struct {
	raw     string
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
// The contents of the error text are not contractual and subject to change.
func (e Error) Error() string {
	return e.raw
}

type Feline struct {
	Hisses *bool `json:"hisses,omitempty"`
	Meows  *bool `json:"meows,omitempty"`
}

type Horse struct {
	Pet
	IsAShowHorse *bool `json:"isAShowHorse,omitempty"`
}

type Kitten struct {
	Cat
	EatsMiceYet *bool `json:"eatsMiceYet,omitempty"`
}

// MultipleInheritanceServiceClientGetCatOptions contains the optional parameters for the MultipleInheritanceServiceClient.GetCat method.
type MultipleInheritanceServiceClientGetCatOptions struct {
	// placeholder for future optional parameters
}

// MultipleInheritanceServiceClientGetFelineOptions contains the optional parameters for the MultipleInheritanceServiceClient.GetFeline method.
type MultipleInheritanceServiceClientGetFelineOptions struct {
	// placeholder for future optional parameters
}

// MultipleInheritanceServiceClientGetHorseOptions contains the optional parameters for the MultipleInheritanceServiceClient.GetHorse method.
type MultipleInheritanceServiceClientGetHorseOptions struct {
	// placeholder for future optional parameters
}

// MultipleInheritanceServiceClientGetKittenOptions contains the optional parameters for the MultipleInheritanceServiceClient.GetKitten method.
type MultipleInheritanceServiceClientGetKittenOptions struct {
	// placeholder for future optional parameters
}

// MultipleInheritanceServiceClientGetPetOptions contains the optional parameters for the MultipleInheritanceServiceClient.GetPet method.
type MultipleInheritanceServiceClientGetPetOptions struct {
	// placeholder for future optional parameters
}

// MultipleInheritanceServiceClientPutCatOptions contains the optional parameters for the MultipleInheritanceServiceClient.PutCat method.
type MultipleInheritanceServiceClientPutCatOptions struct {
	// placeholder for future optional parameters
}

// MultipleInheritanceServiceClientPutFelineOptions contains the optional parameters for the MultipleInheritanceServiceClient.PutFeline method.
type MultipleInheritanceServiceClientPutFelineOptions struct {
	// placeholder for future optional parameters
}

// MultipleInheritanceServiceClientPutHorseOptions contains the optional parameters for the MultipleInheritanceServiceClient.PutHorse method.
type MultipleInheritanceServiceClientPutHorseOptions struct {
	// placeholder for future optional parameters
}

// MultipleInheritanceServiceClientPutKittenOptions contains the optional parameters for the MultipleInheritanceServiceClient.PutKitten method.
type MultipleInheritanceServiceClientPutKittenOptions struct {
	// placeholder for future optional parameters
}

// MultipleInheritanceServiceClientPutPetOptions contains the optional parameters for the MultipleInheritanceServiceClient.PutPet method.
type MultipleInheritanceServiceClientPutPetOptions struct {
	// placeholder for future optional parameters
}

type Pet struct {
	// REQUIRED
	Name *string `json:"name,omitempty"`
}