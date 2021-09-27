package main

import (
	"fmt"
	"net/http"
)
func index(w http.ResponseWriter, r *http.Request)  {
fmt.Fprintf(w,"<h1>This is my first Golang web project </h1>")
}
func check(w http.ResponseWriter, r *http.Request)  {
fmt.Fprintf(w,"<h1>This is the next page</h1>")
}
func main()  {
http.HandleFunc("/",index)
http.HandleFunc("/check_health",check)
fmt.Println("Server Starting")
http.ListenAndServe(":8150", nil)

}