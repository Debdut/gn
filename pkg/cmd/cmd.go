package cmd

type ArgCheck func(*Command, []string) bool
type ArgRun func(*Command, []string) error

type Arg struct {
	Args  []string
	Check ArgCheck
	Run   ArgRun
}

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
