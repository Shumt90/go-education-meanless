package main

import (
	"awesomeProject/src/go/config"
	"awesomeProject/src/go/handler"
	"awesomeProject/src/go/service"
	"net/http"
)

func main() {

	config.ReadProperty()
	config.RedisInit()
	config.KafkaInit()

	go service.ReadUpdateSignal()

	service.UpdateRootPageData()

	http.HandleFunc("/", handler.RootHandler())
	_ = http.ListenAndServe(":8080", nil)

	config.RedisDestroy()
}


