package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Errorf("error: Cannot GET access")
	}

	Client(resp)
}

func Client(resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("error: Cannot read response body")
	}
	defer resp.Body.Close()
	fmt.Printf("%s", body)
}
