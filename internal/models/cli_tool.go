package models

type CLITool struct {
	BaseInfo
	QuickStart      QuickStart       `validate:"required"`
	Installation    Installation     `validate:"required"`
	Usage           Usage            `validate:"required"`
	Commands        []Command        `validate:"min=1,dive"`
	Examples        []Example        `validate:"min=1,dive"`
	Configuration   *Configuration   `validate:"omitempty"` // optional
	Troubleshooting *Troubleshooting `validate:"omitempty"` // optional
}

type QuickStart struct {
	Commands    []string `validate:"min=1,dive,min=1"`
	Description string   `validate:"required,min=10,max=200"`
}

type Installation struct {
	PackageManagers []PackageManager `validate:"omitempty,dive"`
	Binary          *BinaryInstall   `validate:"omitempty"`
	FromSource      *SourceInstall   `validate:"omitempty"`
}

type PackageManager struct {
	Name    string `validate:"required,min=1,max=50"`  // "homebrew", "go", "npm"
	Command string `validate:"required,min=1,max=100"` // e.g., "brew install", "go get", "npm install"
}

type BinaryInstall struct {
	URL          string   `validate:"required,url"`
	Platforms    []string `validate:"required,min=1,dive,oneof=Windows MacOS Linux"`
	Instructions string   `validate:"omitempty,max=500"`
}

type SourceInstall struct {
	RepoURL      string   `validate:"required,url"`
	BuildCmd     string   `validate:"required,min=1,max=200"`
	Requirements []string `validate:"omitempty,dive,min=1,max=100"`
}

type Usage struct {
	BasicUsage  string   `validate:"required,min=10,max=200"`
	Description string   `validate:"required,min=10,max=500"`
	CommonFlags []string `validate:"omitempty,dive,min=1,max=30"`
}

type Command struct {
	Name        string `validate:"required,min=1,max=50"`
	Description string `validate:"required,min=5,max=200"`
	Flags       []Flag `validate:"omitempty,dive"`
}

type Flag struct {
	Name        string `validate:"required,min=1,max=30"`
	Short       string `validate:"omitempty,len=1,alpha"`
	Description string `validate:"required,min=5,max=100"`
	Default     string `validate:"omitempty,max=50"`
	Required    bool
}

type Example struct {
	Title       string   `validate:"required,min=5,max=100"`
	Description string   `validate:"required,min=10,max=300"`
	Commands    []string `validate:"required,min=1,dive,min=1"`
	Output      string   `validate:"omitempty,max=500"`
}

type Configuration struct {
	ConfigFile string          `validate:"required,min=3,max=50"` // e.g., "config.yaml"
	EnvVars    []EnvVar        `validate:"omitempty,dive"`
	Examples   []ConfigExample `validate:"omitempty,dive"`
}

type ConfigExample struct {
	Format  string `validate:"required,oneof=yaml json toml"`
	Content string `validate:"required,min=10,max=1000"`
}

type Troubleshooting struct {
	CommonIssues []Issue `validate:"omitempty,dive"`
	FAQs         []FAQ   `validate:"omitempty,dive"`
}

type Issue struct {
	Problem  string `validate:"required,min=5,max=200"`
	Solution string `validate:"required,min=5,max=500"`
}

type FAQ struct {
	Question string `validate:"required,min=5,max=200"`
	Answer   string `validate:"required,min=5,max=500"`
}
