package templates

import (
	"text/template"

	"github.com/lyft/protoc-gen-star"
	"github.com/lyft/protoc-gen-star/lang/go"

	golang "github.com/wores/protoc-gen-pipeline/templates/go"
	"github.com/wores/protoc-gen-pipeline/templates/gogo"
	"github.com/wores/protoc-gen-pipeline/templates/shared"
)

type RegisterFn func(tpl *template.Template, params pgs.Parameters)
type FilePathFn func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath

func makeTemplate(ext string, fn RegisterFn, params pgs.Parameters) *template.Template {
	tpl := template.New(ext)
	shared.RegisterFunctions(tpl, params)
	fn(tpl, params)
	return tpl
}

func Template(params pgs.Parameters) map[string][]*template.Template {
	return map[string][]*template.Template{
		"go": {makeTemplate("go", golang.Register, params)},
		"gogo": {makeTemplate("go", gogo.Register, params)},
	}
}

func FilePathFor() FilePathFn {
	return func(f pgs.File, ctx pgsgo.Context, tpl *template.Template) *pgs.FilePath {
		out := ctx.OutputPath(f)
		out = out.SetExt(".pipeline." + tpl.Name())
		return &out
	}
}
