package main

import (
  "fmt"
  "os"
  "encoding/json"
)

type Config struct {
  Directories []Directory `json:"directories"`
}

type Directory struct {
  Object string `json:"directory"`
}

func LoadConfiguration(filename string) (Config, error) {
  var config Config
  configFile, err := os.Open(filename)
  defer configFile.Close()
  if err != nil {
    fmt.Println(err.Error())
  }
  jsonParser := json.NewDecoder(configFile)
  jsonParser.Decode(&config)
  return config, err
}