package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/etherealmachine/markdown"
	"gopkg.in/yaml.v2"
)

var (
	preview = flag.Bool("preview", false, "Serve locally for preview")
)

type Site struct {
	Recent        []*Page
	PreviewScript string
}

type Meta struct {
	Title    string
	Summary  string
	Path     string
	FileInfo os.FileInfo
}

type Page struct {
	*Site
	Meta    *Meta
	Content string
}

type generator struct {
	site     *Site
	errors   []error
	pages    map[string]*Page
	template *template.Template
}

type ByModTime []*Page

func (a ByModTime) Len() int      { return len(a) }
func (a ByModTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByModTime) Less(i, j int) bool {
	return a[i].Meta.FileInfo.ModTime().Before(a[j].Meta.FileInfo.ModTime())
}

func (g *generator) generatePageContent() {
	if err := g.walk(); err != nil {
		log.Fatal(err)
	}
}

func (g *generator) walk() error {
	filepath.Walk("pages", g.visit)
	var err error
	g.template, err = template.New("").Funcs(templateFunctions).
		ParseGlob(filepath.Join("templates", "*.html"))
	if err != nil {
		return err
	}
	return nil
}

func (g *generator) visit(path string, fileInfo os.FileInfo, err error) error {
	if strings.HasSuffix(path, ".md") {
		page, err := g.generatePage(path, fileInfo)
		if err != nil {
			g.errors = append(g.errors, err)
			return nil
		}
		g.pages[path] = page
	}
	return nil
}

func (g *generator) gatherSiteData() {
	for _, page := range g.pages {
		g.site.Recent = append(g.site.Recent, page)
	}
	sort.Sort(ByModTime(g.site.Recent))
	if len(g.site.Recent) > 4 {
		g.site.Recent = g.site.Recent[:5]
	}
}

func (g *generator) generateIndex() {
	f, err := os.Create("index.html")
	if err != nil {
		log.Fatal(err)
	}
	g.template.ExecuteTemplate(f, "index.html", g.site)
}

func (g *generator) generatePages() error {
	for _, page := range g.pages {
		f, err := os.Create(url(page))
		if err != nil {
			return err
		}
		g.template.ExecuteTemplate(f, "page.html", page)
	}
	return nil
}

func (g *generator) generatePage(path string, fileInfo os.FileInfo) (*Page, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	buf, meta, err := parseMeta(buf)
	if err != nil {
		return nil, err
	}
	meta.Path = path
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
	return markdown.Markdown(string(buf))
}

func runGenerator() {
	g := &generator{
		site: &Site{
			PreviewScript: previewScript,
		},
		pages: make(map[string]*Page),
	}
	// generate page content
	g.generatePageContent()
	if len(g.errors) > 0 {
		for _, err := range g.errors {
			log.Println(err)
		}
	}
	g.gatherSiteData()
	g.generateIndex()
	g.generatePages()
}

func main() {
	flag.Parse()
	runGenerator()
	if *preview {
		runPreview()
	}
}
