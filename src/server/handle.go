package server

import (
	"database/sql"
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
	defer func() {
		_ = recover()
	}()
	_, err := sql.Open("mysql", "root:wenjiamin@tcp(139.196.74.31:3306)/123_hao_dai")
	if err != nil {
		panic(err)
	}
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
