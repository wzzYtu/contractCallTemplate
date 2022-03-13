package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"

	"contractCallTemplate/model"
)

var Conf model.Conf

func GetConf() {
	if _, err := toml.DecodeFile("./conf/conf.toml", &Conf); err != nil {
		fmt.Println("decode config file err")
	}
}
