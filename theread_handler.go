package main

import (
	"fmt"
	"net/http"

	"github.com/shimataroo/goforum/data"
)

func newThread(writer http.ResponseWriter, request *http.Request) {
	_, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		generateHTML(writer, nil, "layout", "private_navbar", "new_thread")
	}
}

func readThread(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	uuid := vals.Get("id")
	thread, err := data.ThreadByUUID(uuid)
	if err != nil {
		panic(err)
	} else {
		_, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &thread, "layout", "public_navbar", "public_thread")
		} else {
			generateHTML(writer, &thread, "layout", "private_navbar", "private_thread")
		}
	}
}

func createThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			panic(err)
		}
		user, err := sess.User()
		if err != nil {
			panic(err)
		}
		topic := request.PostFormValue("topic")
		if _, err := user.CreateThread(topic); err != nil {
			panic(err)
		}
		http.Redirect(writer, request, "/", 302)
	}
}

func postThread(writer http.ResponseWriter, request *http.Request) {
	sess, err := session(writer, request)
	if err != nil {
		http.Redirect(writer, request, "/login", 302)
	} else {
		err = request.ParseForm()
		if err != nil {
			panic(err)
		}
		user, err := sess.User()
		if err != nil {
			panic(err)
		}
		body := request.PostFormValue("body")
		uuid := request.PostFormValue("uuid")
		thread, err := data.ThreadByUUID(uuid)
		if err != nil {
			panic(err)
		}
		if _, err := user.CreatePost(thread, body); err != nil {
			panic(err)
		}
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer, request, url, 302)
	}
}
