// BitMesh allows you to create only 1 AI agent at a  single  time on a local hardware.
// What I mean by that is at the moment as a user can select only 1 AI model and AI provider.

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
		src.CreateAgent(agent)
	}
}
