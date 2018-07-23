package server

import (
	"123_hao_dai_be/elea"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func registerFileServer(h *elea.Handle) {
	h.Register("/be/manage/product/logo/upload", receiveFile)
	h.Register("/be/file/", serveFile)

	// todo:
	h.Register("/lab/img/upload", receiveImage)
	h.Register("/lab/img/free", serveImage)
}

func receiveImage(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, _ := r.FormFile("img")
	fileBytes, _ := ioutil.ReadAll(file)
	fmt.Println(fileHeader.Filename)
	db.Exec("INSERT INTO test (img) VALUES (?)", fileBytes)
}

func serveImage(w http.ResponseWriter, r *http.Request) {
	var img []byte
	var id int
	row := db.QueryRow("SELECT id, img FROM test WHERE id = 13")
	row.Scan(&id, &img)
	fmt.Println(id)
	fmt.Println(len(img))
	w.Write(img)
}

func receiveFile(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("logo")
	if err != nil {
		w.Write([]byte("failure"))
		return
	}
	separatedFileName := strings.Split(fileHeader.Filename, ".")
	len := len(separatedFileName)
	fileType := separatedFileName[len-1]
	// 使用微秒时间戳给图片文件重命名
	newFileName := strconv.FormatInt(time.Now().UnixNano()/1000, 10) + "." + fileType
	fileDuplicate, _ := os.Create("./src/assets/img/" + newFileName)
	io.Copy(fileDuplicate, file)
	var res UploadResponse
	res.Msg = "success"
	res.Data.Url = "/be/file/img/" + newFileName
	resBytes, _ := json.Marshal(res)
	// 返回图片存放的信息
	w.Write(resBytes)
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	steps := strings.Split(path, "/")
	fileType := steps[len(steps)-2]
	fileName := steps[len(steps)-1]
	file, err := os.Open("/root/go/src/123_hao_dai_be/src/assets/" + fileType + "/" + fileName)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(`<body align=center><h2>404 Not Found</h2><hr>Elea 0.0.1</body>`))
		return
	}
	fileBytes, _ := ioutil.ReadAll(file)
	w.Write(fileBytes)
}
