package testutil

import (
	"strings"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/andreyvit/diff"
)

type Test struct {
	T       *testing.T
	Name    string
	Message string
}

func (t *Test) Equals(actual string, expected string) {
	a := strings.TrimSpace(actual)
	e := strings.TrimSpace(heredoc.Doc(expected))

	if a != e {
		t.T.Errorf(
			"%s test fails: %s\n%v",
			t.Name,
			t.Message,
			diff.LineDiff(a, e),
		)
	}
}
