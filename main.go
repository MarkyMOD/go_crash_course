// word_count102.go
//
// separate and process the words in a given text
// to form a map of word:frequency pairs, use a structure
// to show the result sorted by frequency and word
//
// for detailed info see:
// http://golang.org/pkg/fmt/
// http://golang.org/pkg/regexp/
// http://golang.org/pkg/sort/
// http://golang.org/pkg/strings/
//
// online play at:
// https://play.golang.org/p/E-5p3ZJfs-
//
// tested with Go version 1.4.2   by vegaseat  28may2015
package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func word_count(words []string) map[string]int {
	// iterate over the slice of words and populate
	// the map with word:frequency pairs
	// (where word is the key and frequency is the value)
	word_freq := make(map[string]int)
	// range string slice gives index, word pairs
	// index is not needed, so use blank identifier _
	for _, word := range words {
		// check if word (the key) is already in the map
		_, ok := word_freq[word]
		// if true add 1 to frequency (value of map)
		// else start frequency at 1
		if ok == true {
			word_freq[word] += 1
		} else {
			word_freq[word] = 1
		}
	}
	return word_freq
}

// // nicely simplified algo
// // takes a slice of strings
// // returns a word:frequency map
// func word_count2 (words []string) map[string]int{
//     word_freq := make(map[string]int)
//     for _, word := range words{
//         word_freq[word]++
//     }
//     return word_freq
// }

type word_struct struct {
	freq int
	word string
}

// word_struct will be displayed in this format
func (p word_struct) String() string {
	return fmt.Sprintf("%3d   %s", p.freq, p.word)
}

// by_freq implements sort.Interface for []word_struct based on the freq field
type by_freq []word_struct

func (a by_freq) Len() int           { return len(a) }
func (a by_freq) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_freq) Less(i, j int) bool { return a[i].freq < a[j].freq }

// by_word implements sort.Interface for []word_struct based on the word field
type by_word []word_struct

func (a by_word) Len() int           { return len(a) }
func (a by_word) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a by_word) Less(i, j int) bool { return a[i].word < a[j].word }
func main() {

	// in Go a multiline text is enclosed in back ticks `
	text := `Boy in a letter to Santa: "Send me a baby brother."
Santa in a letter to boy: "Send me your mother."`
	// change text to all lower characters
	text = strings.ToLower(text)
	// match whole words, removes any punctuation/whitespace
	re := regexp.MustCompile("\\w+")
	words := re.FindAllString(text, -1)
	// word_map() returns a map of word:frequency pairs
	word_map := word_count(words)

	// convert the map to a slice of structures for sorting
	// create pointer to an instance of word_struct
	pws := new(word_struct)
	struct_slice := make([]word_struct, len(word_map))
	ix := 0
	for key, val := range word_map {
		pws.freq = val
		pws.word = key
		// test, %+v shows field names
		//fmt.Printf("%v %v  %+v\n", pws.freq, pws.word, *pws)
		struct_slice[ix] = *pws
		ix++
	}
	// testing ...
	//fmt.Printf("%+v\n", struct_slice[0])
	//fmt.Printf("%+v\n", struct_slice[1])
	//fmt.Printf("%v\n", struct_slice[1].freq)
	//fmt.Printf("%v\n", struct_slice[1].word)
	//fmt.Println("-------------")
	fmt.Println("Words in text sorted by frequency:")
	// sorting slice of structers by field freq in place
	sort.Sort(by_freq(struct_slice))
	for ix := 0; ix < len(struct_slice); ix++ {
		fmt.Printf("%v\n", struct_slice[ix])
	}
	fmt.Println("-------------")
	fmt.Println("Words in text sorted by word:")
	// sorting slice of structures by field word in place
	sort.Sort(by_word(struct_slice))
	for ix := 0; ix < len(struct_slice); ix++ {
		fmt.Printf("%v\n", struct_slice[ix])
	}
}

/*
Words in text sorted by frequency:
  1   brother
  1   your
  1   baby
  1   mother
  2   to
  2   send
  2   me
  2   santa
  2   letter
  2   boy
  2   in
  3   a
-------------
Words in text sorted by word:
  3   a
  1   baby
  2   boy
  1   brother
  2   in
  2   letter
  2   me
  1   mother
  2   santa
  2   send
  2   to
  1   your
*/
