package urlgroup

// Code generated by Microsoft (R) AutoRest Code Generator 0.17.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
// "github.com/Azure/go-autorest/autorest"
)

// URIColor enumerates the values for uri color.
type URIColor string

const (
	// Bluecolor specifies the bluecolor state for uri color.
	Bluecolor URIColor = "blue color"
	// Greencolor specifies the greencolor state for uri color.
	Greencolor URIColor = "green color"
	// Redcolor specifies the redcolor state for uri color.
	Redcolor URIColor = "red color"
)

// Error is
type Error struct {
	Status  *int32  `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}
