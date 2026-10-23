// starter 插件生命周期服务,进程异常退出由宿主隔离并按策略重启。
package starter

import (
	"context"

	"google.golang.org/grpc"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"

	"github.com/slaincsuite/modality-gopn-sdk-go/service"
)

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
