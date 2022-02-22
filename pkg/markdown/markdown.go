package markdown

import (
	"strings"
)

type Markdown struct {
	Nodes []MarkdownNode
}

func (md *Markdown) AddNode(node MarkdownNode) {
	md.Nodes = append(md.Nodes, node)
}

func (md *Markdown) AddHeading(level uint8, text string) {
	md.Nodes = append(md.Nodes, &Heading{
		Level: level,
		Text:  []Phrase{&Plain{Text: text}},
	})
}

func (md *Markdown) AddTable(headers []string, rows [][]string, align []string) {
	table := Table{
		Headers: []Text{},
		Rows:    [][]Text{},
		Align:   align,
	}
	numRows := len(rows)
	numCols := len(rows[0])

	// fill the headers
	table.Headers = make([]Text, len(headers))
	for c := 0; c < len(headers); c++ {
		table.Headers[c] = []Phrase{&Plain{Text: headers[c]}}
	}

	// fill the rows
	table.Rows = make([][]Text, numRows)
	for r := 0; r < numRows; r++ {
		table.Rows[r] = make([]Text, numCols)
		for c := 0; c < numCols; c++ {
			table.Rows[r][c] = []Phrase{&Plain{Text: rows[r][c]}}
		}
	}

	// append node to markdown
	md.Nodes = append(md.Nodes, &table)
}

func (md *Markdown) Render() string {
	var text []string
	for _, node := range md.Nodes {
		text = append(text, node.Render())
	}
	return strings.Join(text, "\n")
}
