package main

import (
  "fmt"
	//  "os"
	"flag"
)

func main() {
  verbose := flag.Bool("verbose", false, "Enable verbose output")
  version := flag.Bool("v", false, "Print current version of this software")

  flag.Parse()

  if *verbose == true {
    fmt.Printf("Running in verbose mode\n")
  }

  fmt.Printf("Version: %v\n", *version)

}
