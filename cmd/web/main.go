package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Sakthiel/bookings/pkgs/config"
	"github.com/Sakthiel/bookings/pkgs/handler"
	"github.com/Sakthiel/bookings/pkgs/render"
	"github.com/alexedwards/scs/v2"
)

const port = ":8080"

var app config.Appconfig
var session *scs.SessionManager

func main() {

	app.UseCache = false
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	
	app.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err.Error())
	}

	app.TemplateCache = tc
	render.NewTemplates(&app)

	repo := handler.NewRepo(&app)
	handler.NewHandler(repo)

	// http.HandleFunc("/", handler.Repo.Home)
	// http.HandleFunc("/about", handler.Repo.About)

	// http.ListenAndServe(port, nil)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err.Error())
	}

}
