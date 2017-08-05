package main

import (
    "log"
    "os"
    "io/ioutil"
    "fmt"
    "encoding/json"
    "sort"
)

type Component struct {
     Material bool	`json:"material,omitempty"`
     Materials []string `json:"materials_needed,omitempty"`
     Raw      string  	`json:"raw,omitempty"`
     Somatic  bool    	`json:"somatic,omitempty"`
     Verbal   bool    	`json:"verbal,omitempty"`
}

type Spell struct {
     Casting_t 	  string	`json:"casting_time,omitempty"`
     Classes 	  []string	`json:"classes,omitempty"`
     Components   Component	`json:"components,omitempty"`
     Description  string	`json:"description,omitempty"`
     Duration 	  string	`json:"duration,omitempty"`
     Level 	  string	`json:"level,omitempty"`
     Range 	  string	`json:"range,omitempty"`
     Ritual 	  bool		`json:"ritual,omitempty"`
     School 	  string	`json:"school,omitempty"`
     Tags 	  []string	`json:"tags,omitempty"`
     Type 	  string	`json:"type,omitempty"`
     Name 	string		`json:"name,omitempty"`
}

type Spells []Spell

func (slice Spells) Len() int {
	return len(slice)
}

func (slice Spells) Less(i, j int) bool {
	return slice[i].Name < slice[j].Name;
}

func (slice Spells) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}


func (s Spells) save_JSON(path string){

     if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(path, 0755)
		} else {
			log.Println(err)
		}
	}

	b, err := json.Marshal(s)
	if err != nil { log.Println(err) }

	ioutil.WriteFile(path, b, 0644) 
}

func main() {

     file, err := ioutil.ReadFile(os.Args[1])
     if err != nil {
             log.Fatal(err)
     }

     var spells Spells
     err = json.Unmarshal(file, &spells)
     if err != nil {
             log.Fatal(err)
     }
    
    sort.Sort(spells)
/*    fmt.Println("\nSorted")
    for i, c := range spells {
        fmt.Println(i, c.Name)
    }*/

    b, _ := json.Marshal(spells)
    // Convert bytes to string.
    s := string(b)
    fmt.Println(s)
}
