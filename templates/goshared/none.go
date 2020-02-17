package goshared

const noneTpl = `// no pipeline processes for {{ name .Field }}
	{{- if .Index }}[{{ .Index }}]{{ end }}
	{{- if .OnKey }} (key){{ end }}`

