package goshared

const strTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if $r.ForceHoge }}
	{{ property . }} = "force hoge" 
    {{ end }}

	{{ if $r.Trim }}
	{{ trim . $r.Trim }}
    {{ end }}
`

//  {{ if $r.ForceHoge }}
//  if !{{ lookup $f "ForceHoge" }}.MatchString({{ accessor . }}) {
//  	return {{ err . "value does not match regex pattern " (lit $r.GetPattern) }}
//  }

