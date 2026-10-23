// starter 生命周期服务单测。
package starter

import (
	"context"
	"errors"
	"testing"

	"google.golang.org/grpc"

	"github.com/slaincsuite/modality-gopn-sdk-go/host"
)

func TestServiceHealth(t *testing.T) {
	server := grpc.NewServer()
	t.Cleanup(server.Stop)
	if item := Service(); item.Name != "starter.v1.Health" {
		t.Fatalf("service name = %q", item.Name)
	}
	Service().Register(server)

}

type denyCaller struct{}

func (denyCaller) Call(context.Context, host.Authorization, host.Request) (host.Response, error) {
	return host.Response{}, host.ErrForbidden
}

func TestHostAccessDefaultDeny(t *testing.T) {
	client := NewHostClient(denyCaller{})
	if _, err := client.Call(context.Background(), "module.registry.read", nil); !errors.Is(err, host.ErrForbidden) {
		t.Fatalf("expected ErrForbidden for unauthorized call, got %v", err)
	}
}

func TestQueryModuleRegistryDeny(t *testing.T) {
	client := NewHostClient(denyCaller{})
	if _, err := QueryModuleRegistry(context.Background(), client); err == nil {
		t.Fatal("expected error when host denies")
	}
}

func TestNewHostClientCapabilities(t *testing.T) {
	client := NewHostClient(nil)
	caps := client.Capabilities()
	if len(caps) != 1 || caps[0].Name != "module.registry.read" {
		t.Fatalf("capabilities = %+v", caps)
	}
	// 未声明的能力一律拒绝。
	if _, err := client.Call(context.Background(), "upms.user.write", nil); err == nil {
		t.Fatal("undeclared capability should be denied")
	}
	// caller 为 nil 时已声明能力返回宿主不可用。
	if _, err := client.Call(context.Background(), "module.registry.read", nil); !errors.Is(err, host.ErrHostUnavailable) {
		t.Fatalf("expected ErrHostUnavailable, got %v", err)
	}
}
