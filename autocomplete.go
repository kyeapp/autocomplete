package main

import(
  "fmt"
)

struct trie {
  isWord bool
  next map[rune]trie
}

// adds word into the trie
func (t trie) add(word string) {
  
}

// returns a list of all words that are a possible autocomplete match
func (t trie) autocomplete() []string {

}

//load dictionary into the trie
func loadDictionary(root trie, filename string) trie {

}

func main() {
  
}
