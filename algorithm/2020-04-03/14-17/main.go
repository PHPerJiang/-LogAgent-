package main

import (
	"encoding/json"
	"fmt"
)

type Persion struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonStr := `{"name":"Gopher","age":24}`
	p := &Persion{}
	err := json.Unmarshal([]byte(jsonStr), p)
	if err != nil {
		fmt.Println("unmarshaler failed")
	}
	fmt.Printf("%#v\n", *p)

	encodeP, err := json.Marshal(p)
	if err != nil {
		fmt.Println("marshaler failed")
	}
	fmt.Printf("%#v", string(encodeP))
}
