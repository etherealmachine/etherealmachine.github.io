package main

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/etherealmachine/blackfriday"
	"gopkg.in/yaml.v2"
)

type Meta struct {
	Title    string
	Summary  string
	FileInfo os.FileInfo
}

type Page struct {
	*Site
	Meta    *Meta
	Content string
}

func (g *generator) newPage(path string, fileInfo os.FileInfo) (*Page, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	buf, meta, err := parseMeta(buf)
	if err != nil {
		return nil, err
	}
	meta.FileInfo = fileInfo
	return &Page{
		Site:    g.site,
		Meta:    meta,
		Content: parseContent(buf),
	}, nil
}

func parseMeta(buf []byte) ([]byte, *Meta, error) {
	var meta Meta
	end := bytes.Index(buf, []byte("\n\n"))
	if err := yaml.Unmarshal(buf[:end], &meta); err != nil {
		return nil, nil, err
	}
	return buf[end:], &meta, nil
}

func parseContent(buf []byte) string {
	return string(blackfriday.Markdown(buf,
		blackfriday.HtmlRenderer(blackfriday.CommonHtmlFlags, "", ""),
		blackfriday.CommonExtensions|blackfriday.EXTENSION_PARSE_DIV))
}
