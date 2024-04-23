package main

import "net/http"

type Middleware func(http.Handler) http.Handler

func (app *application) withSession(f func(http.ResponseWriter, *http.Request)) http.Handler {
	return app.sessionManager.LoadAndSave(http.HandlerFunc(f))
}

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.Handle("GET /{$}", app.withSession(app.home))
	mux.Handle("GET /snippet/view/{id}", app.withSession(app.snippetView))

	mux.Handle("GET /user/signup", app.withSession(app.userSignup))
	mux.Handle("POST /user/signup", app.withSession(app.userSignupPost))
	mux.Handle("GET /user/login", app.withSession(app.userLogin))
	mux.Handle("POST /user/login", app.withSession(app.userLoginPost))

	mux.Handle("GET /snippet/create", app.withSession(app.snippetCreate))
	mux.Handle("POST /snippet/create", app.withSession(app.snippetCreatePost))
	mux.Handle("POST /user/logout", app.withSession(app.userLogoutPost))

	middlewares := []Middleware{
		commonHeaders,
		app.logRequest,
		app.recoverPanic,
	}
	handler := http.Handler(mux)
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}
