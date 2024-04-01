package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Sakthiel/bookings/models"
	"github.com/Sakthiel/bookings/pkgs/config"
)

var appConfig *config.Appconfig

func NewTemplates(app *config.Appconfig) {
	appConfig = app
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	var tc map[string]*template.Template

	if appConfig.UseCache {
		tc = appConfig.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("no template found in cache")
	}

	buf := new(bytes.Buffer)

	data = AddDefaultData(data)
	
	err := t.Execute(buf, data)

	if err != nil {
		log.Println(err.Error())
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println(err.Error())
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.gohtml")

	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")

			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
