package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
	- Partindo do código abaixo, utilize unmarshal e demonstre os valores.
*/


type Person struct {
    First string   `json:"First"`
    Last  string   `json:"Last"`
    Age   int      `json:"Age"`
    Sayings []string `json:"Sayings"`
}

var people []Person

func main() {
    jsonStr := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
    err := json.Unmarshal([]byte(jsonStr), &people)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v", people)
	fmt.Println(people[1].Last)
}