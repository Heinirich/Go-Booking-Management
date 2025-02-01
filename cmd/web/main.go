package main

import (
	"github.com/alexedwards/scs/v2"
	"github.com/heinirich/bookings/pkg/config"
	"github.com/heinirich/bookings/pkg/handlers"
	"github.com/heinirich/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const port=":8000"
var app config.AppConfig
var session *scs.SessionManager


func main(){


	// change to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24*time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc,err := render.CreateTemplateCache()

	if err != nil{
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)

	server := &http.Server{
		Addr: port, Handler: routes(&app),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}



