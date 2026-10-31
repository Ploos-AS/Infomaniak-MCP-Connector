package infomaniak

import (
 "context"
 "encoding/json"
 "fmt"
 "io"
 "net/http"
 "strings"
 "time"
)

const DefaultBaseURL = "https://api.infomaniak.com"

type Client struct {
 BaseURL string
 Token string
 HTTPClient *http.Client
}

func NewClient(token string) *Client {
 return &Client{BaseURL: DefaultBaseURL, Token: token, HTTPClient: &http.Client{Timeout: 15 * time.Second}}
}

func (c *Client) Profile(ctx context.Context) (json.RawMessage, error) {
 if strings.TrimSpace(c.Token) == "" { return nil, fmt.Errorf("Infomaniak access token is not configured") }
 base := strings.TrimRight(c.BaseURL, "/")
 if base == "" { return nil, fmt.Errorf("Infomaniak API base URL is empty") }
 req, err := http.NewRequestWithContext(ctx, http.MethodGet, base+"/profile", nil)
 if err != nil { return nil, err }
 req.Header.Set("Authorization", "Bearer "+c.Token)
 req.Header.Set("Accept", "application/json")
 client := c.HTTPClient
 if client == nil { client = &http.Client{Timeout: 15 * time.Second} }
 resp, err := client.Do(req)
 if err != nil { return nil, fmt.Errorf("Infomaniak profile request failed: %w", err) }
 defer resp.Body.Close()
 if resp.StatusCode < 200 || resp.StatusCode >= 300 { return nil, fmt.Errorf("Infomaniak profile request returned HTTP %d", resp.StatusCode) }
 data, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
 if err != nil { return nil, err }
 if !json.Valid(data) { return nil, fmt.Errorf("Infomaniak profile response is not valid JSON") }
 return json.RawMessage(data), nil
}
