package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle(
		"GET /static/",
		http.StripPrefix(
			"/static",
			http.FileServer(http.Dir("./ui/static/")),
		),
	)
	mux.Handle(
		"GET /{$}",
		app.dynamic(http.HandlerFunc(app.home)),
	)
	mux.Handle(
		"GET /snippet/view/{id}",
		app.dynamic(http.HandlerFunc(app.snippetView)),
	)
	mux.Handle(
		"GET /user/signup",
		app.dynamic(http.HandlerFunc(app.userSignup)),
	)
	mux.Handle(
		"POST /user/signup",
		app.dynamic(http.HandlerFunc(app.userSignupPost)),
	)
	mux.Handle(
		"GET /user/login",
		app.dynamic(http.HandlerFunc(app.userLogin)),
	)
	mux.Handle(
		"POST /user/login",
		app.dynamic(http.HandlerFunc(app.userLoginPost)),
	)
	mux.Handle(
		"GET /snippet/create",
		app.dynamicWithAuth(http.HandlerFunc(app.snippetCreate)),
	)
	mux.Handle(
		"POST /snippet/create",
		app.dynamicWithAuth(http.HandlerFunc(app.snippetCreatePost)),
	)
	mux.Handle(
		"POST /user/logout",
		app.dynamicWithAuth(http.HandlerFunc(app.userLogoutPost)),
	)
	return app.standard(http.Handler(mux))
}
