# {{.Title}}

{{.Description}}

{{if .Screenshots}}

## Demo 

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

- **Language:** {{.TechStack.Language}}
{{if .TechStack.Framework}}
- **Framework:** {{.TechStack.Framework}}{{end}}
{{if .TechStack.Database}}
- **Database:** {{.TechStack.Database}}{{end}}
{{if .TechStack.Dependencies}}
- **Dependencies:**
{{range .TechStack.Dependencies}}  - {{.}}
{{end}}
{{end}}
{{end}}

## 📄 License

This project is licensed under the {{.License}} License.

## 👤 Author

**{{.Author.Name}}**
- 📧 Email: {{.Author.Email}}
- 🐙 GitHub: [{{.Author.Name}}]({{.Author.GitHub}})
{{if .Author.Website}}
- 🌐 Website: [{{.Author.Website}}]({{.Author.Website}}){{end}}
