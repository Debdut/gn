package react

import (
	"github.com/debdut/gn/pkg/command"

	"github.com/MakeNowJust/heredoc"
)

func New() *command.Command {
	cmd := &command.Command{
		Name:    "React",
		Command: "react",
		Use:     "gn react <subcommand> ... :modifier:modifier:",
		Aliases: []string{"rct"},
		Short:   "Generate templates for React",
		Long: heredoc.Doc(`
			Generate templates for React components, routes, apis, hooks,
			also react libraries such as redux. Scaffold your whole react
			project.
		`),
	}

	return cmd
}
