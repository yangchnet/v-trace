// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"context"
)

type Querier interface {
	// 创建一个产品类别记录
	CreateClass(ctx context.Context, arg *CreateClassParams) (int64, error)
	// 创建一个产品
	CreateGoods(ctx context.Context, arg *CreateGoodsParams) (int64, error)
	// 创建一个产品批次
	CreateSerial(ctx context.Context, arg *CreateSerialParams) (int64, error)
	// 根据id获取产品类别
	GetClassById(ctx context.Context, id int32) (*Class, error)
	GetGoodsById(ctx context.Context, id int32) (*Good, error)
	// 获取产品所属企业
	GetOrgOfClass(ctx context.Context, id int32) (int32, error)
	// 获取产品所属企业
	GetOrgOfGoods(ctx context.Context, id int32) (int32, error)
	GetOrgOfSerial(ctx context.Context, id int32) (int32, error)
	GetSerialById(ctx context.Context, id int32) (*Serial, error)
	// 列出产品类型
	ListGoodClass(ctx context.Context, arg *ListGoodClassParams) ([]*Class, error)
	ListGoods(ctx context.Context, arg *ListGoodsParams) ([]*ListGoodsRow, error)
	// 列出产品批次
	ListGoodsSerial(ctx context.Context, arg *ListGoodsSerialParams) ([]*ListGoodsSerialRow, error)
	// 更新产品信息
	UpdateGood(ctx context.Context, arg *UpdateGoodParams) error
	// 更新产品类型
	UpdateGoodClass(ctx context.Context, arg *UpdateGoodClassParams) error
	// 更新产品批次
	UpdateGoodSerial(ctx context.Context, arg *UpdateGoodSerialParams) error
}

var _ Querier = (*Queries)(nil)