package main

import (
	"fmt"
	"testing"
)

func TestJSONdecode(t *testing.T) {
	fmt.Println("Test JSON Decoder")
	j := NewJSONComp()
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	m := j.decode(byt)

	fmt.Println(m)
}
