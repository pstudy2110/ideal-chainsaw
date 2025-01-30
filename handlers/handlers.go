package handlers

import (
	"christinaalpha/quickdrop"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func ThisDirPath() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	return path
}

/*
	func LoadBaseDb() (db *nutsdb.DB) {
		filedirect := filepath.Join(ThisDirPath(), "db/Base")
		opt := nutsdb.DefaultOptions
		opt.Dir = filedirect
		db, err := nutsdb.Open(opt)
		if err != nil {
			log.Fatal(err)
		}
		return db
	}
*/
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	quickdrop.ManualParse(w, nil, "layout", "html/layout.html", "html/main/main.html")
}

func Input(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	StoreAllWebsocket(w, r)
}

func Read(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	quickdrop.ManualParse(w, nil, "layout", "html/layout.html", "html/main/text.html")
}

func AllRouters() *httprouter.Router {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("static"))
	router.GET("/", Index)
	router.GET("/home", Input)
	router.GET("/read", Read)
	return router
}
