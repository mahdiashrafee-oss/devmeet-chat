package main

import "fmt"
import "net/http"
import "html/template"

func homepage(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost{
		r.ParseForm()
		username :=r.FormValue("username")
		password :=r.FormValue("password")
		fmt.Println(username)
		fmt.Println(password)
	tmpl, _ := template.ParseFiles("templates/login.html")

	tmpl.Execute(w, nil)

	}
} 
func main(){
	mux := http.NewServeMux()
	mux.Handle("/",http.HandlerFunc(homepage))

	server := http.Server{
			Addr: ":8080",
		  Handler: mux,
	}

	server.ListenAndServe()


}

