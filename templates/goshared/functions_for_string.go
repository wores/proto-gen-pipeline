package goshared

import (
	"fmt"

	"github.com/wores/protoc-gen-pipeline/pipeline"
	"github.com/wores/protoc-gen-pipeline/templates/shared"
)

type goSharedStringFuncs struct{ goSharedFuncs }

// テンプレート内で描画したかったが、trimTypeの比較方法が見つけられなかったためこの中で行っている。
func (fns goSharedStringFuncs) trim(ctx shared.RuleContext, trimType pipeline.Trim) string {
	fnGenStatement := func(position string) string {
		statementFormat := "%s = strings.Trim%sFunc(%s, func(r rune) bool { return unicode.IsSpace(r) })"
		return fmt.Sprintf(statementFormat, fns.property(ctx), position, fns.accessor(ctx))
	}

	switch trimType {

	case pipeline.Trim_TrimBoth:
		return fnGenStatement("")

	case pipeline.Trim_TrimLeft:
		return fnGenStatement("Left")

	case pipeline.Trim_TrimRight:
		return fnGenStatement("Right")

	default:
		panic(fmt.Sprintf("TrimType %d is not exist", trimType))
	}

}

func (fns goSharedStringFuncs) isOmissionPositionLeft(position pipeline.OmissionPosition) bool {
	return position == pipeline.OmissionPosition_OmissionPositionLeft
}

func (fns goSharedStringFuncs) isOmissionPositionCenter(position pipeline.OmissionPosition) bool {
	return position == pipeline.OmissionPosition_OmissionPositionCenter
}

func (fns goSharedStringFuncs) isOmissionPositionRight(position pipeline.OmissionPosition) bool {
	return position == pipeline.OmissionPosition_OmissionPositionRight
}
