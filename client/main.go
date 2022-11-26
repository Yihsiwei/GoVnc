package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/saljam/mjpeg"
	"github.com/vova616/screenshot"
)

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
	首先编译好命令参数如	: client.exe ip port
	例				如		:client.exe 127.0.0.1 9527
	`
)

func main() {
	fmt.Println(logo)
	fmt.Println(tvb)
	if len(os.Args) != 3 {
		fmt.Println(keytishi)
		return
	}
	ip := os.Args[1]
	port := os.Args[2]
	stream := mjpeg.NewStream()
	func(stream *mjpeg.Stream) {
		for {
			img, err := screenshot.CaptureScreen()
			if err != nil {
				continue
			}
			file, err := os.Create("./tmp.jpeg")
			if err != nil {
				continue
			}
			defer file.Close()
			jpeg.Encode(file, img, nil)
			buf, err := ioutil.ReadFile("tmp.jpeg")
			if err != nil {
				continue
			}
			resp, err := http.Post("http://"+ip+":"+port+"/update", "multipart/form-data", bytes.NewReader(buf))
			if err != nil {
				break
			}
			resp.Body.Close()
			stream.UpdateJPEG(buf)
		}
	}(stream)

}
