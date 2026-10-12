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
	AI_Provider string `json:"ai_provider"`
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
		fmt.Println("Ai Provider:", agent.AI_Provider)
		fmt.Println("Model:", agent.Model)
		fmt.Println("Domain:", agent.Domain)
		fmt.Println("Available:", agent.IsAvailable)
		fmt.Println("System Prompt:", agent.SystemPrompt)
		fmt.Println("Guardrails:", agent.Gaurdrails)
		return agent
	}
}

func CreateAgent(agent *Agent) {
 if agent == nil {
   log.Fatal("To create a angent define the agent information in the agent.json file ");
   return;
 }
 if agent.AI_Provider == "openai" {
	 fmt.Println("AI agent provider is open ai");
 } else if agent.AI_Provider == "antropic" {
	 fmt.Println("Ai agent provider is antropic");
 } else if agent.AI_Provider == "gemini" {
	 fmt.Println("AI agent provided by google");
 }
}
