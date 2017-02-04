package main

import (
  "fmt"
  "os"
  "time"

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

  var title string

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
          List()
        }
        return nil
      },
    },
    {
      Name:    "add",
      Aliases: []string{"a"},
      Usage:   "add a new task to the list",
      Flags: []cli.Flag{
        cli.StringFlag{
          Name: "title, t",
          Destination: &title,
        },
      },
      Action:  func(c *cli.Context) error {
        Add(title)
        return nil
      },
    },
  }
  app.Run(os.Args)
}
