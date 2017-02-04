package main

import (
  "fmt"
  "os"
  "time"

  "github.com/tillkahlbrock/todo/commands/list"
  "github.com/urfave/cli"
)

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
          list.Run()
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
