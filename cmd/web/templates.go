package main

import (
	"html/template"
	"path/filepath"
	"time"

	"AlikhanAitbayev.net/snippetbox/pkg/forms"
	"AlikhanAitbayev.net/snippetbox/pkg/models"
)

var functions = template.FuncMap{
	"humanDate": humanDate,
}

type templateData struct {
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	CurrentYear     int
	Form            *forms.Form
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
	// FormData url.Values
	// FormErrors map[string]string
}

func newTemplateCache(dir string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join(dir, "*.html"))
	// fmt.Println(pages)
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}
