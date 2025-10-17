package main

import (
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/tambah", TambahHandler)
	http.HandleFunc("/simpan", SimpanHandler)
	http.HandleFunc("/edit/", EditHandler)
	http.HandleFunc("/update", UpdateHandler)
	http.HandleFunc("/hapus/", HapusHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
