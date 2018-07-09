package elea

import (
	"log"
	"net/http"
	"time"
)

var mux *http.ServeMux

func init() {
	// singleton pattern
	mux = http.NewServeMux()
}

type HttpInterceptor interface {
	Intercept(w http.ResponseWriter, r *http.Request) bool
}

type HandleSet interface {
	GetHandle() map[string]http.HandlerFunc
}

type Server struct {
	Addr        string
	Handle      HandleSet
	Interceptor HttpInterceptor
}

func (sr *Server) ListenAndServer() {
	for k, v := range sr.Handle.GetHandle() {
		mux.HandleFunc(k, v)
	}
	server := http.Server{
		Addr:           sr.Addr,
		Handler:        &scheduler{sr.Interceptor},
		ReadTimeout:    3 * time.Minute,
		WriteTimeout:   3 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Elea Server is listening on " + sr.Addr)
	log.Fatalln(server.ListenAndServe())
}

// 调度器类用于处理所以http请求的顶层逻辑
type scheduler struct {
	interceptor HttpInterceptor
}

func (form *scheduler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if form.interceptor.Intercept(w, r) {
		mux.ServeHTTP(w, r)
	}
}
