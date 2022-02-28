package create

import (
	"github.com/debdut/gn/pkg/command"

	"github.com/MakeNowJust/heredoc"
)

func New() *command.Command {
	cmd := &command.Command{
		Name:    "Create",
		Command: "create",
		Use:     "gn next page create <args> :mod:mod:",
		Short:   "Create Page templates for Next",
		Long: heredoc.Doc(`
			Generate Next Page templates with optional class components,
			hooks, server side rendering, static generation.
		`),
	}

	return cmd
}
