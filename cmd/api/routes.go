package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/status", app.statusHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users/register", app.userRegisterHandler)
	router.HandlerFunc(http.MethodPost, "/v1/users/login", app.userLoginHandler)

	return app.authenticate(router)
}
