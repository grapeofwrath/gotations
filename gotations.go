package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "os"
    "gopkg.in/yaml.v3"

    "github.com/charmbracelet/log"
)

type Quote struct {
    Text  string
    Author string
}

func main() {
    if len(os.Args) < 2 {log.Fatal("Error missing args: expected JSON or YAML quotes filepath")}
    data, err := os.ReadFile(os.Args[1])
        if err != nil {log.Fatal(err)}
    var validJson = json.Valid([]byte(data))
    var quotes []Quote
    if validJson {
        err = json.Unmarshal(data, &quotes)
            if err != nil {log.Fatal(err)}
    } else {
        err = yaml.Unmarshal(data, &quotes)
            if err != nil {log.Fatal(err)}
    }
    var randQuote = quotes[rand.Intn(len(quotes)-0)]
    fmt.Println(randQuote.Text)
    if randQuote.Author != "" {
        fmt.Println(" - " + randQuote.Author)
    }
}
