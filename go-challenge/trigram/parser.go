package trigram

import (
	"bufio"
	// "errors"
	"fmt"
	// "github.com/jdkato/prose/tag"
	// "github.com/jdkato/prose/tokenize"
	"io"
	"os"
	"strings"
	// "github.com/jdkato/prose/summarize"
	"github.com/jdkato/prose/v2"
)

// ParseFile parses all trigrams from a given file.
func ParseFile(path string) ([]Trigram, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()
	return Parse(f)
}

// Parse reads trigrams from an io.Reader.
func Parse(r io.Reader) ([]Trigram, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	// var words []string
	// for scanner.Scan() {
	// 	word := strings.TrimSpace(scanner.Text())
	// 	words = append(words, word)
	// }
	for scanner.Scan() {
		doc, _ := prose.NewDocument(strings.Join([]string(scanner.Text()), " "))
		// fmt.Println(doc)
		sents := doc.Sentences()
		fmt.Println(len(sents)) // 2
		for _, sent := range sents {
			fmt.Println(sent.Text)
			// I can see Mt. Fuji from here.
			// St. Michael's Church is on 5th st. near the light.
		}
	}

	// if err := scanner.Err(); err != nil {
	// 	return nil, err
	// }

	// if len(words) < 3 {
	// 	return nil, errors.New("input has less than 3 words")
	// }
	// for _, word := range words {
	// 	words = tokenize.NewTreebankWordTokenizer().Tokenize(word)

	// 	tagger := tag.NewPerceptronTagger()
	// }

	// for _, tok := range tagger.Tag(words) {
	// 	fmt.Println(tok.Text, tok.Tag)
	// }

	var trigrams []Trigram
	// for i := 0; i < len(words)-2; i++ {
	// 	trigram := Trigram{words[i], words[i+1], words[i+2]}
	// 	trigrams = append(trigrams, trigram)
	// }

	return trigrams, nil
}
