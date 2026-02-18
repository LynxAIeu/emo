package main

import (
	"flag"
	"fmt"

	"github.com/lynxai-team/emo/codegen/core"
	"github.com/lynxai-team/emo/codegen/dart"
	"github.com/lynxai-team/emo/codegen/doc"
	"github.com/lynxai-team/emo/codegen/golang"
	"github.com/lynxai-team/emo/codegen/ts"

	py "github.com/lynxai-team/emo/codegen/python"
)

func main() {
	fmt.Println("[codegen] Generator of emo source code")

	dartFlag := flag.Bool("dart", false, "generate Dart code")
	docFlag := flag.Bool("doc", false, "generate the documentation")
	goFlag := flag.Bool("go", false, "generate Go code")
	pyFlag := flag.Bool("py", false, "generate Python code")
	tsFlag := flag.Bool("ts", false, "generate Typescript code")
	flag.Parse()

	hasFlag := *dartFlag || *docFlag || *goFlag || *pyFlag || *tsFlag
	if !hasFlag {
		*dartFlag, *docFlag, *goFlag, *pyFlag, *tsFlag = true, true, true, true, true
		fmt.Println("[codegen] No flag => generate code for all languages")
	}

	ref := core.GetRef()
	ref = core.Sanitize(ref)

	if *dartFlag {
		fmt.Println("[codegen] Generating Dart code")
		dart.GenCode(ref)
	}

	if *docFlag {
		fmt.Println("[codegen] Generating documentation")
		doc.GenDoc(ref)
	}

	if *goFlag {
		fmt.Println("[codegen] Generating Go code")
		golang.GenGo(ref)
	}

	if *pyFlag {
		fmt.Println("[codegen] Generating Python code")
		py.GenPy(ref)
	}

	if *tsFlag {
		fmt.Println("[codegen] Generating Typescript code")
		ts.GenTs(ref)
	}
}
