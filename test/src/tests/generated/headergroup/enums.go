package headergroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// GreyscaleColors enumerates the values for greyscale colors.
type GreyscaleColors string

const (
	// Black ...
	Black GreyscaleColors = "black"
	// GREY ...
	GREY GreyscaleColors = "GREY"
	// White ...
	White GreyscaleColors = "White"
)

// PossibleGreyscaleColorsValues returns an array of possible values for the GreyscaleColors const type.
func PossibleGreyscaleColorsValues() []GreyscaleColors {
	return []GreyscaleColors{Black, GREY, White}
}