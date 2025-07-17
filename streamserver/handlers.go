package main

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func testPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// io.WriteString(w, "Hello, this is a test page.")
	t, _ := template.ParseFiles("./videos/upload.html")
	t.Execute(w, nil)
}

func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// vid := p.ByName("vid-id") 老师的
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid

	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Error when try to open file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// 限制请求体大小
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	// 限制上传文件大小
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "file is too big")
		return
	}
	// 读取文件数据
	file, _, err := r.FormFile("file")
	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	// 将file读到我们的数据里面,写成二进制的数据 读取文件内容到data里面
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Readfile error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}

	// 保存文件到磁盘
	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666)
	if err != nil {
		log.Printf("Writefile error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	// 设置状态码为201 Created
	w.WriteHeader(http.StatusCreated)

	// 返回上传成功信息
	io.WriteString(w, "Uploaded successfully")
}
