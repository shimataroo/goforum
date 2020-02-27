package main

import (
	"net/http"
	"os"
)

func getListenPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":3000"
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// index
	mux.HandleFunc("/", index)

	// auth_handler.go
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_Account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// thread_handler.go
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)

	// mypage_handler
	mux.HandleFunc("/mypage", mypage)
	mux.HandleFunc("/upload", upload)

	server := http.Server{
		Addr: getListenPort(),
	}
	server.ListenAndServe()
}
