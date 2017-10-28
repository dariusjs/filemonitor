package main

import (
  "encoding/json"
  "fmt"
  "os"
)

func GenerateConfig() {

  data := Config {
    []Directory {
      Directory {
        "/tmp",
      },
    },
  }
  
  data_json, _ := json.Marshal(data)
  
  jsonFile, err := os.Create("./config.json.sample")
  if err != nil {
    panic(err)
  }
  defer jsonFile.Close()

  jsonFile.Write(data_json)
  jsonFile.Close()
  fmt.Println("JSON data written to ", jsonFile.Name())

}
