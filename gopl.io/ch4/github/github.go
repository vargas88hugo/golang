// Package github provides a Go API for the github issue tracker.
// See https://developer.github.com/v3/search/#search-issues
package github

import "time"

// IssuesURL is a constant that defines the url of github issues
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult is a struct that represents a search result
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue is a struct that defines a github issue
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

// User is a struct that represets a github user
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
