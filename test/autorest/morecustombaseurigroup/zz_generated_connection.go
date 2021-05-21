// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package morecustombaseurigroup

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// ConnectionOptions contains configuration settings for the connection's pipeline.
// All zero-value fields will be initialized with their default values.
type ConnectionOptions struct {
	// HTTPClient sets the transport for making HTTP requests.
	HTTPClient azcore.Transport
	// Retry configures the built-in retry policy behavior.
	Retry azcore.RetryOptions
	// Telemetry configures the built-in telemetry policy behavior.
	Telemetry azcore.TelemetryOptions
	// Logging configures the built-in logging policy behavior.
	Logging azcore.LogOptions
	// PerCallPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request.
	PerCallPolicies []azcore.Policy
	// PerRetryPolicies contains custom policies to inject into the pipeline.
	// Each policy is executed once per request, and for each retry request.
	PerRetryPolicies []azcore.Policy
}

func (c *ConnectionOptions) telemetryOptions() *azcore.TelemetryOptions {
	to := c.Telemetry
	if to.Value == "" {
		to.Value = telemetryInfo
	} else {
		to.Value = fmt.Sprintf("%s %s", telemetryInfo, to.Value)
	}
	return &to
}

// Connection - Test Infrastructure for AutoRest
type Connection struct {
	dnsSuffix string
	p         azcore.Pipeline
}

// NewConnection creates an instance of the Connection type with the specified endpoint.
// Pass nil to accept the default options; this is the same as passing a zero-value options.
func NewConnection(dnsSuffix *string, options *ConnectionOptions) *Connection {
	if options == nil {
		options = &ConnectionOptions{}
	}
	policies := []azcore.Policy{
		azcore.NewTelemetryPolicy(options.telemetryOptions()),
	}
	policies = append(policies, options.PerCallPolicies...)
	policies = append(policies, azcore.NewRetryPolicy(&options.Retry))
	policies = append(policies, options.PerRetryPolicies...)
	policies = append(policies, azcore.NewLogPolicy(&options.Logging))
	client := &Connection{
		p:         azcore.NewPipeline(options.HTTPClient, policies...),
		dnsSuffix: "host",
	}
	if dnsSuffix != nil {
		client.dnsSuffix = *dnsSuffix
	}
	return client
}

// DnsSuffix returns part of the parameterized host.
func (c *Connection) DnsSuffix() string {
	return c.dnsSuffix
}

// Pipeline returns the connection's pipeline.
func (c *Connection) Pipeline() azcore.Pipeline {
	return c.p
}
