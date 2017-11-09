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
        "<5",
        "-60s",
        "5s",
        "Look at http://help.me/whatsthis on how to fix this",
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
