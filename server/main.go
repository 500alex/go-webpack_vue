package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../templates/index.html", "../templates/header.html", "../templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Println("get index")
	t.ExecuteTemplate(w, "index", nil)
}
func main() {
	fmt.Println("listening on port: 3000")
	http.Handle("/build/", http.StripPrefix("/build/", http.FileServer(http.Dir("../build/"))))
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}
