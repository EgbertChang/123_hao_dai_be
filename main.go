package main

import (
	"123_hao_dai/elea"
	"123_hao_dai/src/server"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	Server := elea.Server{
		Addr:        ":8080",
		Handle:      server.Handle(),
		Interceptor: &server.HttpInterceptor{},
	}
	Server.ListenAndServer()
}
