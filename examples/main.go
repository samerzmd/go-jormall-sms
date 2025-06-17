package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/samerzmd/go-jormall-sms"
)

func main() {
	// Setup config
	config := sms.Config{
		BaseURL:         "https://www.josms.net", // Use josms.net as per API
		AccountName:     "your_account_name",
		AccountPassword: "your_account_password",
		SenderID:        "your_sender_id",
		RequestTimeout:  5000000,      // Optional, used in bulk sending
	}

	// Optional: Custom HTTP client with timeout
	httpClient := &http.Client{
		Timeout: 15 * time.Second,
	}

	// Initialize the client
	client := sms.NewJormallClient(config, httpClient)

	// Check balance
	balance, err := client.Balance()
	if err != nil {
		fmt.Println("❌ Error checking balance:", err)
	} else {
		fmt.Println("💰 Current balance:", balance)
	}

	// Send a general SMS
	messageID, err := client.Send("9627XXXXXXXX", "📝 General message via API.")
	if err != nil {
		fmt.Println("❌ Error sending SMS:", err)
	} else {
		fmt.Println("✅ SMS sent. Message ID:", messageID)
	}

	// Send bulk SMS
	numbers := []string{"9627XXXXXXXX", "9627YYYYYYYY"}
	bulkMessage := "📢 Bulk SMS Test"
	messageID, err = client.SendBulk(numbers, bulkMessage)
	if err != nil {
		fmt.Println("❌ Error sending bulk SMS:", err)
	} else {
		fmt.Println("✅ Bulk SMS sent. Message ID:", messageID)
	}

	// Send OTP message
	otpMsg := "🔐 Your OTP is: 123456"
	messageID, err = client.SendOtp("9627XXXXXXXX", otpMsg)
	if err != nil {
		fmt.Println("❌ Error sending OTP:", err)
	} else {
		fmt.Println("✅ OTP sent. Message ID:", messageID)
	}
}
