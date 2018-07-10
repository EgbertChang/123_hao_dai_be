package server

import (
	"net/http"

	"123_hao_dai_be/elea"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	// 数据库的打开操作，可以转移到其他地方
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

	h.Register("/be/manage/A/add", addA)
	h.Register("/be/manage/B/add", addB)
}

func Handle() elea.HandleSet {
	h := &elea.Handle{}
	register(h)
	return h
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
	Id         int
	Name       string
	PartyAId   int
	PartyAUrl  string
	PartyBUrl  string
	ClickCount int
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
		_ = rows.Scan(&b.Id, &b.Name, &b.PartyAId, &b.PartyAUrl, &b.PartyBUrl, &b.ClickCount)
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

type addBParams struct {
	PartyAId  int    `json:"partyAId"`
	Name      string `json:"name"`
	PartyAUrl string `json:"partyAUrl"`
	PartyBUrl string `json:"partyBUrl"`
}

func addB(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	var params addBParams
	json.Unmarshal(bodyBytes, &params)
	stmt, err := db.Prepare(insertB)
	defer stmt.Close()
	_, err = stmt.Exec(params.Name, params.PartyAId, params.PartyAUrl, params.PartyBUrl)
	if err != nil {
		log.Println(err)
	}
}

func addA(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)

	fmt.Println("新增A接口")
}
