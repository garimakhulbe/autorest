package optionalgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
// 
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
)    

// ArrayOptionalWrapper is
type ArrayOptionalWrapper struct {
    Value *[]string `json:"value,omitempty"`
}

// ArrayWrapper is
type ArrayWrapper struct {
    Value *[]string `json:"value,omitempty"`
}

// ClassOptionalWrapper is
type ClassOptionalWrapper struct {
    Value *Product `json:"value,omitempty"`
}

// ClassWrapper is
type ClassWrapper struct {
    Value *Product `json:"value,omitempty"`
}

// Error is
type Error struct {
    autorest.Response `json:"-"`
    Status *int32 `json:"status,omitempty"`
    Message *string `json:"message,omitempty"`
}

// IntOptionalWrapper is
type IntOptionalWrapper struct {
    Value *int32 `json:"value,omitempty"`
}

// IntWrapper is
type IntWrapper struct {
    Value *int32 `json:"value,omitempty"`
}

// Product is
type Product struct {
    ID *int32 `json:"id,omitempty"`
    Name *string `json:"name,omitempty"`
}

// StringOptionalWrapper is
type StringOptionalWrapper struct {
    Value *string `json:"value,omitempty"`
}

// StringWrapper is
type StringWrapper struct {
    Value *string `json:"value,omitempty"`
}

