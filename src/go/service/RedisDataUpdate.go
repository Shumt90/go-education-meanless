package service

import (
	"awesomeProject/src/go/config"
	"awesomeProject/src/go/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func UpdateRootPageData()  {
	println("Update root page")

	sportResp, err := http.Get(config.GetConfigByValue("app.proxy.url"))

	if err != nil {
		log.Print(err)
		return
	}

	body, err := ioutil.ReadAll(sportResp.Body)

	if err != nil {
		log.Print(err)
		return
	}
	cache := model.RedisWrap{}
	cache.Data=string(body)
	cache.UpdateTime=time.Now()

	jsonRedis,err := json.Marshal(cache)

	_, _ = config.Conn.Do("SET", "app:root", jsonRedis)

}
