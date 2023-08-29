package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port string `yaml:"port"`
}

func main() {
	// 解析命令行参数
	configFile := flag.String("f", "", "配置文件路径")
	flag.Parse()

	// 读取配置文件
	configData, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("无法读取配置文件：%v", err)
	}

	// 解析配置文件
	var config Config
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("无法解析配置文件：%v", err)
	}

	// 输出配置信息
	fmt.Printf("端口号：%s\n", config.Port)

	// 启动服务
	startServer(config.Port)
}

func startServer(port string) {
	// 注册路由处理函数
	http.HandleFunc("/bar", helloHandler)

	// 启动HTTP服务
	addr := ":" + port
	fmt.Println("服务已启动，监听地址：" + addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("无法启动HTTP服务：%v", err)
	}

	// 服务结束前的清理工作
	defer func() {
		fmt.Println("服务已停止")
	}()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello bar!"))
}
