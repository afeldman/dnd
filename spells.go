package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "encoding/json"
)

/*type component struct {
     material bool	`json:"material"`
     materials []string `json:"materials_needed"`
     raw      string  	`json:"raw"`
     somatic  bool    	`json:"somatic"`
     verbal   bool    	`json:"verbal"`
}

type Spell struct {
     casting_t 	  string	`json:"casting_time"`
     classes 	  []string	`json:"classes"`
     components   component	`json:"components"`
     description  string	`json:"description"`
     duration 	  string	`json:"duration"`
     level 	  string	`json:"level"`
     rang 	  string	`json:"range"`
     ritual 	  bool		`json:"ritual"`
     school 	  string	`json:"school"`
     tags 	  []string	`json:"tags"`
     typ 	  string	`json:"type"`
}

func main() {

    file, e := ioutil.ReadFile(os.Args[1])
    if e != nil {
        fmt.Printf("File error: %v\n", e)
        os.Exit(1)
    }
    fmt.Printf("%s\n", string(file))

    var spells []Spell 
    json.Unmarshal(file, &spells)
    fmt.Printf("Results: %v\n", spells[0])
}*/
