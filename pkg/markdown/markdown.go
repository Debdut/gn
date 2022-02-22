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

func (md *Markdown) Render() string {
	var text []string
	for _, node := range md.Nodes {
		text = append(text, node.Render())
	}
	return strings.Join(text, "\n")
}
