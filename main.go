package main

import (
  "fmt"
  "flag"
)

func main() {
  var username string = parseArgs()
  fmt.Println("Hello notifissue.")
  printIssues([]string{"hoge", "poyo"})
}

func parseArgs() string {
  var username = flag.String("u", "", "username")
  flag.Parse()
  return *username
}

func printIssues(issues []string) {
  for _, issue := range issues {
    fmt.Println(issue)
  }
}
