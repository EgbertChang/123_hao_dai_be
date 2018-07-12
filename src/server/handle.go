package server

import (
	"net/http"

	"123_hao_dai_be/elea"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var db *sql.DB

func init() {
	// 无法在这里初始化数据库的连接状态
}

type HttpInterceptor struct{}

func (form *HttpInterceptor) Intercept(w http.ResponseWriter, r *http.Request) bool {
	return true
}

func registerHandler(h *elea.Handle) {
	// 数据库的打开操作，可以转移到其他地方
	var err error
	db, err = sql.Open("mysql", "root:wenjiamin@tcp(139.196.74.31:3306)/123_hao_dai")
	defer func() {
		if e := recover(); e != nil {
			// recover()检测当前函数是否存在panic()，然后恢复它！
			fmt.Printf("Panicking %s\r\n", e)
		}
	}()
	if err != nil {
		panic(err)
	}

	h.Register("/be/manage/A/add", addA)
	h.Register("/be/manage/A/list", listA)
	h.Register("/be/manage/A/delete/", deleteA)
	h.Register("/be/manage/B/add", addB)
	h.Register("/be/manage/B/list", listB)
	h.Register("/be/manage/B/delete/", deleteB)
}

func Handle() elea.HandleSet {
	h := &elea.Handle{}
	registerHandler(h)
	registerFileServer(h)
	return h
}

func addA(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	var params addAParams
	json.Unmarshal(bodyBytes, &params)
	stmt, err := db.Prepare(insertASql)
	_, err = stmt.Exec(params.Name)
	ret := CreateResult{Msg: "success"}
	retBytes, err := json.Marshal(ret)
	w.Write(retBytes)
	defer func() {
		r.Body.Close()
		stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()
}

func listA(w http.ResponseWriter, r *http.Request) {
	// 根据A的id返回B的数量
	bodyBytes, err := ioutil.ReadAll(r.Body)
	var params listAParams
	json.Unmarshal(bodyBytes, &params)
	rows, err := db.Query(selectAllASql, (params.PageIndex-1)*params.PageSize, params.PageSize)
	var AList []A
	for rows.Next() {
		a := &A{}
		err = rows.Scan(&a.Id, &a.Name, &a.BNum)
		AList = append(AList, *a)
	}
	ret := RetrieveResult{Msg: "success", Data: AList}
	retBytes, err := json.Marshal(ret)
	w.Write(retBytes)
	defer func() {
		rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()
}

func deleteA(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	steps := strings.Split(path, "/")
	id := steps[len(steps)-1]
	stmt, err := db.Prepare(deleteASql)
	_, err = stmt.Exec(id)
	ret := DeleteResult{Msg: "success"}
	retBytes, err := json.Marshal(ret)
	w.Write(retBytes)
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()
}

func addB(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	var params addBParams
	json.Unmarshal(bodyBytes, &params)
	stmt, err := db.Prepare(insertBSql)
	_, err = stmt.Exec(params.Name, params.PartyAId, params.PartyAUrl, params.PartyBUrl)
	ret := CreateResult{Msg: "success"}
	retBytes, err := json.Marshal(ret)
	w.Write(retBytes)
	defer func() {
		r.Body.Close()
		stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()
}

func listB(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	var params listBParams
	json.Unmarshal(bodyBytes, &params)
	rows, err := db.Query(selectBListSql, params.PartyAId, (params.PageIndex-1)*params.PageSize, params.PageSize)
	if err != nil {
		// 数据库抛出的错误
		ret := RetrieveResult{Msg: "failure", Data: []int{}}
		retBytes, _ := json.Marshal(ret)
		w.Write(retBytes)
		return
	}
	var BList []B
	for rows.Next() {
		b := &B{}
		_ = rows.Scan(&b.Id, &b.Name, &b.PartyAUrl, &b.PartyBUrl, &b.ClickCount)
		BList = append(BList, *b)
	}
	// var bs []byte 创建一个 nil slice 直接使用
	// 这里可以使用装饰器模式
	// var ret Result
	ret := RetrieveResult{Msg: "success", Data: BList}
	retBytes, err := json.Marshal(ret)
	w.Write(retBytes)
	defer func() {
		rows.Close()
	}()
}

func deleteB(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	steps := strings.Split(path, "/")
	id := steps[len(steps)-1]
	stmt, err := db.Prepare(deleteBSql)
	_, err = stmt.Exec(id)
	ret := DeleteResult{Msg: "success"}
	retBytes, err := json.Marshal(ret)
	w.Write(retBytes)
	defer func() {
		log.Println(err)
	}()
}
