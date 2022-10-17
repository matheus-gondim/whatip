package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://ip-api.com/json"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("Status code: %d", resp.StatusCode))
	}
	defer resp.Body.Close()

	bodyInByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var body map[string]interface{}

	if err := json.Unmarshal(bodyInByte, &body); err != nil {
		panic(err)
	}

	fmt.Printf("My IP is: %s\n", body["query"])
}
