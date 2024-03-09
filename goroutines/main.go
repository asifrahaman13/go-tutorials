package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals=[]string{"test"}
var wg sync.WaitGroup // Protects
var mut sync.Mutex    // Pointers

func greet(s string) {
	for i := 0; i < len(s); i++ {
		time.Sleep(1)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in the endpoint")
	}else{
		mut.Lock()
		signals=append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
	}
	fmt.Printf("%d 200 status code for the website %s.\n", result.StatusCode, endpoint)
}


func main() {
	// Iniotialize the goroutines using go keyword.
	// go greet("Hello world with goroutine.")
	// greet("Hello world without goroutine.")

	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.dev",
		"https://github.com",
		"https://facebook.com",
	}
	for index := 0; index < len(websitelist); index++ {
		go getStatusCode(websitelist[index])
		wg.Add(1)
	}

	wg.Wait()
}
