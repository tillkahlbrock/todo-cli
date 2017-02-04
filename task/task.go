package task

import "encoding/json"

type Task struct {
  Id string `json:"id"`
  Title string `json:"title"`
  Completed bool `json:"completed"`
}

type TaskList struct {
  Tasks []Task `json:"tasks"`
}

func GetTaskList(body []byte) (*TaskList, error) {
    var s = new(TaskList)
    err := json.Unmarshal(body, &s)
    if(err != nil){
        panic(err.Error())
    }
    return s, err
}
