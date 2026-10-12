// Package commentviewer provides small helpers for normalizing public social comment records.
//
// For a browser-based viewer and CSV or Excel export workflow, see https://commentviewer.net/.
package commentviewer

import "strings"

// Comment is a normalized public comment and its optional parent relationship.
type Comment struct {
	ID       string
	ParentID string
	Text     string
	Author   string
}

// NormalizeText trims surrounding whitespace from a comment body.
func NormalizeText(value string) string { return strings.TrimSpace(value) }
