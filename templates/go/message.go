package golang

const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if .MessageRules.GetSkip }}
		// skipping pipeline for {{ $f.Name }}
	{{ else }}
		if v, ok := interface{}({{ accessor . }}).(interface{ Pipeline() error }); ok {
			if err := v.Pipeline(); err != nil {
				return {{ errCause . "err" "embedded message failed pipeline" }}
			}
		}
	{{ end }}
`
