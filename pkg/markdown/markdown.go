package markdown

import "fmt"

type Markdown struct {
	Nodes *[]MarkdownNode
}

func (md *Markdown) AddNode(node *MarkdownNode) {
	*md.Nodes = append(*md.Nodes, *node)
}

func (md *Markdown) Render() string {
	text := ""
	for _, node := range *md.Nodes {
		text += fmt.Sprintf("%s\n", node.Render())
	}
	return text
}
