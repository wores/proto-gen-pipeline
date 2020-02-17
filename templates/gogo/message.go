package gogo

const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if .MessageRules.GetSkip }}
		// skipping pipeline for {{ $f.Name }}
	{{ else }}
	{
		tmp := {{ accessor . }}
		{{ if .Gogo.Nullable }}
		if v, ok := interface{}(tmp).(interface{ Pipeline() error }); ok {
		{{ else }}
		if v, ok := interface{}(&tmp).(interface{ Pipeline() error }); ok {
		{{ end }}
			if err := v.Pipeline(); err != nil {
				return {{ errCause . "err" "embedded message failed pipeline" }}
			}
		}
	}
	{{ end }}
`
