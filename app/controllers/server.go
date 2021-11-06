package controllers

import (
	"go-lesson-webapp_211103/config"
	"net/http"
)

func StartMainServer() error {
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
