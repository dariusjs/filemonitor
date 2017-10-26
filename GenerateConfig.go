package main

import (
  "encoding/json"
  "fmt"
  "os"
  "strconv"
)

func GenerateConfig() {
  
  dataStr := `{"directories":[{"directory":"/tmp"}]}`
  dataMap := make(map[string]interface{})
  err := json.Unmarshal([]byte(dataStr), &dataMap)

  if err != nil {
    panic(err)
  }

  var oneConfig Config

  oneConfig.Directories = fmt.Sprintf("%s", oneConfig["Directories"])
  // oneConfig.Directories = fmt.Sprintf("%s", dataMap["directory"])
  // oneDirectory.Directory, _ = strconv.Atoi(fmt.Sprintf("%v", dataMap["Directory"]))

  jsonData, err := json.Marshal(oneConfig)

  if err != nil {
    panic(err)
  }
  
  fmt.Println(string(jsonData))

  jsonFile, err := os.Create("./Config1.json")

  if err != nil {
    panic(err)
  }
  defer jsonFile.Close()

  jsonFile.Write(jsonData)
  jsonFile.Close()
  fmt.Println("JSON data written to ", jsonFile.Name())

}
