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
	"time"

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
	Title     string
	Summary   string
	Published bool
	Date      time.Time
	Path      string
	FileInfo  os.FileInfo
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

type ByTime []*Page

func (a ByTime) Len() int      { return len(a) }
func (a ByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool {
	return a[i].Meta.Date.Before(a[j].Meta.Date)
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
		err := g.generatePage(path, fileInfo)
		if err != nil {
			g.errors = append(g.errors, err)
			return nil
		}
	}
	return nil
}

func (g *generator) gatherSiteData() {
	for _, page := range g.pages {
		if page.Meta.Published {
			g.site.Recent = append(g.site.Recent, page)
		}
	}
	sort.Sort(ByTime(g.site.Recent))
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

func (g *generator) generatePage(path string, fileInfo os.FileInfo) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	buf, meta, err := parseMeta(buf)
	if err != nil {
		return err
	}
	meta.Path = path
	meta.FileInfo = fileInfo
	if meta.Date.IsZero() {
		meta.Date = meta.FileInfo.ModTime()
	}
	g.pages[path] = &Page{
		Site:    g.site,
		Meta:    meta,
		Content: parseContent(buf),
	}
	return nil
}

func parseMeta(buf []byte) ([]byte, *Meta, error) {
	var meta Meta
	var metaBytes []byte
	end := bytes.Index(buf, []byte("\n\n"))
	if end > 0 {
		metaBytes = buf[:end]
	} else {
		metaBytes = buf
	}
	if err := yaml.Unmarshal(metaBytes, &meta); err != nil {
		return nil, nil, err
	}
	return buf[len(metaBytes):], &meta, nil
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
