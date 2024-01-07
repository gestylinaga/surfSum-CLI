package grammar

import (
  "fmt"
  "strings"
  "math/rand"

  "surfSum-CLI/surfWords"
  "surfSum-CLI/latinWords"
)

func sentenceBuilder(latin bool) string {
  // Randomizing word count per sentence
  wordCount := rand.Intn(5) + 7
  var sentence []string

  // Appending first word (capitalized; always a surf word)
  word := surfwords.Words()
  sentence = append(sentence, word)

  // Appending lowercase words for remaining word count
  for i := 1; i <= wordCount ; i++ {
    // Chance to use latin words if latin flag is passed
    if latin && rand.Intn(3) == 1 {
      word = latinwords.Words()
    } else {
      word = surfwords.Words()
    }
    // Random chance to insert commas, every 3 words (except last word)
    if i % 3 == 0  && i != wordCount {
      if rand.Intn(4) == 1 {
        sentence = append(sentence, (strings.ToLower(word) + ","))
      }
    } else {
      sentence = append(sentence, strings.ToLower(word))
    }
  }

  // Finalized sentence
  return fmt.Sprintf(strings.Join(sentence, " ") + ". ")
}

func ParagraphBuilder(latin bool) string {
  // Randomizing number of sentences
  paragraphLength := rand.Intn(6) + 5
  var paragraph []string

  for i := 1; i <= paragraphLength; i++ {
    sentence := sentenceBuilder(latin)
    paragraph = append(paragraph, sentence)
  }

  return strings.Join(paragraph, "")
}

