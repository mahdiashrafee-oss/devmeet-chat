package main

import "fmt"
import "net/http"
import "html/template"

var username string
var password string
func login(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodPost{
		r.ParseForm()
		username =r.FormValue("username")
		password =r.FormValue("password")
		fmt.Println(username)
		fmt.Println(password)
		if username == "test" && password == "1234"{
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}
		tmpl, _ := template.ParseFiles("templates/login.html")

	tmpl.Execute(w, nil)

	
} 
func homepage(w http.ResponseWriter, r *http.Request){
	tmpl, _ :=  template.ParseFiles("templates/index.html")

	tmpl.Execute(w, username)

}
func main(){
	mux := http.NewServeMux()
	mux.Handle("/login",http.HandlerFunc(login))
	mux.Handle("/",http.HandlerFunc(homepage))

	server := http.Server{
			Addr: ":8080",
		  Handler: mux,
	}

	server.ListenAndServe()


}

