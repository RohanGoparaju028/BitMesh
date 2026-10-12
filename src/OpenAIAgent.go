// This file focuses on creating AI agent in go using Chatgpt as LLM model
package src 

import (
  "fmt"
  "context"
  "log"
  "os"
  "github.com/joho/godotenv"

)

func AI_Agent() {
  err := godotenv.Load("./.env");
  if err != nil {
	  log.Fatal("Failed to load the .env file");
	  return
  }
  openai_api := os.GetEnv("OPENAI_KEY")
  openai_model := os.GetEnv("OPENAI_KEY") 
  if openai_api == "" || openai_model == "" {
	  log.Fatal("We Require api and model type to create an agent");
	  return
  }
}
