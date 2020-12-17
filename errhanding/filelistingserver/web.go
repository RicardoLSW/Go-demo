package main

import (
	"go-demo/errhanding/filelistingserver/fileListing"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

type userError interface {
	error
	Message() string
}

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Printf("Panic: %v", r)
			}
		}()
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling request: %s", err.Error())
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(fileListing.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
