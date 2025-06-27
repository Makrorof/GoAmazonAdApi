package tests_test

import (
	"context"
	"log"
	"testing"

	"github.com/Makrorof/GoAmazonAdApi/amazon"
	"github.com/joho/godotenv"
)

func TestAccessToken(t *testing.T) {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client := amazon.NewClientWithEnv(context.Background(), nil)

	if client == nil {
		log.Fatalf("Error client is empty")
	}

}
