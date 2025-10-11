package day06

import (
	"strings"

	"github.com/Evokoo/AOC_2019_Go/utils"
)

// ========================
// NODE
// ========================
type Node struct {
	id       string
	parent   *Node
	children []*Node
}

func NewNode(id string) *Node {
	return &Node{id: id}
}

func (n *Node) AddParent(node *Node) {
	(*n).parent = node
}
func (n *Node) AddChild(node *Node) {
	(*n).children = append((*n).children, node)
}

// ========================
// TREE
// ========================

type Tree struct {
	root   *Node
	lookup map[string]*Node
}

func NewTree(rootID string) Tree {
	root := NewNode(rootID)
	return Tree{
		root:   root,
		lookup: map[string]*Node{rootID: root},
	}
}

func (t *Tree) RetriveNode(id string) *Node {
	if node, found := t.lookup[id]; found {
		return node
	} else {
		t.lookup[id] = NewNode(id)
	}
	return t.lookup[id]
}

func (t *Tree) InsertNode(pair []string) {
	parent := t.RetriveNode(pair[0])
	child := t.RetriveNode(pair[1])

	parent.AddChild(child)
	child.AddParent(parent)
}

// ========================
// PARSER
// ========================

func BuildTree(file string) Tree {
	data := utils.ReadFile(file)
	tree := NewTree("COM")

	for line := range strings.SplitSeq(data, "\n") {
		pair := strings.Split(line, ")")
		tree.InsertNode(pair)
	}

	return tree
}
