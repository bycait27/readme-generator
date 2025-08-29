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

{{if .Architecture}}
## 🏗️ Architecture

**Pattern:** {{.Architecture.Pattern}}

{{if .Architecture.DataFlow}}
**Data Flow:** {{.Architecture.DataFlow}}
{{end}}

{{if .Architecture.Components}}
### Components
{{range .Architecture.Components}}
- **{{.Name}}** ({{.Technology}}) - {{.Description}}
  {{if .Path}}Path: `{{.Path}}`{{end}}
{{end}}
{{end}}
{{end}}

## 🚀 Getting Started

{{if .GettingStarted.Prerequisites}}
### Prerequisites
{{range .GettingStarted.Prerequisites}}
- {{.}}
{{end}}
{{end}}

{{if .GettingStarted.EnvSetup}}
### Environment Setup
{{range .GettingStarted.EnvSetup}}
1. {{.}}
{{end}}
{{end}}

### Running the Application
{{range .GettingStarted.RunCommands}}
```bash
{{.}}
```
{{end}}

{{if .GettingStarted.Notes}}
**Note:** {{.GettingStarted.Notes}}
{{end}}

{{if .EnvVars}}
## ⚙️ Environment Variables

| Variable | Description | Required | Default | Example |
|----------|-------------|----------|---------|---------|
{{range .EnvVars}}| `{{.Name}}` | {{.Description}} | {{if .Required}}✅{{else}}❌{{end}} | {{.Default}} | {{.Example}} |
{{end}}
{{end}}

{{if .AppStructure}}
## 📁 Project Structure

{{if .AppStructure.Overview}}
{{.AppStructure.Overview}}
{{end}}

{{if .AppStructure.Frontend}}
### Frontend ({{.AppStructure.Frontend.Framework}})

{{if .AppStructure.Frontend.EntryPoint}}
**Entry Point:** `{{.AppStructure.Frontend.EntryPoint}}`
{{end}}

{{if .AppStructure.Frontend.Structure}}
```
{{range .AppStructure.Frontend.Structure}}{{.}}
{{end}}
```
{{end}}
{{end}}

{{if .AppStructure.Backend}}
### Backend ({{.AppStructure.Backend.Framework}})

{{if .AppStructure.Backend.EntryPoint}}
**Entry Point:** `{{.AppStructure.Backend.EntryPoint}}`
{{end}}

{{if .AppStructure.Backend.Structure}}
```
{{range .AppStructure.Backend.Structure}}{{.}}
{{end}}
```
{{end}}
{{end}}

{{if .AppStructure.Database}}
### Database ({{.AppStructure.Database.Type}})

{{if .AppStructure.Database.Migrations}}
**Migrations:** {{.AppStructure.Database.Migrations}}
{{end}}

{{if .AppStructure.Database.Schema}}
**Schema:**
{{range .AppStructure.Database.Schema}}
- {{.}}
{{end}}
{{end}}
{{end}}
{{end}}

{{if .Development}}
## 🛠️ Development

{{if .Development.DevServer}}
**Development Server:**
```bash
{{.Development.DevServer}}
```
{{end}}

{{if .Development.HotReload}}
**Hot Reload:** Enabled ✅
{{else}}
**Hot Reload:** Disabled ❌
{{end}}

{{if .Development.DevDatabase}}
**Development Database:** {{.Development.DevDatabase}}
{{end}}

{{if .Development.TestData}}
**Test Data:** {{.Development.TestData}}
{{end}}

{{if .Development.Notes}}
**Development Notes:** {{.Development.Notes}}
{{end}}
{{end}}

{{if .Testing}}
## 🧪 Testing

**Framework:** {{.Testing.TestFramework}}

**Run Tests:**
```bash
{{.Testing.TestCommand}}
```

{{if .Testing.CoverageCmd}}
**Coverage:**
```bash
{{.Testing.CoverageCmd}}
```
{{end}}

{{if .Testing.Notes}}
**Notes:** {{.Testing.Notes}}
{{end}}
{{end}}

{{if .Deployment}}
## 🚀 Deployment

**Platform:** {{.Deployment.Platform}}

{{if .Deployment.BuildCmd}}
**Build:**
```bash
{{.Deployment.BuildCmd}}
```
{{end}}

{{if .Deployment.DeployCmd}}
**Deploy:**
```bash
{{.Deployment.DeployCmd}}
```
{{end}}

{{if .Deployment.HealthCheck}}
**Health Check:**
```bash
{{.Deployment.HealthCheck}}
```
{{end}}

{{if .Deployment.Notes}}
**Notes:** {{.Deployment.Notes}}
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
