package routers

import (
	v1 "main/server/routers/v1"
	"net/http"
)

//import ("main/v1")

func GetRouter(name string) http.Handler {
	var defaultRouter = v1.V1Router()
	switch name {
	default:
		return defaultRouter
	}
}
