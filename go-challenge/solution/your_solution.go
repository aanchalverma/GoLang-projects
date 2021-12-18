package solution

import (
	"errors"

	"fraugster.com/interview-challenge/trigram"
)

type Solution struct{}

type WordFreqMatch struct {
	WordMap map[string]string
	Freq    int
}

var frequency []WordFreqMatch

func New() *Solution {
	return &Solution{
		// TODO: initialize your store here
	}
}

func (s *Solution) Add(t trigram.Trigram) []WordFreqMatch {
	// TODO: store the trigram in a way that you can later use to generate the text
	/*

	 */
	prefix := t.Word1 + " " + t.Word2
	// fmt.Println(prefix)
	word := t.Word3
	var w WordFreqMatch
	w.WordMap = make(map[string]string)
	w.WordMap[word] = prefix
	w.Freq = 1
	// s1 := t.Word1 + " " + t.Word2 + " " + t.Word3
	// w.freq = 1
	// fmt.Println(w)
	// m := make(map[string]int)
	// m[s1] = w.freq
	// fmt.Println(m)
	frequency = append(frequency, w)
	// fmt.Println(frequency)
	for i, b := range frequency {
		if b.WordMap[prefix] == w.WordMap[prefix] {
			if _, ok := b.WordMap[word]; ok {
				b.Freq = b.Freq + 1
				frequency = append(frequency[:i], frequency[i+1:]...)
				frequency = append(frequency, b)
			}
		}

		// fmt.Println(frequency)
	}
	return frequency
}

func (s *Solution) Generate(numWords int) (string, error) {
	// TODO: generate a new text based on the trigrams that were previously added.

	// 1. Select the first trigram randomly from the previously added ones.
	//    This will give you the first two words of your generated text

	// 2. To generate the missing numWords-2 words:
	//    a) select each next word based on the previous last two words you created
	//    b) if you want you can use the trigram.WeightedRandom() function that is already provided
	//    c) when you have the new word, remember to update the last 2 words now
	//    d) repeat until you generated all words, the result is the concatenation
	//       of all words (with a space in between words).

	return "", errors.New("TODO")
}
