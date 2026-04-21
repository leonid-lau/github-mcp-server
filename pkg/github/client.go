// Package github provides a GitHub API client and related utilities
// for use with the GitHub MCP server.
package github

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/v67/github"
	"golang.org/x/oauth2"
)

// Client wraps the GitHub API client with additional configuration.
type Client struct {
	*github.Client
	token string
	baseURL string
}

// ClientOption is a functional option for configuring a Client.
type ClientOption func(*Client)

// WithBaseURL sets a custom base URL for the GitHub API client.
// This is useful for GitHub Enterprise Server deployments.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// NewClient creates a new GitHub API client using the provided token.
// If token is empty, it falls back to the GITHUB_TOKEN environment variable.
// Also checks GH_TOKEN as a fallback, which is used by the GitHub CLI.
// Additionally checks GITHUB_PERSONAL_ACCESS_TOKEN as a fallback, which
// is a common name used in personal scripts and dotfiles.
func NewClient(token string, opts ...ClientOption) (*Client, error) {
	if token == "" {
		token = os.Getenv("GITHUB_TOKEN")
	}
	if token == "" {
		token = os.Getenv("GH_TOKEN")
	}
	if token == "" {
		token = os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
	}
	if token == "" {
		return nil, fmt.Errorf("GitHub token is required: set GITHUB_TOKEN environment variable or provide a token")
	}

	c := &Client{
		token: token,
	}

	for _, opt := range opts {
		opt(c)
	}

	httpClient := oauth2.NewClient(
		context.Background(),
		oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
	)

	var ghClient *github.Client
	if c.baseURL != "" {
		var err error
		ghClient, err = github.NewClient(httpClient).WithEnterpriseURLs(c.baseURL, c.baseURL)
		if err != nil {
			return nil, fmt.Errorf("failed to create GitHub enterprise client: %w", err)
		}
	} else {
		ghClient = github.NewClient(httpClient)
	}

	c.Client = ghClient
	return c, nil
}

// NewHTTPClient returns an *http.Client authenticated with the given token.
// This is useful when a raw HTTP client is needed for custom requests.
func NewHTTPClient(token string) *http.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	return oauth2.NewClient(context.Background(), ts)
}

// IsNotFound returns true if the error is a GitHub 404 Not Found error.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	errResp, ok := err.(*github.ErrorResponse)
	return ok && errResp.Response != nil && errResp.Response.StatusCode == http.StatusNotFound
}

// IsRateLimited returns true if the error is a GitHub rate limit error.
func IsRateLimited(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*github.RateLimitError)
	return ok
}
