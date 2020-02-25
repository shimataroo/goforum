package main

import (
	"fmt"
	"net/http"

	"github.com/shimataroo/goforum/data"
)

func login(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "layout", "navbar", "login")
}

func signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "layout", "navbar", "signup")
}

func authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := data.UserByEmail(request.PostFormValue("email"))
	if err != nil {

	}
	if user.Password == request.PostFormValue("password") {
		settion, err := user.CreateSettion()
		if err != nil {
			fmt.Println(err)
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    settion.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}

func logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		session := data.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
		http.Redirect(writer, request, "/", 302)
	}
	http.Redirect(writer, request, "/signup", 302)
}
