package root

import (
	"github.com/debdut/gn/pkg/command"

	"github.com/MakeNowJust/heredoc"
)

func New() *command.Command {
	cmd := &command.Command{
		Name:    "Generator",
		Command: "gn",
		Use:     "gn <command> <subcommand> ... :modifier:modifier:",

		Short: "The Next React Scaffolder & Generator",
		Long: heredoc.Doc(`
			A Code Generator and Scaffolding tool for the JS, TS ecosystem
			Scaffold your whole project, and generate components, files,
			file groups in your project for frameworks like Next, React.
			Never write boilerplate or repetative code again!
		`),
	}

	return cmd
}
