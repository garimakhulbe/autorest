package report

// Code generated by Microsoft (R) AutoRest Code Generator 0.14.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
)    

// Error is
type Error struct {
    Status *int32 `json:"status,omitempty"`
    Message *string `json:"message,omitempty"`
}

// SetInt32 is
type SetInt32 struct {
    autorest.Response `json:"-"`
    Value *map[string]*int32 `json:"value,omitempty"`
}

