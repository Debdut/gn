package command

type CommandExample struct {
	Command string
	Output  string
}

type CommandMeta struct {
	Short    string // Short Description
	Long     string // Long Description
	Aliases  []string
	Examples []CommandExample
}

type CommandRun func(*Command, []string) error

type Command struct {
	Name        string
	SubCommands [](*Command)
	Parent      *Command
	Args        [](*Arg)
	Modifiers   []string
	Meta        CommandMeta
	Run         CommandRun
}
