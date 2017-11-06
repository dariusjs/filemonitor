package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "flag"
  "os"
  "encoding/json"
  "time"
  "strings"
  "sync"
)

type Config struct {
  Directories []Directory `json:"directories"`
}

type Directory struct {
  Name string `json:"directory"`
  Count int `json:"count"`
  Age string `json:"age"`
  Frequency int `json:"frequency"`
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

// ListObjects will run something like ls on unix
func ListObjects(config Config) {
  for _, dir := range config.Directories {
    files, err := ioutil.ReadDir(dir.Name)
    if err != nil {
      log.Fatal(err)
    }
    for _, file := range files {
      fmt.Println(file.Mode(), file.ModTime(), file.Size, file.Name())
    }
  }
}

// Timer will watch directories specifcally as file monitors will be separate to this
func Timer(dir Directory) {
  for t := range (time.NewTicker(time.Duration(dir.Frequency)*time.Second).C) {

    fmt.Println("Gday", t)

    splitAge := strings.SplitAfterN(dir.Age, "", 2)
    fmt.Println(splitAge[1])
    fmt.Println(time.Now())

    if (splitAge[0] == ">") {
      fmt.Println("greater than")
      fmt.Println(dir.ErrorMsg)
    } else if (splitAge[0] == "<") {
      fmt.Println("less than")
    } else {
      // implement some error catch here
      fmt.Println("Some error stuff")
    }
  }
}

// Monitor will watch the monitored file system objects 
func Monitor(config Config) {
  for _, dir := range config.Directories {
    fmt.Println(dir)
    go Timer(dir)
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