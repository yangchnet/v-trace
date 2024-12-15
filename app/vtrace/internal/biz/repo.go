package biz

import (
	"context"
	"time"

	algoV1 "gitee.com/qciip-icp/v-trace/api/algo/v1"
	circV1 "gitee.com/qciip-icp/v-trace/api/circ/v1"
	goodsV1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	iamV1 "gitee.com/qciip-icp/v-trace/api/iam/v1"
	transV1 "gitee.com/qciip-icp/v-trace/api/trans/v1"
	"gitee.com/qciip-icp/v-trace/app/vtrace/internal/data"
)

type VTraceRepo interface {
	iamRepo
	caRepo
	transRepo
	goodsRepo
	circRepo
	algoRepo
}

type iamRepo interface {
	// iam
	CreateUser(ctx context.Context, nickname, passwd, phone string) (*iamV1.User, error)
	// Token
	GetToken(ctx context.Context, phone, passwd string) (string, error)
	// UpdateToken 更新令牌
	UpdateToken(ctx context.Context, token string) (string, error)
	// GetUserInfo 获取用户信息
	GetUserInfo(ctx context.Context, username string) (*iamV1.User, error)
	// CreateIdentity 记录用户实名
	CreateIdentity(ctx context.Context, username, realname, idcard string) (*iamV1.User, error)
	// CreateOrg 创建企业
	CreateOrg(ctx context.Context, owner, orgName, orgCode, legalName, legalPhone string, canProduce bool, orgInfo []byte) (*iamV1.Org, error)
	// GetOrg 获取企业信息
	GetOrg(ctx context.Context, orgId int32) (*iamV1.Org, error)
	// AddMember 企业增加成员
	AddMember(ctx context.Context, orgId int32, username string) error
	// 删除用户
	DeleteUser(ctx context.Context, username string) error
	// UpdateUser 用户信息更新
	UpdateUser(ctx context.Context, user *iamV1.User) (*iamV1.User, error)

	// GetOrgUser 查询用户所属企业
	GetOrgOfUser(ctx context.Context, username string) (*iamV1.Org, error)

	// RemoveOrgMember 企业删除成员
	OrgRemoveMember(ctx context.Context, id int32, username string) error

	// ListOrgMember 企业查询成员列表
	ListOrgMember(ctx context.Context, org_id, offset, limit int32) ([]*iamV1.User, error)

	// UpdateOrg 企业信息更新
	UpdateOrg(ctx context.Context, org *iamV1.Org) (*iamV1.Org, error)
}

type caRepo interface {
	// 签署用户证书
	CreateCert(ctx context.Context, username string) (*transV1.Identity, error)
	// 获取用户证书
	GetCert(ctx context.Context, username string) (*transV1.Identity, error)
}

type transRepo interface {
	// 调用合约
	CallContract(ctx context.Context, params *data.CallContractParams) error

	// 根据transId获取交易信息
	GetTxByTransId(ctx context.Context, transId string) (*transV1.TransRecord, error)
}

type goodsRepo interface {
	// 创建商品类别
	CreateGoodsClass(ctx context.Context, name string, des []byte, materials, orgId int32, tm string) (int32, error)

	// 获取类别
	GetClass(ctx context.Context, class_id int32) (*goodsV1.Class, error)

	// 列出商品类别
	ListClass(ctx context.Context, offset, limit, orgId int32) ([]*goodsV1.Class, error)

	// 更新商品类别
	UpdateClass(ctx context.Context, class *goodsV1.Class) error

	// 创建商品批次
	CreateGoodsSerial(ctx context.Context, productTime time.Time, class_id int32) (int32, error)

	// 获得商品批次
	GetGoodsSerial(ctx context.Context, serial_id int32) (*goodsV1.Serial, error)

	// 列出产品批次
	ListGoodsSerial(ctx context.Context, offset, limit, orgId int32) ([]*goodsV1.Serial, error)

	// 更新产品批次
	UpdateGoodsSerial(ctx context.Context, serial *goodsV1.Serial) error

	// 批量创建商品
	BatchCreateGoods(ctx context.Context, serial_id int32, sum int32) ([]int32, error)

	// 获取商品
	GetGoodsByGoodsId(ctx context.Context, goods_id int32) (*goodsV1.Goods, error)

	// 列出商品
	ListGoods(ctx context.Context, offset, limit, orgId int32) ([]*goodsV1.Goods, error)

	// 更新商品
	UpdateGoods(ctx context.Context, goods *goodsV1.Goods) error

	// 获取所属企业
	GetOrgOfClass(ctx context.Context, classId int32) (int32, error)
	GetOrgOfSerial(ctx context.Context, serialId int32) (int32, error)
	GetOrgOfGoods(ctx context.Context, goodsId int32) (int32, error)
}

type circRepo interface {
	// 创建流转
	CreateCirc(ctx context.Context, transId string, circType string, operator, from, to string, fromInfo []byte) (int32, error)
	// 获取流转历史
	GetCircByGoodsId(ctx context.Context, goods_id int32) ([]*circV1.CircRecord, error)
	// 生成TransId
	TransId(ctx context.Context, goodsId int32) (string, error)
	// 批量生成TransId
	BatchTransId(ctx context.Context, goodsIds []int32) (map[int32]string, error)
	// 批量流转
	BatchCirc(ctx context.Context, transIds []string, circType string, operator, from, to string, fromInfo []byte) ([]int32, error)
}

type algoRepo interface {
	// 获取原料列表
	GetMaterials(ctx context.Context) ([]*algoV1.Material, error)
	// ListAlgoModels 列出所有算法模型
	ListAlgoModels(ctx context.Context) ([]*algoV1.Model, error)
	// Predict 根据数据进行预测
	Predict(ctx context.Context, modelName string, data []byte) (*algoV1.Material, error)
}
