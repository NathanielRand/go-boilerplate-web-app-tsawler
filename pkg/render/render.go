package render

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

// RenderTemplate renders templates
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	
	parsedTemplate, err := template.ParseFiles("./../../templates/" + tmpl)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error executing parsed template: ", err)
		return
	}
}

// CreateTemplateCache creates a template cache as a map.
func CreateTemplateCache() (map[string]*template.Template, error) {
	
	// 	Create a cache to store our parsed templates.
	myCache := map[string]*template.Template{}

	// 	Find pages in template folder.
	pages, err := filepath.Glob("./templates/*.pages.html")
	if err != nil {
		return myCache, err
	}

	// 	Range through pages found in template folder.
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
