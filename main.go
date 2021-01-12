package main

import (
    "fmt"

    "md2html/cmd"
)

func main() {
    err := cmd.RootCmd.Execute()
    if err != nil {
        fmt.Printf("err: %s\n", err)
    }
}
