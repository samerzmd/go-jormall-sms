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

// --- Helpers ---

func parseGatewayResponse(respBody []byte) (string, error) {
	parts := strings.Split(string(respBody), " = ")
	if len(parts) < 2 {
		return "", errors.New("invalid response format")
	}
	return parts[1], nil
}

func sanitizeMessage(msg string) string {
	msg = strings.ReplaceAll(msg, "%", "%25")
	msg = strings.ReplaceAll(msg, "&", "%26")
	return msg
}

// --- API Methods ---

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

	balanceStr := strings.Trim(string(body), "\"")
	balance, err := strconv.Atoi(balanceStr)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

// Send sends a single general SMS message.
func (client *JormallClient) Send(number, message string) (string, error) {
	endPoint := fmt.Sprintf("%s/SMSServices/Clients/Prof/RestSingleSMS_General/SendSMS", client.Config.BaseURL)
	reqData := url.Values{}
	reqData.Set("AccName", client.Config.AccountName)
	reqData.Set("AccPass", client.Config.AccountPassword)
	reqData.Set("senderid", client.Config.SenderID) // Use service-prefixed sender
	reqData.Set("numbers", number)
	reqData.Set("msg", sanitizeMessage(message))

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

	return parseGatewayResponse(body)
}

// SendOtp sends an OTP message.
func (client *JormallClient) SendOtp(number, otp string) (string, error) {
	endPoint := fmt.Sprintf("%s/SMSServices/Clients/Prof/RestSingleSMS/SendSMS", client.Config.BaseURL)
	reqData := url.Values{}
	reqData.Set("AccName", client.Config.AccountName)
	reqData.Set("AccPass", client.Config.AccountPassword)
	reqData.Set("senderid", client.Config.SenderID) // OTP sender, no "SRV" prefix
	reqData.Set("numbers", number)
	reqData.Set("msg", sanitizeMessage(otp))

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

	return parseGatewayResponse(body)
}

// SendBulk sends messages to multiple numbers (max 120).
func (client *JormallClient) SendBulk(numbers []string, message string) (string, error) {
	endPoint := fmt.Sprintf("%s/sms/api/SendBulkMessages.cfm", client.Config.BaseURL)
	reqData := url.Values{}
	reqData.Set("AccName", client.Config.AccountName)
	reqData.Set("AccPass", client.Config.AccountPassword)
	reqData.Set("senderid", client.Config.SenderID) // Use "SRV" sender for general bulk
	reqData.Set("numbers", strings.Join(numbers, ","))
	reqData.Set("msg", sanitizeMessage(message))
	reqData.Set("requesttimeout", strconv.Itoa(client.Config.RequestTimeout))

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

	return parseGatewayResponse(body)
}
