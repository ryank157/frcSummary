package statbotics

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// NewClient creates a new Statbotics API client.
func NewClient(baseURL string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	return &Client{
		baseURL:    baseURL,
		httpClient: httpClient,
	}
}

// GetDefault makes a request to the default route of the Statbotics API.
// It returns a DefaultResponse struct and an error, if any.
func (c *Client) GetDefault(ctx context.Context) (*DefaultResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response DefaultResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &response, nil
}

func (c *Client) GetMatch(ctx context.Context, matchKey string) (*MatchResponse, error) {
	// Construct the URL with the match key.  Important to URL encode the match key!
	endpoint := fmt.Sprintf("match/%s", url.PathEscape(matchKey))
	url := c.baseURL + endpoint

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response MatchResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &response, nil
}
