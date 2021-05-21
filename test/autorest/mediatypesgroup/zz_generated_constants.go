// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package mediatypesgroup

const telemetryInfo = "azsdk-go-mediatypesgroup/<version>"

// ContentType - Content type for upload
type ContentType string

const (
	// ContentTypeApplicationPDF - Content Type 'application/pdf'
	ContentTypeApplicationPDF ContentType = "application/pdf"
	// ContentTypeImageJPEG - Content Type 'image/jpeg'
	ContentTypeImageJPEG ContentType = "image/jpeg"
	// ContentTypeImagePNG - Content Type 'image/png'
	ContentTypeImagePNG ContentType = "image/png"
	// ContentTypeImageTIFF - Content Type 'image/tiff'
	ContentTypeImageTIFF ContentType = "image/tiff"
)

// PossibleContentTypeValues returns the possible values for the ContentType const type.
func PossibleContentTypeValues() []ContentType {
	return []ContentType{
		ContentTypeApplicationPDF,
		ContentTypeImageJPEG,
		ContentTypeImagePNG,
		ContentTypeImageTIFF,
	}
}

// ToPtr returns a *ContentType pointing to the current value.
func (c ContentType) ToPtr() *ContentType {
	return &c
}
