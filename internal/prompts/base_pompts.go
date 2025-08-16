package prompts

import (
	"errors"
	"strings"

	"github.com/bycait27/readme-generator/internal/models"
	"github.com/bycait27/readme-generator/internal/validation"

	"github.com/manifoldco/promptui"
)

func PromptBaseInfo() (*models.BaseInfo, error) {
	// prompt for title
	title, err := promptTitle()
	if err != nil {
		return nil, err
	}

	// prompt for description
	description, err := promptDescription()
	if err != nil {
		return nil, err
	}

	// prompt for author info
	author, err := promptAuthorInfo()
	if err != nil {
		return nil, err
	}

	// prompt for license
	license, err := promptLicense()
	if err != nil {
		return nil, err
	}

	// prompt for optional sections
	screenshots, techStack, err := promptOptionalSections()
	if err != nil {
		return nil, err
	}

	// create BaseInfo
	baseInfo := &models.BaseInfo{
		Title:       title,
		Description: description,
		License:     license,
		Author:      *author,
		Screenshots: screenshots, // optional
		TechStack:   techStack,   // optional
	}

	// final validation
	if err := validation.ValidateStruct(baseInfo); err != nil {
		return nil, err
	}

	return baseInfo, nil
}

func promptAuthorInfo() (*models.AuthorInfo, error) {
	// prompt for name
	name, err := promptName()
	if err != nil {
		return nil, err
	}

	// prompt for email
	email, err := promptEmail()
	if err != nil {
		return nil, err
	}

	// prompt for github
	github, err := promptGithub()
	if err != nil {
		return nil, err
	}

	// prompt for website
	website, err := promptWebsite()
	if err != nil {
		return nil, err
	}

	//  create AuthorInfo
	authorInfo := &models.AuthorInfo{
		Name:    name,
		Email:   email,
		GitHub:  github,
		Website: website,
	}

	// final validation
	if err := validation.ValidateStruct(authorInfo); err != nil {
		return nil, err
	}

	return authorInfo, nil
}

func promptOptionalSections() (*models.Screenshots, *models.TechStack, error) {
	// prompt for optional sections

	// ask about screenshots
	screenshots, err := promptScreenshotsInfo()
	if err != nil {
		return nil, nil, err
	}

	// ask about tech stack
	techStack, err := promptTechStackInfo()
	if err != nil {
		return screenshots, nil, err
	}

	return screenshots, techStack, nil
}

