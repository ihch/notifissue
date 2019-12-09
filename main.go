package main

import (
  "fmt"
  "flag"
  "net/http"
)

func main() {
  var username string = parseArgs()
  fmt.Println("Hello notifissue.")
  printIssues([]string{"hoge", "poyo"})
  fmt.Println(fetchUserInfo(username))
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

func fetchUserInfo(username string) (resp *http.Response, err error) {
  var url string = "https://api.github.com/users/" + username
  return http.Get(url)
}
