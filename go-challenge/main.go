package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"fraugster.com/interview-challenge/solution"
	"fraugster.com/interview-challenge/trigram"
)

const usage = `
Usage: %s [corpus.txt] [n]

This program parses all words from a given text corpus to randomly generate a
new text that is looking similar to the provided text.
`

type Solution interface {
	Add(trigram.Trigram) []solution.WordFreqMatch
	Generate(n int) (string, error)
}

func main() {
	// Initialize the RNG so we get a different output on each run.
	rand.Seed(time.Now().UnixNano())

	corpusFile, n := parseFlags()
	fmt.Println("n value is", n)
	trigrams, err := trigram.ParseFile(corpusFile)
	fmt.Println(trigrams)
	if err != nil {
		log.Fatal(err)
	}

	// This is where you should create a new instance of your solution.
	// It must implement the trigram.Solution interface.
	var s Solution = solution.New()
	// var frequency []solution.WordFreqMatch
	// We pass all parsed trigrams to your solution so it can learn its model
	// based on the provided corpus of words.
	// for _, t := range trigrams {
	// 	frequency = s.Add(t)
	// }
	fmt.Println("DONE WITH TRIAGRAM========================================================")
	// for _, freq := range frequency {
	// 	fmt.Println("map: word", freq.WordMap)
	// 	// fmt.Println("map: prefix", freq.wordMap[word])
	// 	fmt.Println(freq.Freq)
	// }

	// We now want to generate a new text based on the trigrams that we added above.
	result, err := s.Generate(n)
	if err != nil {
		log.Fatal(err)
	}

	// Finally we print the result to stdout.
	fmt.Println(result)
}

func parseFlags() (corpusFile string, n int) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, usage, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	args := flag.Args()

	corpusFile = "corpus/test.txt"
	if len(args) > 0 {
		corpusFile = args[0]
	}

	n = 100
	if len(args) > 1 {
		var err error
		n, err = strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("second argument is not an integer: %v", err)
		}
	}

	return corpusFile, n
}
