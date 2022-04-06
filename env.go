package main

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

func Init() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}

func GetENV(key string) string {
	return os.Getenv(key)
}