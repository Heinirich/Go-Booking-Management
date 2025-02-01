package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func writeToConsole(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w,r)
	})
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		Path: "/",
		HttpOnly: true,
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler

}

func SessionLoad(next http.Handler) http.Handler{
	return session.LoadAndSave(next)
}