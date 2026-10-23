// starter 生命周期服务单测。
package starter

import (
	"testing"

	"google.golang.org/grpc"
)

func TestServiceHealth(t *testing.T) {
	server := grpc.NewServer()
	t.Cleanup(server.Stop)
	if item := Service(); item.Name != "starter.v1.Health" {
		t.Fatalf("service name = %q", item.Name)
	}
	Service().Register(server)

}
