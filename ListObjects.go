package main

import (
  "fmt"
  "io/ioutil"
  "log"
)

func ListObjects () {
  config, _ := LoadConfiguration("config.json")

  for _, dir := range config.Directories {
    files, err := ioutil.ReadDir(dir.Object)
    if err != nil {
      log.Fatal(err)
    }
    for _, file := range files {
      fmt.Println(file.Name())
    }
  }
}