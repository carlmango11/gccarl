package parser

import "encoding/json"

type Node struct {
	ID     int
	Key    RuleKey
	Values []*Value
	parent *Node
}

func (n *Node) Clone() *Node {
	if n.parent != nil {
		parent := n.parent.Clone()
		return parent.Values[len(parent.Values)-1].Node
	}

	return n.CloneDown()
}

func (n *Node) CloneDown() *Node {
	clone := &Node{
		Key: n.Key,
	}

	var vals []*Value
	for _, v := range n.Values {
		if v.Node == nil {
			vals = append(vals, &Value{
				Token: v.Token,
			})
		} else {
			node := v.Node.CloneDown()
			node.parent = clone

			vals = append(vals, &Value{
				Node: node,
			})
		}
	}

	clone.Values = vals
	return clone
}

func (n *Node) DeepClone(parent *Node) *Node {
	clone := &Node{
		Key:    n.Key,
		parent: parent,
	}

	var vals []*Value
	for _, v := range n.Values {
		if v.Node == nil {
			vals = append(vals, &Value{
				Token: v.Token,
			})
		} else {
			vals = append(vals, &Value{
				Node: v.Node.DeepClone(clone),
			})
		}
	}

	clone.Values = vals
	return clone
}

func (n *Node) String() string {
	s, _ := json.MarshalIndent(n, "", "\t")
	return string(s)
}
