package ts

import (
	"github.com/lynxai-team/emo/codegen/core"
)

func GenTs(ref []core.Ref) {
	code := codeStart
	for _, item := range ref {
		code += genFunc(item.Name, item.Emoji)
	}
	code += codeEnd

	core.Write("lang/typescript/src/emo_gen.ts", code)
}

func genFunc(name, emoji string) string {
	name = core.Uncapitalized(name)

	return "\n\t" + name + `(...obj: any[]): string { return this.emo("` + emoji + `", obj); }
`
}
