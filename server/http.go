package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gigamons/common/logger"
	"github.com/gorilla/mux"
)

// StartServer to start the HTTP Server.
func StartServer(host string, port int16) {
	r := mux.NewRouter()

	logger.Info(" Serverlist is listening on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%v", host, port), r))
}
