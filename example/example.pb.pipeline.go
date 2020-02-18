// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: example/example.proto

package example

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = unicode.Hiragana
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = time.Duration(0)
	_ = ptypes.DynamicAny{}
)

func (m *StringAllExample) Pipeline() error {
	if m == nil {
		return nil
	}

	m.Text = strings.TrimFunc(m.GetText(), func(r rune) bool { return unicode.IsSpace(r) })

	m.Text = strings.ReplaceAll(m.GetText(), "q", "")

	m.Text = strings.ReplaceAll(m.GetText(), "'", "\"")

	return nil
}

// StringAllExamplePipelineError is the pipeline error returned by
// StringAllExample.Pipeline if the designated constraints aren't met.
type StringAllExamplePipelineError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringAllExamplePipelineError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringAllExamplePipelineError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringAllExamplePipelineError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringAllExamplePipelineError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringAllExamplePipelineError) ErrorName() string { return "StringAllExamplePipelineError" }

// Error satisfies the builtin error interface
func (e StringAllExamplePipelineError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringAllExample.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringAllExamplePipelineError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringAllExamplePipelineError{}

func (m *StringTrimExample) Pipeline() error {
	if m == nil {
		return nil
	}

	m.Text = strings.TrimLeftFunc(m.GetText(), func(r rune) bool { return unicode.IsSpace(r) })

	if wrapper := m.GetWrapText(); wrapper != nil {

		wrapper.Value = strings.TrimFunc(wrapper.GetValue(), func(r rune) bool { return unicode.IsSpace(r) })

	}

	for idx, item := range m.GetTexts() {
		_, _ = idx, item

		m.Texts[idx] = strings.TrimFunc(item, func(r rune) bool { return unicode.IsSpace(r) })

	}

	{
		tmp := m.GetInner()

		if v, ok := interface{}(tmp).(interface{ Pipeline() error }); ok {

			if err := v.Pipeline(); err != nil {
				return StringTrimExamplePipelineError{
					field:  "Inner",
					reason: "embedded message failed pipeline",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// StringTrimExamplePipelineError is the pipeline error returned by
// StringTrimExample.Pipeline if the designated constraints aren't met.
type StringTrimExamplePipelineError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringTrimExamplePipelineError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringTrimExamplePipelineError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringTrimExamplePipelineError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringTrimExamplePipelineError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringTrimExamplePipelineError) ErrorName() string { return "StringTrimExamplePipelineError" }

// Error satisfies the builtin error interface
func (e StringTrimExamplePipelineError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringTrimExample.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringTrimExamplePipelineError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringTrimExamplePipelineError{}

func (m *StringRemoveExample) Pipeline() error {
	if m == nil {
		return nil
	}

	m.Text = strings.ReplaceAll(m.GetText(), "-", "")

	return nil
}

// StringRemoveExamplePipelineError is the pipeline error returned by
// StringRemoveExample.Pipeline if the designated constraints aren't met.
type StringRemoveExamplePipelineError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringRemoveExamplePipelineError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringRemoveExamplePipelineError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringRemoveExamplePipelineError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringRemoveExamplePipelineError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringRemoveExamplePipelineError) ErrorName() string {
	return "StringRemoveExamplePipelineError"
}

// Error satisfies the builtin error interface
func (e StringRemoveExamplePipelineError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringRemoveExample.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringRemoveExamplePipelineError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringRemoveExamplePipelineError{}

func (m *StringReplaceExample) Pipeline() error {
	if m == nil {
		return nil
	}

	m.Text = strings.ReplaceAll(m.GetText(), "*", "%")

	return nil
}

// StringReplaceExamplePipelineError is the pipeline error returned by
// StringReplaceExample.Pipeline if the designated constraints aren't met.
type StringReplaceExamplePipelineError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringReplaceExamplePipelineError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringReplaceExamplePipelineError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringReplaceExamplePipelineError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringReplaceExamplePipelineError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringReplaceExamplePipelineError) ErrorName() string {
	return "StringReplaceExamplePipelineError"
}

// Error satisfies the builtin error interface
func (e StringReplaceExamplePipelineError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringReplaceExample.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringReplaceExamplePipelineError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringReplaceExamplePipelineError{}

func (m *StringTrimExample_Inner) Pipeline() error {
	if m == nil {
		return nil
	}

	m.Text = strings.TrimRightFunc(m.GetText(), func(r rune) bool { return unicode.IsSpace(r) })

	if wrapper := m.GetWrapText(); wrapper != nil {

		wrapper.Value = strings.TrimFunc(wrapper.GetValue(), func(r rune) bool { return unicode.IsSpace(r) })

	}

	for idx, item := range m.GetTexts() {
		_, _ = idx, item

		m.Texts[idx] = strings.TrimFunc(item, func(r rune) bool { return unicode.IsSpace(r) })

	}

	{
		tmp := m.GetInner()

		if v, ok := interface{}(tmp).(interface{ Pipeline() error }); ok {

			if err := v.Pipeline(); err != nil {
				return StringTrimExample_InnerPipelineError{
					field:  "Inner",
					reason: "embedded message failed pipeline",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// StringTrimExample_InnerPipelineError is the pipeline error returned by
// StringTrimExample_Inner.Pipeline if the designated constraints aren't met.
type StringTrimExample_InnerPipelineError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StringTrimExample_InnerPipelineError) Field() string { return e.field }

// Reason function returns reason value.
func (e StringTrimExample_InnerPipelineError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StringTrimExample_InnerPipelineError) Cause() error { return e.cause }

// Key function returns key value.
func (e StringTrimExample_InnerPipelineError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StringTrimExample_InnerPipelineError) ErrorName() string {
	return "StringTrimExample_InnerPipelineError"
}

// Error satisfies the builtin error interface
func (e StringTrimExample_InnerPipelineError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStringTrimExample_Inner.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StringTrimExample_InnerPipelineError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StringTrimExample_InnerPipelineError{}
