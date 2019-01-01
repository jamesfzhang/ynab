package ynab

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	API_BASE_URL = "https://api.youneedabudget.com/v1"
)

type Service struct {
	Client *Client
}

type Client struct {
	BaseURL     string
	AccessToken string

	// Internal HTTP client
	HttpClient *http.Client

	// API services
	BudgetSummaryService *BudgetSummaryService
}

func NewClient(accessToken string) *Client {
	client := &Client{
		BaseURL:     API_BASE_URL,
		AccessToken: accessToken,
	}

	client.HttpClient = &http.Client{}
	client.BudgetSummaryService = &BudgetSummaryService{client}

	return client
}

func (client *Client) Get(
	path string,
	result *BudgetSummaryResponse,
) (err error) {

	// Create request
	req, err := http.NewRequest("GET", client.BaseURL+path, nil)
	if err != nil {
		fmt.Println("Error: %v", err)
		return
	}

	// Add auth
	req.Header.Add("Authorization", "Bearer "+client.AccessToken)

	// Make request
	resp, err := client.HttpClient.Do(req)
	if err != nil {
		fmt.Println("Error: %v", err)
		return
	}

	// Ready response body & defer close
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: %v", err)
		return
	}

	// Parse JSON
	err = json.Unmarshal(body, result)
	if err != nil {
		fmt.Println("Error: %v", err)
		return
	}

	return
}
