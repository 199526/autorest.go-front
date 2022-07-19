//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package validationgroup

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ChildProduct.
func (c ChildProduct) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["constProperty"] = "constant"
	populate(objectMap, "count", c.Count)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ChildProduct.
func (c *ChildProduct) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "constProperty":
			err = unpopulate(val, "ConstProperty", &c.ConstProperty)
			delete(rawMsg, key)
		case "count":
			err = unpopulate(val, "Count", &c.Count)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ConstantProduct.
func (c ConstantProduct) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["constProperty"] = "constant"
	objectMap["constProperty2"] = "constant2"
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ConstantProduct.
func (c *ConstantProduct) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "constProperty":
			err = unpopulate(val, "ConstProperty", &c.ConstProperty)
			delete(rawMsg, key)
		case "constProperty2":
			err = unpopulate(val, "ConstProperty2", &c.ConstProperty2)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Product.
func (p Product) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "capacity", p.Capacity)
	populate(objectMap, "child", p.Child)
	populate(objectMap, "constChild", p.ConstChild)
	objectMap["constInt"] = 0
	objectMap["constString"] = "constant"
	objectMap["constStringAsEnum"] = "constant_string_as_enum"
	populate(objectMap, "display_names", p.DisplayNames)
	populate(objectMap, "image", p.Image)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Product.
func (p *Product) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "capacity":
			err = unpopulate(val, "Capacity", &p.Capacity)
			delete(rawMsg, key)
		case "child":
			err = unpopulate(val, "Child", &p.Child)
			delete(rawMsg, key)
		case "constChild":
			err = unpopulate(val, "ConstChild", &p.ConstChild)
			delete(rawMsg, key)
		case "constInt":
			err = unpopulate(val, "ConstInt", &p.ConstInt)
			delete(rawMsg, key)
		case "constString":
			err = unpopulate(val, "ConstString", &p.ConstString)
			delete(rawMsg, key)
		case "constStringAsEnum":
			err = unpopulate(val, "ConstStringAsEnum", &p.ConstStringAsEnum)
			delete(rawMsg, key)
		case "display_names":
			err = unpopulate(val, "DisplayNames", &p.DisplayNames)
			delete(rawMsg, key)
		case "image":
			err = unpopulate(val, "Image", &p.Image)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
