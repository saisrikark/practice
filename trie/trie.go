package trie

import (
	"errors"
	"fmt"
)

var (
	ErrNoWordSupplied  = fmt.Errorf("no word supplied")
	ErrFoundButNotWord = fmt.Errorf("found but not word")
	ErrNoPrefixGiven   = fmt.Errorf("no prefix given")
	ErrNoCompletions   = fmt.Errorf("no completions")
)

type TrieTree struct {
	root *trieNode
}

type trieNode struct {
	isWord bool
	letter string
	index  map[string]*trieNode
}

func NewTrieTree() *TrieTree {
	return &TrieTree{
		root: newNode(),
	}
}

func newNode() *trieNode {
	return &trieNode{
		index: map[string]*trieNode{},
	}
}

func (tt *TrieTree) Insert(words ...string) error {

	var errs []error

	for _, word := range words {
		if err := tt.insert(tt.root, word, 0); err != nil {
			errs = append(errs, fmt.Errorf("unable to insert word %s err: %w", word, err))
		}
	}

	return errors.Join(errs...)
}

func (tt *TrieTree) insert(node *trieNode, word string, index int) error {

	if index+1 > len(word) {
		return nil
	}

	letter := string(word[index])
	letterNode, isPresent := node.index[letter]
	if !isPresent {
		letterNode = newNode()
		letterNode.letter = letter
		node.index[letter] = letterNode
	}

	if index+1 == len(word) {
		letterNode.isWord = true
	}

	return tt.insert(letterNode, word, index+1)
}

func (tt *TrieTree) Search(word string) (bool, error) {

	if len(word) == 0 {
		return false, ErrNoWordSupplied
	}

	return tt.search(tt.root, word, 0)
}

func (tt *TrieTree) search(node *trieNode, word string, index int) (bool, error) {

	if index+1 > len(word) {
		if node.isWord {
			return true, nil
		}
		return false, ErrFoundButNotWord
	}

	letter := string(word[index])
	letterNode, isPresent := node.index[letter]
	if !isPresent {
		return false, fmt.Errorf("did not find letter %s", letter)
	}

	return tt.search(letterNode, word, index+1)
}

func (tt *TrieTree) Completions(prefix string) ([]string, error) {

	if len(prefix) == 0 {
		return nil, ErrNoPrefixGiven
	}

	return tt.completions(tt.root, prefix, 0)
}

func (tt *TrieTree) completions(node *trieNode, prefix string, index int) ([]string, error) {

	if index+1 > len(prefix) {
		words := node.findWords(prefix, true)
		if len(words) == 0 {
			return nil, ErrNoCompletions
		}
		return words, nil
	}

	letter := string(prefix[index])
	letterNode, isPresent := node.index[letter]
	if !isPresent {
		return []string{}, fmt.Errorf("did not find letter %s", letter)
	}

	return tt.completions(letterNode, prefix, index+1)
}

func (tn *trieNode) findWords(prefix string, ignoreRoot bool) []string {

	words := []string{}

	if tn.isWord && !ignoreRoot {
		words = append(words, prefix+tn.letter)
	}

	for _, node := range tn.index {
		words = append(words, node.findWords(prefix+tn.letter, false)...)
	}

	return words
}
