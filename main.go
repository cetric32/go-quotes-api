package main

import (
	"io"
	"log"
	"net/http"
)

func getQuotesHandle(resp http.ResponseWriter, req *http.Request) {
	// check if the request method is GET
	if req.Method != "GET" {
		http.Error(resp, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	const url = "https://type.fit/api/quotes"

	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error making GET request:", err)
	}

	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	// Set the response content type
	resp.Header().Set("Content-Type", "application/json")
	// Set the response status code
	resp.WriteHeader(http.StatusOK)
	// Write the response body to the client
	resp.Write(body)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	},
	)

	http.HandleFunc("/quotes", getQuotesHandle)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
