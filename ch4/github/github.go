// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 110.
//!+

// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

const IssuesURL = "https://xkcd.com/571/info.0.json"

type IssuesSearchResult struct {
	Month      string
	Num        int
	Year       string
	Transcript string
	Title      string
}

//!-
