package util

type Fn func(interface{}) (string, error)

type Node struct {
	Name     string
	Children []Node
	Fn       Fn
}

var help = func(d interface{}) (string, error) {
	return "Help!", nil
}

var commands = Node{
	Name: "generate",
	Children: []Node{
		{
			Name:     "next",
			Children: []Node{},
			Fn:       help,
		},
	},
	Fn: help,
}

// func GetNode(n *Node, path []string) (*Node) {
// 	if len(path) == 0 {
// 		return n
// 	}
// 	for _, child := range n.Children {

// 	}
// 	return nil
// }
