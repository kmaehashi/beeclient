package main

import (
  "fmt"
  "os"
  "github.com/sensorbee/sensorbee/client"
)

func main() {
  req, err := client.NewRequester("http://127.0.0.1:15601", "v1")
  if err != nil {
    os.Exit(1)
  }
  res, err := req.Do(client.Get, "/topologies", nil)
  if err != nil {
    os.Exit(2)
  }
  raw, err := res.Body()
  if err != nil {
    os.Exit(3)
  }
  fmt.Print(string(raw))
}
