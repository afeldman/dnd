package main

import (  
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type component struct {
     material bool
     raw string
     somatic bool
     verbal bool
}

type Spell struct {
     casting_time string
     classes []string
     components component
     description string
     duration string
     level string
     rang string `json:"range"`
     ritual bool
     school string
     tags []string
     typ string `json:"type"`
}

func (s Spell) toString() string {
    return toJson(s)
}

func toJson(s interface{}) string {
    bytes, err := json.Marshal(s)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    return string(bytes)
}

func main() {

    spells := getSpells(os.Args[1])
    for _, s := range spells {
        fmt.Println(s.toString())
    }
    fmt.Println(toJson(spells))	
}

func getSpells(filename string) []Spell {
    raw, err := ioutil.ReadFile(filename)

    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var c []Spell
    json.Unmarshal(raw, &c)

        fmt.Println(c)	

    return c
}
