package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "flag"
  "os"
  "encoding/json"
  "time"
  "sync"
)

type Config struct {
  Directories []Directory `json:"directories"`
}

type Directory struct {
  Name string `json:"directory"`
  Count int `json:"count"`
  Mtime string `json:"mtime"`
  Frequency string `json:"frequency"`
  ErrorMsg string `json:"errormsg"`
}

// LoadConfiguration will load the json config from file into memory
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

func ListObjects(config Config) {
  for _, dir := range config.Directories {

    files, err := ioutil.ReadDir(dir.Name)
    if err != nil {
      log.Fatal(err)
    }
    for _, file := range files {
      targetTime, err := time.ParseDuration(dir.Mtime)
      if err != nil {
        log.Fatal(err)
      }
      timeDiff := file.ModTime().Sub(time.Now())

      if (timeDiff >= targetTime) {
        fmt.Println(file.Mode(), file.ModTime(), file.Size(), file.Name())
      }
    }
  }
}

// Timer will watch directories specifcally as file monitors will be separate to this
func Watcher(dir Directory) {
  // for t := range (time.NewTicker(time.Duration(dir.Frequency)*time.Second).C) {
  scanFreq, err := time.ParseDuration(dir.Frequency)
  if err != nil {
    log.Fatal(err)
  }
  for t := range (time.NewTicker(scanFreq).C) {
    fmt.Println("Gday", t)
  }
}

// Monitor will watch the monitored file system objects
func Monitor(config Config) {
  for _, dir := range config.Directories {
    fmt.Println(dir)
    go Watcher(dir)
  }
}

func main() {
  var wg sync.WaitGroup
  config, _ := LoadConfiguration("config.json")
  configLength := len(config.Directories)

  genConfig := flag.Bool("g", false, "Used for generating config files")
  listObjects := flag.Bool("l", false, "Execute the default config.json config file")
  daemonize := flag.Bool("d", false, "Run the filemonitor")

  flag.Parse()

  if *genConfig == true {
    fmt.Println("Generate Config:", *genConfig)
    GenerateConfig()
  }

  if *listObjects == true {
    fmt.Println("List Config:", *listObjects)
    ListObjects(config)
  }

  if *daemonize == true {
    fmt.Println("Running the Filemonitor")
    wg.Add(configLength)
    Monitor(config)
    wg.Wait()
  }
}