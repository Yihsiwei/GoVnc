package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/saljam/mjpeg"
)

var (
	stream = mjpeg.NewStream()
)

//服务端编写的业务逻辑处理程序 —— 回调函数
func myHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("method = ", r.Method) //请求方法
	//fmt.Println("URL = ", r.URL)       // 浏览器发送请求文件路径
	//fmt.Println("header = ", r.Header) // 请求头
	//fmt.Println("body = ", r.Body)     // 请求包体
	//fmt.Println(r.RemoteAddr, "连接成功")  //客户端网络地址

	//len := r.ContentLength
	//fmt.Println(len)
	bodys, _ := ioutil.ReadAll(r.Body)
	////fmt.Println(body)

	stream.UpdateJPEG(bodys)
	//f, _ := os.Create("1.jpeg")
	//_, _ = f.Write(bodys)
	//f.Close()
}

var (
	logo = `
	\__  |   |   |  |__   _____|__|_  _  __ ____ |__|
	/   |   |   |  |  \ /  ___/  \ \/ \/ // __ \|  |
	\____   |   |   Y  \\___ \|  |\     /\  ___/|  |
	/ ______|___|___|  /____  >__| \/\_/  \___  >__|
	\/               \/     \/                \/   
	`
	tvb = "这是我的频道欢迎投稿学习:https://space.bilibili.com/353948151	"

	keytishi = `
	首先编译好命令参数如	: server.exe port
	例				如		:server.exe 9527
	`
)

func main() {
	fmt.Println(logo)
	fmt.Println(tvb)
	if len(os.Args) != 2 {
		fmt.Println(keytishi)
		return
	}
	port := os.Args[1]
	http.HandleFunc("/update", myHandler) // 注册处理函数
	http.HandleFunc("/mjpeg", stream.ServeHTTP)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content := "<img width='100%' height='100%' src='/mjpeg'>"
		w.Header().Add("Content-Type", "text/html;charset=utf-8")
		w.Write([]byte(content))
	})

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
