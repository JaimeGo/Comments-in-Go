package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
	"time"
)

type Comment struct{
	Content string
	Hour time.Time
	Ip string
}

var allComments = make([]Comment, 0)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/comment", comment)
	fmt.Println("Server running in port 8080")
	http.ListenAndServe(":80", nil)
}

func index(w http.ResponseWriter, _ *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", allComments)
	HandleError(w, err)
}

func comment(w http.ResponseWriter, r *http.Request) {
    
    r.ParseForm()
    // logic part of log in
    content:=r.Form["comment"][0]
    t := time.Now()
    ip:=r.RemoteAddr
    comment:=Comment{Content:content,Hour:t,Ip:ip}
    fmt.Println("Comment: ", comment.Content)
    allComments=append(allComments, comment)
    http.Redirect(w, r, "/", http.StatusSeeOther)
    
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}