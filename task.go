package main

import "encoding/json"

type Task struct {
  Id string `json:"id"`
  Title string `json:"title"`
  Completed bool `json:"completed"`
}

type TaskList struct {
  Tasks []Task `json:"tasks"`
}

func GetTaskList(body []byte) (*TaskList) {
    var taskList = new(TaskList)
    err := json.Unmarshal(body, &taskList)
    Check(err)

    return taskList
}
