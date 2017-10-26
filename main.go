package main

import (
  "fmt"
  "io/ioutil"
  "log"
)

func main() {
  config, _ := LoadConfiguration("config.json")
  fmt.Println(config)
  fmt.Println(config.Directories)

  for _, dir := range config.Directories {
  fmt.Println(dir.Directory)

  files, err := ioutil.ReadDir(dir.Directory)
    if err != nil {
      log.Fatal(err)
    }
    for _, file := range files {
      fmt.Println(file.Name())
    }
  }

  GenerateConfig()
}
