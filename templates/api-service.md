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

### Running the Service
{{range .GettingStarted.RunCommands}}
```bash
{{.}}
```
{{end}}

{{if .GettingStarted.Notes}}
**Note:** {{.GettingStarted.Notes}}
{{end}}

{{if .Database}}
## 🗄️ Database

**Type:** {{.Database.Type}}

{{if .Database.Schema}}
### Schema
```sql
{{.Database.Schema}}
```
{{end}}

{{if .Database.Migrations}}
### Migrations
{{.Database.Migrations}}
{{end}}

{{if .Database.SeedData}}
### Seed Data
{{.Database.SeedData}}
{{end}}
{{end}}

{{if .EnvVars}}
## ⚙️ Environment Variables

| Variable | Description | Required | Default | Example |
|----------|-------------|----------|---------|---------|
{{range .EnvVars}}| `{{.Name}}` | {{.Description}} | {{if .Required}}✅{{else}}❌{{end}} | `{{.Default}}` | `{{.Example}}` |
{{end}}
{{end}}

{{if .Auth}}
## 🔐 Authentication

**Method:** {{.Auth.Method}}
{{if .Auth.TokenFormat}}
**Token Format:** {{.Auth.TokenFormat}}
{{end}}

{{if .Auth.ExampleUsage}}
### Usage Example
```bash
{{.Auth.ExampleUsage}}
```
{{end}}

{{if .Auth.Endpoints}}
### Auth Endpoints
{{range .Auth.Endpoints}}
- {{.}}
{{end}}
{{end}}
{{end}}

## 📚 API Documentation

**Base URL:** {{.APIDocs.BaseURL}}

**Authentication:** {{.APIDocs.Authentication}}

### Endpoints

{{range .APIDocs.Endpoints}}
#### {{.Method}} {{.Path}}

{{.Description}}

**Parameters:**
{{range .Parameters}}
- `{{.Name}}` ({{.Type}}) {{if .Required}}*required*{{else}}*optional*{{end}} - {{.Description}}
  {{if .Example}}Example: `{{.Example}}`{{end}}
{{end}}

**Response:**
```json
{{.Response}}
```

---
{{end}}

{{if .APIDocs.ErrorHandling}}
### Error Handling

**Format:** {{.APIDocs.ErrorHandling.Format}}

{{if .APIDocs.ErrorHandling.ErrorResponse.Structure}}
**Error Response Structure:**
```json
{{.APIDocs.ErrorHandling.ErrorResponse.Structure}}
```
{{end}}

{{if .APIDocs.ErrorHandling.StatusCodes}}
**Status Codes:**
{{range .APIDocs.ErrorHandling.StatusCodes}}
- `{{.Code}}` - {{.Description}}
  {{if .Example}}Example: `{{.Example}}`{{end}}
{{end}}
{{end}}

{{if .APIDocs.ErrorHandling.CommonErrors}}
**Common Errors:**
{{range .APIDocs.ErrorHandling.CommonErrors}}
- **{{.Code}}** - {{.Message}}
  - Description: {{.Description}}
  - Solution: {{.Solution}}
{{end}}
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

{{if .Monitoring}}
## 📊 Monitoring & Health

{{if .Monitoring.HealthCheck}}
**Health Check:**
```bash
{{.Monitoring.HealthCheck}}
```
{{end}}

{{if .Monitoring.Metrics}}
**Metrics:**
{{range .Monitoring.Metrics}}
- {{.}}
{{end}}
{{end}}

{{if .Monitoring.Logging}}
**Logging:**
- Level: {{.Monitoring.Logging.Level}}
- Format: {{.Monitoring.Logging.Format}}
- Output: {{.Monitoring.Logging.Output}}
{{end}}

{{if .Monitoring.Alerts}}
**Alerts:**
{{range .Monitoring.Alerts}}
- {{.}}
{{end}}
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
