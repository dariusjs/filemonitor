package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "flag"
  "os"
  "encoding/json"
  "time"
)

type Config struct {
  Directories []Directory `json:"directories"`
}

type Directory struct {
  Name string `json:"directory"`
  Count int `json:"count"`
  Find string `json:"find"`
  Frequency int `json:"frequency"`
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

func ListObjects () {
  config, _ := LoadConfiguration("config.json")

  for _, dir := range config.Directories {
    files, err := ioutil.ReadDir(dir.Name)
    if err != nil {
      log.Fatal(err)
    }
    for _, file := range files {
      fmt.Println(file.Name())
    }
  }
}

func Monitor() {
  for t := range time.NewTicker(2 * time.Second).C {
    fmt.Println("Gday", t)
  }
}

func main() {
  //loadConf := flag.String("c", "config.json", "Used for loading config files.")
  genConfig := flag.Bool("g", false, "Used for generating config files")
  listObjects := flag.Bool("l", false, "Execute the default config.json config file")
  daemonize := flag.Bool("d", false, "Daemonise the filemonitor")

  flag.Parse()

  if *genConfig == true {
    fmt.Println("Generate Config:", *genConfig)
    GenerateConfig()
  }

  if *listObjects == true {
    fmt.Println("List Config:", *listObjects)
    ListObjects()
  }

  if *daemonize == true {
    fmt.Println("Daemonising the Filemonitor")
    Monitor()
  }
}