package main

import (
	"net/http"

	"github.com/saber59/base/function"
)

func main() {
	http.HandleFunc("/look", function.FLook)
	http.HandleFunc("/change", function.Fchange)
	http.HandleFunc("/write", function.Fwrite)
	http.HandleFunc("/gai", function.Fgai)
	http.HandleFunc("/upup", function.Fupup)
	// http.HandleFunc("/subcheck", function.Fsubcheck)
	// http.HandleFunc("/upAllStatus", function.FupAllStatus)
	// http.HandleFunc("/check", function.Fcheck)
	go function.GrpcLook()
	go function.Fsend()
	http.ListenAndServe(":9091", nil)
}
