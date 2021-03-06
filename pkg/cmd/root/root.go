package root

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/debdut/gn/pkg/command"

	"github.com/debdut/gn/pkg/cmd/next"
	"github.com/debdut/gn/pkg/cmd/react"
)

func New() *command.Command {
	cmd := &command.Command{
		Name:    "Generator",
		Command: "gn",
		Use:     "gn <command> <subcommand> ... :mod:mod:",
		Aliases: []string{"generate"},
		Short:   "The Next React Scaffolder & Generator",
		Long: heredoc.Doc(`
			A Code Generator and Scaffolding tool for the JS, TS ecosystem
			Scaffold your whole project, and generate components, files,
			file groups in your project for frameworks like Next, React.
			Never write boilerplate or repetative code again!
		`),
		Commands: []*command.Command{},
	}

	cmd.AddCommands(
		react.New(),
		next.New(),
	)

	return cmd
}
