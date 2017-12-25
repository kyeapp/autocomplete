package main

import(
  "fmt"
)

type trie struct {
  isWord bool
  trieChar map[byte]*trie
}

func (t *trie) init() {
	t.isWord = false
	t.trieChar = make(map[byte]*trie)
}

// adds word into the trie
func (t *trie) add(word []byte) {
  
  //if end of word add new trie and mark
  fmt.Println("word", word)
  if len(word) == 0 {
  	t.isWord = true
  	return
  }

  letter := word[0]
  fmt.Println("letter", letter)
  var next *trie
  var ok bool
  if next, ok = t.trieChar[letter]; !ok {
  	newT := new(trie)
  	newT.init()
  	t.trieChar[letter] = newT
  	next = t.trieChar[letter]
  }

  next.add(word[1:])


  //recursion

}

// returns a list of all words that are a possible autocomplete match
func (t *trie) autocomplete() []string {
	
	return nil
}

//load dictionary into the trie
func loadDictionary(root *trie, filename string) *trie {
	return nil
}

func initTrie() (t *trie) {
	t = new(trie)
	t.isWord = false
	t.trieChar = make(map[byte]*trie)
	return
}

func main() {
	root := new(trie)
	root.init()
	root.add([]byte{'a', 'l', 'l'})
	root.add([]byte{'a', 't'})

	fmt.Println(root.trieChar)

  
}
