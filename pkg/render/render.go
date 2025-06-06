package render

import (
	"bytes"
	"fmt"
	"github.com/heinirich/bookings/pkg/config"
	"github.com/heinirich/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)


var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplate sets config for templatecache
func NewTemplate(a *config.AppConfig){
	app = a
}


func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	return td
}

// RenderTemplate render HTML template
func RenderTemplate(w http.ResponseWriter,tmpl string,td *models.TemplateData){
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	}else{
		tc,_ = CreateTemplateCache()
	}

	t,ok := tc[tmpl]

	if !ok{
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf,td)

	_,err := buf.WriteTo(w)

	if err != nil{
		fmt.Println("Error writing template on browser")
	}


	parsedTemplate,_ := template.ParseFiles("./templates/"+tmpl)

	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		return
	}

}

// CreateTemplateCache renders template cache as a map
func CreateTemplateCache() (map[string]*template.Template,error) {

	myCache := map[string]*template.Template{}

	// get all matching files for the given pattern
	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache,err
	}



	for _,page := range pages{

		// Get last element of the path
		name := filepath.Base(page)

		fmt.Println(name)

		// Create a template with the provided name
		// Pass functions
		// Parses the named files and associates the resulting templates with t
		ts,err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache,err
		}
		// Glob returns the names of all files matching pattern
		matches,err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return myCache,err
		}
		if len(matches) > 0 {
			// Get parses the template definitions in the files
			ts,err = ts.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return myCache,err
			}
		}

		myCache[name] = ts

	}

	return myCache,nil

}
