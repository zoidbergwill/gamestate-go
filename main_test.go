package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type MessageStream struct {
	Items []Message `json:"items"`
}

func TestSchema(t *testing.T) {
	file, err := ioutil.ReadFile("./samples.json.stream")
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
