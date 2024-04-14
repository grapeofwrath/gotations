package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "os"
    "strings"
)

type Quotation struct {
    Quote  string
    Author string
}

func main() {
    filePath := os.Args[1]
    data, err := os.ReadFile(filePath)
    if err != nil {
        fmt.Println(err)
        return
    }

    var quotations []Quotation

    err = json.NewDecoder(strings.NewReader(string(data))).Decode(&quotations)
    if err != nil {
        fmt.Println(err)
        return
    }
    var randQuote = quotations[rand.Intn(len(quotations)-0)]

    fmt.Println(randQuote.Quote)
    fmt.Println(" - " + randQuote.Author)
}
