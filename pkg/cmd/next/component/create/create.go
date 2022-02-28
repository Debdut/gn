package create

import (
	"github.com/debdut/gn/pkg/command"

	"github.com/MakeNowJust/heredoc"
)

func New() *command.Command {
	cmd := &command.Command{
		Name:    "Create",
		Command: "create",
		Use:     "gn next component create <args> :mod:mod:",
		Short:   "Create Components for Next",
		Long: heredoc.Doc(`
			Generate components for Next with optional class components,
			hooks, prebuilt html generics and redux integrations.
		`),
	}

	return cmd
}
