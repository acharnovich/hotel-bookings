package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/acharnovich/hotel-bookings/pkg/config"
	"github.com/acharnovich/hotel-bookings/pkg/models"
)

var app *config.AppConfig
func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData)*models.TemplateData{
	
	return td
}

// Renders html templates
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get template
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not create template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	// render template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// create a map of pointers to template.Template
	myCache := map[string]*template.Template{}
	
	// get all files named .page.tmpl under the templates folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files with .page.tmpl
	for _, page := range pages {
		// name is only the filename and strips rest of the path
		name := filepath.Base(page)
		// pointer to template.Template with filename
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// gets layout files
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		// if layout file is found, parse it
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		// adds template to cache
		myCache[name] = ts
	}
	return myCache, err
}