package main

import (
	"BitMesh/src"
	"fmt"
)

func main() {
	agent := src.ParseAgent()
	if agent == nil {
		fmt.Println("Error while parsing agent.json")
	} else {
		fmt.Println("Successfully parsed the agent.json")
	}
}
