package main

import "fmt"

// SLLNode ...
type SLLNode struct {
	next  *SLLNode
	value int
}

// SetValue ...
func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

// GetValue ...
func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

// NewSLLNode ...
func NewSLLNode() *SLLNode {
	return new(SLLNode)
}

func main() {
	node := NewSLLNode()
	node.SetValue(4)
	fmt.Println("Node is of value:", node.GetValue())
}
