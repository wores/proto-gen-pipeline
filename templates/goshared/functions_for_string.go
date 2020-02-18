package goshared

import (
	"fmt"

	"github.com/wores/protoc-gen-pipeline/pipeline"
	"github.com/wores/protoc-gen-pipeline/templates/shared"
)

type goSharedStringFuncs struct{ goSharedFuncs }

func (fns goSharedStringFuncs) trim(ctx shared.RuleContext, trimType pipeline.TrimType) string {
	genStatement := func(trimType string) string {
		statementFormat := "%s = strings.Trim%sFunc(%s, func(r rune) bool { return unicode.IsSpace(r) })"
		return fmt.Sprintf(statementFormat, fns.property(ctx), trimType, fns.accessor(ctx))
	}

	switch trimType {

	case pipeline.TrimType_TrimTypeBoth:
		return genStatement("")

	case pipeline.TrimType_TrimTypeLeft:
		return genStatement("Left")

	case pipeline.TrimType_TrimTypeRight:
		return genStatement("Right")

	}

	return fmt.Sprintf("%s = strings.TrimSpace(%s)", fns.property(ctx), fns.accessor(ctx))
}

