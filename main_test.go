package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
	"net/http/httptest"
	"net/http"
	"bytes"
)

type MessageStream struct {
	Items []Message `json:"items"`
}

func TestSchema(t *testing.T) {
	file, err := ioutil.ReadFile("./samples.json")
	if err != nil {
		t.Error("File error: %v\n", err)
	}

	var stream MessageStream
	err = json.Unmarshal(file, &stream)
	if err != nil {
		t.Error("JSON error: %v\n", err)
	}

	//fmt.Printf("%v\n", stream.Items)
	for _, message := range stream.Items {
		fmt.Printf("%d: %d\n", message.Provider.Appid, message.Provider.Timestamp)
	}
}

func TestMessageHandler(t *testing.T) {
	file, err := ioutil.ReadFile("./samples.json")
	if err != nil {
		t.Error("File error: %v\n", err)
	}

	var stream MessageStream
	err = json.Unmarshal(file, &stream)
	if err != nil {
		t.Error("JSON error: %v\n", err)
	}

	//fmt.Printf("%v\n", stream.Items)
	for _, message := range stream.Items {
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		body, err := json.Marshal(message)
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest("POST", "/", bytes.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(MessageHandler)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := `Message Received!`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}