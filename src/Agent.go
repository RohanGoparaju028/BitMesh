package src

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Config struct {
	Agent Agent `json:"agent"`
}
type Agent struct {
	Model        string `json:"model"`
	Domain       string `json:"domain"`
	IsAvailable  bool   `json:"is_available"`
	SystemPrompt string `json:"system_prompt"`
	Gaurdrails   string `json:"gaurdrails"`
}

const agent_json = "./agent.json"

func ParseAgent() *Agent {
	if _, err := os.Stat(agent_json); os.IsNotExist(err) {
		fmt.Println("agent.json does not exist")
		return nil
	} else {
		fmt.Println("agent.json file exist")
		agent := &Agent{}
		file, err := os.Open(agent_json)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		defer file.Close()
		data, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		var config Config
		if err := json.Unmarshal(data, &config); err != nil {
			log.Fatal(err)
			return nil
		}
		agent = &config.Agent
		fmt.Println("Model:", agent.Model)
		fmt.Println("Domain:", agent.Domain)
		fmt.Println("Available:", agent.IsAvailable)
		fmt.Println("System Prompt:", agent.SystemPrompt)
		fmt.Println("Guardrails:", agent.Gaurdrails)
		return agent
	}
}
