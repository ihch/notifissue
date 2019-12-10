package main

import (
  "fmt"
  "flag"
  "log"
  "net/http"
  "io/ioutil"
)

func main() {
  var username string = parseArgs()
  fmt.Println("Hello notifissue.")
  printIssues([]string{"hoge", "poyo"})
  body, err := fetchUserInfo(username)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(body))
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

func fetchUserInfo(username string) (body string, err error) {
  var url string = "https://api.github.com/users/" + username
  return fetch(url)
}

func fetch(url string) (body string, err error) {
  var resp *http.Response
  resp, err = http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  var bytes []byte
  bytes, err = ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  body = string(bytes)
  return body, err
}
