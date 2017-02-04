package main

import (
    "fmt"
    "net/http"
    "net/url"
    "io/ioutil"
    "bytes"
)

func Add(title string) {
  var buffer bytes.Buffer
  buffer.WriteString("http://localhost:9884/task?title=")
  buffer.WriteString(url.QueryEscape(title))
  urlString := buffer.String()

  resp, _ := http.Post(urlString, "text/plain", nil)
  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
    panic(err.Error)
  }

  fmt.Println(string(body))
}
