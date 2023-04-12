package main

import (
	"log"
)

type Trie struct {
	letter map[string]Trie
	end    bool
}

func wordToBranch(t Trie, s string) {
	// halt condition
	if len(s) == 0 {
		return
	}

	child, ok := t.letter[string(s[0])]
	// end of word
	if len(s) == 1 {
		child.end = true
	}

	if !ok {
		// child becomes t in the next recursive call
		// if it is new compiler throws:
		// assigment to entry in nil map error
		child.letter = make(map[string]Trie)
		t.letter[string(s[0])] = child
	}
	// recursive call
	wordToBranch(child, s[1:])
}

func wordsToTrie(t Trie, sl []string) {
	if len(sl) == 0 {
		return
	}
	wordToBranch(t, sl[0])
	wordsToTrie(t, sl[1:])
}

func extractWord(t Trie, s string) string {
	if t.end == true {
		return s
	}
	for letter, child := range t.letter {
		return s + extractWord(child, letter)
	}

	return s

}

func main() {
	// initialize map
	letter := make(map[string]Trie)
	trie := Trie{letter: letter, end: false}
	wordsToTrie(trie, []string{"abacaxi", "abacaxa", "babacaxe"})
	log.Println(trie)
	log.Println(extractWord(trie, ""))

}
