package main

import (
  "os"
  "fmt"
)

func Check(err error) {
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
