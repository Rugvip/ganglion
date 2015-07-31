package api

import (
	"log"
	"net/http"
)

var Root *http.ServeMux

type TestRequest struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

type TestResponse struct {
	OkStatus
	Message string `json:"message"`
}

func handleGet(req *http.Request) WithStatus {
	log.Println("got get", req.URL)
	return &TestResponse{
		Message: "i got get",
	}
}

func handlePost(req *http.Request, body *TestRequest) WithStatus {
	log.Println("got body", body.Code, body.Message)
	return &TestResponse{
		Message: "hello",
	}
}

func init() {
	Root = http.NewServeMux()
	Root.Handle("/", NewJsonHandler(func(req *http.Request) WithStatus {
		return defaultUnrecognizedError
	}))
	Root.Handle("/test", Resource{
		Get:  NewJsonHandler(handleGet),
		Post: NewJsonHandler(handlePost),
	})
	Root.Handle("/login", loginResource)
	Root.Handle("/register", registerResource)
}
