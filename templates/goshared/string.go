package goshared

const strTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if $r.Trim }}
	{{ trim . $r.Trim }}
    {{ end }}

	{{ if $r.Remove }}
	{{ property . }} = strings.ReplaceAll({{ accessor . }}, "{{ $r.Remove }}", "")
    {{ end }}

	{{ if $r.Replace }}
	{{ property . }} = strings.ReplaceAll({{ accessor . }}, "{{ $r.Replace.Old }}", "{{ $r.Replace.New }}")
    {{ end }}
    
`

//  {{ if $r.ForceHoge }}
//  if !{{ lookup $f "ForceHoge" }}.MatchString({{ accessor . }}) {
//  	return {{ err . "value does not match regex pattern " (lit $r.GetPattern) }}
//  }

