// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package errorsgroup

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Animal struct {
	AniType *string `json:"aniType,omitempty"`
}

type AnimalNotFound struct {
	NotFoundErrorBase
	Name *string `json:"name,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AnimalNotFound.
func (a *AnimalNotFound) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			if val != nil {
				err = json.Unmarshal(*val, &a.Name)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return a.NotFoundErrorBase.unmarshalInternal(rawMsg)
}

type BaseError struct {
	SomeBaseProp *string `json:"someBaseProp,omitempty"`
}

func (b *BaseError) unmarshalInternal(rawMsg map[string]*json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "someBaseProp":
			if val != nil {
				err = json.Unmarshal(*val, &b.SomeBaseProp)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

type LinkNotFound struct {
	NotFoundErrorBase
	WhatSubAddress *string `json:"whatSubAddress,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LinkNotFound.
func (l *LinkNotFound) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "whatSubAddress":
			if val != nil {
				err = json.Unmarshal(*val, &l.WhatSubAddress)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return l.NotFoundErrorBase.unmarshalInternal(rawMsg)
}

// NotFoundErrorBaseClassification provides polymorphic access to related types.
// Call the interface's GetNotFoundErrorBase() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *NotFoundErrorBase, *AnimalNotFound, *LinkNotFound
type NotFoundErrorBaseClassification interface {
	error
	// GetNotFoundErrorBase() returns the NotFoundErrorBase content of the underlying type.
	GetNotFoundErrorBase() *NotFoundErrorBase
}

type NotFoundErrorBase struct {
	BaseError
	Reason       *string `json:"reason,omitempty"`
	WhatNotFound *string `json:"whatNotFound,omitempty"`
}

// Error implements the error interface for type NotFoundErrorBase.
func (e NotFoundErrorBase) Error() string {
	msg := ""
	if e.Reason != nil {
		msg += fmt.Sprintf("Reason: %v\n", *e.Reason)
	}
	if e.WhatNotFound != nil {
		msg += fmt.Sprintf("WhatNotFound: %v\n", *e.WhatNotFound)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}

// GetNotFoundErrorBase implements the NotFoundErrorBaseClassification interface for type NotFoundErrorBase.
func (n *NotFoundErrorBase) GetNotFoundErrorBase() *NotFoundErrorBase { return n }

// UnmarshalJSON implements the json.Unmarshaller interface for type NotFoundErrorBase.
func (n *NotFoundErrorBase) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return n.unmarshalInternal(rawMsg)
}

func (n *NotFoundErrorBase) unmarshalInternal(rawMsg map[string]*json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "reason":
			if val != nil {
				err = json.Unmarshal(*val, &n.Reason)
			}
			delete(rawMsg, key)
		case "whatNotFound":
			if val != nil {
				err = json.Unmarshal(*val, &n.WhatNotFound)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return n.BaseError.unmarshalInternal(rawMsg)
}

type Pet struct {
	Animal
	// READ-ONLY; Gets the Pet by id.
	Name *string `json:"name,omitempty" azure:"ro"`
}

type PetAction struct {
	// action feedback
	ActionResponse *string `json:"actionResponse,omitempty"`
}

// PetActionErrorClassification provides polymorphic access to related types.
// Call the interface's GetPetActionError() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *PetActionError, *PetHungryOrThirstyError, *PetSadError
type PetActionErrorClassification interface {
	error
	// GetPetActionError() returns the PetActionError content of the underlying type.
	GetPetActionError() *PetActionError
}

type PetActionError struct {
	// the error message
	ErrorMessage *string `json:"errorMessage,omitempty"`
	ErrorType    *string `json:"errorType,omitempty"`
}

// Error implements the error interface for type PetActionError.
func (e PetActionError) Error() string {
	msg := ""
	if e.ErrorMessage != nil {
		msg += fmt.Sprintf("ErrorMessage: %v\n", *e.ErrorMessage)
	}
	if e.ErrorType != nil {
		msg += fmt.Sprintf("ErrorType: %v\n", *e.ErrorType)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}

// GetPetActionError implements the PetActionErrorClassification interface for type PetActionError.
func (p *PetActionError) GetPetActionError() *PetActionError { return p }

// UnmarshalJSON implements the json.Unmarshaller interface for type PetActionError.
func (p *PetActionError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return p.unmarshalInternal(rawMsg)
}

func (p *PetActionError) unmarshalInternal(rawMsg map[string]*json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "errorMessage":
			if val != nil {
				err = json.Unmarshal(*val, &p.ErrorMessage)
			}
			delete(rawMsg, key)
		case "errorType":
			if val != nil {
				err = json.Unmarshal(*val, &p.ErrorType)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// PetActionResponse is the response envelope for operations that return a PetAction type.
type PetActionResponse struct {
	PetAction *PetAction

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PetDoSomethingOptions contains the optional parameters for the Pet.DoSomething method.
type PetDoSomethingOptions struct {
	// placeholder for future optional parameters
}

// PetGetPetByIDOptions contains the optional parameters for the Pet.GetPetByID method.
type PetGetPetByIDOptions struct {
	// placeholder for future optional parameters
}

type PetHungryOrThirstyError struct {
	PetSadError
	// is the pet hungry or thirsty or both
	HungryOrThirsty *string `json:"hungryOrThirsty,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PetHungryOrThirstyError.
func (p *PetHungryOrThirstyError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "hungryOrThirsty":
			if val != nil {
				err = json.Unmarshal(*val, &p.HungryOrThirsty)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return p.PetSadError.unmarshalInternal(rawMsg)
}

// PetResponse is the response envelope for operations that return a Pet type.
type PetResponse struct {
	Pet *Pet

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PetSadErrorClassification provides polymorphic access to related types.
// Call the interface's GetPetSadError() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *PetSadError, *PetHungryOrThirstyError
type PetSadErrorClassification interface {
	PetActionErrorClassification
	// GetPetSadError() returns the PetSadError content of the underlying type.
	GetPetSadError() *PetSadError
}

type PetSadError struct {
	PetActionError
	// why is the pet sad
	Reason *string `json:"reason,omitempty"`
}

// GetPetSadError implements the PetSadErrorClassification interface for type PetSadError.
func (p *PetSadError) GetPetSadError() *PetSadError { return p }

// UnmarshalJSON implements the json.Unmarshaller interface for type PetSadError.
func (p *PetSadError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	return p.unmarshalInternal(rawMsg)
}

func (p *PetSadError) unmarshalInternal(rawMsg map[string]*json.RawMessage) error {
	for key, val := range rawMsg {
		var err error
		switch key {
		case "reason":
			if val != nil {
				err = json.Unmarshal(*val, &p.Reason)
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return p.PetActionError.unmarshalInternal(rawMsg)
}