package main

import (
	"fmt"
	"io"
	"os"
)

type fileReader struct {}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Please provide argument")
        os.Exit(1)
    }
    fileName := os.Args[1]

    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println("Problem reading file")
        os.Exit(1)
    }

    io.Copy(os.Stdout, file)
}

func (fr fileReader) Read(b []byte) (n int, err error)  {
    byteLength := len(b)

    fmt.Println(string(b))
    fmt.Printf("Just wrote %v bytes", byteLength)

    return byteLength, nil
}
