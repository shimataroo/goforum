package main

import (
	"io"
	"net/http"
	"os"

	"github.com/shimataroo/goforum/data"
)

func mypage(writer http.ResponseWriter, request *http.Request) {
	_, id := data.ReadSession()
	user, _ := data.UserNameById(id)
	generateHTML(writer, user, "layout", "private_navbar", "mypage")
}

func upload(writer http.ResponseWriter, request *http.Request) {
	request.ParseMultipartForm(32 << 20)
	file, _, err := request.FormFile("uploadfile")
	if err != nil {
		http.Redirect(writer, request, "/", 302)
	}
	defer file.Close()
	f, err := os.OpenFile("./public/upload/"+"sample.png", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	io.Copy(f, file)

	http.Redirect(writer, request, "/mypage", 302)
}
