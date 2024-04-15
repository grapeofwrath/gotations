package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "os"
    "strings"
    "log"
)

type Quote struct {
    Text  string
    Author string
}

func main() {
    quotesPath := os.Args[1]
    data, err := os.ReadFile(quotesPath)
    if err != nil {
        log.Fatal("Error during Readfile(): ",err)
    }
    var quotes []Quote
    err = json.NewDecoder(strings.NewReader(string(data))).Decode(&quotes)
    if err != nil {
        log.Fatal("Error during NewDecoder(): ",err)
    }
    var randQuote = quotes[rand.Intn(len(quotes)-0)]
    fmt.Println(randQuote.Text)
    if randQuote.Author != "" {
        fmt.Println(" - " + randQuote.Author)
        os.Exit(0)
    }
}
