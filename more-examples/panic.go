package main

import "os"

func main() {

    panic("error here")

    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}
