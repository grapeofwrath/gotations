package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "os"
    "log"
    "path/filepath"
    "gopkg.in/yaml.v3"
)

type Quote struct {
    Text  string
    Author string
}

func main() {
    dataExt := filepath.Ext(os.Args[1])
    data, err := os.Open(os.Args[1])
    if err != nil {log.Fatal("Error during os.Open(): ",err)}
    defer data.Close()
    var quotes []Quote
    if dataExt == ".json" {
        err = json.NewDecoder(data).Decode(&quotes)
        if err != nil {log.Fatal("Error during json.NewDecoder(): ",err)}
    }
    if dataExt == ".yaml" {
        err = yaml.NewDecoder(data).Decode(&quotes)
        if err != nil {log.Fatal("Error during yaml.NewDecoder(): ",err)}
    }
    var randQuote = quotes[rand.Intn(len(quotes)-0)]
    fmt.Println(randQuote.Text)
    if randQuote.Author != "" {
        fmt.Println(" - " + randQuote.Author)
        os.Exit(0)
    }
}
