package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

// An example showing how to unmarshal embedded
// structs from YAML

// StructA is an exported type that
// contains a string
type StructA struct {
	A string `yaml:"a"`
}

// StructB is an exported type that
// contains a string and a StructA.
type StructB struct {
	// Embedded structs are not treated as embedded in YAML by default. To do that,
	// add the ",inline" annotation below
	StructA `yaml:",inline"`
	B       string `yaml:"b"`
}

var data = `
a: a string from struct A
b: a string from struct B
`

func main() {
	var b StructB

	err := yaml.Unmarshal([]byte(data), &b)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	fmt.Println(b.A)
	fmt.Println(b.B)
}
