package main

import (
	"html/template"
	"log"
	"net/http"
)

// Fungsi handler untuk homepage
func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	data := map[string]string{
		"Title":       "Selamat datang di Store Maulana Nur Anfajm",
		"Description": "Happy Shopping",
	}
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", homePage)
	log.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
