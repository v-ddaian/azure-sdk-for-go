// Package creator implements the Azure ARM Creator service API version .
//
//
package creator

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultGeography is the default value for geography
	DefaultGeography = Us
)

// BaseClient is the base client for Creator.
type BaseClient struct {
	autorest.Client
	Geography       GeographicResourceLocation
	XMsClientID     string
	SubscriptionKey string
}

// New creates an instance of the BaseClient client.
func New(xMsClientID string, SubscriptionKey string) BaseClient {
	return NewWithoutDefaults(xMsClientID, DefaultGeography, SubscriptionKey)
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults(xMsClientID string, geography GeographicResourceLocation, SubscriptionKey string) BaseClient {
	return BaseClient{
		Client:          autorest.NewClientWithUserAgent(UserAgent()),
		Geography:       geography,
		XMsClientID:     xMsClientID,
		SubscriptionKey: SubscriptionKey,
	}
}
