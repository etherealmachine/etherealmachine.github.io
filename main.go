package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

var (
	preview = flag.Bool("preview", false, "Serve locally for preview")
)

type generator struct {
	site     *Site
	errors   []error
	pages    map[string]*Page
	template *template.Template
}

func (g *generator) walk(path string) error {
	filepath.Walk(filepath.Join(path, "pages"), g.visit)
	var err error
	g.template, err = template.New("").Funcs(templateFunctions).
		ParseGlob(filepath.Join(path, "templates", "*.html"))
	if err != nil {
		return err
	}
	return nil
}

func (g *generator) visit(path string, fileInfo os.FileInfo, err error) error {
	if strings.HasSuffix(path, ".md") {
		page, err := g.newPage(path, fileInfo)
		if err != nil {
			g.errors = append(g.errors, err)
			return nil
		}
		g.pages[path] = page
	}
	return nil
}

type ByTime []*Page

func (a ByTime) Len() int      { return len(a) }
func (a ByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool {
	return a[i].Meta.FileInfo.ModTime().Before(a[j].Meta.FileInfo.ModTime())
}

func (g *generator) generate() error {
	for _, page := range g.pages {
		f, err := os.Create(url(page))
		if err != nil {
			return err
		}
		g.template.ExecuteTemplate(f, "page.html", page)
		g.site.Recent = append(g.site.Recent, page)
	}
	sort.Sort(ByTime(g.site.Recent))
	if len(g.site.Recent) > 4 {
		g.site.Recent = g.site.Recent[:5]
	}
	f, err := os.Create("index.html")
	if err != nil {
		return err
	}
	g.template.ExecuteTemplate(f, "index.html", g.site)
	return nil
}

func runGenerator() {
	g := &generator{
		site: &Site{
			PreviewScript: previewScript,
		},
		pages: make(map[string]*Page),
	}
	if err := g.walk("."); err != nil {
		log.Fatal(err)
	}
	if len(g.errors) > 0 {
		for _, err := range g.errors {
			log.Println(err)
		}
	}
	if err := g.generate(); err != nil {
		log.Fatal(err)
	}
}

type Site struct {
	Recent        []*Page
	PreviewScript string
}

func main() {
	flag.Parse()
	runGenerator()
	if *preview {
		runPreview()
	}
}
