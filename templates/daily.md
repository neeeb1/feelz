# {{ .Date }}

## Daily Notes

{{ range .Prompts}}
### {{ .Name }}

#### {{ .Prompt }}

{{ .FinalText }}

{{end}}