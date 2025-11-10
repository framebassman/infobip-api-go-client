package examples

import (
	"context"
	"fmt"
	"testing"

	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip"
	"github.com/infobip/infobip-api-go-client/v3/pkg/infobip/api"
)

func TestSendEmail(t *testing.T) {
	configuration := infobip.NewConfiguration()
	configuration.Host = "<YOUR_BASE_URL>"

	infobipClient := api.NewAPIClient(configuration)

	auth := context.WithValue(
		context.Background(),
		infobip.ContextAPIKeys,
		map[string]infobip.APIKey{
			"APIKeyHeader": {Key: "<YOUR_API_KEY>", Prefix: "<API_PREFIX>"},
		},
	)

	apiResponse, httpResponse, err := infobipClient.
		EmailAPI.
		SendEmail(auth).
		To([]string{"<DESTINATION1>", "<DESTINATION2>"}).
		From("<FROM>").
		Subject("<SUBJECT>").
		Text("<TEXT>").
		Execute()

	// Check for errors
	if err != nil {
		t.Fatalf("Failed to send Email: %v", err) // Fail the test with the error message
	}

	// Output response details for debugging
	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	// Validate response
	if apiResponse == nil || apiResponse.Messages == nil || len(apiResponse.Messages) == 0 {
		t.Fatalf("Invalid response: expected messages, but got: %+v", apiResponse)
	}
}

func TestCreateEmailTemplate(t *testing.T) {
	configuration := infobip.NewConfiguration()
	configuration.Host = "<YOUR_BASE_URL>"

	infobipClient := api.NewAPIClient(configuration)

	auth := context.WithValue(
		context.Background(),
		infobip.ContextAPIKeys,
		map[string]infobip.APIKey{
			"APIKeyHeader": {Key: "<YOUR_API_KEY>", Prefix: "<API_PREFIX>"},
		},
	)

	apiResponse, httpResponse, err := infobipClient.
		EmailAPI.
		CreateEmailTemplate(auth).
		Name("Welcome email").
		From("<FROM>").
		ReplyTo("<REPLY_TO>").
		Subject("Welcome to Infobip").
		Preheader("Welcome to Infobip").
		Html("<html><head></head><body><h2>Welcome to Infobip</h2></body></html>").
		LandingPage("1_2345").
		Execute()

	// Check for errors
	if err != nil {
		t.Fatalf("Failed to create Email template: %v", err) // Fail the test with the error message
	}

	// Output response details for debugging
	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	// Validate response
	if apiResponse == nil || apiResponse.ID == 0 || apiResponse.Name == "" {
		t.Fatalf("Invalid response: expected id, but got: %+v", apiResponse)
	}
}

func TestGetAllEmailTemplates(t *testing.T) {
	configuration := infobip.NewConfiguration()
	configuration.Host = "<YOUR_BASE_URL>"

	infobipClient := api.NewAPIClient(configuration)

	auth := context.WithValue(
		context.Background(),
		infobip.ContextAPIKeys,
		map[string]infobip.APIKey{
			"APIKeyHeader": {Key: "<YOUR_API_KEY>", Prefix: "<API_PREFIX>"},
		},
	)

	apiResponse, httpResponse, err := infobipClient.
		EmailAPI.
		GetAllEmailTemplates(auth).
		Execute()

	// Check for errors
	if err != nil {
		t.Fatalf("Failed to create Email template: %v", err) // Fail the test with the error message
	}

	// Output response details for debugging
	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	// Validate response
	if apiResponse == nil {
		t.Fatalf("Invalid response: expected id, but got: %+v", apiResponse)
	}
}

func TestGetEmailTemplate(t *testing.T) {
	configuration := infobip.NewConfiguration()
	configuration.Host = "<YOUR_BASE_URL>"

	infobipClient := api.NewAPIClient(configuration)

	auth := context.WithValue(
		context.Background(),
		infobip.ContextAPIKeys,
		map[string]infobip.APIKey{
			"APIKeyHeader": {Key: "<YOUR_API_KEY>", Prefix: "<API_PREFIX>"},
		},
	)

	apiResponse, httpResponse, err := infobipClient.
		EmailAPI.
		GetEmailTemplate(auth).
		ID(205000000016538).
		Execute()

	// Check for errors
	if err != nil {
		t.Fatalf("Failed to create Email template: %v", err) // Fail the test with the error message
	}

	// Output response details for debugging
	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	// Validate response
	if apiResponse == nil {
		t.Fatalf("Invalid response: expected id, but got: %+v", apiResponse)
	}
}

func TestDeleteEmailTemplate(t *testing.T) {
	configuration := infobip.NewConfiguration()
	configuration.Host = "<YOUR_BASE_URL>"

	infobipClient := api.NewAPIClient(configuration)

	auth := context.WithValue(
		context.Background(),
		infobip.ContextAPIKeys,
		map[string]infobip.APIKey{
			"APIKeyHeader": {Key: "<YOUR_API_KEY>", Prefix: "<API_PREFIX>"},
		},
	)

	httpResponse, err := infobipClient.
		EmailAPI.
		RemoveEmailTemplate(auth).
		ID(205000000016538).
		Execute()

	// Check for errors
	if err != nil {
		t.Fatalf("Failed to remove Email template: %v", err) // Fail the test with the error message
	}

	// Output response details for debugging
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)
}

func TestUpdateEmailTemplate(t *testing.T) {
	configuration := infobip.NewConfiguration()
	configuration.Host = "<YOUR_BASE_URL>"

	infobipClient := api.NewAPIClient(configuration)

	auth := context.WithValue(
		context.Background(),
		infobip.ContextAPIKeys,
		map[string]infobip.APIKey{
			"APIKeyHeader": {Key: "<YOUR_API_KEY>", Prefix: "<API_PREFIX>"},
		},
	)

	apiResponse, httpResponse, err := infobipClient.
		EmailAPI.
		UpdateEmailTemplate(auth).
		ID(205000000016536).
		Name("Welcome email1").
		From("Romashov <noreply@romashov.tech>").
		ReplyTo("support@example.com").
		Subject("Welcome to Infobip").
		Preheader("Welcome to Infobip").
		Html("<html><head></head><body><h2>Welcome to Infobip</h2></body></html>").
		LandingPage("1_2345").
		Execute()

	// Check for errors
	if err != nil {
		t.Fatalf("Failed to create Email template: %v", err) // Fail the test with the error message
	}

	// Output response details for debugging
	fmt.Printf("Response: %+v\n", apiResponse)
	fmt.Printf("HTTP Response Details: %+v\n", httpResponse)

	// Validate response
	if apiResponse == nil || apiResponse.ID == 0 || apiResponse.Name == "" {
		t.Fatalf("Invalid response: expected id, but got: %+v", apiResponse)
	}
}
