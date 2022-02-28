package next

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/debdut/gn/pkg/command"

	"github.com/debdut/gn/pkg/cmd/next/api"
	"github.com/debdut/gn/pkg/cmd/next/component"
	"github.com/debdut/gn/pkg/cmd/next/page"
)

func New() *command.Command {
	cmd := &command.Command{
		Name:    "Next",
		Command: "next",
		Use:     "gn next <subcommand> ... :mod:mod:",
		Aliases: []string{"nxt"},
		Short:   "Generate templates for Next",
		Long: heredoc.Doc(`
			Generate templates for Next pages, components, routes, apis, hooks,
			also next libraries such as redux. Scaffold your whole next project.
		`),
		SubCommands: []*command.Command{},
	}

	cmd.SubCommands = append(cmd.SubCommands, api.New())
	cmd.SubCommands = append(cmd.SubCommands, component.New())
	cmd.SubCommands = append(cmd.SubCommands, page.New())

	return cmd
}
