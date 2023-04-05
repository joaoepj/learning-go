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

func printstrTree(t Node) {
	fmt.Printf("%d ", t.data)
	if t.nodes != nil {
		for _, n := range *t.nodes {
			printstrTree(n)
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
	printstrTree(t)
}

func test2() {
	count := 9
	t := New(-1)
	for i := 0; i < count; i++ {
		t.Insert(New(i))
	}
	printstrTree(t)
}

/*
implementing tree using map


*/

type strTree map[string]strTree

type intTree map[int]intTree

type byteTree map[byte]byteTree

func test3() strTree {
	stree := strTree{
		"root": strTree{
			"sibiling1": strTree{
				"child4": nil,
			},
			"sibiling2": strTree{
				"child4": nil,
			},
			"sibiling3": strTree{
				"child4": nil,
			},
		},
	}

	return stree
}

func (it *intTree) sibiling(parent int, x int) {
	sib := (*it)[parent]
	sib[x] = intTree{x: nil}
	//(*it) = sib
}

func main() {
	//test2()

	itree := make(intTree)
	itree[-1] = nil
	itree[-1] = intTree{1: nil}
	sib := itree[-1]
	sib[2] = nil
	// disnecessary
	//itree[-1] = sib

	itree.sibiling(-1, 3)
	itree[-1] = intTree{4: nil}

	fmt.Println(len(itree))
	//fmt.Printf("%v\n", test3())
	fmt.Printf("%v\n", itree)
	var r byte = 'a'
	fmt.Println("🐘", string(r))
}
