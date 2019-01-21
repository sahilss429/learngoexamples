package main

import (
  "fmt"
  "os"
  "encoding/json"
  "io/ioutil"
)

type shipObject struct {
  Name string
  Source string
}

func main() {
  filePath := "./pattern.json";
  fmt.Printf( "// reading file %s\n", filePath )
  file, err1 := ioutil.ReadFile( filePath )
  if err1 != nil {
      fmt.Printf( "// error while reading file %s\n", filePath )
      fmt.Printf("File error: %v\n", err1)
      os.Exit(1)
  }

  fmt.Println( "// defining array of struct shipObject" )
  var ships []shipObject

  err2 := json.Unmarshal(file, &ships)
  if err2 != nil {
    fmt.Println("error:", err2)
    os.Exit(1)
  }

  fmt.Println( "// loop over array of structs of shipObject" )
  for k := range ships {
    fmt.Printf( "The ship '%s' first appeared on '%s'\n", ships[k].Name, ships[k].Source );
  }
}
