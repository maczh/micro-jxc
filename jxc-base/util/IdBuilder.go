package util

import (
	"fmt"
	"ququ.im/common/config"
	"strconv"
)

func GenerateId(prefix, snType, shopId string, l int) string {
	key := "sn:" + snType + ":" + shopId
	s := config.Redis.Incr(key).Val()
	sn := prefix + fmt.Sprintf("%0"+strconv.Itoa(l-len(prefix))+"d", s)
	return sn
}
