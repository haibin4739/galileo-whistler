package main

import (
	"fmt"
	"net/http"

	"guthub.com/haibin4739/galileo-whistler/conf/setting"
	"guthub.com/haibin4739/galileo-whistler/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.Server.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.Server.ReadTimeout,
		WriteTimeout:   setting.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
