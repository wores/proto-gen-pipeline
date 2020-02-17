package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	"github.com/wores/protoc-gen-pipeline/module"
)

func main() {
	pgs.Init(pgs.DebugEnv("PIPELINE_DEBUG")).
		RegisterModule(module.Pipeline()).
		RegisterPostProcessor(pgsgo.GoFmt()).
		Render()
}
