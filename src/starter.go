// starter 插件生命周期与宿主访问服务;进程异常退出由宿主隔离并按策略重启。
package starter

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

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

// starterService 返回插件对宿主暴露的健康检查服务。
func Service() service.Service {
	return service.NewService("starter.v1.Health", func(registrar grpc.ServiceRegistrar) {
		healthpb.RegisterHealthServer(registrar, healthServer{})
	})
}

type healthServer struct {
	healthpb.UnimplementedHealthServer
}

// Check 宿主健康检查探针。
func (healthServer) Check(context.Context, *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}
