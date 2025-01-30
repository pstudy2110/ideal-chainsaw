package quickdrop

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var version string = "0.01"

func QuickdropVersion() {
	fmt.Println(version)
}

func RunApp(router *httprouter.Router) {
	server := http.Server{
		Addr:    ":" + ReturnPort(),
		Handler: router,
	}
	fmt.Printf("App host on %v", server.Addr)
	server.ListenAndServe()
}

func RunAppLocal(router *httprouter.Router) {
	server := http.Server{
		Addr:    "localhost:8089",
		Handler: router,
	}
	fmt.Printf("App host on %v", server.Addr)
	Open("http://" + server.Addr)
	server.ListenAndServe()
}
