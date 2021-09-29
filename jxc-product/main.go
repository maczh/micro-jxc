package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"ququ.im/common"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"syscall"
	"time"
)

const config_file = "jxc-product.yml"

//@title	蛐蛐侠进销存商品信息管理模块API
//@version 	1.0.0(jxc-product)
//@description	蛐蛐侠进销存接口

func main() {
	//初始化配置，自动连接数据库和Nacos服务注册
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	config.InitConfig(path + "/" + config_file)

	//GIN的模式，生产环境可以设置成release
	gin.SetMode("debug")

	engine := setupRouter()

	server := &http.Server{
		Addr:    ":" + config.GetConfigString("go.application.port"),
		Handler: engine,
	}

	common.PrintLogo()
	fmt.Println("|-----------------------------------|")
	fmt.Println("|     蛐蛐侠进销存产品模块 0.0.1     |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|  Go Http Server Start Successful  |")
	fmt.Println("|    Port:" + config.GetConfigString("go.application.port") + "     Pid:" + fmt.Sprintf("%d", os.Getpid()) + "        |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Error("HTTP server listen: {}", err.Error())
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-signalChan
	logs.Error("Get Signal:{}", sig.String())
	logs.Error("Shutdown Server ...")
	config.SafeExit()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logs.Error("Server Shutdown:{}", err.Error())
	}
	logs.Error("Server exiting")

}
