# {{.Title}}

{{.Description}}

{{if .Screenshots}}

## 📸 Demo 

{{if .Screenshots.Demo}}
![Demo]({{.Screenshots.Demo}})
{{end}}

{{.Screenshots.Description}}

{{if .Screenshots.Images}}
### Additional Screenshots
{{range .Screenshots.Images}}
![Screenshot]({{.}})
{{end}}
{{end}}
{{end}}

{{if .TechStack}}
## 🛠️ Tech Stack

**Language:** {{.TechStack.Language}}
{{if .TechStack.Framework}}

**Framework:** {{.TechStack.Framework}}{{end}}
{{if .TechStack.Database}}

**Database:** {{.TechStack.Database}}{{end}}
{{if .TechStack.Dependencies}}

**Dependencies:**
{{range .TechStack.Dependencies}}  - {{.}}
{{end}}
{{end}}
{{end}}

## ⚡ Quick Start

{{.QuickStart.Description}}

{{if .QuickStart.Commands}}
{{range .QuickStart.Commands}}
```bash
{{.}}
```
{{end}}
{{end}}

## 📦 Installation

{{if .Installation.PackageManagers}}
### Package Managers
{{range .Installation.PackageManagers}}
**{{.Name}}:**
```bash
{{.Command}}
```
{{end}}
{{end}}

{{if .Installation.Binary}}
### Binary Download

{{if .Installation.Binary.Platforms}}
**Supported Platforms:** {{range $i, $platform := .Installation.Binary.Platforms}}{{if $i}}, {{end}}{{$platform}}{{end}}
{{end}}

**Download:** [{{.Installation.Binary.URL}}]({{.Installation.Binary.URL}})

{{.Installation.Binary.Instructions}}
{{end}}

{{if .Installation.FromSource}}
### From Source

**Repository:** {{.Installation.FromSource.RepoURL}}

{{if .Installation.FromSource.Requirements}}
**Requirements:**
{{range .Installation.FromSource.Requirements}}
- {{.}}
{{end}}
{{end}}

**Build:**
```bash
{{.Installation.FromSource.BuildCmd}}
```
{{end}}

## 🚀 Usage

### Basic Usage

```bash
{{.Usage.BasicUsage}}
```

{{.Usage.Description}}

{{if .Usage.CommonFlags}}
### Common Flags
{{range .Usage.CommonFlags}}
- {{.}}
{{end}}
{{end}}

## 📋 Commands

{{range .Commands}}
### `{{.Name}}`

{{.Description}}

{{if .Flags}}
**Flags:**
{{range .Flags}}
- `{{if .Short}}-{{.Short}}, {{end}}--{{.Name}}`{{if .Required}} *(required)*{{end}} - {{.Description}}
{{if .Default}}  Default: `{{.Default}}`{{end}}
{{end}}
{{end}}

---
{{end}}

{{if .Examples}}
## 💡 Examples

{{range .Examples}}
### {{.Title}}

{{.Description}}

{{range .Commands}}
```bash
{{.}}
```
{{end}}

{{if .Output}}
**Output:**
```
{{.Output}}
```
{{end}}

---
{{end}}
{{end}}

{{if .Configuration}}
## ⚙️ Configuration

{{if .Configuration.ConfigFile}}
**Config File:** `{{.Configuration.ConfigFile}}`
{{end}}

{{if .Configuration.EnvVars}}
### Environment Variables

| Variable | Description | Required | Default | Example |
|----------|-------------|----------|---------|---------|
{{range .Configuration.EnvVars}}| `{{.Name}}` | {{.Description}} | {{if .Required}}✅{{else}}❌{{end}} | {{.Default}} | {{.Example}} |
{{end}}
{{end}}

{{if .Configuration.Examples}}
### Configuration Examples

{{range .Configuration.Examples}}
**{{.Format}}:**
```{{.Format}}
{{.Content}}
```
{{end}}
{{end}}
{{end}}

{{if .Troubleshooting}}
## 🔧 Troubleshooting

{{if .Troubleshooting.CommonIssues}}
### Common Issues

{{range .Troubleshooting.CommonIssues}}
**Problem:** {{.Problem}}

**Solution:** {{.Solution}}

---
{{end}}
{{end}}

{{if .Troubleshooting.FAQs}}
### Frequently Asked Questions

{{range .Troubleshooting.FAQs}}
**Q: {{.Question}}**

A: {{.Answer}}

---
{{end}}
{{end}}
{{end}}

## 📄 License

This project is licensed under the {{.License}} License.

## Contact

**{{.Author.Name}}**

**📧 Email:** {{.Author.Email}}

**🐙 GitHub:** [{{.Author.Name}}]({{.Author.GitHub}})

{{if .Author.Website}}  
**🌐 Website:** [{{.Author.Website}}]({{.Author.Website}}){{end}}
