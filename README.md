# go-jormall-sms

go-jormall-sms is a Go package for interacting with the Jormall SMS Gateway. It provides methods to send SMS messages, check account balance, and more.

## Installation

You can install this package using Go modules:

```shell
go get github.com/samerzmd/go-jormall-sms
```

## Usage

```go
package main

import (
    "fmt"
    "go-jormall-sms"
)

func main() {
// Create a JormallConfig with your configuration
 config := &sms.JormallConfig{
  BaseURL:         "https://api.jormallsms.com",
  AccountName:     "your_account_name",
  AccountPassword: "your_account_password",
  SenderID:        "your_sender_id",
 }

 // Create a JormallClient
 client := sms.NewJormallClient(config)

 // Check balance
 balance, err := client.Balance()
 if err != nil {
  fmt.Println("Error checking balance:", err)
 } else {
  fmt.Println("Current balance:", balance)
 }

 // Send an SMS
 messageID, err := client.Send("recipient_number", "Hello from go-jormall-sms!")
 if err != nil {
  fmt.Println("Error sending SMS:", err)
 } else {
  fmt.Println("SMS sent successfully. Message ID:", messageID)
 }

 // Send bulk SMS
 numbers := []string{"recipient_number_1", "recipient_number_2"}
 message := "Hello from go-jormall-sms!"
 messageID, err = client.SendBulk(numbers, message)
 if err != nil {
  fmt.Println("Error sending bulk SMS:", err)
 } else {
  fmt.Println("Bulk SMS sent successfully. Message ID:", messageID)
 }

 // SendOTP
 messageId, err := client.SendOtp("recipient_number","11")
 if err != nil {
  fmt.Println("Error sending OTP:", err)
 } else {
  fmt.Println("OTP sent successfully. Message ID:", messageId)
 }
}
```

## License

This package is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
