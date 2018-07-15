package server

import (
	"123_hao_dai_be/elea"
	"encoding/json"
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
	file, err := os.Open("./src/assets/" + fileType + "/" + fileName)
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(`<body align=center><h2>404 Not Found</h2><hr>Elea 0.0.1</body>`))
		return
	}
	fileBytes, _ := ioutil.ReadAll(file)
	w.Write(fileBytes)
}
