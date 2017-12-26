package main

import (
	"fmt"
	"errors"
)

type trie struct {
	base     string
	isWord   bool
	trieChar map[byte]*trie
}

func (t *trie) init() {
	t.isWord = false
	t.trieChar = make(map[byte]*trie)
}

// adds word into the trie
func (t *trie) add(word []byte) {

	//if end of word add new trie and mark
	if len(word) == 0 {
		t.isWord = true
		return
	}

	letter := word[0]
	var next *trie
	var ok bool
	if next, ok = t.trieChar[letter]; !ok {
		newT := new(trie)
		newT.init()
		newT.base = t.base + string(letter)
		t.trieChar[letter] = newT
		next = t.trieChar[letter]
	}

	next.add(word[1:])
}


// find the top level trie of possible matchine autocomplete words
func (t *trie) findRoot(w []byte) (*trie, error) {
	if len(w) == 0 {
		return t, nil
	}
	nextLetter := w[0]
	nextTrie := t.trieChar[nextLetter]
	if nextTrie == nil {
		return nil, errors.New("no word possible matches found")
	}
	return nextTrie.findRoot(w[1:])
}

// returns a list of all words that are a possible autocomplete match
func (t *trie) autocomplete(w []byte) []string {
	start, err := t.findRoot(w)
	if err != nil {
		return []string{}
	}

	return start.listWords()
}

//lists words from the current trie
func (t *trie) listWords() (list []string) {
	if t.isWord {
		list = append(list, t.base)
	}
	
	for _, childTrie := range t.trieChar {
		list = append(list, childTrie.listWords()...)
	}
	
	return list
}

//load dictionary into the trie
func loadDictionary(root *trie, filename string) *trie {
	return nil
}

func main() {
	root := new(trie)
	root.init()
	root.add([]byte{'a', 'l', 'l'})
	root.add([]byte{'a', 't'})
	root.add([]byte{'b', 'a', 't'})

	list := root.autocomplete([]byte{'b', 'r'})
	fmt.Println(list)

}
