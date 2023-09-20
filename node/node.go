package node

type Node struct {
	name     string
	incoming chan Message
	outgoing []*chan Message
	main     func(*Node)
}

func initNode(name string, main func(*Node)) *Node {
	r := Node{
		name,
		make(chan Message),
		[]Message{},
		main,
	}
	return r
}

func addOutgoing(node *Node, toAdd *Node) *Node {
	return node
}