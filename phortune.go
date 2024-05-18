package main

import (
    "fmt"
    "log"
    "math/rand"
    "os"
    "path/filepath"
    "strings"
    "gopkg.in/yaml.v3"
)

type Config struct {
    Fortune bool `yaml:"fortune"`
}
type Phortune struct {
    Text  string
    Source string
}

func main() {
    var err error
    var cfg Config
    var phData []byte
    var phRand string
    var ph []string
    var phPath string
    fortune := true

    configDir, err := os.UserConfigDir()
    if err != nil {
        log.Printf("ERROR: %v",err)
    }
    configCheck, err := filepath.Glob(filepath.Join(configDir, "phortune", "config.yaml"))
    if err != nil {
        log.Printf("ERROR: %v",err)
    }
    if len(configCheck) > 0 {
        cfgData, err := os.ReadFile(configCheck[0])
        if err != nil {
            log.Printf("ERROR: %v",err)
        }
        err = yaml.Unmarshal(cfgData, &cfg)
        if err != nil {
            log.Printf("ERROR: %v",err)
        }
        var nil bool
        if cfg.Fortune != nil {
            fortune = cfg.Fortune
        }
    }

    phPath = filepath.Join(configDir, "phortune", "phortunes")
    if len(os.Args) > 1 {
        phPath = os.Args[1]
    }
    phPathCheck, err := filepath.Glob(phPath)
    if err != nil {
        log.Printf("ERROR: %v",err)
    }
    if phPathCheck == nil {
        log.Println("ERROR: missing or incorrect quotes filepath")
        os.Exit(1)
    }
    phData, err = os.ReadFile(phPath)
    if err != nil {
        log.Printf("ERROR: %v",err)
    }

    if fortune == false {
        var phRand Phortune
        var ph []Phortune
        err = yaml.Unmarshal(phData, &ph)
        if err != nil {
            log.Printf("ERROR: %v",err)
        }
        phRand = ph[rand.Intn(len(ph)-0)]
        fmt.Println(phRand.Text)
        if phRand.Source != "" {
            fmt.Println(" - " + phRand.Source)
        }
        os.Exit(0)
    }
    sData := string(phData)
    ph = strings.Split(sData, "%")
    phRand = ph[rand.Intn(len(ph)-0)]
    fmt.Println(strings.TrimSpace(phRand))
}
