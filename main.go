package main

import (
  "encoding/json"
  "fmt"
  "flag"
  "log"
  "net/http"
  "io/ioutil"
)

type Color struct {
  Green string
  Red string
  End string
}

var color Color = Color{"\x1b[32;1m", "\x1b[31;1m", "\x1b[m"}

type Issue struct {
  Title string `json:"title"`
  UpdatedAt string `json:"updated_at"`
}

type PullRequest struct {
  Title string `json:"title"`
  UpdatedAt string `json:"updated_at"`
}

type Event struct {
  EventType string `json:"type"`
  Payload struct {
    Action string `json:"action"`
    Issue Issue `json:"issue"`
    PullRequest PullRequest `json:"pull_request"`
  } `json:"payload"`
}

func main() {
  var username string = parseArgs()
  bytes, err := fetchEvents(username)
  if err != nil {
    log.Fatal(err)
  }

  var events []Event
  if err := json.Unmarshal(bytes, &events); err != nil {
    log.Fatal(err)
  }

  fmt.Print("Recent activities\n\n")
  printLine()
  printEvents(events)
  printLine()
}

func parseArgs() string {
  var username = flag.String("u", "", "username")
  flag.Parse()
  return *username
}

func printLine() {
  fmt.Println("+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+")
}

func printEvents(events []Event) {
  for _, event := range events {
    switch event.EventType {
    case "IssuesEvent": printIssue(event)
    case "PullRequestEvent": printPullRequest(event)
    }
  }
}

func printIssue(event Event) {
  var action string
  if event.Payload.Action == "opened" {
    action = color.Green + event.Payload.Action + color.End
  } else if event.Payload.Action == "closed" {
    action = color.Red + event.Payload.Action + color.End
  }
  fmt.Printf(
    "+ updated at %s %s %s: %s\n",
    event.Payload.Issue.UpdatedAt,
    action,
    event.EventType,
    event.Payload.Issue.Title,
  )
}

func printPullRequest(event Event) {
  var action string
  if event.Payload.Action == "opened" {
    action = color.Green + event.Payload.Action + color.End
  } else if event.Payload.Action == "closed" {
    action = color.Red + event.Payload.Action + color.End
  }
  fmt.Printf(
    "+ updated at %s %s %s:  %s\n",
    event.Payload.PullRequest.UpdatedAt,
    action,
    event.EventType,
    event.Payload.PullRequest.Title,
  )
}

func fetch(url string) ([]byte, error) {
  resp, err := http.Get(url)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()

  return ioutil.ReadAll(resp.Body)
}

func fetchEvents(username string) (bytes []byte, err error) {
  var url string = "https://api.github.com/users/" + username + "/events/public"
  return fetch(url)
}
