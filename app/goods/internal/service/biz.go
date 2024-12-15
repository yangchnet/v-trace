package service

import (
	"context"
	"time"

	"gitee.com/qciip-icp/v-trace/app/goods/internal/data/db"
)

// GoodsCaseInterface 完成有关产品的业务操作s.
type GoodsCaseInterface interface {
	// 创建产品种类
	CreateClass(ctx context.Context, name, creator string, des []byte, materialId int32, orgId int32, tm string) (*db.Class, error)
	// 创建产品批次
	CreateSerial(ctx context.Context, productTime time.Time, classId int32) (*db.Serial, error)
	// 批量创建产品
	BatchCreateGoods(ctx context.Context, serialId, sum int32) ([]int32, error)
	// 获取商品
	GetGoods(ctx context.Context, goods_id int32) (*db.Good, *db.Serial, *db.Class, error)

	//获取种类
	GetClass(ctx context.Context, class_id int32) (*db.Class, error)

	//获取批次
	GetSerial(ctx context.Context, serial_id int32) (*db.Serial, *db.Class, error)

	//列出产品类型
	ListGoodsClass(ctx context.Context, offset, limit, orgId int32) ([]*db.Class, error)

	//列出产品批次
	ListGoodsSerial(ctx context.Context, offset int32, limit, orgID int32) ([]*db.Serial, []*db.Class, error)

	//列出产品
	ListGoods(ctx context.Context, offset int32, limit, orgID int32) ([]*db.Good, []*db.Serial, []*db.Class, error)

	//更新产品类型信息
	UpdateGoodsClass(ctx context.Context, class *db.Class) (*db.Class, error)

	//更新产品
	UpdateGoods(ctx context.Context, good *db.Good) (*db.Good, error)

	//更新产品批次
	UpdateGoodsSerial(ctx context.Context, serial *db.Serial) (*db.Serial, error)

	// 获取所属企业
	GetOrgOfX(ctx context.Context, x string, id int32) (int32, error)
}
