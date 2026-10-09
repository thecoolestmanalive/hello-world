package main

import
 (
	"fmt"
	"net/http"
)

func api(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "success")
}

func main(){
	http.HandleFunc("/get", api)

	fmt.Println("api on http://localhost:3000")

	http.ListenAndServe(":3000", nil)
}