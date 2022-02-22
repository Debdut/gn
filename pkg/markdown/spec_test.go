package markdown

import (
	"strings"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/andreyvit/diff"
)

func TestRenderHeading(t *testing.T) {
	h2 := Heading{
		Level: 2,
		Text:  []Phrase{&Plain{Text: "Heading"}, &Italic{Text: "Two"}},
	}

	mdH2 := h2.Render()
	mdH2Expected := "## Heading *Two*"

	if mdH2 != mdH2Expected {
		t.Errorf(
			"Heading doesn't render correctly\n%v",
			diff.LineDiff(mdH2, mdH2Expected),
		)
	}
}

func TestRenderList(t *testing.T) {
	list := List{
		List: []ListItem{
			{Text: []Phrase{&Plain{Text: "Item One"}}},
			{Text: []Phrase{&Plain{Text: "Item Two"}}},
			{
				Text: []Phrase{&Plain{Text: "Item Three"}},
				List: []ListItem{
					{Text: []Phrase{&Plain{Text: "Sub Item 1"}}},
					{Text: []Phrase{&Plain{Text: "Sub Item 2"}}},
				},
			},
			{Text: []Phrase{&Plain{Text: "Item Four"}}},
		},
	}

	mdList := list.Render()
	mdListExpected := heredoc.Doc(`
		* Item One
		* Item Two
		* Item Three
			* Sub Item 1
			* Sub Item 2
		* Item Four
	`)

	actual := strings.TrimSpace(mdList)
	expected := strings.TrimSpace(mdListExpected)

	if actual != expected {
		t.Fatalf(
			"List doesn't render correctly\n%v",
			diff.LineDiff(actual, expected),
		)
	}
}

func TestRenderTable(t *testing.T) {
	table := Table{
		Headers: []Text{
			[]Phrase{&Plain{Text: "Name"}},
			[]Phrase{&Plain{Text: "Country"}},
			[]Phrase{&Plain{Text: "Age"}},
		},
		Rows: [][]Text{
			{
				[]Phrase{&Plain{Text: "Debdut"}, &Italic{Text: "Karmakar"}},
				[]Phrase{&Plain{Text: "India"}},
				[]Phrase{&Plain{Text: "24"}},
			},
			{
				[]Phrase{&Plain{Text: "Sayan"}, &Bold{Text: "Tan"}},
				[]Phrase{&Plain{Text: "Turkiye"}},
				[]Phrase{&Plain{Text: "-80"}},
			},
		},
		Align: []string{"left", "center", "right"},
	}
	mdTable := table.Render()
	mdTableExpected := heredoc.Doc(`
		| Name               | Country  | Age  |
		| ------------------ | :------: | ---: |
		| Debdut *Karmakar*  | India    | 24   |
		| Sayan **Tan**      | Turkiye  | -80  |	
	`)

	if mdTable != mdTableExpected {
		t.Fatal("Table doesn't render to markdown correctly")
	}
}
