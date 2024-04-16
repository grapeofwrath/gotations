package main

import (
    "fmt"
    "math/rand"
    "os"
    "log"
    "gopkg.in/yaml.v3"
)

type Quote struct {
    Text  string
    Author string
}
type YAMLFile struct {
    Quotes []Quote `yaml:"quotes"`
}

func main() {
    data, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal("Error during os.Open(): ",err)
    }
    defer data.Close()
    var quotes YAMLFile
    err = yaml.NewDecoder(data).Decode(&quotes)
    if err != nil {
        log.Fatal("Error during yaml.NewDecoder(): ",err)
    }
    var randQuote = quotes.Quotes[rand.Intn(len(quotes.Quotes)-0)]
    fmt.Println(randQuote.Text)
    if randQuote.Author != "" {
        fmt.Println(" - " + randQuote.Author)
        os.Exit(0)
    }
}
