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
	for i := 0; i < 1000; i++ {
		randomWord := words[rand.Intn(len(words))]

		t.Run(fmt.Sprintf("test %s", randomWord), func(t *testing.T) {
			found, err := tt.Search(randomWord)
			if err != nil {
				t.Errorf("error while trying to find word %s", err.Error())
				return
			} else if !found {
				t.Errorf("unable to find word")
				return
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
			if err != nil && err != trie.ErrNoCompletions {
				t.Errorf("unable to get completions %s", err.Error())
				return
			}

			if len(trieTreeCompletions) != len(bruteForceCompletions) {
				t.Errorf("trieTreeCompletions %d bruteForceCompletions %d prefix %s", len(trieTreeCompletions), len(bruteForceCompletions), prefix)
				return
			}
		})
	}
}
