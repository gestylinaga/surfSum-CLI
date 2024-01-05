package main

import (
	"fmt"
  "flag"
	"math/rand"
	"strings"

	"surfSum-CLI/surfWords"
)

func main() {
  // Shaka Brah! starting text bool flag
  startText := flag.Bool("sb", false, "start with \"Shaka Brah!\"")
  flag.Parse()
  if *startText {
    fmt.Printf("Shaka Brah! ")
  }

  // Sentence building (6-9 words each)
  wordCount := rand.Intn(5) + 5
  var sentence []string

  // Appending first word (capitalized)
  word := surfwords.Words()
  sentence = append(sentence, word)

  // Appending lowercase words for remaining word count
  for i := 1; i <= wordCount ; i++ {
    sentence = append(sentence, strings.ToLower(surfwords.Words()))
  }

  // Final print
  fmt.Println(strings.Join(sentence, " ") + ".")
}
