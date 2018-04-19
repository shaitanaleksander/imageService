package handler

import (
	"net/http"
	"github.com/gorilla/mux"
	"imageService/repository"
	"log"
	"bytes"
	"io"
)

func DeleteImage(writer http.ResponseWriter, request *http.Request) {
	name := mux.Vars(request)["name"]
	err := repository.DeleteImage(name)
	if err != nil {
		log.Println(err)
	}
	writer.WriteHeader(http.StatusAccepted)
}

func GetImage(writer http.ResponseWriter, request *http.Request) {
	name := mux.Vars(request)["name"]
	err := repository.DownloadFile(name, writer)
	if err != nil {
		log.Println(err)
		http.Error(writer, "not found", 404)
		return
	}
	return
}

func SaveImage(writer http.ResponseWriter, request *http.Request) {
	var buf bytes.Buffer
	file, header, err := request.FormFile("file")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()
	io.Copy(&buf, file)
	err = repository.UploadFile(header.Filename, &buf)
	if err != nil {
		log.Println(err)
		http.Error(writer, "cannot save image", 500)
	}
}
