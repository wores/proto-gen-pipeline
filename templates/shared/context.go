package shared

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/golang/protobuf/proto"
	pgs "github.com/lyft/protoc-gen-star"

	"github.com/wores/protoc-gen-pipeline/gogoproto"
	"github.com/wores/protoc-gen-pipeline/pipeline"
)

type RuleContext struct {
	Field        pgs.Field
	Rules        proto.Message
	MessageRules *pipeline.MessageProcesses
	Gogo         Gogo

	Typ        string
	WrapperTyp string

	OnKey            bool
	Index            string
	AccessorOverride string
	PropertyOverride string
}

type Gogo struct {
	Nullable    bool
	//StdDuration bool
	//StdTime     bool
}

func rulesContext(f pgs.Field) (out RuleContext, err error) {
	out.Field = f

	out.Gogo.Nullable = true
	_, _ = f.Extension(gogoproto.E_Nullable, &out.Gogo.Nullable)
	//f.Extension(gogoproto.E_Stdduration, &out.Gogo.StdDuration)
	//f.Extension(gogoproto.E_Stdtime, &out.Gogo.StdTime)

	var rules pipeline.FieldProcesses
	if _, err = f.Extension(pipeline.E_Processes, &rules); err != nil {
		return
	}

	var wrapped bool

	if out.Typ, out.Rules, out.MessageRules, wrapped = resolveRules(f.Type(), &rules); wrapped {
		out.WrapperTyp = out.Typ
		out.Typ = "wrapper"
	}

	if out.Typ == "error" {
		err = fmt.Errorf("unknown rule type (%T)", rules.Type)
	}

	return
}

func (ctx RuleContext) Elem(name, idx, property string) (out RuleContext, err error) {
	out.Field = ctx.Field
	out.AccessorOverride = name
	out.Index = idx
	out.PropertyOverride = fmt.Sprintf("%s[%s]", property, idx)
	out.Gogo = ctx.Gogo

	var rules *pipeline.FieldProcesses
	switch r := ctx.Rules.(type) {
	case *pipeline.RepeatedProcesses:
		rules = r.GetItems()
	default:
		err = fmt.Errorf("cannot get Elem RuleContext from %T", ctx.Field)
		return
	}

	var wrapped bool
	if out.Typ, out.Rules, out.MessageRules, wrapped = resolveRules(ctx.Field.Type().Element(), rules); wrapped {
		out.WrapperTyp = out.Typ
		out.Typ = "wrapper"
	}

	if out.Typ == "error" {
		err = fmt.Errorf("unknown rule type (%T)", rules)
	}

	return
}

func (ctx RuleContext) Unwrap(name string) (out RuleContext, err error) {
	if ctx.Typ != "wrapper" {
		err = fmt.Errorf("cannot unwrap non-wrapper type %q", ctx.Typ)
		return
	}

	return RuleContext{
		Field:            ctx.Field,
		Rules:            ctx.Rules,
		MessageRules:     ctx.MessageRules,
		Typ:              ctx.WrapperTyp,
		AccessorOverride: name,
		PropertyOverride: name,
		Gogo:             ctx.Gogo,
	}, nil
}

func Render(tpl *template.Template) func(ctx RuleContext) (string, error) {
	return func(ctx RuleContext) (string, error) {
		var b bytes.Buffer
		err := tpl.ExecuteTemplate(&b, ctx.Typ, ctx)
		return b.String(), err
	}
}

func resolveRules(typ interface{ IsEmbed() bool }, rules *pipeline.FieldProcesses) (
	ruleType string,
	rule proto.Message,
	messageRule *pipeline.MessageProcesses,
	wrapped bool,
) {
	switch r := rules.GetType().(type) {
	case *pipeline.FieldProcesses_String_:
		ruleType, rule, wrapped = "string", r.String_, typ.IsEmbed()
	case *pipeline.FieldProcesses_Repeated:
		ruleType, rule, wrapped = "repeated", r.Repeated, false
	case nil:
		if ft, ok := typ.(pgs.FieldType); ok && ft.IsRepeated() {
			return "repeated", &pipeline.RepeatedProcesses{}, rules.GetMessage(), false
		} else if typ.IsEmbed() {
			return "message", rules.GetMessage(), rules.GetMessage(), false
		}
		return "none", nil, nil, false
	default:
		ruleType, rule, wrapped = "error", nil, false
	}

	return ruleType, rule, rules.GetMessage(), wrapped
}
