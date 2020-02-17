package goshared

const noneTpl = `// no pipeline tasks for {{ name .Field }}
	{{- if .Index }}[{{ .Index }}]{{ end }}
	{{- if .OnKey }} (key){{ end }}`

