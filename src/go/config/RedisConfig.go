package config

import (
	"github.com/gomodule/redigo/redis"
)


var Conn redis.Conn

func RedisInit()  {
	var err error
	Conn, err = redis.Dial("tcp", GetConfigByValue("redis.host"))
	if err != nil {
		panic(err.Error())
	}
}

func RedisDestroy()  {
	_ = Conn.Close()
}