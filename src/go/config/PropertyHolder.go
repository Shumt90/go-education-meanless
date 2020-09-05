package config

import (
	"github.com/magiconair/properties"
)

var prop *properties.Properties

func ReadProperty()  {
	prop = properties.MustLoadFile("./conf.properties", properties.UTF8)
}

func GetConfigByValue(key string) string  {
	get, ok := prop.Get(key)
	if !ok {
		panic("Can't read property: "+key)
	}
	return get
}