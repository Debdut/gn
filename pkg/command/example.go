package command

import (
	"fmt"
	"strings"
)

type CommandExample struct {
	Command string
	Output  string
	Short   string
}

func (ex *CommandExample) String() string {
	var text []string
	if len(ex.Short) > 0 {
		text = append(text, fmt.Sprintf("# %s", ex.Short))
	}
	text = append(text, fmt.Sprintf("$ %s", ex.Command))
	if len(ex.Output) > 0 {
		text = append(text, ex.Output)
	}

	if len(text) > 1 {
		text = append(text, "")
	}

	return strings.Join(text, "\n")
}
