package doc

import (
	"fmt"
	"strings"

	"github.com/lynxai-team/emo/codegen/core"
)

func GenDoc(ref []core.Ref) {
	var code strings.Builder
	code.WriteString(fileStart)
	for _, item := range ref {
		code.WriteString(genFunc(item.Name, item.Emoji, item.IsError))
	}

	core.Write("doc/events/README.md", code.String())
}

var fileStart = `
# Emo event types

| Name          |  Emoji |  IsError |
|---------------|:------:|:--------:|
`

func genFunc(name, emoji string, isError bool) string {
	errStr := " "
	if isError {
		errStr = "✔️"
	}
	return fmt.Sprintf("| %-13s |   "+emoji+"   |     "+errStr+"    |"+"\n", name)
}
