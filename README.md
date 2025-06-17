# go-jormall-sms
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/ebe183f7ba7244ab821f18e09d00d172)](https://app.codacy.com/gh/samerzmd/go-jormall-sms?utm_source=github.com&utm_medium=referral&utm_content=samerzmd/go-jormall-sms&utm_campaign=Badge_Grade)  
[![Codacy Badge](https://app.codacy.com/project/badge/Coverage/be38289930ec403e9b74eb576de8530c)](https://app.codacy.com/gh/samerzmd/go-jormall-sms/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_coverage)

`go-jormall-sms` is a Go client for the [Jormall SMS Gateway (josms.net)](https://www.josms.net), supporting:

- 💬 General SMS sending
- 🔢 OTP SMS delivery
- 📢 Bulk messaging (up to 120 recipients)
- 💰 Account balance checking

---

## 📦 Installation

```sh
go get github.com/samerzmd/go-jormall-sms
```

---

## 🚀 Usage

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/samerzmd/go-jormall-sms"
)

func main() {
	config := sms.Config{
		BaseURL:         "https://www.josms.net",
		AccountName:     "your_account_name",
		AccountPassword: "your_account_password",
		SenderID:        "your_sender_id",
		RequestTimeout:  5000000, // Only needed for bulk sending
	}

	httpClient := &http.Client{
		Timeout: 15 * time.Second,
	}

	client := sms.NewJormallClient(config, httpClient)

	// Check balance
	balance, err := client.Balance()
	if err != nil {
		fmt.Println("Error checking balance:", err)
	} else {
		fmt.Println("Current balance:", balance)
	}

	// Send a general SMS
	messageID, err := client.Send("9627XXXXXXXX", "Hello from go-jormall-sms!")
	if err != nil {
		fmt.Println("Error sending SMS:", err)
	} else {
		fmt.Println("SMS sent successfully. Message ID:", messageID)
	}

	// Send a bulk SMS
	numbers := []string{"9627XXXXXXXX", "9627YYYYYYYY"}
	message := "This is a bulk message."
	messageID, err = client.SendBulk(numbers, message)
	if err != nil {
		fmt.Println("Error sending bulk SMS:", err)
	} else {
		fmt.Println("Bulk SMS sent. Message ID:", messageID)
	}

	// Send an OTP message
	messageID, err = client.SendOtp("9627XXXXXXXX", "Your OTP is: 123456")
	if err != nil {
		fmt.Println("Error sending OTP:", err)
	} else {
		fmt.Println("OTP sent. Message ID:", messageID)
	}
}
```

---

## 📌 Configuration Notes

| Field         | Description |
|---------------|-------------|
| `BaseURL`     | Use `https://www.josms.net` for the production gateway |
| `AccountName` | Provided by SMS provider |
| `AccountPassword` | Provided by SMS provider |
| `SenderID`    | Must be pre-approved (e.g., `App` or `srvApp`) |
| `RequestTimeout` | Optional for bulk sending timeout |

---

## 📄 License

This package is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
