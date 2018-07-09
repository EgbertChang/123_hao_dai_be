package server

import (
	"net/http"

	"123_hao_dai/elea"
)

type HttpInterceptor struct{}

func (form *HttpInterceptor) Intercept(w http.ResponseWriter, r *http.Request) bool {
	return true
}

func register(h *elea.Handle) {
	h.Register("/path1", Path1)
	h.Register("/path2", Path2)
	h.Register("/path3", Path3)
}

func Path1(w http.ResponseWriter, r *http.Request) {
	ua := r.Header.Get("user-agent")
	w.Write([]byte("<body><h1>user-agent is: </h1></body>"))
	w.Write([]byte(ua))
}

func Path2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Path2"))
}

func Path3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<body><h1>Hello Path3</h1></body>"))
}

func Handle() elea.HandleSet {
	h := &elea.Handle{}
	register(h)
	return h
}
