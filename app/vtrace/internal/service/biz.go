package service

import (
	"context"
	"time"

	algoV1 "gitee.com/qciip-icp/v-trace/api/algo/v1"
	circV1 "gitee.com/qciip-icp/v-trace/api/circ/v1"
	goodsV1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	iamV1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
)

type VTraceCaseInterface interface {
	iamInterface
	goodsInterface
	circInterface
	algoInterface
	fileInterface
}

type iamInterface interface {
	// 注册
	Register(ctx context.Context, nickname, passwd, phone string) (*iamV1.User, string, error)

	// 获取token
	GetToken(ctx context.Context, phone, passwd string) (string, error)

	// UpdateToken 更新令牌
	UpdateToken(ctx context.Context) (string, error)

	// Profile 获取个人信息
	Profile(ctx context.Context) (*iamV1.User, error)

	// IdentityAuth 个人实名认证
	IdentityAuth(ctx context.Context, realname, idcard string) (*iamV1.User, error)

	// OrgAuth 企业认证
	OrgAuth(ctx context.Context, legalName, orgName, code, legalPhone string, info []byte) (*iamV1.Org, error)

	// 企业增加成员
	Member(ctx context.Context, orgId int32, username string) error

	// 删除用户
	DeleteUser(ctx context.Context, username string) error

	// UpdateUser 用户信息更新
	UpdateUser(ctx context.Context, user *iamV1.User) (*iamV1.User, error)

	// GetOrgUser 查询用户所属企业
	GetOrgUser(ctx context.Context, username string) (*iamV1.Org, error)

	// OrgRemoveMember 企业删除成员
	OrgRemoveMember(ctx context.Context, id int32, username string) error

	// ListOrgMember 企业查询成员列表
	ListOrgMember(ctx context.Context, org_id, offset, limit int32) ([]*iamV1.User, error)

	// UpdateOrg 企业信息更新
	UpdateOrg(ctx context.Context, org *iamV1.Org) (*iamV1.Org, error)
}

type goodsInterface interface {
	// 创建商品类别
	CreateGoodsClass(ctx context.Context, name string, des []byte, materials int32, tm string) (int32, error)

	//获取类型
	GetClass(ctx context.Context, class_id int32) (*goodsV1.Class, error)

	// 列出产品类型
	ListClass(ctx context.Context, offset, limit int32) ([]*goodsV1.Class, error)

	// 更新产品类型
	UpdateClass(ctx context.Context, class *goodsV1.Class) error

	// 创建商品批次
	CreateGoodsSerial(ctx context.Context, productTime time.Time, class_id int32) (int32, error)

	// 获得商品批次
	GetGoodsSerial(ctx context.Context, serial_id int32) (*goodsV1.Serial, error)

	// 列出产品批次
	ListGoodsSerial(ctx context.Context, offset, limit int32) ([]*goodsV1.Serial, error)

	// 更新产品批次
	UpdateGoodsSerial(ctx context.Context, serial *goodsV1.Serial) error

	// 批量创建商品
	BatchCreateGoods(ctx context.Context, serial_id int32, sum int32) ([]int32, error)

	// 获得产品
	GetGoods(ctx context.Context, goods_id int32) (*goodsV1.Goods, error)

	// 列出商品
	ListGoods(ctx context.Context, offset, limit int32) ([]*goodsV1.Goods, error)

	// 更新产品
	UpdateGoods(ctx context.Context, goods *goodsV1.Goods) error
}

type circInterface interface {
	// 创建流转历史
	CreateCirc(ctx context.Context, goodsId int32, circType string, from, to string, formValue []byte) (int32, string, error)
	// 批量流转
	BatchCirc(ctx context.Context, goodsIds []int32, circType, from, to string, formValue []byte) (map[int32]string, error)
	// 获取流转历史
	GetGoodsAndCircs(ctx context.Context, goosId int32) (*goodsV1.Goods, []*circV1.CircRecord, error)
}

type algoInterface interface {
	// 获取原料列表
	GetMaterials(ctx context.Context) ([]*algoV1.Material, error)

	// ListAlgoModels 列出所有算法模型
	ListAlgoModels(ctx context.Context) ([]*algoV1.Model, error)

	// Predict 根据数据进行预测
	Predict(ctx context.Context, modelName string, data []byte) (*algoV1.Material, error)
}

type fileInterface interface {
	// 存储文件
	Store(ctx context.Context, content []byte, meta map[string]string) (string, error)
}
