package main

import (
	"encoding/json"
	"fmt"
)

type JSONComp struct{}

func NewJSONComp() *JSONComp {
	return &JSONComp{}
}

func (j *JSONComp) decode(inp chan []byte, out chan map[string]interface{}) {
	for i := range inp {
		var m map[string]interface{}

		if err := json.Unmarshal(i, &m); err != nil {
			fmt.Println(err)
		}
		out <- m
	}
}
