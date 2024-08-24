package xmlx

func GetDocNodeValue(doc *Document, ns string, name string) string {
	node := doc.SelectNode(ns, name)
	if node != nil {
		return node.GetValue()
	}
	return ""
}

func GetNodeValue(node *Node, ns string, name string) string {
	_node := node.SelectNode(ns, name)
	if _node != nil {
		return _node.GetValue()
	}
	return ""
}
