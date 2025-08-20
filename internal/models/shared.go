package models

type GettingStarted struct {
	Prerequisites []string `validate:"omitempty,dive,min=1"`
	EnvSetup      []string `validate:"omitempty,dive,min=1"`
	RunCommands   []string `validate:"min=1,dive,min=1"`
	Notes         *string  `validate:"omitempty,max=500"`
}

type EnvVar struct {
	Name        string `validate:"required,min=1,max=50"`
	Description string `validate:"required,min=5,max=200"`
	Required    bool
	Default     string `validate:"omitempty,max=100"`
	Example     string `validate:"required,max=100"`
}

type APIDocs struct {
	BaseURL        string         `validate:"required,url"`
	Authentication string         `validate:"required"`
	Endpoints      []Endpoint     `validate:"min=1,dive"`
	ErrorHandling  *ErrorHandling `validate:"omitempty"`
}

type Endpoint struct {
	Method      string      `validate:"required,oneof=GET POST PUT DELETE PATCH OPTIONS HEAD"`
	Path        string      `validate:"required,startswith=/"`
	Description string      `validate:"required,min=5,max=200"`
	Parameters  []Parameter `validate:"omitempty,dive"`
	Response    string      `validate:"required,min=1"`
}

type Parameter struct {
	Name        string `validate:"required,min=1,max=50"`
	Type        string `validate:"required,oneof=string int bool"`
	Required    bool
	Description string `validate:"required,min=5,max=200"`
	Example     string `validate:"required,max=100"`
}
type ErrorHandling struct {
	Format          string        `validate:"required,oneof=json xml plain"`
	StatusCodes     []StatusCode  `validate:"omitempty,dive"`
	ErrorResponse   ErrorResponse `validate:"omitempty"`
	CommonErrors    []CommonError `validate:"omitempty,dive"`
	ValidationRules string        `validate:"omitempty,max=500"`
}

type StatusCode struct {
	Code        int    `validate:"required,min=100,max=600"`
	Description string `validate:"required,min=5,max=200"`
	Example     string `validate:"omitempty,max=200"`
}

type ErrorResponse struct {
	Structure string `validate:"required"` // JSON structure example
	Example   string `validate:"required"` // actual example response
}

type CommonError struct {
	Code        int    `validate:"required,min=100,max=600"`
	Message     string `validate:"required,min=5,max=200"`
	Description string `validate:"required,min=5,max=200"`
	Solution    string `validate:"required,min=5,max=500"`
}

type Database struct {
	Type       string `validate:"required,oneof=PostgreSQL MySQL MongoDB Redis SQLite Cassandra MariaDB DynamoDB Firebase"`
	Schema     string `validate:"required,min=5,max=2000"`
	Migrations string `validate:"omitempty,min=5,max=500"`
	SeedData   string `validate:"omitempty,min=5,max=500"`
}

type Testing struct {
	TestCommand   string `validate:"required,min=1,max=200"`
	CoverageCmd   string `validate:"omitempty,max=200"`
	TestFramework string `validate:"omitempty,min=1,max=50"`
	Notes         string `validate:"omitempty,max=500"`
}

type Deployment struct {
	Platform    string `validate:"required,oneof=Docker Heroku AWS GCP Azure Vercel Netlify"`
	BuildCmd    string `validate:"omitempty,max=200"`
	DeployCmd   string `validate:"omitempty,max=200"`
	HealthCheck string `validate:"omitempty,max=200"`
	Notes       string `validate:"omitempty,max=500"`
}
