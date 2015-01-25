package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

// vim:sw=4:ts=4

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: sensor your_email_address")
	}
	user := os.Args[1]
	client := http.Client{}

	sendEntry(client, user)
	for _ = range time.Tick(15 * time.Second) {
		sendEntry(client, user)
	}
}

func sendEntry(client http.Client, user string) {
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
	req.Header.Set("X-Panopticon-Token", user)

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
