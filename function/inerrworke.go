package function

import (
	"log"
)

//错误处理
func inerrworke(code int64, err error, mes string) {
	if err == nil {
		return
	}

	log.Print(mes)

}
