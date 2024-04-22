package main

import "net/http"

type Middleware func(http.Handler) http.Handler

func (app *application) dynamic(f func(http.ResponseWriter, *http.Request)) http.Handler {
	return app.sessionManager.LoadAndSave(http.HandlerFunc(f))
}

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.Handle("GET /{$}", app.dynamic(app.home))
	mux.Handle("GET /snippet/view/{id}", app.dynamic(app.snippetView))
	mux.Handle("GET /snippet/create", app.dynamic(app.snippetCreate))
	mux.Handle("POST /snippet/create", app.dynamic(app.snippetCreatePost))

	mux.Handle("GET /user/signup", app.dynamic(app.userSignup))
	mux.Handle("POST /user/signup", app.dynamic(app.userSignupPost))
	mux.Handle("GET /user/login", app.dynamic(app.userLogin))
	mux.Handle("POST /user/login", app.dynamic(app.userLoginPost))
	mux.Handle("POST /user/logout", app.dynamic(app.userLogoutPost))

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
