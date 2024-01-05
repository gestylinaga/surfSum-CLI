package grammar

import (
  "fmt"
  "strings"
  "math/rand"

  "surfSum-CLI/surfWords"
)

func sentenceBuilder() string{
  // Sentence building (6-9 words each)
  wordCount := rand.Intn(5) + 5
  var sentence []string

  // Appending first word (capitalized)
  word := surfwords.Words()
  sentence = append(sentence, word)

  // Appending lowercase words for remaining word count
  for i := 1; i <= wordCount ; i++ {
    // Random chance to insert commas, every 3 words
    if i % 3 == 0 {
      if rand.Intn(3) == 1 {
        sentence = append(sentence, strings.ToLower(surfwords.Words()) + ",")
      }
    }
    sentence = append(sentence, strings.ToLower(surfwords.Words()))
  }

  // Final print
  return fmt.Sprintf(strings.Join(sentence, " ") + ". ")
}

func ParagraghBuilder() string{
  paragraghLength := rand.Intn(5) + 5
  var paragragh []string
  for i := 1; i <= paragraghLength; i++ {
    sentence := sentenceBuilder()
    paragragh = append(paragragh, sentence)
  }
  return strings.Join(paragragh, "")
}

