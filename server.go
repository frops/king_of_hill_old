package main

import (
    "os"
    "encoding/json"
    "./utils/config"
    "fmt"
)

func main() {
    fmt.Println(config)
}

func loadConfig() {
    file, _ := os.Open("conf.json")
    decoder := json.NewDecoder(file)

    err := decoder.Decode(&configuration)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Println(configuration.Users)

}