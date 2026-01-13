package dart

import (
	"strings"

	"github.com/lynxai-team/emo/codegen/core"
)

func GenCode(ref []core.Ref) {
	var code strings.Builder
	code.WriteString(codeStart)
	for _, item := range ref {
		code.WriteString(genFunc(item.Name, item.Emoji))
	}
	code.WriteString(codeEnd)

	core.Write("lang/dart/lib/src/debug.dart", code.String())
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
