package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type IdentityClient struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func NewIdentityClient(baseURL, apiKey string, httpClient *http.Client) *IdentityClient {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 3 * time.Second}
	}
	return &IdentityClient{
		baseURL:    baseURL,
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

type UserPermissions struct {
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

func (c *IdentityClient) GetUserPermissions(ctx context.Context, userID string) (*UserPermissions, error) {
	url := fmt.Sprintf("%s/internal/permissions?user_id=%s", c.baseURL, userID)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Internal-Api-Key", c.apiKey)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("identity service returned status %d", res.StatusCode)
	}

	var result UserPermissions
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
