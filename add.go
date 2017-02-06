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

  resp, err := http.Post(urlString, "text/plain", nil)
  Check(err)

  body, err := ioutil.ReadAll(resp.Body)
  Check(err)

  fmt.Println(string(body))
}