func promptScreenshotsInfo() (*models.Screenshots, error) {
	prompt := promptui.Select{
		Label: "Do you want to include screenshots/demo?",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	// if no, return nil
	if result == "No" {
		return nil, nil
	}

	// if yes, prompt for screenshot details
	demo, err := promptDemo()
	if err != nil {
		return nil, err
	}

	description, err := promptDemoDescription()
	if err != nil {
		return nil, err
	}

	images, err := promptImages()
	if err != nil {
		return nil, err
	}

	screenshotsInfo := &models.Screenshots{
		Demo:        demo,
		Description: description,
		Images:      images,
	}

	return screenshotsInfo, nil
}

func promptTechStackInfo() (*models.TechStack, error) {
	prompt := promptui.Select{
		Label: "Do you want to include tech stack details?",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	// if no, return nil
	if result == "No" {
		return nil, nil
	}

	// if yes, prompt for tack stack details
	language, err := promptLanguage()
	if err != nil {
		return nil, err
	}

	framework, err := promptFramework()
	if err != nil {
		return nil, err
	}

	database, err := promptDatabase()
	if err != nil {
		return nil, err
	}

	dependencies, err := promptDependencies()
	if err != nil {
		return nil, err
	}

	techStackInfo := &models.TechStack{
		Language:     language,
		Framework:    framework,
		Database:     database,
		Dependencies: dependencies,
	}

	return techStackInfo, nil
}

func promptTitle() (string, error) {
	prompt := promptui.Prompt{
		Label: "Project Title",
		Validate: func(input string) error {
			if len(strings.TrimSpace(input)) < 1 {
				return errors.New("title is required")
			}
			if len(input) > 100 {
				return errors.New("title must be 100 characters or less")
			}
			return nil
		},
	}
	return prompt.Run()
}

func promptDescription() (string, error) {
	prompt := promptui.Prompt{
		Label: "Project Description",
		Validate: func(input string) error {
			if len(strings.TrimSpace(input)) < 10 {
				return errors.New("description must be at least 10 characters")
			}
			if len(input) > 500 {
				return errors.New("description must be 500 characters or less")
			}
			return nil
		},
	}
	return prompt.Run()
}

func promptLicense() (string, error) {
	prompt := promptui.Select{
		Label: "Choose Project License",
		Items: []string{"MIT", "Apache-2.0", "GPL-3.0", "BSD-3-Clause", "ISC", "Unlicense"},
	}
	_, result, err := prompt.Run()
	return result, err
}

func promptName() (string, error) {
	prompt := promptui.Prompt{
		Label: "Author Name",
		Validate: func(input string) error {
			if len(strings.TrimSpace(input)) < 2 {
				return errors.New("name must be at least 2 characters")
			}
			if len(input) > 50 {
				return errors.New("name must be 50 characters or less")
			}
			return nil
		},
	}
	return prompt.Run()
}

func promptEmail() (string, error) {
	prompt := promptui.Prompt{
		Label: "Author Email",
		Validate: func(input string) error {
			return nil
		},
	}
	return prompt.Run()
}

func promptGithub() (string, error) {
	prompt := promptui.Prompt{
		Label: "Author GitHub",
		Validate: func(input string) error {
			return nil
		},
	}
	return prompt.Run()
}

func promptWebsite() (string, error) {
	prompt := promptui.Prompt{
		Label: "Author Website",
		Validate: func(input string) error {
			return nil
		},
	}
	return prompt.Run()
}

func promptDemo() (string, error) {
	prompt := promptui.Prompt{
		Label: "Project Demo",
		Validate: func(input string) error {
			return nil
		},
	}
	return prompt.Run()
}

func promptDemoDescription() (string, error) {
	prompt := promptui.Prompt{
		Label: "Project Demo Description",
		Validate: func(input string) error {
			if len(strings.TrimSpace(input)) < 5 {
				return errors.New("demo description must be at least 5 characters")
			}
			if len(input) > 200 {
				return errors.New("demo description must be 200 characters or less")
			}
			return nil
		},
	}
	return prompt.Run()
}

func promptImages() ([]string, error) {
	prompt := promptui.Select{
		Label: "Do you have any additional screenshots?",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	if result == "No" {
		return nil, nil
	}

	// collect images as comma-separated
	imagePrompt := promptui.Prompt{
		Label: "Additional screenshot URLs/paths (comma-separated)",
		Validate: func(input string) error {
			if len(strings.TrimSpace(input)) < 1 {
				return errors.New("at least one image path/URL is required")
			}
			return nil
		},
	}
	images, err := imagePrompt.Run()
	if err != nil {
		return nil, err
	}

	// split by comma and clean up
	imageList := strings.Split(images, ",")
	for i, img := range imageList {
		imageList[i] = strings.TrimSpace(img)
	}

	return imageList, nil
}

func promptLanguage() (string, error) {
	prompt := promptui.Prompt{
		Label: "Project Language",
		Validate: func(input string) error {
			if len(strings.TrimSpace(input)) < 1 {
				return errors.New("langauge must be at least 1 characters")
			}
			if len(input) > 50 {
				return errors.New("language must be 50 characters or less")
			}
			return nil
		},
	}
	return prompt.Run()
}

func promptFramework() (string, error) {
	prompt := promptui.Prompt{
		Label: "Project Framework",
		Validate: func(input string) error {
			if len(input) > 50 {
				return errors.New("framework must be 50 characters or less")
			}
			return nil
		},
	}
	return prompt.Run()
}

func promptDatabase() (string, error) {
	prompt := promptui.Prompt{
		Label: "Project Database",
		Validate: func(input string) error {
			if len(input) > 50 {
				return errors.New("database must be 50 characters or less")
			}
			return nil
		},
	}
	return prompt.Run()
}

func promptDependencies() ([]string, error) {
	prompt := promptui.Select{
		Label: "Do you want to list key dependencies?",
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	if result == "No" {
		return nil, nil
	}

	// collect dependencies as comma-separated
	depPrompt := promptui.Prompt{
		Label: "Key dependencies (comma-separated)",
		Validate: func(input string) error {
			if len(strings.TrimSpace(input)) < 1 {
				return errors.New("at least one dependency is required")
			}
			return nil
		},
	}
	deps, err := depPrompt.Run()
	if err != nil {
		return nil, err
	}

	// split by comma and clean up
	dependencies := strings.Split(deps, ",")
	for i, dep := range dependencies {
		dependencies[i] = strings.TrimSpace(dep)
	}

	return dependencies, nil
}
