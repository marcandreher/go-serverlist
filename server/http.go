package server

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/Gigamons/common/logger"
	"github.com/gorilla/mux"
)

func errHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("------------ERROR------------")
				fmt.Println(err)
				fmt.Println("---------ERROR TRACE---------")
				fmt.Println(string(debug.Stack()))
				fmt.Println("----------END ERROR----------")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// StartServer to start the HTTP Server.
func StartServer(host string, port int16) {
	r := mux.NewRouter()
	r.Use(errHandler)

	logger.Info(" Serverlist is listening on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%v", host, port), r))
}
