package contractCallTemplate

import (
	"contractCallTemplate/conf"
	"contractCallTemplate/service"
	"testing"
)

func TestT(t *testing.T) {
	conf.GetConf()
	service.SendETH("0xE9C16613F962d5C27E385927BC3C22Ef25cb9afB", "0x17aa924282bf0Fa76CF3e80eC3EaA2a54C649186", 1.111111)
	// bl := service.GetBalance("0xE9C16613F962d5C27E385927BC3C22Ef25cb9afB")
	// fmt.Println("balance is:", bl)
}
