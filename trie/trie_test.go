package trie_test

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
	"trie"

	"math/rand"
)

func TestTrieTree(t *testing.T) {

	file, err := os.Open("words.txt")
	if err != nil {
		t.Fatalf(fmt.Sprintf("unable to open file %s", err.Error()))
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	words := []string{}
	for fileScanner.Scan() {
		words = append(words, fileScanner.Text())
	}

	file.Close()

	// test word insert
	tt := trie.NewTrieTree()
	for _, word := range words {
		if err := tt.Insert(word); err != nil {
			t.Fatalf(fmt.Sprintf("unable to push word %s", word))
		}
	}

	// test word search
	randomWord := words[rand.Intn(len(words))]
	found, _, err := tt.Search(randomWord)
	if err != nil {
		t.Errorf("error while trying to find word %s", err.Error())
	} else if !found {
		t.Errorf("unable to find word")
	}

	// test word completions
	prefix := words[rand.Intn(len(words))]
	bruteForceCompletions := []string{}
	for _, word := range words {
		if word != prefix && strings.HasPrefix(word, prefix) {
			bruteForceCompletions = append(bruteForceCompletions, word)
		}
	}

	trieTreeCompletions, err := tt.Completions(prefix)
	if err != nil {
		t.Errorf("unable to get completions %s", err.Error())
		return
	}

	if len(trieTreeCompletions) != len(bruteForceCompletions) {
		t.Errorf("trieTreeCompletions %d bruteForceCompletions %d prefix %s", len(trieTreeCompletions), len(bruteForceCompletions), prefix)
	}
}
