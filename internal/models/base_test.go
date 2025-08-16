package models

import (
	"testing"

	"github.com/bycait27/readme-generator/internal/validation"
)

func TestBaseInfo_Valid(t *testing.T) {
	base := &BaseInfo{
		Title:       "My Awesome Project",
		Description: "This is a really cool project that does amazing things.",
		License:     "MIT",
		Author: AuthorInfo{
			Name:   "John Doe",
			Email:  "john@example.com",
			GitHub: "https://github.com/johndoe",
		},
	}

	err := validation.ValidateStruct(base)
	if err != nil {
		t.Errorf("Valid BaseInfo should pass validation, got: %v", err)
	}
}

func TestBaseInfo_Invalid(t *testing.T) {
	base := &BaseInfo{
		Title:       "",          // required but empty
		Description: "Too short", // less than 10 characters
		License:     "INVALID",   // not in the oneof list
		Author: AuthorInfo{
			Name:   "",
			Email:  "bad-email",
			GitHub: "not-a-url",
		},
	}

	err := validation.ValidateStruct(base)
	if err == nil {
		t.Errorf("Invalid BaseInfo should fail validation")
	}

	t.Logf("Validation errors: %v", err)
}
