// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azurespecialsgrouptest

import (
	"context"
	"generatortests/autorest/generated/azurespecialsgroup"
	"generatortests/helpers"
	"net/http"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/to"
)

func getOdataOperations(t *testing.T) azurespecialsgroup.OdataOperations {
	client, err := azurespecialsgroup.NewDefaultClient(nil)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}
	return client.OdataOperations()
}

// GetWithFilter - Specify filter parameter with value '$filter=id gt 5 and name eq 'foo'&$orderby=id&$top=10'
func TestGetWithFilter(t *testing.T) {
	client := getOdataOperations(t)
	result, err := client.GetWithFilter(context.Background(), &azurespecialsgroup.OdataGetWithFilterOptions{
		Filter:  to.StringPtr("id gt 5 and name eq 'foo'"),
		Orderby: to.StringPtr("id"),
		Top:     to.Int32Ptr(10),
	})
	if err != nil {
		t.Fatal(err)
	}
	helpers.VerifyStatusCode(t, result, http.StatusOK)
}