package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)

	b, err := json.Marshal(hitesh) // Encoding producs byte format of data.

	if err != nil {
		panic(err)
	}

	fmt.Println(b)

	var v User
	err = json.Unmarshal(b, &v)
	fmt.Println(v)

	// What if we do not know the data structure of the data stored. Arbitary data can be serialized using the next steps.
	p:= []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err=json.Unmarshal(p, &f)
	if err!=nil{
		panic(err)
	}
	fmt.Println(f)

	g,err:=json.Marshal(f)
	fmt.Println(g)


	// What about the encoding and decoding of json data. 
}
