package markdown

import (
	"testing"

	"github.com/MakeNowJust/heredoc"
)

func TestRenderHeading(t *testing.T) {
	h2 := Heading{
		Level: 2,
		Text:  []Phrase{&Plain{Text: "Heading"}, &Italic{Text: "Two"}},
	}

	mdH2 := h2.Render()
	mdH2Expected := "## Heading *Two*"

	if mdH2 != mdH2Expected {
		t.Fatalf("Heading doesn't render correctly\n%s\n%s\n", mdH2, mdH2Expected)
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
