package main

import (
    "fmt"
    "net/http"
    "net/url"
    "io/ioutil"
)

func Add(title, serverUrl string) {
  url := fmt.Sprintf("%s/task?title=%s", serverUrl, url.QueryEscape(title))
  resp, err := http.Post(url, "text/plain", nil)
  Check(err)

  body, err := ioutil.ReadAll(resp.Body)
  Check(err)

  fmt.Println(string(body))
}
