package impl

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/xingcanli132/rest-api-demo/apps/host"
)

// 静态检查，确保 HostServiceImpl 实现了 Service 接口
var _ host.Service = (*HostServiceImpl)(nil)

// HostServiceImpl host app service 实现
type HostServiceImpl struct {
	l *logrus.Logger
}

// CreateHost 创建主机
func (service *HostServiceImpl) CreateHost(ctx context.Context, host *host.Host) (*host.Host, error) {
	service.l.Println("CreateHost")
	return nil, nil
}

// QueryHost 查询主机
func (*HostServiceImpl) QueryHost(ctx context.Context, host *host.Host) (*host.HostSet, error) {
	return nil, nil
}

// DescribeHost 查询主机详情
func (*HostServiceImpl) DescribeHost(ctx context.Context, host *host.Host) (*host.Host, error) {
	return nil, nil
}

// UpdateHost 更新主机信息
func (*HostServiceImpl) UpdateHost(ctx context.Context, host *host.Host) (*host.Host, error) {
	return nil, nil
}

// DeleteHost 删除主机信息
func (*HostServiceImpl) DeleteHost(ctx context.Context, host *host.Host) (*host.Host, error) {
	return nil, nil
}

func NewHostService() host.Service {
	return &HostServiceImpl{l: logrus.New()}
}
