package function

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//错误处理
func errworke(code int64, err error, mes string, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}
	var reok Ok
	reok.Code = code
	reok.Message = "error"
	re, err := json.Marshal(reok)
	log.Print(err)
	log.Print(mes)
	fmt.Fprint(w, re)
	return true
}
