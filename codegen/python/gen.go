package py

import (
	"strings"

	"github.com/lynxai-team/emo/codegen/core"
)

func GenPy(ref []core.Ref) {
	var code strings.Builder
	code.WriteString(codeStart)
	for _, item := range ref {
		code.WriteString(genFunc(item.Name, item.Emoji))
	}

	core.Write("lang/python/pyemo/emo_gen.py", code.String())
}

func genFunc(name, emoji string) string {
	name = core.Uncapitalized(name)
	name = core.SnakeCase(name)

	return `
    def ` + name + `(self, *args):
        return self.emo("` + emoji + `", list(args))
`
}
