package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var mut sync.Mutex
var signal = []string{"test"}

func main() {
	websites := []string{
		"https://devparas.com",
		"https://google.com",
		"https://fb.com",
	}
	for _, weblink := range websites {
		go getStatusCode(weblink)
		wg.Add(1)
	}
	wg.Wait()
}
func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	} else {
		mut.Lock()
		signal = append(signal, endpoint)
		for _, each := range signal {
			fmt.Println(each)
		}
		fmt.Println(signal)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
