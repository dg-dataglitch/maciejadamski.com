# L-002: HTTP API Client Pattern in Go

Date: 2026-01-25
Category: Pattern
Related: L-001

## 1. Context (Why I Asked)

Building gopress will require calling external APIs. I needed to understand the production-quality pattern for making HTTP requests, handling errors, and structuring API client code.

## 2. The Concept

A well-structured API client encapsulates all HTTP communication with an external service. It handles authentication, request building, error parsing, and response decoding in one place. The pattern uses a dedicated struct with an internal `do()` method that all public methods call.

## 3. Key Points

- Never use `http.DefaultClient` - it has no timeout
- Always pass `context.Context` as first parameter
- Create dedicated client struct holding: baseURL, credentials, httpClient, logger
- Use internal `do()` method for DRY request handling
- Define custom error type that captures status code and API error message
- Use `errors.As()` to check for specific API errors in caller code
- Log at boundaries: request start, response status, duration
- Make client testable by allowing injection of httpClient and baseURL

## 4. Example
```go
package github

type Client struct {
    baseURL    string
    token      string
    httpClient *http.Client
    logger     *slog.Logger
}

func NewClient(token string, logger *slog.Logger) *Client {
    return &Client{
        baseURL: "https://api.github.com",
        token:   token,
        httpClient: &http.Client{
            Timeout: 30 * time.Second,
        },
        logger: logger,
    }
}

// Internal method handles all HTTP details
func (c *Client) do(ctx context.Context, method, path string, body, result interface{}) error {
    var reqBody io.Reader
    if body != nil {
        jsonBytes, err := json.Marshal(body)
        if err != nil {
            return fmt.Errorf("marshaling request: %w", err)
        }
        reqBody = bytes.NewReader(jsonBytes)
    }

    req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, reqBody)
    if err != nil {
        return fmt.Errorf("creating request: %w", err)
    }

    req.Header.Set("Authorization", "Bearer "+c.token)
    if body != nil {
        req.Header.Set("Content-Type", "application/json")
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return fmt.Errorf("executing request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        return c.parseError(resp)
    }

    if result != nil {
        if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
            return fmt.Errorf("decoding response: %w", err)
        }
    }
    return nil
}

// Public methods are clean and focused
func (c *Client) GetRepository(ctx context.Context, owner, repo string) (*Repository, error) {
    var result Repository
    if err := c.do(ctx, http.MethodGet, fmt.Sprintf("/repos/%s/%s", owner, repo), nil, &result); err != nil {
        return nil, err
    }
    return &result, nil
}

// Custom error type
type APIError struct {
    StatusCode int
    Message    string `json:"message"`
}

func (e *APIError) Error() string {
    return fmt.Sprintf("github api (status=%d): %s", e.StatusCode, e.Message)
}

func (c *Client) parseError(resp *http.Response) error {
    body, _ := io.ReadAll(resp.Body)
    var apiErr APIError
    json.Unmarshal(body, &apiErr)
    apiErr.StatusCode = resp.StatusCode
    return &apiErr
}
```

**Using the client:**
```go
func (s *Service) SyncRepo(ctx context.Context, owner, repo string) error {
    repository, err := s.github.GetRepository(ctx, owner, repo)
    if err != nil {
        var apiErr *github.APIError
        if errors.As(err, &apiErr) && apiErr.StatusCode == 404 {
            return ErrRepoNotFound
        }
        return fmt.Errorf("fetching repository: %w", err)
    }
    // use repository...
}
```

## 5. When to Use / When to Avoid

**Use this pattern when:**
- Calling any external API more than once
- API requires authentication
- You need consistent error handling
- You want testable code

**Avoid when:**
- One-off script with single API call
- Prototyping (but refactor before production)

## 6. Sources

- AI discussion during gopress development
- Go standard library: net/http
- Production patterns from real-world Go services