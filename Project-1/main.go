package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SupplyResponse struct {
	Amount string `json:"amount"`
}

func getTotalSupply(w http.ResponseWriter, r *http.Request, url string) {
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Error fetching data from external service", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "External service returned non-OK status", http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response body", http.StatusInternalServerError)
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		http.Error(w, "Error parsing JSON response", http.StatusInternalServerError)
		return
	}

	// extract the total number of NGL tokens (denom = "aNGL")
	supplyArray := result["supply"].([]interface{})
	nglData := supplyArray[0].(map[string]interface{})
	amount := nglData["amount"].(string)

	// prepare the response
	supplyResponse := SupplyResponse{Amount: amount}
	jsonResponse, err := json.Marshal(supplyResponse)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	http.HandleFunc("GET /getTotalSupply", func(w http.ResponseWriter, r *http.Request) {
		getTotalSupply(w, r, config.Url)
	})

	log.Printf("Starting server on port %s...", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Port), nil))
}
