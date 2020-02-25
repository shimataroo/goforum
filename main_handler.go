package main

import (
	"fmt"
	"net/http"

	"github.com/shimataroo/goforum/data"
)

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	_, err = session(writer, request)
	if err != nil {
		generateHTML(writer, threads, "layout", "public_navbar", "index")
	} else {
		generateHTML(writer, threads, "layout", "private_navbar", "index")
	}
}

func signupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		fmt.Println("errr")
	}
	user := data.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Create(); err != nil {
		panic(err)
	}
	http.Redirect(writer, request, "/login", 302)
}
