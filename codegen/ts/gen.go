package ts

import (
	"strings"

	"github.com/lynxai-team/emo/codegen/core"
)

func GenTs(ref []core.Ref) {
	var code strings.Builder
	code.WriteString(codeStart)
	for _, item := range ref {
		code.WriteString(genFunc(item.Name, item.Emoji))
	}
	code.WriteString(codeEnd)

	core.Write("lang/typescript/src/emo_gen.ts", code.String())
}

func genFunc(name, emoji string) string {
	name = core.Uncapitalized(name)

	return "\n\t" + name + `(...obj: any[]): string { return this.emo("` + emoji + `", obj); }
`
}
