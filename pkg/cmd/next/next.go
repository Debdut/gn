package next

import (
	"github.com/debdut/gn/pkg/command"

	"github.com/MakeNowJust/heredoc"
)

func New() *command.Command {
	cmd := &command.Command{
		Name:    "Next",
		Command: "next",
		Use:     "gn next <subcommand> ... :modifier:modifier:",
		Aliases: []string{"nxt"},
		Short:   "Generate templates for Next",
		Long: heredoc.Doc(`
			Generate templates for Next pages, components, routes, apis, hooks,
			also next libraries such as redux. Scaffold your whole next project.
		`),
	}

	return cmd
}
