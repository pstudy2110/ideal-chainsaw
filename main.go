package main

import (
	"christinaalpha/handlers"
	"christinaalpha/quickdrop"
)

func main() {
	quickdrop.RunApp(handlers.AllRouters())
}
