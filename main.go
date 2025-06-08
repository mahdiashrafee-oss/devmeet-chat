package main

import "net/http"
import "html/template"

func homepage(w http.ResponseWriter, r *http.Request){
	tmpl, _ := template.ParseFiles("templates/login.html")

	tmpl.Execute(w, nil)

	

} 
func main(){
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(homepage))

	server := http.Server{
			Addr: ":8080",
		  Handler: mux,
	}

	server.ListenAndServe()


}

