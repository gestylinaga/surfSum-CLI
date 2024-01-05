package main

import (
	"fmt"
  "flag"

	"surfSum-CLI/grammar"
)

func main() {
  startText := flag.Bool("sb", false, "start with \"Shaka Brah!\"")
  paragraghNum := flag.Int("p", 1, "number of paragraghs (max 10)")
  flag.Parse()

  if *startText {
    fmt.Printf("Shaka Brah! ")
  }

  // Max number of paragraghs returned
  if *paragraghNum > 10 {
    *paragraghNum = 10
  }

  // Print suggestions if no flags passed
  if !*startText && *paragraghNum == 1 {
    fmt.Println("Hey kook! Try passing some flags! --help/-h to see options.")
    fmt.Printf("\n")
  }

  for i := 1; i <= *paragraghNum; i++ {
    paragragh := grammar.ParagraghBuilder()
    fmt.Println(paragragh)
    fmt.Printf("\n")
  }
}
