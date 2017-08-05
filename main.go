package main

import (
       "fmt"
       "os"
       "dnd/spells"
)

func main() {

     spell := spells.From_JSON(os.Args[1])

     for _, v := range spell {
	fmt.Println(v.Name)
     }
}
