package main

import (
	"log"
	"net/http"
	"golangweb/hendler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", hendler.HendlerHome)
	mux.HandleFunc("/hello", hendler.HendlerHello)
	mux.HandleFunc("/yosa", hendler.HendlerYosa)
	mux.HandleFunc("/produc", hendler.HendlerProduc)

	FileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", FileServer))

	log.Println("Strating Web On Port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

