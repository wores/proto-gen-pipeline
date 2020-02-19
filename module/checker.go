package module

import (
	"unicode/utf8"

	pgs "github.com/lyft/protoc-gen-star"

	"github.com/wores/protoc-gen-pipeline/pipeline"
)

type FieldType interface {
	ProtoType() pgs.ProtoType
	Embed() pgs.Message
}

type Repeatable interface {
	IsRepeated() bool
}

func (m *Module) CheckRules(msg pgs.Message) {
	m.Push("msg: " + msg.Name().String())
	defer m.Pop()

	//var disabled bool
	//_, err := msg.Extension(validate.E_Disabled, &disabled)
	//m.CheckErr(err, "unable to read validation extension from message")

	//if disabled {
	//	m.Debug("validation disabled, skipping checks")
	//	return
	//}

	for _, f := range msg.Fields() {
		m.Push(f.Name().String())

		var rules pipeline.FieldProcesses
		_, err := f.Extension(pipeline.E_Processes, &rules)
		m.CheckErr(err, "unable to read validation rules from field")

		if rules.GetMessage() != nil {
			m.MustType(f.Type(), pgs.MessageT, pgs.UnknownWKT)
			m.CheckMessage(f, &rules)
		}

		m.CheckFieldRules(f.Type(), &rules)

		m.Pop()
	}
}

func (m *Module) CheckFieldRules(typ FieldType, rules *pipeline.FieldProcesses) {
	if rules == nil {
		return
	}

	switch r := rules.Type.(type) {
	case *pipeline.FieldProcesses_String_:
		m.MustType(typ, pgs.StringT, pgs.StringValueWKT)
		m.CheckString(r.String_)
	case *pipeline.FieldProcesses_Repeated:
		m.CheckRepeated(typ, r.Repeated)
	//case *validate.FieldRules_Map:
	//	m.CheckMap(typ, r.Map)
	case nil: // noop
	default:
		m.Failf("unknown rule type (%T)", rules.Type)
	}
}

func (m *Module) MustType(typ FieldType, pt pgs.ProtoType, wrapper pgs.WellKnownType) {
	if emb := typ.Embed(); emb != nil && emb.IsWellKnown() && emb.WellKnownType() == wrapper {
		m.MustType(emb.Fields()[0].Type(), pt, pgs.UnknownWKT)
		return
	}

	if typ, ok := typ.(Repeatable); ok {
		m.Assert(!typ.IsRepeated(),
			"repeated rule should be used for repeated fields")
	}

	m.Assert(typ.ProtoType() == pt,
		" expected rules for ",
		typ.ProtoType().Proto(),
		" but got ",
		pt.Proto(),
	)
}

func (m *Module) CheckString(r *pipeline.StringProcesses) {
	if r.Omission != nil {
		m.Assert(*r.Omission.Len > uint64(utf8.RuneCountInString(*r.Omission.Replace) + 3), "`replace` length exceeds the `len`")
	}
	//if r.MaxLen != nil {
	//	max := int(r.GetMaxLen())
	//	m.Assert(utf8.RuneCountInString(r.GetPrefix()) <= max, "`prefix` length exceeds the `max_len`")
	//	m.Assert(utf8.RuneCountInString(r.GetSuffix()) <= max, "`suffix` length exceeds the `max_len`")
	//	m.Assert(utf8.RuneCountInString(r.GetContains()) <= max, "`contains` length exceeds the `max_len`")
	//
	//	m.Assert(
	//		r.MaxBytes == nil || r.GetMaxBytes() >= r.GetMaxLen(),
	//		"`max_len` cannot exceed `max_bytes`")
	//}
	//
	//if r.MaxBytes != nil {
	//	max := int(r.GetMaxBytes())
	//	m.Assert(len(r.GetPrefix()) <= max, "`prefix` length exceeds the `max_bytes`")
	//	m.Assert(len(r.GetSuffix()) <= max, "`suffix` length exceeds the `max_bytes`")
	//	m.Assert(len(r.GetContains()) <= max, "`contains` length exceeds the `max_bytes`")
	//}
}

func (m *Module) CheckMessage(f pgs.Field, rules *pipeline.FieldProcesses) {
	m.Assert(f.Type().IsEmbed(), "field is not embedded but got message rules")
	emb := f.Type().Embed()
	if emb != nil && emb.IsWellKnown() {
		switch emb.WellKnownType() {
		case pgs.AnyWKT:
			m.Failf("Any rules should be used for Any fields")
		case pgs.DurationWKT:
			m.Failf("Duration rules should be used for Duration fields")
		case pgs.TimestampWKT:
			m.Failf("Timestamp rules should be used for Timestamp fields")
		}
	}

	if rules.Type != nil && rules.GetMessage().GetSkip() {
		m.Failf("Skip should not be used with WKT scalar rules")
	}
}

func (m *Module) CheckRepeated(ft FieldType, r *pipeline.RepeatedProcesses) {
	typ := m.mustFieldType(ft)

	m.Assert(typ.IsRepeated(), "field is not repeated but got repeated rules")

	//if r.GetUnique() {
	//	m.Assert(
	//		!typ.Element().IsEmbed(),
	//		"unique rule is only applicable for scalar types")
	//}

	m.Push("items")
	m.CheckFieldRules(typ.Element(), r.Items)
	m.Pop()
}

func (m *Module) mustFieldType(ft FieldType) pgs.FieldType {
	typ, ok := ft.(pgs.FieldType)
	if !ok {
		m.Failf("unexpected field type (%T)", ft)
	}

	return typ
}
