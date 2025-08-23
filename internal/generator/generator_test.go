package generator

import (
	"os"
	"strings"
	"testing"

	"github.com/bycait27/readme-generator/internal/models"
)

func TestGenerateREADME(t *testing.T) {
	// Create test data
	baseInfo := &models.BaseInfo{
		Title:       "Test Project",
		Description: "This is a test project for README generation",
		License:     "MIT",
		Author: models.AuthorInfo{
			Name:   "Test Author",
			Email:  "test@example.com",
			GitHub: "https://github.com/testuser",
		},
		TechStack: &models.TechStack{
			Language:     "Go",
			Framework:    "Gin",
			Dependencies: []string{"gin", "gorm"},
		},
	}

	// Generate README
	err := GenerateREADME("basic", baseInfo)
	if err != nil {
		t.Fatalf("GenerateREADME failed: %v", err)
	}

	// Verify file was created
	content, err := os.ReadFile("README.md")
	if err != nil {
		t.Fatalf("Failed to read generated README: %v", err)
	}

	// Verify content
	contentStr := string(content)
	if !strings.Contains(contentStr, "Test Project") {
		t.Error("Generated README doesn't contain project title")
	}
	if !strings.Contains(contentStr, "This is a test project") {
		t.Error("Generated README doesn't contain description")
	}
	if !strings.Contains(contentStr, "Test Author") {
		t.Error("Generated README doesn't contain author name")
	}
	if !strings.Contains(contentStr, "**Language:** Go") {
		t.Error("Generated README doesn't contain tech stack")
	}

	// Cleanup
	os.Remove("README.md")
}

func TestValidateTemplate(t *testing.T) {
	// Test existing template
	err := ValidateTemplate("basic")
	if err != nil {
		t.Errorf("basic template should exist: %v", err)
	}

	// Test non-existing template
	err = ValidateTemplate("nonexistent")
	if err == nil {
		t.Error("nonexistent template should return error")
	}
}

func TestListAvailableTemplates(t *testing.T) {
	templates, err := ListAvailableTemplates()
	if err != nil {
		t.Fatalf("Failed to list templates: %v", err)
	}

	// Should at least contain "basic"
	found := false
	for _, tmpl := range templates {
		if tmpl == "basic" {
			found = true
			break
		}
	}

	if !found {
		t.Error("basic template should be in available templates list")
	}
}
