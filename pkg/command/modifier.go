package command

type ModRun func(*Modifier) error

type Modifier struct {
	Name     string
	Modifier string

	Parent *Command
	Run    ModRun

	Short    string // Short Description
	Long     string // Long Description
	Aliases  []string
	Examples []string
}
