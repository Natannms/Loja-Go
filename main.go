package main

import (
	"goweb/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/create", controllers.Create)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/updatePage", controllers.UpdatePge)
	http.HandleFunc("/delete", controllers.Destroy)
	http.ListenAndServe(":8000", nil)
}
