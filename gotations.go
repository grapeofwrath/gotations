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
    if len(os.Args) < 2 {
        log.Fatal("Error during runtime: expected JSON or YAML quotes filepath")
    }
    dataString, err := os.ReadFile(os.Args[1])
        if err != nil {log.Fatal(err)}
    validJson := json.Valid([]byte(dataString))
    data, err := os.Open(os.Args[1])
        if err != nil {log.Fatal(err)}
    defer data.Close()
    var quotes []Quote
    if validJson {
        err = json.NewDecoder(data).Decode(&quotes)
            if err != nil {log.Fatal(err)}
    } else {
        err = yaml.NewDecoder(data).Decode(&quotes)
            if err != nil {log.Fatal(err)}
    }
    var randQuote = quotes[rand.Intn(len(quotes)-0)]
    fmt.Println(randQuote.Text)
    if randQuote.Author != "" {
        fmt.Println(" - " + randQuote.Author)
    }
    os.Exit(0)
}
