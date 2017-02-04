package list

import (
    "fmt"
    "net/http"
    "io/ioutil"

    t "github.com/tillkahlbrock/todo/task"
)

func Run() {
  resp, _ := http.Get("http://localhost:9884/tasks")
  body, _ := ioutil.ReadAll(resp.Body)

  taskList, err := t.GetTaskList([]byte(body))
  if err != nil {
    panic(err.Error())
  }
  for _,task := range taskList.Tasks {
    fmt.Println(task.Id)
    fmt.Println(task.Title)
    fmt.Println(task.Completed)
  }
}
