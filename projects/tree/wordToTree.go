package main

import (
	"fmt"
	"log"
)

// as each node is a map you need to carefully reserve a value in your key space for root
// if you try to instantiate a no-value root tree you'll end up with no tree at all
// as an empty map is referenced by a nil pointer. so every function must validate that
// the tree root has the choosed value

// ########################################################################
type Tree map[string]Tree

func main() {

	t := Tree{}
	log.SetFlags(log.Lshortfile)
	log.Println("wordsTotree", t, []string{"a", "the", "their", "abacate", "abacadabra"})
	wordsTotree(t, []string{"a", "the", "their", "abacate", "abacadabra"})
	printBranch(t, 10)
}

func wordsTotree(t Tree, sl []string) Tree {
	if len(sl) != 0 {
		// for each word
		for _, w := range sl {
			// create a branch in tree
			//log.Println("wordTobranch", t, w)
			wordTobranch(t, w)
			//log.Println("wordsTotree", t, sl[1:len(sl)])
			wordsTotree(t, sl[1:len(sl)])
		}
	}
	return t
}

// a more elegant version
func wordTobranch(tree Tree, s string) {
	if len(s) == 0 {
		return
	}
	child, ok := tree[string(s[0])]
	if !ok {
		child = Tree{}
		tree[string(s[0])] = child
	}

	wordTobranch(child, s[1:])

}

func padding(p int) string {
	b := make([]byte, p)
	for i := 0; i < p-1; i++ {
		b = append(b, 32)
	}

	return string(b)
}

func printBranch(t Tree, p int) {
	for node, child := range t {

		fmt.Println(padding(p), node)

		printBranch(child, p-1)

	}

}

func wordBranch(t Tree, s string) {
	if len(s) != 0 {
		t[string(s[0])] = Tree{}
		wordBranch(t, s[1:len(s)])
	}

	return
}
