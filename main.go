package main

import
 (
	"fmt"
	"net/http"
)

func api(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "success")
}

func page(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r, "index.html")
}

func get (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "main page")
}

func main(){
	http.HandleFunc("/get", api)

	http.HandleFunc("/page", page)

	http.HandleFunc("/", get)

	fmt.Println("api on http://localhost:3000")

	http.ListenAndServe(":3000", nil)
}