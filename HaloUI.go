package HaloUI

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"
	"time"
)

var Pars []string
var WebInput []string
var RunFunc func(input []string)
var Sync bool

func Run() {
	HaloUI_output = ""                 //还原输出
	gin.DefaultWriter = ioutil.Discard //关闭gin的日志
	ginServer := gin.Default()

	ginServer.POST("/Run", gin_post)
	ginServer.GET("/Run", gin_get)

	//打印主界面UI
	ginServer.GET("/", func(c *gin.Context) {
		c.Header("content-type", "text/html; charset=utf-8")
		c.String(200, HaloUI_index())
	})

	port, err := GetFreePort() //获取空闲端口
	if err != nil {            //错误处理
		log.Print(err)
	}
	WebAddress := fmt.Sprintf("127.0.0.1:%d", port) //定义监听端口
	fmt.Println("UI界面:  ", "http://"+WebAddress)
	go OpenBrowser("http://" + WebAddress)
	err = ginServer.Run(WebAddress) //启动HaloUI的http服务（基于gin）
	if err != nil {                 //错误处理
		log.Print(err)
	}
}

// 一键运行的入口
func gin_post(c *gin.Context) {
	WebInput = []string{}
	fmt.Println("[+]Run...")
	NoFinsh()
	for i := 0; i < len(Pars); i++ {
		WebInput = append(WebInput, c.PostForm(Pars[i]))
		fmt.Println("[+]入参", Pars[i], " : ", WebInput[i])
	}
	if Sync == true {
		RunFunc(WebInput)
	} else {
		go RunFunc(WebInput)
	}
	c.Header("content-type", "text/html; charset=utf-8")
	c.String(200, HaloUI_run())
}

// 更新页面的入口
func gin_get(c *gin.Context) {
	c.Header("content-type", "text/html; charset=utf-8")
	c.String(200, HaloUI_run())
}

// 打开指定url
func OpenBrowser(url string) error {
	var cmd string
	var args []string
	time.Sleep(1 * time.Second) // 睡眠4秒
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
