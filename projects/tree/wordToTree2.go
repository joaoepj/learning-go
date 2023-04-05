package main

import "fmt"

type Tree map[string]Tree

func newNode() Tree {
	t := make(Tree)
	return t
}

func main() {
	tree := Tree{}
	tree["root"] = nil
	tree["root"] = Tree{"a": nil}
	tree2 := Tree{}
	fmt.Println(tree, tree2)
}
