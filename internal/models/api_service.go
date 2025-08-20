package models

type APIService struct {
	BaseInfo
	GettingStarted GettingStarted `validate:"required"`
	Database       *Database      `validate:"omitempty"`      // optional
	EnvVars        []EnvVar       `validate:"omitempty,dive"` // optional
	APIDocs        APIDocs        `validate:"required"`
	Auth           *Auth          `validate:"omitempty"` // optional
	Testing        *Testing       `validate:"omitempty"` // optional
	Deployment     *Deployment    `validate:"omitempty"` // optional
	Monitoring     *Monitoring    `validate:"omitempty"` // optional
}

type Auth struct {
	Method       string   `validate:"required,oneof=jwt oauth api-key basic none"` // "jwt", "oauth", "api-key"
	TokenFormat  string   `validate:"omitempty,max=200"`
	ExampleUsage string   `validate:"omitempty,max=500"`
	Endpoints    []string `validate:"omitempty,dive,startswith=/"` // auth-related endpoints
}

type Monitoring struct {
	HealthCheck string        `validate:"omitempty,max=200"`
	Metrics     []string      `validate:"omitempty,dive,min=1,max=100"`
	Logging     LoggingConfig `validate:"omitempty"`
	Alerts      []string      `validate:"omitempty,dive,min=1,max=100"`
}

type LoggingConfig struct {
	Level  string `validate:"omitempty,oneof=DEBUG INFO WARN ERROR"`
	Format string `validate:"omitempty,oneof=json text"`
	Output string `validate:"omitempty,max=200"` // e.g., file path or stdout
}
