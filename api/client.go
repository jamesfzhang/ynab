package api

import (
	"encoding/json"
	"github.com/jamesfzhang/ynab/model"
	"io/ioutil"
	"net/http"
)

const (
	API_BASE_URL = "https://api.youneedabudget.com/v1"
)

type ApiService struct {
	Client *Client
}

type AccountService ApiService
type BudgetService ApiService
type UserService ApiService

type Client struct {
	baseURL     string
	accessToken string

	// Internal HTTP client
	httpClient *http.Client

	// API services
	AccountService *AccountService
	BudgetService  *BudgetService
	UserService    *UserService
}

// NewClient returns an API client using the specified access token.
func NewClient(accessToken string) *Client {
	client := &Client{
		baseURL:     API_BASE_URL,
		accessToken: accessToken,
	}

	client.httpClient = &http.Client{}
	client.AccountService = &AccountService{client}
	client.BudgetService = &BudgetService{client}
	client.UserService = &UserService{client}

	return client
}

func (client Client) get(
	path string,
	result interface{},
) (err error) {

	// Create request
	req, err := http.NewRequest("GET", client.baseURL+path, nil)
	if err != nil {
		return
	}

	// Add headers
	req.Header.Add("Authorization", "Bearer "+client.accessToken)
	req.Header.Add("Content-Type", "application/json")

	// Make request
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return
	}

	// Close body after returning
	defer resp.Body.Close()

	// Check error
	err = client.validateResponse(resp)
	if err != nil {
		return
	}

	// Read body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// Parse JSON
	err = json.Unmarshal(body, result)
	if err != nil {
		return
	}

	return
}

// validateResponse checks that the response has status code 2xx.
// Otherwise, it parses the API error from the response body.
func (client Client) validateResponse(resp *http.Response) error {
	status := resp.StatusCode
	if status >= 200 && status <= 299 {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	apiError := &model.ApiError{Response: resp}
	err = json.Unmarshal(body, apiError)

	return apiError
}
