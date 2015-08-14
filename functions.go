package main

import (
	"path/filepath"
	"strings"
	"time"
)

var (
	templateFunctions = map[string]interface{}{
		"displaytime": displaytime,
		"url":         url,
	}
)

func displaytime(t time.Time) string {
	return t.Format("Jan 2, 2006")
}

func url(p *Page) string {
	if filepath.Ext(p.Meta.Path) == ".md" {
		return strings.Replace(p.Meta.Path, ".md", ".html", 1)
	}
	return p.Meta.Path
}
