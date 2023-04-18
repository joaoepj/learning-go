package main

import (
	"fmt"
	"log"
)

// Trie node state
type State byte

const (
	Undiscovered State = iota
	Discovered
	Processed
)

type Trie struct {
	childs map[string]Trie
	end    bool
	State
}

// WordToBranch takes a word and builds a trie branch
// First letter is root's child and last letter is trie's leaf
func WordToBranch(t *Trie, s string) {
	// halt condition
	if len(s) == 0 {
		return
	}

	// useless if Undiscovered is zero value
	t.State = Undiscovered
	child, ok := t.childs[string(s[0])]
	// end of word

	if !ok {
		// child becomes t in the next recursive call
		// if it is nil compiler throws:
		// assigment to entry in nil map error
		child.childs = make(map[string]Trie)
		if len(s) == 1 {
			child.end = true
		}
		// useless if Undiscovered is zero value
		child.State = Undiscovered
		t.childs[string(s[0])] = child

	}
	// recursive call
	WordToBranch(&child, s[1:])
}

// WordsToTrie takes a word list
// Each time it call WordToBranch it passes a word to it
// Each diferent word makes a different branch on trie
func WordsToTrie(t *Trie, sl []string) {
	if len(sl) == 0 {
		return
	}
	WordToBranch(t, sl[0])
	WordsToTrie(t, sl[1:])
}

func extractWord(t Trie, s string) string {
	if t.end == true {
		return s
	}
	for childs, child := range t.childs {
		return s + extractWord(child, childs)
	}
	return s
}

func dfs(t *Trie, stack *[]*Trie, s string) []string {
	var sl []string

	if t.end == true {
		// TODO: append the word into a string list
		t.State = Processed

		*stack = (*stack)[:len(*stack)-1]
		log.Println("final")
		log.Println("s", s, "Trie:\t", *t)
		log.Println("Stack:", stack)
		log.Println("-------------------------------------------")

		sl = append(sl, s)
		// Easy, very dangerous
		//dfs((*stack)[len(*stack)-1], stack, s)
	}

	t.State = Discovered
	// push into stack
	*stack = append(*stack, t)
	for str, trie := range t.childs {
		if trie.State == Undiscovered {
			//sl = append(sl, dfs(&trie, stack, s+str)...)
			log.Println("recursive")
			log.Println("sl:", sl, "s:", s, "str:", str, "Trie:\t", trie)
			log.Println("Stack:", stack)
			log.Println("-------------------------------------------")
			sl = append(sl, dfs(&trie, stack, s+str)...)

			log.Println("sl:", sl, "s:", s, "str:", str, "Trie:\t", trie)
			log.Println("Stack:", stack)
			log.Println("-------------------------------------------")
		}
	}
	return sl
}

func deepestBranch() {
	i := 0
	fmt.Println(i)
	i++
	deepestBranch()
}

func oldmain() {

	//log.SetOutput(ioutil.Discard)
	log.SetFlags(log.Lshortfile)
	// initialize map
	stack := make([]*Trie, 0)
	childs := make(map[string]Trie)
	trie := Trie{childs: childs, end: false, State: Undiscovered}
	WordsToTrie(&trie, []string{"abacaxi", "abacaxa", "babacaxe"})
	//WordsToTrie(&trie, []string{"ab", "abc"})
	log.Println("after WordsToTrie:", trie)
	//log.Println(extractWord(trie, ""))
	fmt.Println(dfs(&trie, &stack, ""))
	log.Println("after dfs:", trie)
	log.Println(len(stack))
}

func main() {
	deepestBranch()
}
