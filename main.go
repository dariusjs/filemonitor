package main

import "flag"
import "fmt"

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
  }
}