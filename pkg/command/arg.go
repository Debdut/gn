package command

type ArgCheck func(*Command, []string) bool
type ArgRun func(*Command, []string) error

type Arg struct {
	Args  []string
	Check ArgCheck
	Run   ArgRun
}
