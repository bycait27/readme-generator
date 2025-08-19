package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/bycait27/readme-generator/internal/models"
)

// findTemplatesDir searches for the templates directory
func findTemplatesDir() string {
	// possible template directory locations
	possiblePaths := []string{
		"templates",          // current directory
		"../templates",       // one level up
		"../../templates",    // two levels up (for tests in subdirs)
		"../../../templates", // three levels up
	}

	for _, path := range possiblePaths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	// default to current directory
	return "templates"
}

// GenerateREADME generates a README file from BaseInfo and template
func GenerateREADME(templateName string, info *models.BaseInfo) error {
	templatesDir := findTemplatesDir()
	templatePath := filepath.Join(templatesDir, templateName+".md")

	// load the template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("failed to load template %s: %w", templateName, err)
	}

	// create the README file
	file, err := os.Create("README.md")
	if err != nil {
		return fmt.Errorf("failed to create README.md: %w", err)
	}
	defer file.Close()

	// execute the template with the data
	if err := tmpl.Execute(file, info); err != nil {
		return fmt.Errorf("failed to generate README: %w", err)
	}

	return nil
}

// GenerateREADMEToFile generates README to a custom file path
func GenerateREADMEToFile(templateName string, info *models.BaseInfo, filePath string) error {
	templatesDir := findTemplatesDir()
	templatePath := filepath.Join(templatesDir, templateName+".md")

	// load the template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return fmt.Errorf("failed to load template %s: %w", templateName, err)
	}

	// create the custom file
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	// execute the template with the data
	if err := tmpl.Execute(file, info); err != nil {
		return fmt.Errorf("failed to generate README: %w", err)
	}

	return nil
}

// ValidateTemplate checks if a template exists
func ValidateTemplate(templateName string) error {
	templatesDir := findTemplatesDir()
	filePath := filepath.Join(templatesDir, templateName+".md")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("template %s does not exist at %s", templateName, filePath)
	}
	return nil
}

// ListAvailableTemplates returns a list of available template names
func ListAvailableTemplates() ([]string, error) {
	templatesDir := findTemplatesDir()

	files, err := os.ReadDir(templatesDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read templates directory: %w", err)
	}

	var templates []string
	for _, file := range files {
		if !file.IsDir() && len(file.Name()) > 3 && file.Name()[len(file.Name())-3:] == ".md" {
			// remove .md extension
			templateName := file.Name()[:len(file.Name())-3]
			templates = append(templates, templateName)
		}
	}

	return templates, nil
}
