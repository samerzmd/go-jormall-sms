package sms

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Balance retrieves the current balance from the Jormall SMS Gateway.
func (client *JormallClient) Balance() (int, error) {
    endPoint := fmt.Sprintf("%s/SMS/API/GetBalance", client.Config.BaseURL)
    reqData := url.Values{}
    reqData.Set("AccName", client.Config.AccountName)
    reqData.Set("AccPass", client.Config.AccountPassword)

    req, err := http.NewRequest("GET", endPoint, nil)
    if err != nil {
        return 0, err
    }
    req.URL.RawQuery = reqData.Encode()

    resp, err := client.HTTPClient.Do(req)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return 0, err
    }

    // Remove double quotes and convert to int
    balanceStr := strings.Trim(string(body), "\"")
    balance, err := strconv.Atoi(balanceStr)
    if err != nil {
        return 0, err
    }

    return balance, nil
}

// Send sends a single SMS message to the specified number.
func (client *JormallClient) Send(number, message string) (string, error) {
    endPoint := fmt.Sprintf("%s/SMSServices/Clients/Prof/RestSingleSMS_General/SendSMS", client.Config.BaseURL)
    reqData := url.Values{}
    reqData.Set("AccName", client.Config.AccountName)
    reqData.Set("AccPass", client.Config.AccountPassword)
    reqData.Set("senderid", client.Config.SenderID)
    reqData.Set("numbers", number) // Assume number is already formatted
    reqData.Set("msg", message)    // Assume message is already formatted

    req, err := http.NewRequest("GET", endPoint, nil)
    if err != nil {
        return "", err
    }
    req.URL.RawQuery = reqData.Encode()

    resp, err := client.HTTPClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    // Extracting message ID from the response
    parts := strings.Split(string(body), " = ")
    if len(parts) < 2 {
        return "", errors.New("invalid response format")
    }
    messageID := parts[1]

    return messageID, nil
}

func (client *JormallClient) SendBulk(numbers []string, message string) (string, error) {
    endPoint := fmt.Sprintf("%s/sms/api/SendBulkMessages.cfm", client.Config.BaseURL)
    timeout := 5000000 // Adjust as necessary

    reqData := url.Values{}
    reqData.Set("AccName", client.Config.AccountName)
    reqData.Set("AccPass", client.Config.AccountPassword)
    reqData.Set("senderid", client.Config.SenderID)
    reqData.Set("numbers", strings.Join(numbers, ",")) // Join numbers with a comma
    reqData.Set("msg", message)
    reqData.Set("requesttimeout", fmt.Sprintf("%d", timeout))

    req, err := http.NewRequest("GET", endPoint, nil)
    if err != nil {
        return "", err
    }
    req.URL.RawQuery = reqData.Encode()

    resp, err := client.HTTPClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    parts := strings.Split(string(body), " = ")
    if len(parts) < 2 {
        return "", errors.New("invalid response format")
    }
    messageID := parts[1]

    return messageID, nil
}

// SendOtp sends an OTP message to the specified phone number.
func (client *JormallClient) SendOtp(number, otp string) (string, error) {
    endPoint := fmt.Sprintf("%s/SMSServices/Clients/Prof/RestSingleSMS/SendSMS", client.Config.BaseURL)
    reqData := url.Values{}
    reqData.Set("AccName", client.Config.AccountName)
    reqData.Set("AccPass", client.Config.AccountPassword)
    reqData.Set("senderid", client.Config.SenderID)
    reqData.Set("numbers", number) // Assume number is already formatted
    reqData.Set("msg", otp)        // OTP message

    req, err := http.NewRequest("GET", endPoint, nil)
    if err != nil {
        return "", err
    }
    req.URL.RawQuery = reqData.Encode()

    resp, err := client.HTTPClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)

    if err != nil {
        return "", err
    }

    parts := strings.Split(string(body), " = ")
    if len(parts) < 2 {
        return "", errors.New("invalid response format")
    }
    messageID := parts[1]

    return messageID, nil
}
