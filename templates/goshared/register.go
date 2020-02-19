package goshared

import (
	"fmt"
	"reflect"
	"strings"
	"text/template"

	"github.com/lyft/protoc-gen-star"
	"github.com/lyft/protoc-gen-star/lang/go"

	"github.com/wores/protoc-gen-pipeline/templates/shared"
)

func Register(tpl *template.Template, params pgs.Parameters) {
	fns := goSharedFuncs{pgsgo.InitContext(params)}
	fnsForString := goSharedStringFuncs{fns}

	tpl.Funcs(map[string]interface{}{
		"accessor":    fns.accessor,
		"property":    fns.property,
		"cmt":         pgs.C80,
		"err":         fns.err,
		"errCause":    fns.errCause,
		"errIdx":      fns.errIdx,
		"errIdxCause": fns.errIdxCause,
		"errname":     fns.errName,
		"lit":         fns.lit,
		"lookup":      fns.lookup,
		"msgTyp":      fns.msgTyp,
		"name":        fns.Name,
		"oneof":       fns.oneofTypeName,
		"pkg":         fns.PackageName,
		"typ":         fns.Type,
		"unwrap":      fns.unwrap,
	})

	funcsForString := map[string]interface{}{
		"trim":                     fnsForString.trim,
		"isOmissionPositionLeft":   fnsForString.isOmissionPositionLeft,
		"isOmissionPositionCenter": fnsForString.isOmissionPositionCenter,
		"isOmissionPositionRight":  fnsForString.isOmissionPositionRight,
	}

	template.Must(tpl.New("msg").Parse(msgTpl))

	template.Must(tpl.New("none").Parse(noneTpl))
	template.Must(tpl.New("string").Funcs(funcsForString).Parse(strTpl))

	template.Must(tpl.New("repeated").Parse(repTpl))

	template.Must(tpl.New("wrapper").Parse(wrapperTpl))
}

type goSharedFuncs struct{ pgsgo.Context }

func (fns goSharedFuncs) accessor(ctx shared.RuleContext) string {
	if ctx.AccessorOverride != "" {
		return ctx.AccessorOverride
	}

	return fmt.Sprintf("m.Get%s()", fns.Name(ctx.Field))
}

func (fns goSharedFuncs) property(ctx shared.RuleContext) string {
	if ctx.PropertyOverride != "" {
		return ctx.PropertyOverride
	}

	return fmt.Sprintf("m.%s", fns.Name(ctx.Field))
}

func (fns goSharedFuncs) errName(m pgs.Message) pgs.Name {
	return fns.Name(m) + "PipelineError"
}

func (fns goSharedFuncs) errIdxCause(ctx shared.RuleContext, idx, cause string, reason ...interface{}) string {
	f := ctx.Field
	n := fns.Name(f)

	var fld string
	if idx != "" {
		fld = fmt.Sprintf(`fmt.Sprintf("%s[%%v]", %s)`, n, idx)
	} else if ctx.Index != "" {
		fld = fmt.Sprintf(`fmt.Sprintf("%s[%%v]", %s)`, n, ctx.Index)
	} else {
		fld = fmt.Sprintf("%q", n)
	}

	causeFld := ""
	if cause != "nil" && cause != "" {
		causeFld = fmt.Sprintf("cause: %s,", cause)
	}

	keyFld := ""
	if ctx.OnKey {
		keyFld = "key: true,"
	}

	return fmt.Sprintf(`%s{
		field: %s,
		reason: %q,
		%s%s
	}`,
		fns.errName(f.Message()),
		fld,
		fmt.Sprint(reason...),
		causeFld,
		keyFld)
}

func (fns goSharedFuncs) err(ctx shared.RuleContext, reason ...interface{}) string {
	return fns.errIdxCause(ctx, "", "nil", reason...)
}

func (fns goSharedFuncs) errCause(ctx shared.RuleContext, cause string, reason ...interface{}) string {
	return fns.errIdxCause(ctx, "", cause, reason...)
}

func (fns goSharedFuncs) errIdx(ctx shared.RuleContext, idx string, reason ...interface{}) string {
	return fns.errIdxCause(ctx, idx, "nil", reason...)
}

func (fns goSharedFuncs) lookup(f pgs.Field, name string) string {
	return fmt.Sprintf(
		"_%s_%s_%s",
		fns.Name(f.Message()),
		fns.Name(f),
		name,
	)
}

func (fns goSharedFuncs) lit(x interface{}) string {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Interface {
		val = val.Elem()
	}

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.String:
		return fmt.Sprintf("%q", x)
	case reflect.Uint8:
		return fmt.Sprintf("0x%X", x)
	case reflect.Slice:
		els := make([]string, val.Len())
		for i, l := 0, val.Len(); i < l; i++ {
			els[i] = fns.lit(val.Index(i).Interface())
		}
		return fmt.Sprintf("%T{%s}", val.Interface(), strings.Join(els, ", "))
	default:
		return fmt.Sprint(x)
	}
}

func (fns goSharedFuncs) oneofTypeName(f pgs.Field) pgsgo.TypeName {
	return pgsgo.TypeName(fns.OneofOption(f)).Pointer()
}

func (fns goSharedFuncs) unwrap(ctx shared.RuleContext, name string) (shared.RuleContext, error) {
	ctx, err := ctx.Unwrap("wrapper")
	if err != nil {
		return ctx, err
	}

	overrideName := pgsgo.PGGUpperCamelCase(ctx.Field.Type().Embed().Fields()[0].Name())
	ctx.AccessorOverride = fmt.Sprintf("%s.Get%s()",
		name,
		overrideName,
	)
	ctx.PropertyOverride = fmt.Sprintf("%s.%s", name, overrideName)

	return ctx, nil
}

func (fns goSharedFuncs) msgTyp(message pgs.Message) pgsgo.TypeName {
	return pgsgo.TypeName(fns.Name(message))
}
