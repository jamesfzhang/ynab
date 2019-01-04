package api

import (
	"encoding/json"
	"github.com/jamesfzhang/ynab/model"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

const (
	API_BASE_URL = "https://api.youneedabudget.com/v1"
	H_RATE_LIMIT = "X-Rate-Limit"
)

type Client struct {
	baseURL     string
	accessToken string

	// Internal HTTP client
	httpClient *http.Client

	// Utils
	remainingRequestsCount     int
	remainingRequestsCountLock *sync.RWMutex

	// API services
	AccountService *AccountService
	BudgetService  *BudgetService
	UserService    *UserService
}

type ApiService struct {
	Client *Client
}

type AccountService ApiService
type BudgetService ApiService
type UserService ApiService

// NewClient returns an API client using the specified access token.
func NewClient(accessToken string) *Client {
	client := &Client{
		baseURL:     API_BASE_URL,
		accessToken: accessToken,
	}

	client.httpClient = &http.Client{}
	client.remainingRequestsCountLock = &sync.RWMutex{}
	client.AccountService = &AccountService{client}
	client.BudgetService = &BudgetService{client}
	client.UserService = &UserService{client}

	return client
}

func (client *Client) RemainingRequestsCount() int {
	client.remainingRequestsCountLock.RLock()
	defer client.remainingRequestsCountLock.RUnlock()
	return client.remainingRequestsCount
}

func (client *Client) get(path string, result interface{}) (err error) {

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

	// Always close body
	defer resp.Body.Close()
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return
	}

	// Parse rate limit from response header
	if err := client.parseRateLimit(&resp.Header); err != nil {
		return err
	}

	// Check upstream error
	if err = client.validateResponse(resp); err != nil {
		return err
	}

	// Read body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	// Parse JSON
	if err = json.Unmarshal(body, result); err != nil {
		return err
	}

	return
}

// parseRateLimit reads the response headers and saves number of requests
// remaining before rate limiting starts.
func (client *Client) parseRateLimit(header *http.Header) (err error) {

	client.remainingRequestsCountLock.Lock()
	defer client.remainingRequestsCountLock.Unlock()

	// Rate limit is formatted like X/200 meaning X requests have already been
	// made out of 200 requests allowed. The limit resets hourly.
	val := header.Get(http.CanonicalHeaderKey(H_RATE_LIMIT))
	limits := strings.Split(val, "/")

	if len(limits) != 2 {
		return
	}

	reqCount, err := strconv.Atoi(limits[0])
	if err != nil {
		return
	}

	totalCount, err := strconv.Atoi(limits[1])
	if err != nil {
		return
	}

	client.remainingRequestsCount = totalCount - reqCount
	return
}

// validateResponse checks that the response has status code 2xx.
// Otherwise, it parses the API error from the response body.
func (client *Client) validateResponse(resp *http.Response) error {
	status := resp.StatusCode
	if status >= 200 && status <= 299 {
		return nil
	}

	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		return err
	} else {
		apiError := &model.ApiError{Response: resp}
		err = json.Unmarshal(body, apiError)

		return apiError
	}
}
