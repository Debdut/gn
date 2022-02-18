package util

type Node struct {
	Name     string
	Children [](*Node)
	Data     interface{}
}

func PreOrder(node *Node, order *[](*Node)) {
	*order = append(*order, node)
	for _, childNode := range node.Children {
		PreOrder(childNode, order)
	}
}

func PostOrder(node *Node, order *[](*Node)) {
	for _, childNode := range node.Children {
		PostOrder(childNode, order)
	}
	*order = append(*order, node)
}

func LevelOrder(node *Node, order *[](*Node)) {
	*order = append(*order, node.Children...)
	for _, childNode := range node.Children {
		LevelOrder(childNode, order)
	}
}
