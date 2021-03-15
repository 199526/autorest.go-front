// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package errorsgroup

import "encoding/json"

type notFoundErrorBase struct {
	wrapped NotFoundErrorBaseClassification
}

func (n *notFoundErrorBase) UnmarshalJSON(data []byte) (err error) {
	n.wrapped, err = unmarshalNotFoundErrorBaseClassification((*json.RawMessage)(&data))
	return
}

func unmarshalNotFoundErrorBaseClassification(rawMsg *json.RawMessage) (NotFoundErrorBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(*rawMsg, &m); err != nil {
		return nil, err
	}
	var b NotFoundErrorBaseClassification
	switch m["whatNotFound"] {
	case "AnimalNotFound":
		b = &AnimalNotFound{}
	case "InvalidResourceLink":
		b = &LinkNotFound{}
	default:
		b = &NotFoundErrorBase{}
	}
	return b, json.Unmarshal(*rawMsg, &b)
}

func unmarshalNotFoundErrorBaseClassificationArray(rawMsg *json.RawMessage) (*[]NotFoundErrorBaseClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(*rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]NotFoundErrorBaseClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalNotFoundErrorBaseClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

type petActionError struct {
	wrapped PetActionErrorClassification
}

func (p *petActionError) UnmarshalJSON(data []byte) (err error) {
	p.wrapped, err = unmarshalPetActionErrorClassification((*json.RawMessage)(&data))
	return
}

func unmarshalPetActionErrorClassification(rawMsg *json.RawMessage) (PetActionErrorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(*rawMsg, &m); err != nil {
		return nil, err
	}
	var b PetActionErrorClassification
	switch m["errorType"] {
	case "PetHungryOrThirstyError":
		b = &PetHungryOrThirstyError{}
	case "PetSadError":
		b = &PetSadError{}
	default:
		b = &PetActionError{}
	}
	return b, json.Unmarshal(*rawMsg, &b)
}

func unmarshalPetActionErrorClassificationArray(rawMsg *json.RawMessage) (*[]PetActionErrorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(*rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]PetActionErrorClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalPetActionErrorClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

func unmarshalPetSadErrorClassification(rawMsg *json.RawMessage) (PetSadErrorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var m map[string]interface{}
	if err := json.Unmarshal(*rawMsg, &m); err != nil {
		return nil, err
	}
	var b PetSadErrorClassification
	switch m["errorType"] {
	case "PetHungryOrThirstyError":
		b = &PetHungryOrThirstyError{}
	default:
		b = &PetSadError{}
	}
	return b, json.Unmarshal(*rawMsg, &b)
}

func unmarshalPetSadErrorClassificationArray(rawMsg *json.RawMessage) (*[]PetSadErrorClassification, error) {
	if rawMsg == nil {
		return nil, nil
	}
	var rawMessages []*json.RawMessage
	if err := json.Unmarshal(*rawMsg, &rawMessages); err != nil {
		return nil, err
	}
	fArray := make([]PetSadErrorClassification, len(rawMessages))
	for index, rawMessage := range rawMessages {
		f, err := unmarshalPetSadErrorClassification(rawMessage)
		if err != nil {
			return nil, err
		}
		fArray[index] = f
	}
	return &fArray, nil
}

func strptr(s string) *string {
	return &s
}
