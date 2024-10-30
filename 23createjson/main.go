package main

import (
	"encoding/json"
	"fmt"
)

type data struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   *int   `json:"age,omitempty"` //if there is no value don't show
}

func main() {
	//encodeJson()
	decodeJson()
}
func encodeJson() {
	age := 14
	record := []data{
		{"paras", "pp@", &age},
		{"shalu", "dg@", nil},
		{"happy", "jh@", nil},
	}
	finalJson, err := json.MarshalIndent(record, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}
func decodeJson() {
	jsondatafromweb := []byte(`
	    [
        {
                "name": "paras",
                "email": "pp@",
                "age": 14
        },
        {
                "name": "shalu",
                "email": "dg@"
        },
        {
                "name": "happy",
                "email": "jh@"
        }
]
	`)
	var newrecord data
	checkValid := json.Valid(jsondatafromweb)
	if checkValid {
		fmt.Println("data is valid")
		json.Unmarshal(jsondatafromweb, &newrecord)
		fmt.Printf("%#v", newrecord)
	} else {
		fmt.Println("Json wasn't valid")
	}
}
