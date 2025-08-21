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
	title, err := promptRequiredText("Project Title", 1, 100)
	if err != nil {
		return nil, err
	}

	// prompt for description
	description, err := promptRequiredText("Project Description", 10, 500)
	if err != nil {
		return nil, err
	}

	// prompt for author info
	author, err := promptAuthorInfo()
	if err != nil {
		return nil, err
	}

	// prompt for license
	license, err := promptFromOptions("Choose Project License", []string{"MIT", "Apache-2.0", "GPL-3.0", "BSD-3-Clause", "ISC", "Unlicense"})
	if err != nil {
		return nil, err
	}

	// prompt for screenshots
	var screenshots *models.Screenshots
	wantScreenshots, err := promptYesNo("Do you want to include screenshots/demo?")
	if err != nil {
		return nil, err
	}
	if wantScreenshots {
		screenshots, err = promptScreenshotsInfo()
		if err != nil {
			return nil, err
		}
	}

	// prompt for tech stack
	var techStack *models.TechStack
	wantTechStack, err := promptYesNo("Do you want to include tech stack details?")
	if err != nil {
		return nil, err
	}
	if wantTechStack {
		techStack, err = promptTechStackInfo()
		if err != nil {
			return nil, err
		}
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
	name, err := promptRequiredText("Author Name", 2, 50)
	if err != nil {
		return nil, err
	}

	// prompt for email
	email, err := promptEmail("Author Email")
	if err != nil {
		return nil, err
	}

	// prompt for github
	github, err := promptURL("Author GitHub", true)
	if err != nil {
		return nil, err
	}

	// prompt for website
	website, err := promptURL("Author Website", false)
	if err != nil {
		return nil, err
	}

	//  create AuthorInfo
	authorInfo := &models.AuthorInfo{
		Name:    name,
		Email:   email,
		GitHub:  github,
		Website: website, // optional
	}

	// final validation
	if err := validation.ValidateStruct(authorInfo); err != nil {
		return nil, err
	}

	return authorInfo, nil
}

func promptScreenshotsInfo() (*models.Screenshots, error) {
	// prompt for demo
	demo, err := promptDemo()
	if err != nil {
		return nil, err
	}

	// prompt for demo description
	description, err := promptRequiredText("Demo Description", 5, 200)
	if err != nil {
		return nil, err
	}

	// prompt for additional screenshots
	var images []string
	wantImages, err := promptYesNo("Do you want to add additional screenshots?")
	if err != nil {
		return nil, err
	}
	if wantImages {
		images, err = promptStringList("additional screenshot paths", false)
		if err != nil {
			return nil, err
		}
	}

	// create ScreenshotsInfo
	screenshotsInfo := &models.Screenshots{
		Demo:        demo,
		Description: description,
		Images:      images, // optional
	}

	return screenshotsInfo, nil
}

func promptTechStackInfo() (*models.TechStack, error) {
	// prompt for primary language
	language, err := promptFromOptions("Primary Programming Language", []string{"Go", "Python", "JavaScript", "Java", "C#", "Ruby", "PHP", "C++", "TypeScript", "Swift", "Kotlin"})
	if err != nil {
		return nil, err
	}

	// prompt for primary framework
	var framework string
	wantFramework, err := promptYesNo("Do you want to specify a framework?")
	if err != nil {
		return nil, err
	}
	if wantFramework {
		framework, err = promptFromOptions("Primary Framework",
			[]string{
				"Gin", "Echo", "Fiber", "Gorilla Mux", "Chi", // Go frameworks
				"Express", "Fastify", "Koa", // Node.js
				"Django", "Flask", "FastAPI", // Python
				"Spring", "Spring Boot", // Java
				"ASP.NET", "ASP.NET Core", // C#
				"React", "Angular", "Vue", "Svelte", // Frontend
				"Ruby on Rails", "Sinatra", // Ruby
				"Laravel", "Symfony", // PHP
			})
		if err != nil {
			return nil, err
		}
	}

	// prompt for primary database
	var database string
	wantDatabase, err := promptYesNo("Do you want to specify a database?")
	if err != nil {
		return nil, err
	}
	if wantDatabase {
		database, err = promptFromOptions("Primary Database",
			[]string{"PostgreSQL", "MySQL", "MongoDB", "Redis", "SQLite", "Cassandra", "MariaDB", "OracleDB", "DynamoDB", "Firebase"})
		if err != nil {
			return nil, err
		}
	}

	// prompt for key dependencies/libraries
	var dependencies []string
	wantDependencies, err := promptYesNo("Do you want to list key dependencies/libraries?")
	if err != nil {
		return nil, err
	}
	if wantDependencies {
		dependencies, err = promptStringList("Key Dependencies/Libraries", false)
		if err != nil {
			return nil, err
		}
	}

	// create TechStackInfo
	techStackInfo := &models.TechStack{
		Language:     language,
		Framework:    framework,    // optional
		Database:     database,     // optional
		Dependencies: dependencies, // optional
	}

	return techStackInfo, nil
}

func promptDemo() (string, error) {
	prompt := promptui.Prompt{
		Label: "Project Demo (file path to GIF/video)",
		Validate: func(input string) error {
			trimmed := strings.TrimSpace(input)
			if len(trimmed) < 1 {
				return errors.New("demo file path is required")
			}
			if len(trimmed) > 200 {
				return errors.New("file path must be 200 characters or less")
			}

			// basic file path validation
			if strings.Contains(trimmed, "http://") || strings.Contains(trimmed, "https://") {
				return errors.New("please provide a file path, not a URL")
			}

			// check for valid file extensions
			validExtensions := []string{".gif", ".mp4", ".mov", ".webm", ".avi"}
			hasValidExtension := false
			lowerPath := strings.ToLower(trimmed)

			for _, ext := range validExtensions {
				if strings.HasSuffix(lowerPath, ext) {
					hasValidExtension = true
					break
				}
			}

			if !hasValidExtension {
				return errors.New("file must be a video/GIF (.gif, .mp4, .mov, .webm, .avi)")
			}

			return nil
		},
	}
	return prompt.Run()
}
