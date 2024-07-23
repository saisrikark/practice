package trie

type TrieTree struct {
	root *trieNode
}

type trieNode struct {
	isWord bool
	// each letter can be mapped to a list of other nodes
	index map[string][]*trieNode
}

func NewTrieTree() *TrieTree {
	return &TrieTree{
		root: &trieNode{},
	}
}

func (tt *TrieTree) Insert(word ...string) error {
	return nil
}

func (tt *TrieTree) insert(word string) error {
	return nil
}

func (tt *TrieTree) Search(word string) (bool, []string, error) {
	return false, []string{}, nil
}

func (tt *TrieTree) Completions(prefix string) ([]string, error) {
	return []string{}, nil
}
