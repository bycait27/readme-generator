package models

type BaseInfo struct {
	Title       string       `validate:"required,min=1,max=100"`
	Description string       `validate:"required,min=10,max=500"`
	Screenshots *Screenshots `validate:"omitempty"` // optional
	TechStack   *TechStack   `validate:"omitempty"` // optional for CLI tools
	License     string       `validate:"required,oneof=MIT Apache-2.0 GPL-3.0 BSD-3-Clause ISC Unlicense"`
	Author      AuthorInfo   `validate:"required"`
}

type Screenshots struct {
	Demo        string   `validate:"required,min=1,max=200"`       // path to demo gif/video
	Images      []string `validate:"omitempty,dive,min=1,max=200"` // optional additional screenshots
	Description string   `validate:"required,min=5,max=200"`       // caption for demo
}

type TechStack struct {
	Language     string   `validate:"required,min=1,max=50"`
	Framework    string   `validate:"omitempty,max=50"`             // optional
	Database     string   `validate:"omitempty,max=50"`             // optional
	Dependencies []string `validate:"omitempty,dive,min=1,max=100"` // optional
}

type AuthorInfo struct {
	Name    string `validate:"required,min=2,max=50"`
	Email   string `validate:"required,email"`
	GitHub  string `validate:"required,url"`  // GitHub profile url
	Website string `validate:"omitempty,url"` // optional personal website
}
