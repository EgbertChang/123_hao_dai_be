package main

import (
	"123_hao_dai_be/elea"
	"123_hao_dai_be/src/server"

	_ "github.com/go-sql-driver/mysql"

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
