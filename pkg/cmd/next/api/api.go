package api

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/debdut/gn/pkg/command"

	"github.com/debdut/gn/pkg/cmd/next/api/create"
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
		Commands: []*command.Command{},
	}

	cmd.AddCommands(
		create.New(),
	)

	return cmd
}
