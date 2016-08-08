package dategroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for
// license information.
// 
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
    "fmt"
)

const (
    major = "3"
    minor = "1"
    patch = "0"
    // Always begin a "tag" with a dash (as per http://semver.org)
    tag   = "-beta"
    semVerFormat = "%s.%s.%s%s"
    userAgentFormat = "Azure-SDK-for-Go/%s arm-%s/%s"
)

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
    return fmt.Sprintf(userAgentFormat, Version(), "dategroup", "1.0.0")
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
    return fmt.Sprintf(semVerFormat, major, minor, patch, tag)
}
