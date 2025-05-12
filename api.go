package main

import (
	"net/http"
)

type API interface {
	Mount(http.Handler)
	Health(http.ResponseWriter, *http.Request)
	CreateUser(http.ResponseWriter, *http.Request)
	LoginUser(http.ResponseWriter, *http.Request)
	AuthUser(http.Handler) http.Handler
	GetStorage(http.ResponseWriter, *http.Request)
	Run(http.Handler)
}
