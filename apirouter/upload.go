package apirouter

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_api/resp"
	"io"
	"log"
	"net/http"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	//设置内存大小
	r.ParseMultipartForm(1024 * 1024 * 1024 * 10) // 10GB

	//获取上传文件
	file, handler, err := r.FormFile("file")
	if err != nil {
		log.Println("获取上传文件失败", err)
		return
	}
	defer file.Close()

	// 创建上传目录
	os.Mkdir("./uploads", os.ModePerm)

	// 创建上传文件
	f, err := os.Create("./uploads/" + handler.Filename)
	if err != nil {
		fmt.Println("创建上传文件失败", err)
		return
	}
	defer f.Close()

	// 复制上传文件
	io.Copy(f, file)
	w.WriteHeader(http.StatusCreated)

	// 响应
	resp.Success(w, nil)
}
