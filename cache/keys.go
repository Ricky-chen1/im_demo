package cache

import (
	"fmt"
)

func GetMsgExpiredTimeKey(id string) string {
	return fmt.Sprintf("msgExpiredTime:%s", id)
}
