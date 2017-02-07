package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
)

func List(serverUrl string) {
  url := fmt.Sprintf("%s/tasks", serverUrl)
  resp, err := http.Get(url)
  Check(err)

  body, err := ioutil.ReadAll(resp.Body)
  Check(err)

  taskList := GetTaskList([]byte(body))

  for _,task := range taskList.Tasks {
    fmt.Println(task.Id)
    fmt.Println(task.Title)
    fmt.Println(task.Completed)
  }
}
