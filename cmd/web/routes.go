package main

import "net/http"

type Middleware func(http.Handler) http.Handler

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	middlewares := []Middleware{
		commonHeaders,
		app.logRequest,
		app.recoverPanic,
	}

	var handler http.Handler = mux
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}
