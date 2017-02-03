package main

import (
  "fmt"
  "os"
  "time"
  "net/http"
  "io/ioutil"
  "encoding/json"

  "github.com/urfave/cli"
)

type Task struct {
  Id string `json:"id"`
  Title string `json:"title"`
  Completed bool `json:"completed"`
}

type TaskList struct {
  Tasks []Task `json:"tasks"`
}

func getTaskList(body []byte) (*TaskList, error) {
    var s = new(TaskList)
    err := json.Unmarshal(body, &s)
    if(err != nil){
        fmt.Println("whoops:", err)
    }
    return s, err
}

func main() {
  app := cli.NewApp()
  app.Name = "todo"
  app.Usage = "brain extension for forgetful people"
  app.Version = "0.0.1"
  app.Compiled = time.Now()
  app.Authors = []cli.Author{
    cli.Author{
      Name:  "Till Kahlbrock",
      Email: "till.kahlbrock@gmail.com",
    },
  }

  app.Commands = []cli.Command{
    {
      Name:    "list",
      Aliases: []string{"l"},
      Usage:   "list all open tasks",
      Flags: []cli.Flag{
        cli.BoolFlag{Name: "completed, c"},
      },
      Action:  func(c *cli.Context) error {
        if (c.Bool("completed")) {
          fmt.Println("not implemented yet")
        } else {
          resp, _ := http.Get("http://localhost:9884/tasks")
          body, _ := ioutil.ReadAll(resp.Body)

          taskList, err := getTaskList([]byte(body))
          if err != nil {
            panic(err.Error())
          }
          for _,task := range taskList.Tasks {
            fmt.Println(task.Id)
            fmt.Println(task.Title)
            fmt.Println(task.Completed)
          }
        }
        return nil
      },
    },
    {
      Name:    "add",
      Aliases: []string{"l"},
      Usage:   "add a new task to the list",
      Action:  func(c *cli.Context) error {
        fmt.Println("Hello blurb")
        return nil
      },
    },
  }
  app.Run(os.Args)
}
