package component

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/debdut/gn/pkg/command"

	"github.com/debdut/gn/pkg/cmd/next/component/create"
)

func New() *command.Command {
	cmd := &command.Command{
		Name:    "Component",
		Command: "component",
		Use:     "gn next component <subcommand> :mod:mod:",
		Aliases: []string{"comp"},
		Short:   "Generate components for Next",
		Long: heredoc.Doc(`
			Generate components for Next with optional class components,
			hooks, prebuilt html generics and redux integrations.
		`),
		Commands: []*command.Command{},
	}

	cmd.AddCommands(
		create.New(),
	)

	return cmd
}
