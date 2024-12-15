package constants

// Grpc server.
// const (
// 	EchoServiceName   = "vtrace-echo-service"
// 	IamServiceName    = "vtrace-iam-service"
// 	CircServiceName   = "vtrace-circ-service"
// 	CAServiceName     = "vtrace-ca-service"
// 	GoodsServiceName  = "vtrace-goods-service"
// 	TransServiceName  = "vtrace-trans-service"
// 	AlgoServiceName   = "vtrace-algo-service"
// 	VTraceServiceName = "vtrace-service"

// 	EchoGrpcPort   = 10101
// 	IamGrpcPort    = 10102
// 	CircGrpcPort   = 10103
// 	CAGrpcPort     = 10105
// 	TransGrpcPort  = 10106
// 	GoodsGrpcPort  = 10107
// 	AlgoGrpcPort   = 10108
// 	VTraceGrpcPort = 10110
// )

const (
	RegistryNamespace = "v-trace"

	Iam    = "iam"
	Circ   = "circ"
	CA     = "ca"
	Goods  = "goods"
	Trans  = "trans"
	Algo   = "algo"
	VTrace = "vtrace"
)

// RoleName.
const (
	// 管理员用户.
	AdminRole = "admin"
	// 生产者用户.
	ProducerRole = "producer"
	// 参与者用户.
	TransporterRole = "transporter"
	// 普通注册用户.
	NormalRole = "normal"
	// 检验者用户.
	ExaminerRole = "examiner"
	// 企业所有者
	BossRole = "boss"
)

// 每个角色都有一个编码，使用编码的二进制来实现不同角色的包含关系
const (
	// NormalRoleCode (000001)_2.
	NormalRoleCode = 1 << iota

	// ExaminerCode (000011)_2.
	// 检验者包含普通用户的权限，但当一个用户成为参与者或生产者，其就不再具有检验者权限
	ExaminerCode = (1<<iota | NormalRoleCode)

	// TransporterRoleCode (000111)_2.
	TransporterRoleCode = (1<<iota | ExaminerCode)

	// ProducerRoleCode (001111)_2.
	ProducerRoleCode = (1<<iota | TransporterRoleCode)

	// BossRoleCode (011111)_2
	BossRoleCode = (1<<iota | ProducerRoleCode)

	// AdminRoleCode (111111)_2.
	AdminRoleCode = (1<<iota | BossRoleCode)
)

func Role2Code(role string) int {
	switch role {
	case AdminRole:
		return AdminRoleCode
	case ProducerRole:
		return ProducerRoleCode
	case TransporterRole:
		return TransporterRoleCode
	case NormalRole:
		return NormalRoleCode
	case ExaminerRole:
		return ExaminerCode
	case BossRole:
		return BossRoleCode
	default:
		return NormalRoleCode
	}
}

func Code2Role(code int) string {
	switch code {
	case 1:
		return NormalRole
	case 3:
		return ExaminerRole
	case 7:
		return TransporterRole
	case 15:
		return ProducerRole
	case 31:
		return BossRole
	case 63:
		return AdminRole
	default:
		return NormalRole
	}
}

func ShouldRole(roles []string) string {
	if len(roles) <= 0 {
		return NormalRole
	}

	r := Role2Code(roles[0])
	for i := 1; i < len(roles); i++ {
		r |= Role2Code(roles[i])
	}

	return Code2Role(r)
}

// Contract.
const (
	OrgId             = "org1"
	ChainId           = "chain1"
	SignUsage         = "sign"
	CrtUserTypeClient = "client"
	CrtUserTypeAdmin  = "admin"
	ContractAuthType  = "permissionedwithcert"
)

// Mq.
const (
	MqStream   = "vtrace-stream"
	MqEmptyId  = ""
	MqKey      = "vtrace-task"
	MqGroup    = "g1"
	MqConsumer = "c1"
)
