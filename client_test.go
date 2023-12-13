package sms

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBalance(t *testing.T) {
    // Create a mock HTTP server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Mock response for balance endpoint
        w.Write([]byte(`"100"`)) // Assuming the API returns a string number
    }))
    defer server.Close()

    // Create a JormallClient with the mock server URL
    client := NewJormallClient(&JormallConfig{
        BaseURL:         server.URL, // Use mock server URL
        AccountName:     "test",
        AccountPassword: "password",
        SenderID:        "sender",
    })

    // Call the Balance method
    balance, err := client.Balance()

    // Assert the balance
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if balance != 100 {
        t.Errorf("Expected balance 100, got %d", balance)
    }
}

func TestSendBulk(t *testing.T) {
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Define the mock response here
        w.Write([]byte("message_id = 12345"))
    }))
    defer server.Close()
    // Create a JormallClient with the mock server URL
    client := NewJormallClient(&JormallConfig{
        BaseURL:         server.URL, // Use mock server URL
        AccountName:     "test",
        AccountPassword: "password",
        SenderID:        "sender",
    })

    // Define test numbers and message
    numbers := []string{"12345", "67890"}
    message := "Test message"

    // Execute the SendBulk method
    messageID, err := client.SendBulk(numbers, message)

    // Check if the test passed or failed
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if messageID == "" {
        t.Errorf("Expected message_id, got empty string")
    }
}


func TestSend(t *testing.T) {
    // Setup mock server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Simulate API response for sending an SMS
        w.Write([]byte(`key = message_id`)) // Mock response format
    }))
    defer server.Close()

    // Setup client with mock server URL
    client := NewJormallClient(&JormallConfig{
        BaseURL:         server.URL, // Use mock server URL
        AccountName:     "test",
        AccountPassword: "password",
        SenderID:        "sender",
    })

    // Call the Send method
    messageID, err := client.Send("1234567890", "Test message")

    // Assert the expected message ID and no error
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if messageID != "message_id" {
        t.Errorf("Expected message_id, got %s", messageID)
    }
}

func TestSendOtp(t *testing.T) {
    // Create a mock HTTP server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Define the mock response here
        // Simulate a successful response
        w.Write([]byte("message_id = 12345"))
    }))
    defer server.Close()

    // Create a JormallClient with the mock server URL
    client := NewJormallClient(&JormallConfig{
        BaseURL:         server.URL, // Use mock server URL
        AccountName:     "test",
        AccountPassword: "password",
        SenderID:        "sender",
    })

    // Test sending OTP
    messageID, err := client.SendOtp("123456789", "123456")
    if err != nil {
        t.Errorf("Error sending OTP: %v", err)
    }

    // Check the message ID received
    if messageID != "12345" {
        t.Errorf("Expected message ID: 12345, Got: %s", messageID)
    }
}
