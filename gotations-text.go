package main

import (
    "bufio"
    "fmt"
    "os"
    "math/rand"
)

func main() {
    filePath := os.Args[1]
    readFile, err := os.Open(filePath)

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string

    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }

    readFile.Close()

    var randLine = rand.Intn(len(fileLines)-0)

    fmt.Println(fileLines[randLine])
}
