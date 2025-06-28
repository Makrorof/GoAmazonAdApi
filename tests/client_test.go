package tests_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/Makrorof/GoAmazonAdApi/amazon"
	"github.com/joho/godotenv"
)

func TestAccessToken(t *testing.T) {
	amazon.DEBUG_PRINT_API_BODY = true

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client := amazon.NewClientWithEnv(context.Background(), nil)

	if client == nil {
		log.Fatalf("Error client is empty")
	}

	accounts := amazon.NewAccounts(client, amazon.AMAZON_ENDPOINTS_NA)

	data, err := accounts.Profiles(context.Background())
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	fmt.Println("Data:", data)
}
