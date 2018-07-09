package server

import (
	"net/http"

	"123_hao_dai_be/elea"
	"bytes"
	"database/sql"
	"encoding/json"
)

var db *sql.DB

func init() {
	// 无法在这里初始化数据库的连接状态
}

type HttpInterceptor struct{}

func (form *HttpInterceptor) Intercept(w http.ResponseWriter, r *http.Request) bool {
	return true
}

func register(h *elea.Handle) {
	var err error
	db, err = sql.Open("mysql", "root:wenjiamin@tcp(139.196.74.31:3306)/123_hao_dai")
	defer func() {
		_ = recover()
	}()
	if err != nil {
		panic(err)
	}
	h.Register("/path1", Path1)
	h.Register("/path2", Path2)
	h.Register("/path3", Path3)
}

type A struct {
	Id   int
	Name string
}

func Path1(w http.ResponseWriter, r *http.Request) {
	defer func() {
		recover()
	}()
	a := &A{}
	err := db.QueryRow(selectAInfo).Scan(&a.Id, &a.Name)
	if err != nil {
		panic(err)
		return
	}
	aBytes, _ := json.Marshal(a)
	w.Write(aBytes)
}

type B struct {
	Id   int
	Name string
	AID  int
	Url  string
}

func Path2(w http.ResponseWriter, r *http.Request) {
	defer func() {
		recover()
	}()
	rows, err := db.Query(selectBInfoList, 0, 20)
	if err != nil {
		panic(err)
		return
	}
	defer rows.Close()
	var BList [][]byte
	for rows.Next() {
		b := &B{}
		_ = rows.Scan(&b.Id, &b.Name, &b.AID, &b.Url)
		bBytes, _ := json.Marshal(b)
		BList = append(BList, bBytes)
	}
	// 创建一个 nil slice 直接使用
	var bs []byte
	bs = append(bs, []byte("[")...)
	bs = append(bs, bytes.Join(BList, []byte(","))...)
	bs = append(bs, []byte("]")...)
	w.Write(bs)
}

func Path3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<body><h1>Hello Path3</h1></body>"))
}

func Handle() elea.HandleSet {
	h := &elea.Handle{}
	register(h)
	return h
}
