package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// vim:sw=4:ts=4

func main() {
	client := http.Client{}

	sendEntry(client)
	for _ = range time.Tick(15 * time.Second) {
		sendEntry(client)
	}
}

func sendEntry(client http.Client) {
	e, err := MakeEntry()
	if err != nil {
		log.Fatalf("MakeEntry: %v", err)
	}

	jsonSampleEntry, err := json.Marshal(&e)
	if err != nil {
		log.Fatalf("Couldn't marshal sampleEntry: %v", err)
	}
	req, err := http.NewRequest("PUT",
		"http://localhost:8081/api/v1/add_entry",
		bytes.NewBuffer(jsonSampleEntry))
	if err != nil {
		log.Fatalf("Failed to create req: %v", err)
	}
	req.Header.Set("X-Panopticon-Token", "larry@theclapp.org")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error Do'ing the request: %v", err)
	}

	if resp.StatusCode != 200 {
		log.Printf("resp.StatusCode expected to be 200, not %v", resp.StatusCode)
	} else {
		if _, ok := resp.Header["Location"]; !ok {
			log.Fatalf("No Location header in the response")
		}
		log.Printf("Response: %v", resp)
	}
}
