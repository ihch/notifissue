package main

import (
  "fmt"
)

func main() {
  fmt.Println("Hello notifissue.")
  printIssues([]string{"hoge", "poyo"})
}

func printIssues(issues []string) {
  for _, issue := range issues {
    fmt.Println(issue)
  }
}
