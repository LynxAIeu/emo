package py

import (
	"emo/codegen/core"
)

func GenPy(ref []core.Ref) {
	code := codeStart
	for _, item := range ref {
		code += genFunc(item.Name, item.Emoji)
	}

	core.Write("lang/python/pyemo/emo_gen.py", code)
}

func genFunc(name, emoji string) string {
	name = core.Uncapitalized(name)
	name = core.SnakeCase(name)

	return `
    def ` + name + `(self, *args):
        return self.emo("` + emoji + `", list(args))
`
}
