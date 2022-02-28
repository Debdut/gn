package page

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/debdut/gn/pkg/command"

	"github.com/debdut/gn/pkg/cmd/next/page/create"
)

func New() *command.Command {
	cmd := &command.Command{
		Name:    "Page",
		Command: "page",
		Use:     "gn next page <subcommand> :mod:mod:",
		Short:   "Generate components for Next",
		Long: heredoc.Doc(`
			Generate Next Page templates with optional class components,
			hooks, server side rendering, static generation.
		`),
		Commands: []*command.Command{},
	}

	cmd.AddCommands(
		create.New(),
	)

	return cmd
}
