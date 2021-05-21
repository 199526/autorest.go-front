// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package xmlgroup

import "net/http"

// AppleBarrelResponse is the response envelope for operations that return a AppleBarrel type.
type AppleBarrelResponse struct {
	// A barrel of apples.
	AppleBarrel *AppleBarrel `xml:"AppleBarrel"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BananaArrayResponse is the response envelope for operations that return a []*Banana type.
type BananaArrayResponse struct {
	// Array of Banana
	Bananas []*Banana `xml:"banana"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BananaResponse is the response envelope for operations that return a Banana type.
type BananaResponse struct {
	// A banana.
	Banana *Banana `xml:"banana"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// JSONOutputResponse is the response envelope for operations that return a JSONOutput type.
type JSONOutputResponse struct {
	JSONOutput *JSONOutput

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ListBlobsResponseResponse is the response envelope for operations that return a ListBlobsResponse type.
type ListBlobsResponseResponse struct {
	// An enumeration of blobs
	EnumerationResults *ListBlobsResponse `xml:"EnumerationResults"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ListContainersResponseResponse is the response envelope for operations that return a ListContainersResponse type.
type ListContainersResponseResponse struct {
	// An enumeration of containers
	EnumerationResults *ListContainersResponse `xml:"EnumerationResults"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ModelWithBytePropertyResponse is the response envelope for operations that return a ModelWithByteProperty type.
type ModelWithBytePropertyResponse struct {
	ModelWithByteProperty *ModelWithByteProperty `xml:"ModelWithByteProperty"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ModelWithURLPropertyResponse is the response envelope for operations that return a ModelWithURLProperty type.
type ModelWithURLPropertyResponse struct {
	ModelWithURLProperty *ModelWithURLProperty `xml:"ModelWithURLProperty"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ObjectWithXMsTextPropertyResponse is the response envelope for operations that return a ObjectWithXMsTextProperty type.
type ObjectWithXMsTextPropertyResponse struct {
	// Contans property
	Data *ObjectWithXMsTextProperty `xml:"Data"`

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// RootWithRefAndMetaResponse is the response envelope for operations that return a RootWithRefAndMeta type.
type RootWithRefAndMetaResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// I am root, and I ref a model WITH meta
	RootWithRefAndMeta *RootWithRefAndMeta `xml:"RootWithRefAndMeta"`
}

// RootWithRefAndNoMetaResponse is the response envelope for operations that return a RootWithRefAndNoMeta type.
type RootWithRefAndNoMetaResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// I am root, and I ref a model with no meta
	RootWithRefAndNoMeta *RootWithRefAndNoMeta `xml:"RootWithRefAndNoMeta"`
}

// SignedIdentifierArrayResponse is the response envelope for operations that return a []*SignedIdentifier type.
type SignedIdentifierArrayResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// a collection of signed identifiers
	SignedIdentifiers []*SignedIdentifier `xml:"SignedIdentifier"`
}

// SlideshowResponse is the response envelope for operations that return a Slideshow type.
type SlideshowResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Data about a slideshow
	Slideshow *Slideshow `xml:"slideshow"`
}

// StorageServicePropertiesResponse is the response envelope for operations that return a StorageServiceProperties type.
type StorageServicePropertiesResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Storage Service Properties.
	StorageServiceProperties *StorageServiceProperties `xml:"StorageServiceProperties"`
}

// XMLGetHeadersResponse contains the response from method XML.GetHeaders.
type XMLGetHeadersResponse struct {
	// CustomHeader contains the information returned from the Custom-Header header response.
	CustomHeader *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}
