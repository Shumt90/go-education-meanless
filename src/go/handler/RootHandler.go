package handler

import (
	"awesomeProject/src/go/config"
	"awesomeProject/src/go/service"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"log"
	"net/http"
)

func RootHandler() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		var err error
		var fromRedis string

		fromRedis, err = redis.String(config.Conn.Do("GET", "app:root"))

		if err!=nil {

			print(err.Error())

			respBytes, err := json.Marshal(map[string]string{"error":"no data in redis or impossible to extract"})
			if err != nil {
				log.Print(err)
				return
			}
			_, err = writer.Write(respBytes)
			return
		}

		_, err = writer.Write([]byte(fromRedis))

		service.PushUpdateSignal()

	}
}

