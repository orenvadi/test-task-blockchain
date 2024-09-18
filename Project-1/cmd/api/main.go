package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func getJsonDataFromUrl() {
	resp, err := http.Get("https://cosmos.entangle.fi/cosmos/bank/v1beta1/supply")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func main() {
	getJsonDataFromUrl()
}
