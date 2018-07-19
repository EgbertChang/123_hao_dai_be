package server

import (
	"net/http"

	"123_hao_dai_be/elea"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
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

	h.Register("/be/manage/product/add", addProduct)
	h.Register("/be/manage/product/delete/", deleteProduct)
	h.Register("/be/manage/product/edit/", editProduct)
	h.Register("/be/manage/product/detail/", productDetail)
	h.Register("/be/manage/product/list", listProduct)
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
	ret := CreateResponse{Msg: "success"}
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
	if AList == nil {
		AList = []A{}
	}
	ret := RetrieveResponse{Msg: "success", Data: AList}
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
	ret := DeleteResponse{Msg: "success"}
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
	partyAId, _ := strconv.Atoi(params.PartyAId)
	_, err = stmt.Exec(params.Name, partyAId, params.PartyAUrl, params.PartyBUrl)
	ret := CreateResponse{Msg: "success"}
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
	partyAId, _ := strconv.Atoi(params.PartyAId)
	rows, err := db.Query(selectBListSql, partyAId, (params.PageIndex-1)*params.PageSize, params.PageSize)
	if err != nil {
		// 数据库抛出的错误
		ret := RetrieveResponse{Msg: "failure", Data: []int{}}
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
	if BList == nil {
		BList = []B{}
	}
	// var bs []byte 创建一个 nil slice 直接使用
	// 这里可以使用装饰器模式
	// var ret Result
	ret := RetrieveResponse{Msg: "success", Data: BList}
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
	ret := DeleteResponse{Msg: "success"}
	retBytes, err := json.Marshal(ret)
	w.Write(retBytes)
	defer func() {
		log.Println(err)
	}()
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	var params product
	json.Unmarshal(bodyBytes, &params)
	categorySting := strings.Join(params.Type, ",")
	personalQualificationSting := strings.Join(params.PersonalQualification, ",")
	termString := strings.Join(params.Term, ",")
	var res CreateResponse
	stmt, err := db.Prepare(insertProductSql)
	if err != nil {
		res.Msg = "failure"
		retBytes, _ := json.Marshal(res)
		w.Write(retBytes)
		return
	}
	InterestBytes, _ := json.Marshal(params.Interest)
	LendingRateBytes, _ := json.Marshal(params.LendingRate)
	_, err = stmt.Exec(params.Name,
		params.Url,
		categorySting,
		personalQualificationSting,
		params.LimitMin,
		params.LimitMax,
		params.LogoUrl,
		params.Slogan,
		params.ApplyNumber,
		termString,
		InterestBytes,
		LendingRateBytes,
		params.Credit,
		params.AuditType,
		params.AccountInType,
		params.ApplyStrategy)
	if err != nil {
		res.Msg = "failure"
	} else {
		// ret.LastInsertId()  // 如果是数据库返回err，执行这行代码会导致接口崩溃
		res.Msg = "success"
	}
	retBytes, _ := json.Marshal(res)
	w.Write(retBytes)
	defer func() {
		r.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	steps := strings.Split(path, "/")
	id := steps[len(steps)-1]
	stmt, err := db.Prepare(deleteProductSql)
	_, err = stmt.Exec(id)
	ret := DeleteResponse{Msg: "success"}
	retBytes, err := json.Marshal(ret)
	w.Write(retBytes)
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()
}

func editProduct(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	steps := strings.Split(path, "/")
	id := steps[len(steps)-1]
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	var params product
	json.Unmarshal(bodyBytes, &params)
	categorySting := strings.Join(params.Type, ",")
	personalQualificationSting := strings.Join(params.PersonalQualification, ",")
	termString := strings.Join(params.Term, ",")
	res := UpdateResponse{}
	stmt, err := db.Prepare(updateProductSql)
	if err != nil {
		res.Msg = "failure"
		retBytes, _ := json.Marshal(res)
		w.Write(retBytes)
		return
	}
	InterestBytes, _ := json.Marshal(params.Interest)
	LendingRateBytes, _ := json.Marshal(params.LendingRate)
	_, err = stmt.Exec(
		params.Name,
		params.Url,
		categorySting,
		personalQualificationSting,
		params.LimitMin,
		params.LimitMax,
		params.LogoUrl,
		params.Slogan,
		params.ApplyNumber,
		termString,
		InterestBytes,
		LendingRateBytes,
		params.Credit,
		params.AuditType,
		params.AccountInType,
		params.ApplyStrategy,
		id)
	if err != nil {
		res.Msg = "failure"
	} else {
		// ret.LastInsertId()  // 如果是数据库返回err，执行这行代码会导致接口崩溃
		res.Msg = "success"
	}
	res.Msg = "success"
	resBytes, _ := json.Marshal(res)
	w.Write(resBytes)
}

func productDetail(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	steps := strings.Split(path, "/")
	id := steps[len(steps)-1]
	rows, err := db.Query(selectProductDetailSql, id)
	if err != nil {
		// 数据库抛出的错误
		ret := RetrieveResponse{Msg: "failure", Data: []int{}}
		retBytes, _ := json.Marshal(ret)
		w.Write(retBytes)
		return
	}
	productSearch := productInfo{}
	for rows.Next() {
		var categorySting string
		var personalQualification string
		var term string
		var temp1 []byte
		var interest interest
		var temp2 []byte
		var lendingRate lendingRate

		_ = rows.Scan(&productSearch.Id,
			&productSearch.Name,
			&productSearch.Url,
			&categorySting,
			&personalQualification,
			&productSearch.LimitMin,
			&productSearch.LimitMax,
			&productSearch.LogoUrl,
			&productSearch.Slogan,
			&productSearch.ApplyNumber,
			&term,
			&temp1,
			&temp2,
			&productSearch.Credit,
			&productSearch.AuditType,
			&productSearch.AccountInType,
			&productSearch.ApplyStrategy)

		json.Unmarshal(temp1, &interest)
		json.Unmarshal(temp2, &lendingRate)
		productSearch.Term = strings.Split(term, ",")
		productSearch.Type = strings.Split(categorySting, ",")
		productSearch.PersonalQualification = strings.Split(personalQualification, ",")
		productSearch.Interest = interest
		productSearch.LendingRate = lendingRate
	}

	res := RetrieveResponse{Msg: "success", Data: productSearch}
	retBytes, err := json.Marshal(res)
	w.Write(retBytes)
	defer func() {
		rows.Close()
	}()
}

func listProduct(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(selectProductListSql)
	if err != nil {
		// 数据库抛出的错误
		ret := RetrieveResponse{Msg: "failure", Data: []int{}}
		retBytes, _ := json.Marshal(ret)
		w.Write(retBytes)
		return
	}
	var productSearchList []productSearch
	for rows.Next() {
		p := &productSearch{}
		var temp []byte
		var interest interest
		_ = rows.Scan(&p.Id, &p.Name, &p.LimitMin, &p.LimitMax, &temp)
		json.Unmarshal(temp, &interest)
		p.Interest = interest
		productSearchList = append(productSearchList, *p)
	}
	if productSearchList == nil {
		productSearchList = []productSearch{}
	}
	ret := RetrieveResponse{Msg: "success", Data: productSearchList}
	retBytes, err := json.Marshal(ret)
	w.Write(retBytes)
	defer func() {
		rows.Close()
	}()
}
