package main

import (
    "fmt"
    "os"
    "flag"
)

func main() {
  flag.Parse()
	dirname := flag.Arg(0)
    if err := os.MkdirAll(dirname, 0755); err != nil {
        fmt.Println(err)
    }
}
