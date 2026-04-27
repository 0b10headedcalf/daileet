package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/0b10headedcalf/daileet/internal/api/queries"
)

// Client wraps Go's built-in http.Client.
// Keeping it in a struct lets us attach config (session cookies, headers)
// without changing every function signature later.
type Client struct {
	http    *http.Client
	url     string
	session string // LeetCode session cookie for authenticated requests
}

// NewClient is Go's idiomatic constructor pattern.
// Go has no classes — instead, a function returns a pointer to your struct.
// The caller gets *Client (pointer) so mutations to it persist.
func NewClient() *Client {
	return &Client{
		http: &http.Client{},
		url:  LeetCodeGQLURL,
	}
}

// SetSession stores the LeetCode session cookie on the client.
// Once set, every subsequent Query call will send it automatically.
func (c *Client) SetSession(token string) {
	c.session = token
}

// HasSession reports whether a session cookie has been set.
func (c *Client) HasSession() bool {
	return c.session != ""
}

// gqlRequest is the JSON shape LeetCode's GraphQL endpoint expects.
// The `json:"..."` struct tags control the key names during serialization —
// without them, encoding/json uses the field name as-is (e.g. "Query" not "query").
type gqlRequest struct {
	Query     string         `json:"query"`
	Variables map[string]any `json:"variables"`
}

// Query sends a GraphQL query and returns the raw JSON response as a string.
// vars can be nil if the query takes no arguments.
//
// Note: GraphQL always returns HTTP 200, even on errors. Query errors come
// back inside the JSON body as {"errors": [...]} — parse those yourself.
func (c *Client) Query(q string, vars map[string]any) (string, error) {
	// json.Marshal converts our Go struct into a JSON byte slice.
	// gqlRequest{Query: "..."} → []byte(`{"query":"...","variables":null}`)
	body, err := json.Marshal(gqlRequest{Query: q, Variables: vars})
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	// bytes.NewBuffer wraps the byte slice in an io.Reader.
	// http.NewRequest expects a Reader so it can stream the body.
	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("build request: %w", err)
	}

	// LeetCode requires this header to accept our JSON body.
	req.Header.Set("Content-Type", "application/json")

	// Inject the session cookie if one has been set via SetSession.
	if c.session != "" {
		req.AddCookie(&http.Cookie{Name: "LEETCODE_SESSION", Value: c.session})
	}

	// c.http.Do sends the request and returns the response.
	// resp.Body is a stream — defer closes it when this function returns,
	// preventing a TCP connection leak.
	resp, err := c.http.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// io.ReadAll drains the response stream into a byte slice.
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}

	return string(data), nil
}

// ProblemMeta holds minimal metadata for a LeetCode problem.
type ProblemMeta struct {
	Title      string
	TitleSlug  string
	Difficulty string
	IsPaidOnly bool
}

// FetchUserSolved hits LeetCode's REST endpoint and returns metadata for every
// problem the authenticated user has solved (status == "ac").
func (c *Client) FetchUserSolved() ([]ProblemMeta, error) {
	req, err := http.NewRequest("GET", LeetCodeURL+"/api/problems/all/", nil)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	if c.session != "" {
		req.AddCookie(&http.Cookie{Name: "LEETCODE_SESSION", Value: c.session})
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	var payload struct {
		UserName string `json:"user_name"`
		Pairs    []struct {
			Status     string `json:"status"`
			Difficulty struct {
				Level int `json:"level"`
			} `json:"difficulty"`
			Stat struct {
				Title     string `json:"question__title"`
				TitleSlug string `json:"question__title_slug"`
			} `json:"stat"`
		} `json:"stat_status_pairs"`
	}
	if err := json.Unmarshal(data, &payload); err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}

	if payload.UserName == "" {
		return nil, fmt.Errorf("not authenticated: no valid LEETCODE_SESSION")
	}

	var out []ProblemMeta
	for _, p := range payload.Pairs {
		if p.Status != "ac" {
			continue
		}
		meta := ProblemMeta{
			Title: p.Stat.Title,
			TitleSlug: p.Stat.TitleSlug,
		}
		switch p.Difficulty.Level {
		case 1:
			meta.Difficulty = "Easy"
		case 2:
			meta.Difficulty = "Medium"
		case 3:
			meta.Difficulty = "Hard"
		default:
			meta.Difficulty = "Medium"
		}
		out = append(out, meta)
	}
	return out, nil
}

// FetchProblemMeta queries LeetCode for a problem by title slug.
func (c *Client) FetchProblemMeta(titleSlug string) (ProblemMeta, error) {
	var meta ProblemMeta
	resp, err := c.Query(queries.SelectProblem, map[string]any{"titleSlug": titleSlug})
	if err != nil {
		return meta, err
	}

	// Lightweight parse: we only need a few fields from the nested JSON.
	var payload struct {
		Data struct {
			Question struct {
				Title      string `json:"title"`
				Difficulty string `json:"difficulty"`
				IsPaidOnly bool   `json:"isPaidOnly"`
			} `json:"question"`
		} `json:"data"`
	}
	if err := json.Unmarshal([]byte(resp), &payload); err != nil {
		return meta, fmt.Errorf("parse response: %w", err)
	}
	meta.Title = payload.Data.Question.Title
	meta.Difficulty = payload.Data.Question.Difficulty
	meta.IsPaidOnly = payload.Data.Question.IsPaidOnly
	if meta.Title == "" {
		return meta, fmt.Errorf("problem not found: %s", titleSlug)
	}
	return meta, nil
}
