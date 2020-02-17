package goshared

const msgTpl = `
func (m {{ (msgTyp .).Pointer }}) Pipeline() error {
	if m == nil { return nil }

	{{ range .NonOneOfFields }}
		{{ render (context .) }}
	{{ end }}

	{{ range .OneOfs }}
		switch m.{{ name . }}.(type) {
			{{ range .Fields }}
				case {{ oneof . }}:
					{{ render (context .) }}
			{{ end }}
		}
	{{ end }}

	return nil
}

{{ cmt (errname .) " is the pipeline error returned by " (msgTyp .) ".Pipeline if the designated constraints aren't met." -}}
type {{ errname . }} struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e {{ errname . }}) Field() string { return e.field }

// Reason function returns reason value.
func (e {{ errname . }}) Reason() string { return e.reason }

// Cause function returns cause value.
func (e {{ errname . }}) Cause() error { return e.cause }

// Key function returns key value.
func (e {{ errname . }}) Key() bool { return e.key }

// ErrorName returns error name.
func (e {{ errname . }}) ErrorName() string { return "{{ errname . }}" }

// Error satisfies the builtin error interface
func (e {{ errname . }}) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %s{{ (msgTyp .) }}.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = {{ errname . }}{}

var _ interface{
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = {{ errname . }}{}

{{ range .Fields }}{{ with (context .) }}{{ $f := .Field }}
	{{ if has .Rules "Pattern"}}{{ if .Rules.Pattern }}
		var {{ lookup .Field "Pattern" }} = regexp.MustCompile({{ lit .Rules.GetPattern }})
	{{ end }}{{ end }}

	{{ if has .Rules "Items"}}{{ if .Rules.Items }}
	{{ if has .Rules.Items.GetString_ "Pattern" }} {{ if .Rules.Items.GetString_.Pattern }}
		var {{ lookup .Field "Pattern" }} = regexp.MustCompile({{ lit .Rules.Items.GetString_.GetPattern }})
	{{ end }}{{ end }}
	{{ end }}{{ end }}

{{ end }}{{ end }}
`
