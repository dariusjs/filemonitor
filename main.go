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
    i := 0
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

      if(timeDiff >= targetTime) {
        // fmt.Println(file.Mode(), file.ModTime(), file.Size(), file.Name())
        i += 1
      }
    }
    fmt.Println("Total:", i)
    if (i > dir.Count) {
      fmt.Println("(╯°□°）╯︵ ┻━┻)")
    }
  }
}

func ListObjects2(dir Directory, config Config) {
  i := 0
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

    if(timeDiff >= targetTime) {
      // fmt.Println(file.Mode(), file.ModTime(), file.Size(), file.Name())
      i += 1
    }
  }
  fmt.Println("Total:", i)
  if (i > dir.Count) {
    fmt.Println("(╯°□°）╯︵ ┻━┻)")
  }
}

// Timer will watch directories specifcally as file monitors will be separate to this
func Watcher(dir Directory, config Config) {
  // for t := range (time.NewTicker(time.Duration(dir.Frequency)*time.Second).C) {
  scanFreq, err := time.ParseDuration(dir.Frequency)
  if err != nil {
    log.Fatal(err)
  }
  for t := range (time.NewTicker(scanFreq).C) {
    fmt.Println("Gday", t)
    ListObjects2(dir, config)
  }
}

// Monitor will watch the monitored file system objects
func Monitor(config Config) {
  for _, dir := range config.Directories {
    fmt.Println(dir)
    go Watcher(dir, config)
  }
}

func main() {
  var wg001 sync.WaitGroup
  var wg002 sync.WaitGroup
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
    // ListObjects(config)
    wg002.Add(configLength)
    for _, dir := range config.Directories {
      fmt.Println(dir)
      go ListObjects2(dir, config)
    }
    // wg002.Wait()
    time.Sleep(2 * time.Second)
  }
  
  if *daemonize == true {
    fmt.Println("Running the Filemonitor")
    wg001.Add(configLength)
    Monitor(config)
    wg001.Wait()
  }
}