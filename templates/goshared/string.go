package goshared

const strTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if $r.Trim }}
	{{ trim . $r.Trim }}
    {{ end }}

	{{ if $r.RemoveAll }}
	{{ property . }} = strings.ReplaceAll({{ accessor . }}, "{{ $r.RemoveAll }}", "")
    {{ end }}

	{{ if $r.Replace }}
	{{ property . }} = strings.ReplaceAll({{ accessor . }}, "{{ $r.Replace.Old }}", "{{ $r.Replace.New }}")
    {{ end }}
    
	{{ if $r.Omission }}
    if utf8.RuneCountInString({{ accessor . }}) > {{ $r.Omission.Len }} {
    	runes := []rune({{ accessor . }})
		replace := "{{ $r.Omission.Replace }}"
		lenExcludeReplace := {{ $r.Omission.Len }} - utf8.RuneCountInString(replace)

    	{{ if isOmissionPositionLeft $r.Omission.Position }}
		l := len(runes) - lenExcludeReplace
		{{ property . }} = replace + string(runes[l:])
    	{{- end -}}

    	{{ if isOmissionPositionCenter $r.Omission.Position }}
		lenLeft := lenExcludeReplace / 2
		lenRight := lenLeft
		if lenExcludeReplace % 2 != 0 {
			lenLeft += 1
		}
		rightStart := len(runes) - lenRight
		{{ property . }} = fmt.Sprintf("%s%s%s", string(runes[0:lenLeft]), replace, string(runes[rightStart:]))
    	{{- end -}}

    	{{ if isOmissionPositionRight $r.Omission.Position }}
		{{ property . }} = string(runes[:lenExcludeReplace]) + replace
    	{{- end -}}
	}
    {{ end }}
`

