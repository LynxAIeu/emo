package dart

import (
	"emo/codegen/core"
)

func GenCode(ref []core.Ref) {
	code := codeStart
	for _, item := range ref {
		code += genFunc(item.Name, item.Emoji)
	}
	code += codeEnd

	core.Write("lang/dart/lib/src/debug.dart", code)
}

func genFunc(name, emoji string) string {
	name = core.Uncapitalized(name)

	return `
  /// A debug message for ` + name + `
  ///
  /// emoji: ` + emoji + `
  String ` + name + `([dynamic obj, String? domain]) => emo("` + emoji + `", obj, domain);
`
}
