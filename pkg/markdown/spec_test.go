package markdown

import (
	"fmt"
	"strings"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/debdut/gn/pkg/testutil"
)

func TestRenderHeading(t *testing.T) {
	h2 := Heading{
		Level: 2,
		Text:  []Phrase{Plain("Heading"), Italic("Two")},
	}

	mdH2 := h2.Render()
	mdH2Expected := "## Heading *Two*"
	test := testutil.Test{
		Name:    "Heading",
		Message: "Heading doesn't render correctly",
		T:       t,
	}

	test.Equals(mdH2, mdH2Expected)
}

func TestRenderList(t *testing.T) {
	list := List{
		List: []ListItem{
			{Text: []Phrase{Plain("Item One")}},
			{Text: []Phrase{Plain("Item Two")}},
			{
				Text: []Phrase{Plain("Item Three")},
				List: []ListItem{
					{Text: []Phrase{Plain("Sub Item 1")}},
					{Text: []Phrase{Plain("Sub Item 2")}},
				},
			},
			{Text: []Phrase{Plain("Item Four")}},
		},
	}

	mdList := list.Render()
	mdListExpected := `
		* Item One
		* Item Two
		* Item Three
			* Sub Item 1
			* Sub Item 2
		* Item Four
	`
	test := testutil.Test{
		Name:    "List",
		Message: "List doesn't render correctly",
		T:       t,
	}

	test.Equals(mdList, mdListExpected)
}

func TestRenderOrderedList(t *testing.T) {
	list := List{
		List: []ListItem{
			{Text: []Phrase{Plain("Item One")}},
			{Text: []Phrase{Plain("Item Two")}},
			{
				Text: []Phrase{Plain("Item Three")},
				List: []ListItem{
					{Text: []Phrase{Plain("Sub Item 1")}},
					{Text: []Phrase{Plain("Sub Item 2")}},
				},
				Order: true,
			},
			{Text: []Phrase{Plain("Item Four")}},
		},
		Order: true,
	}

	mdList := list.Render()
	mdListExpected := `
		1. Item One
		2. Item Two
		3. Item Three
			1. Sub Item 1
			2. Sub Item 2
		4. Item Four
	`
	test := testutil.Test{
		Name:    "Ordered List",
		Message: "Ordered List doesn't render correctly",
		T:       t,
	}

	test.Equals(mdList, mdListExpected)
}

func TestRenderTable(t *testing.T) {
	table := Table{
		Headers: []Text{
			[]Phrase{Plain("Name")},
			[]Phrase{Plain("Country")},
			[]Phrase{Plain("Age")},
		},
		Rows: [][]Text{
			{
				[]Phrase{Plain("Debdut"), Italic("Karmakar")},
				[]Phrase{Plain("India")},
				[]Phrase{Plain("24")},
			},
			{
				[]Phrase{Plain("Sayan"), Bold("Tan")},
				[]Phrase{Plain("Turkiye")},
				[]Phrase{Plain("-80")},
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

	test := testutil.Test{
		Name:    "Table",
		Message: "Table doesn't render correctly",
		T:       t,
	}

	test.Equals(mdTable, mdTableExpected)
}

func TestRenderCode(t *testing.T) {
	text := strings.TrimSpace(
		heredoc.Doc(`
			let expensive_closure = |num| {
				println!("calculating slowly...");
				thread::sleep(Duration::from_secs(2));
				num
			};
		`),
	)
	code := Code{
		Lang: "js",
		Text: text,
	}

	mdCode := code.Render()
	mdCodeExpected := fmt.Sprintf("```js\n%s\n```", text)
	test := testutil.Test{
		Name:    "Code",
		Message: "Code doesn't render correctly",
		T:       t,
	}

	test.Equals(mdCode, mdCodeExpected)
}
