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
  fmt.Println(string(data_json))
  
  jsonFile, err := os.Create("./Config1.json")
  if err != nil {
    panic(err)
  }

  defer jsonFile.Close()

  jsonFile.Write(data_json)
  jsonFile.Close()
  fmt.Println("JSON data written to ", jsonFile.Name())

}
