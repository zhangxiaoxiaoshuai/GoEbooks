package main

import (
	"log"
	"net/http"
	"html/template"
)

func body(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("test.html"))
	t.Execute(w, `<script>alert(123);</script>`)
}


func main() {
	http.HandleFunc("/", body)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
