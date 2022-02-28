package create

import (
	"github.com/debdut/gn/pkg/command"

	"github.com/MakeNowJust/heredoc"
)

func New() *command.Command {
	cmd := &command.Command{
		Name:    "Create",
		Command: "create",
		Use:     "gn next api create <args> :mod:mod:",
		Short:   "Create API templates for Next",
		Long: heredoc.Doc(`
			Generate API templates for Next with optional Typescript data types,
			various HTTP methods.
		`),
	}

	return cmd
}
