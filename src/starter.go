// starter 插件生命周期与宿主访问服务;进程异常退出由宿主隔离并按策略重启。
package starter

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/slaincsuite/modality-gopn-sdk-go/host"
	"github.com/slaincsuite/modality-gopn-sdk-go/service"
)

// manifestGrants 与 plugin.json grants 字段保持一致;宿主仍逐项重新鉴权。
var manifestGrants = []host.Capability{
	{Name: "module.registry.read", Module: "module", Mode: host.AccessRead},
}

// NewHostClient 创建宿主访问客户端;caller 为 nil 时所有调用返回宿主不可用。
// go-plugin broker 接入后,由 service.WithBroker 回调创建 BrokerCaller 并传入。
func NewHostClient(caller host.Caller) *host.Client {
	return host.New(host.Config{
		PluginName:   "modality-plugin-starter",
		Capabilities: manifestGrants,
	}, caller)
}

// QueryModuleRegistry 演示宿主只读能力调用;宿主默认拒绝未授权访问。
func QueryModuleRegistry(ctx context.Context, client *host.Client) (any, error) {
	resp, err := client.Call(ctx, "module.registry.read", nil)
	if err != nil {
		return nil, fmt.Errorf("starter: module.registry.read: %w", err)
	}
	return resp.Data, nil
}

// Service 返回插件对宿主暴露的服务声明:健康检查由 go-plugin 的 gRPC server
// **内建**提供(宿主经 gRPC health 客户端调用),插件不得再次注册同名服务——
// 重复注册会触发 gRPC 的 FATAL 并导致插件进程退出。业务服务请照此模式
// 注册自己的 proto 服务(register 内调用 protoc-gen-go-grpc 生成的 RegisterXxxServer)。
func Service() service.Service {
	return service.NewService("grpc.health.v1.Health", func(grpc.ServiceRegistrar) {})
}
