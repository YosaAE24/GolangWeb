package hendler

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HendlerHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Terjadi kesalahna mohon tenang", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{} {
		"title": "Aku sedang belajar golang web bersama pak angga",
		"content": "Belajar golang itu sangat asik sekali",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Terjadi kesalahna mohon tenang", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Home"))
}

func HendlerHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello dunia, sayang sedang belajar golang web"))
}

func HendlerYosa(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hallo Nama Saya Yosa Adytia Pratana"))
}

func HendlerProduc(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNum, err := strconv.Atoi(id)

	if err != nil || idNum < 1 {
		http.NotFound(w, r)
		return
	}
	// fmt.Fprintf(w, "Proudc page : %d", idNum)
	data := map[string]interface{} {
		// "title": "Aku sedang belajar golang web bersama pak angga",
		"content": idNum,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "produc.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Terjadi kesalahna mohon tenang", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Terjadi kesalahna mohon tenang", http.StatusInternalServerError)
		return
	}
}