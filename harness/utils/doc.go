// Package utils provides utility functions for the Harness Go SDK.
//
// This package contains general-purpose utilities used throughout the SDK
// including string manipulation, type conversions, and common helper functions.
//
// # Available Utilities
//
// Environment variable helpers:
//
//	value := utils.GetEnv("MY_VAR", "default-value")
//
// String utilities:
//
//	slug := utils.ToSlug("My Resource Name")  // Returns: my-resource-name
//
// Type utilities:
//
//	ptr := utils.StringPtr("value")  // Returns pointer to string
package utils
