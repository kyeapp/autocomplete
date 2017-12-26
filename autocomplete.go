package main

import (
	"errors"
	"fmt"
	"bufio"
	"os"
	"time"
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
func (t *trie) autocomplete(baseWord string) []string {
	defer timeTrack(time.Now(), "autocomplete")
	start, err := t.findRoot([]byte(baseWord))
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
func loadDictionary(root *trie, filename string) {
	defer timeTrack(time.Now(), "load Dictionary")
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		root.add([]byte(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func timeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    fmt.Printf("%s took %s\n", name, elapsed)
}

func main() {
	root := new(trie)
	root.init()
	loadDictionary(root, "words.txt")

	_ = root.autocomplete("brin")
	//fmt.Println(list)


	fmt.Println()

}
