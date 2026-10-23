// starter 插件进程入口:注册生命周期服务并等待宿主调用。
package main

import (
	"github.com/slaincsuite/modality-gopn-sdk-go/service"
	"github.com/slaincsuite/modality-plugin-starter/src"
)

func main() {
	service.MustServe([]service.Service{starter.Service()})
}
