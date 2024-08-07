package main

import (
	"encoding/json"
	"fmt"
)

type JSONComp struct{}

func NewJSONComp() *JSONComp {
	return &JSONComp{}
}

func (j *JSONComp) decode(data []byte) map[string]interface{} {
	var m map[string]interface{}

	if err := json.Unmarshal(data, &m); err != nil {
		panic(err)
	}
	fmt.Println(m)

	return m
}
