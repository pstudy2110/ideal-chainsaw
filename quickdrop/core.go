package quickdrop

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"text/template"
)

// total implementation of a full running app

func ReturnWD() string {
	// Return working directory
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func Static() {
	// Serve Static file , very useful for local development
	files := http.FileServer(http.Dir(ReturnWD()))
	http.Handle("/static/", files)
}

func ReturnPort() string {
	// return the connect host else localhost:8072
	port := os.Getenv("PORT")
	if port == "" {
		port = "8072"
	}
	return port
}

//BASIC PARSING

func ManualParse(w http.ResponseWriter, data interface{}, definedtemplate string, files ...string) {
	// One of the many questionable thing that i dont understand.
	// So if you dont have any data for the interface simply put "nil",
	// For definedtemplate "simple string of the -layout-" from {{define "layout"}}
	// For files , must add all files including -base layout file- starting from root working directory
	// Let say your layout is  (w , nil, "base","wd/layout.html", "wd/partA.html", "wd/partB.html")
	var filecontain []string
	for _, file := range filecontain {
		files = append(files, fmt.Sprint(file))
	}
	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(w, definedtemplate, data)
}

//open browser

func Open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
