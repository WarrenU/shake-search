// package main serves an endpoint `GET /autocomplete?term=X` that returns the top 25 used words in a `shakespeare-complete.txt` file.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Dictionary to hold 26 SubDictionaries (26 characters in the alphabet)
var Dictionary []*SubDictionary

// SubDictionary holds a map of words (string) and their occurences (int) as values
type SubDictionary struct {
	Words map[string]int
}

// keyVal is used to Sort a resultant map[string]int
type keyVal struct {
	Key string // word
	Val int    // occurrence
}

// getTopWords returns the 25 most popular words matching a prefix
func getTopWords(dict *SubDictionary, prefix string) []string {
	topWords := make([]string, 0, 25)
	var relevantPairs []keyVal

	for word, occurrence := range dict.Words {
		if strings.HasPrefix(word, prefix) {
			relevantPairs = append(relevantPairs, keyVal{word, occurrence})
		}
	}

	sort.Slice(relevantPairs, func(i, j int) bool {
		return relevantPairs[i].Val > relevantPairs[j].Val
	})

	var i int
	for _, kv := range relevantPairs {
		topWords = append(topWords, kv.Key)
		i++
		if i == 25 {
			break
		}
	}

	return topWords
}

// addToSubDict adds a word to a SubDictionaries Words map or increments it occurrence within a specified SubDictionary.
func addToSubDict(word string) {
	subDict := getSubDict(word)
	if val, ok := subDict.Words[word]; ok {
		subDict.Words[word] = val + 1
	} else {
		subDict.Words[word] = 1
	}
}

// getSubDict will return a SubDictionary that contains words that match the first letter in term.
func getSubDict(term string) *SubDictionary {
	asciiLetter := strings.ToLower(term)[0]
	subDict := Dictionary[asciiLetter-97]
	return subDict
}

// StripWord returns a string without punctuation or digits
func StripWord(word string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		fmt.Print(err)
	}
	return strings.ToLower(reg.ReplaceAllString(word, ""))
}

func autocompleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	term := strings.Join(r.URL.Query()["term"], "")
	var body map[string][]string

	if term != "" {
		subDict := getSubDict(term)
		words := getTopWords(subDict, term)
		body = map[string][]string{"words": words, "amount": []string{strconv.Itoa(len(words))}}
	} else {
		body = map[string][]string{"words": []string{""}}
	}

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Println(err)
	}
}

// init reads a txt file and extracts the words into a SubDictionary.Words via the addToSubDict function.
func init() {
	file, err := os.Open("shakespeare-complete.txt")
	if err != nil {
		fmt.Println("File did not load")
	}
	defer file.Close()

	// Init Dictionary with 26 SubDictionaries for each letter in the alphabet.
	for i := 0; i < 26; i++ {
		subDict := new(SubDictionary)
		subDict.Words = map[string]int{}
		Dictionary = append(Dictionary, subDict)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for _, word := range words {
			word = StripWord(word)
			if word == "" {
				continue
			} else {
				addToSubDict(word)
			}
		}
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/autocomplete", autocompleteHandler)
	log.Println("Starting server on Port 9000")
	err := http.ListenAndServe(":9000", mux)
	log.Fatal(err)
}
