package config

import (
    "encoding/json"
    "io/ioutil"
    "log"
)

type Server struct {
    Port string `json:"port"`
}

type ConfigStruct struct {
    Server Server
}

var (
    Config ConfigStruct
)

func Load() bool {
    file, err := ioutil.ReadFile("./configs/main.json")

    if err != nil {
        log.Fatal("Config File Missing. ", err)
        return false
    }
    err = json.Unmarshal(file, &Config)
    if err != nil {
        log.Fatal("Config Parse Error: ", err)
        return false
    }

    return true
}