package main

import (
	"time"

	"github.com/extemporalgenome/slug"
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
	return slug.Slug(p.Meta.Title) + ".html"
}
