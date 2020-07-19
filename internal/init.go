package internal

import (
	"fmt"
	"net/http"

	"geektest/config"
	"geektest/internal/configuration"
	"geektest/internal/logs"
	"geektest/internal/routing"
)

var	appConfig *configuration.Config
var Lang string

func init () {
	appConfig = config.GetConfig()
}

func StartServer() {
	fmt.Println("Starting server at ", appConfig.Port)
	router:= routing.GetRouter()
	err := http.ListenAndServe(
		fmt.Sprintf(":%d", appConfig.Port), http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
			router.ServeHTTP(w,r)
		}))
	if err != nil {
		logs.Critical("Server FAILED: ", err)
	}
}