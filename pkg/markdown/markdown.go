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

func (md *Markdown) AddList(list []string, subList [][]string, order bool) {
	l := List{
		List:  []ListItem{},
		Order: order,
	}
	// populate list
	l.List = make([]ListItem, len(list))
	for i := 0; i < len(list); i++ {
		li := ListItem{
			List:  []ListItem{},
			Text:  Text{&Plain{Text: list[i]}},
			Order: order,
		}
		if len(subList) > i {
			if len(subList[i]) > 0 {
				li.List = make([]ListItem, len(subList))
				for j := 0; j < len(subList); j++ {
					li.List = append(li.List, ListItem{
						Text: Text{&Plain{Text: subList[i][j]}},
					})
				}
			}
		}
		l.List = append(l.List, li)
	}

	md.Nodes = append(md.Nodes, &l)
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

func (md *Markdown) AddCode(lang string, text string) {
	code := Code{
		Lang: lang,
		Text: text,
	}

	md.Nodes = append(md.Nodes, &code)
}

func (md *Markdown) AddParagraph(text string) {
	paragraph := Paragraph{
		Text: []Phrase{&Plain{Text: text}},
	}

	md.Nodes = append(md.Nodes, &paragraph)
}

func (md *Markdown) AddBlockQuote(text string) {
	quote := BlockQuote{
		Text: []Phrase{&Plain{Text: text}},
	}

	md.Nodes = append(md.Nodes, &quote)
}

func (md *Markdown) AddLine() {
	md.Nodes = append(md.Nodes, &Line{})
}

func (md *Markdown) AddNewLine() {
	md.Nodes = append(md.Nodes, &NewLine{})
}

func (md *Markdown) Render() string {
	var text []string
	for _, node := range md.Nodes {
		text = append(text, node.Render())
	}
	return strings.Join(text, "\n")
}
