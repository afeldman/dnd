package main

import (  
    "fmt"
    "io/ioutil"
    "os"
    "github.com/Jeffail/gabs"
)

func main() {

    raw, err := ioutil.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    s := string(raw)

   jsonParsed, err := gabs.ParseJSON([]byte(s))
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    fmt.Println(jsonParsed.StringIndent("", "  "))
}
