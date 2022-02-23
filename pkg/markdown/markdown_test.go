package markdown

import (
	"testing"

	"github.com/debdut/gn/pkg/testutil"
)

func TestRenderMarkdown(t *testing.T) {
	md := Markdown{}

	md.AddHeading(1, "Heading")
	md.AddHeading(2, "Subheading")
	md.AddHeading(3, "Heading Three")
	md.AddLine()
	md.AddBlockQuote("This is a block quote")
	md.AddNewLine()
	md.AddParagraph("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	md.AddNewLine()
	md.AddParagraph("This is a new paragraph")
	md.AddHeading(3, "List of Tasks")
	md.AddList(
		[]string{
			"Buy 10 Mangoes",
			"Peel them following the steps below",
			"Lick the juicy mangoes",
			"Easy *peeze*",
		},
		[][]string{
			{},
			{
				"Grab a peeler under 5$",
				"Hold a mongo on the other hand",
				"Peel slowly",
			},
		},
		true,
	)

	markdown := md.Render()
	markdownExpected := `

	`
	test := testutil.Test{
		Name:    "Markdown",
		Message: "Markdown doesn't generate or render correctly",
		T:       t,
	}

	test.Equals(markdown, markdownExpected)
}
