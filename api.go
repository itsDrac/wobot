package main

import (
	"net/http"
)

type API interface {
	Mount(http.Handler)
	Health(http.ResponseWriter, *http.Request)
	Run(http.Handler)
}
