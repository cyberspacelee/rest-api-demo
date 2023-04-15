package host

import "context"

// Service host app service 接口定义
type Service interface {
	// CreateHost 创建主机
	CreateHost(context.Context, *Host) (*Host, error)
	// QueryHost 查询主机
	QueryHost(context.Context, *Host) (*HostSet, error)
	// DescribeHost 查询主机详情
	DescribeHost(context.Context, *Host) (*Host, error)
	// UpdateHost 更新主机信息
	UpdateHost(context.Context, *Host) (*Host, error)
	// DeleteHost 删除主机信息
	DeleteHost(context.Context, *Host) (*Host, error)
}
