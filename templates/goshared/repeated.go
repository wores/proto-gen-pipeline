package goshared

const repTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	{{ if ne (.Elem "" "" "").Typ "none" }}
		for idx, item := range {{ accessor . }} {
			_, _ = idx, item

			{{$p := property . }}
			{{ render (.Elem "item" "idx" $p) }}
		}
	{{ end }}
`
