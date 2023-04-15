package impl

import (
	"context"
	"github.com/xingcanli132/rest-api-demo/apps/host"
	"testing"
)

var (
	service host.Service
)

func TestNewHostService(t *testing.T) {
	service = NewHostService()
}

func TestHostServiceImpl_CreateHost(t *testing.T) {
	service.CreateHost(context.Background(), &host.Host{
		Resource: &host.Resource{Name: "test"},
	})
}

func init() {
	service = NewHostService()
}
