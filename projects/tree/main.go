package main

import "fmt"

type Node struct {
	nodes *[]Node
	data  int
}

func (n *Node) Data() int {
	return n.data
}

func (n *Node) SetData(d int) {
	n.data = d
}

func New(d int) Node {
	ln := make([]Node, 0, 5)
	return Node{nodes: &ln, data: d}
}

func (n *Node) Insert(i Node) {
	*n.nodes = append(*n.nodes, i)
}

func (n *Node) Childs() []Node {
	if n.nodes != nil {
		return *n.nodes
	}
	return []Node{}
}

func printTree(t Node) {
	fmt.Printf("%d ", t.data)
	if t.nodes != nil {
		for _, n := range *t.nodes {
			printTree(n)
		}
	}

}

func test() {
	t := New(-1)
	t.Insert(New(1))
	x := t.Childs()
	x[0].Insert(New(4))
	t.Insert(New(2))
	t.Insert(New(3))

	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", t.Childs())
	printTree(t)
}

func test2() {
	count := 9
	t := New(-1)
	for i := 0; i < count; i++ {
		t.Insert(New(i))
	}
	printTree(t)
}

/*
implementing tree using map


*/

func test3() Tree {
	tree := Tree{"a": nil}
	tree["b"] = nil
	//override nodes a and b
	//tree = Tree{"c": nil}
	return tree
}

type Tree map[string]Tree

func main() {
	//test2()

	/* 	tree := Tree{
	   		"root": Tree{
	   			"sibiling1": Tree{
	   				"child4": nil,
	   			},
	   			"sibiling2": Tree{
	   				"child4": nil,
	   			},
	   			"sibiling3": Tree{
	   				"child4": nil,
	   			},
	   		},
	   	}
	*/
	fmt.Printf("%v\n", test3())
}
