package api

import (
	"github.com/debdut/gn/pkg/command"

	"github.com/MakeNowJust/heredoc"
)

func New() *command.Command {
	cmd := &command.Command{
		Name:    "Api",
		Command: "api",
		Use:     "gn next api <subcommand> :mod:mod:",
		Short:   "Generate API templates for Next",
		Long: heredoc.Doc(`
			Generate API templates for Next with optional Typescript data types,
			various HTTP methods.
		`),
	}

	return cmd
}
