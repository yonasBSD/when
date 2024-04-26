package main

import (
  "time"
  "fmt"
  "os"
  "strings"

  "github.com/olebedev/when"
  "github.com/olebedev/when/rules/en"
  "github.com/olebedev/when/rules/common"
)

func main() {
  w := when.New(nil)
  w.Add(en.All...)
  w.Add(common.All...)

  text := strings.Join(os.Args[1:], " ")
  r, err := w.Parse(text, time.Now())
  if err != nil {
    // an error has occurred
    fmt.Println(err)
    os.Exit(1)
  }
  if  r == nil {
    // no matches found
    fmt.Println("No matches found")
    os.Exit(1)
  }

  fmt.Println(
    "the time",
    r.Time.String(),
    "mentioned in",
    text[r.Index:r.Index+len(r.Text)],
  )
}
