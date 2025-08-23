package models

type FullStackApp struct {
	BaseInfo
	Architecture   *Architecture  `validate:"omitempty"` // optional
	GettingStarted GettingStarted `validate:"required"`
	EnvVars        []EnvVar       `validate:"omitempty,dive"` // optional
	AppStructure   *AppStructure  `validate:"omitempty"`      // optional
	Development    *Development   `validate:"omitempty"`      // optional
	Testing        *Testing       `validate:"omitempty"`      // optional
	Deployment     *Deployment    `validate:"omitempty"`      // optional
}

type Architecture struct {
	Pattern    string      `validate:"required,oneof=monolithic microservices serverless mvc mvp mern"`
	Components []Component `validate:"omitempty,dive"`
	DataFlow   string      `validate:"omitempty,max=500"`
}

type Component struct {
	Name        string `validate:"required,min=1,max=50"`
	Description string `validate:"required,min=5,max=200"`
	Technology  string `validate:"required,min=1,max=50"`
	Path        string `validate:"omitempty,max=100"`
}

type AppStructure struct {
	Frontend *FrontendStructure `validate:"omitempty"`
	Backend  *BackendStructure  `validate:"omitempty"`
	Database *DatabaseStructure `validate:"omitempty"`
	Overview string             `validate:"omitempty,max=500"`
}

type FrontendStructure struct {
	Framework  string   `validate:"required,oneof=React Angular Vue Svelte Next.js Nuxt.js"`
	Structure  []string `validate:"omitempty,dive,min=1,max=100"`
	EntryPoint string   `validate:"omitempty,max=100"`
}

type BackendStructure struct {
	Framework  string   `validate:"required,oneof=Express Django Flask Spring Flask FastAPI Gin Echo Fiber"`
	Structure  []string `validate:"omitempty,dive,min=1,max=100"`
	EntryPoint string   `validate:"omitempty,max=100"`
}

type DatabaseStructure struct {
	Type       string   `validate:"required,oneof=SQL NoSQL Graph InMemory"`
	Schema     []string `validate:"omitempty,dive,min=1,max=100"`
	Migrations string   `validate:"omitempty,max=200"`
}

type Development struct {
	DevServer   string `validate:"omitempty,max=200"`
	HotReload   bool
	DevDatabase string `validate:"omitempty,max=100"`
	TestData    string `validate:"omitempty,max=200"`
	Notes       string `validate:"omitempty,max=500"`
}
