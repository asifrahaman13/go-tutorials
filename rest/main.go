package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Define the struct data type.
type Product struct {
	Name        string   `json:"coursename"`
	Password    string   `json:"-"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
}

func EncodeJson() { // Function to encode the product in json format.

	// Define the struct data.
	lcoCourses := []Product{
		{
			"Acer Nitr 5", "acer", 2451, "One of the best affordable laptop", "https://i.pravatar.cc", "electronics", []string{"electronics", "laptop"},
		},
		{
			"Asus Rog", "asus", 2651, "One of the best affordable laptop", "https://i.pravatar.cc", "electronics", []string{"electronics", "laptop"},
		},
	}

	// Convert to json encoded data.
	finalJson, err := json.MarshalIndent(lcoCourses, "", " ")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"id": 1,
		"title": "Fjallraven - Foldsack No. 1 Backpack, Fits 15 Laptops",
		"price": 109.95,
		"description": "Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday",
		"category": "men's clothing",
		"image": "https://fakestoreapi.com/img/81fPKd-2AYL._AC_SL1500_.jpg",
		"rating": {
		  "rate": 3.9,
		  "count": 120
		}
	  }
	`)

	var fakeProduct Product // Define the interface out of the struct.

	checkValid := json.Valid(jsonDataFromWeb) // Check if the json data is a valid json or not.
	if checkValid {
		fmt.Println("Json was valid")

		// Decode the json data. Pass the data and the interface.
		json.Unmarshal(jsonDataFromWeb, &fakeProduct)
		fmt.Printf("%#v\n", fakeProduct)
	} else {
		fmt.Println("Json was not valid")
	}

	// map the json data in key value pairs.
	var myOnlineData map[string]interface{}

	// Decode the json data. Pass the data and the interface.
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key is: %v and the value is: %v\n", key, value)
	}
}

func PerformGetReqest() string {
	const myurl = "https://fakestoreapi.com/products"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code:", response.Status)
	fmt.Println("Content length:", response.ContentLength)

	// Define the response string builder.
	var responseString strings.Builder

	// Read out all the content.
	content, _ := ioutil.ReadAll(response.Body)

	// Next return the byte content.
	byteContent, _ := responseString.Write(content)
	fmt.Println(byteContent)

	// Now you can return the string easily.
	return responseString.String()
}

func PerformPostRequest() string {
	const myurl = "https://fakestoreapi.com/products"
	// Make a new reader out of the response.
	requestBody := strings.NewReader(`
	{
			"title": "test product",
			"price": 13.5,
			"description": "lorem ipsum set",
			"image": "https://i.pravatar.cc",
			"category": "electronic"
		}
   `)

	// Response or error will be thrown.
	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	// Close the body.
	defer response.Body.Close()

	// Read all the body.
	content, _ := ioutil.ReadAll(response.Body)

	// Set the status code.
	fmt.Println("The post request was successful with status code", response.StatusCode)

	// Return the string content.
	return string(content)
}

func main() {
	// fmt.Println(PerformGetReqest())
	// fmt.Println(PerformPostRequest())
	// EncodeJson()
	DecodeJson()
}
