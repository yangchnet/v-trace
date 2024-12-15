package contract

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"strings"
	"testing"

	"chainmaker.org/chainmaker/common/v2/evmutils"
	cmSdk "chainmaker.org/chainmaker/sdk-go/v2"
	v1 "gitee.com/qciip-icp/v-trace/api/ca/v1"
	"gitee.com/qciip-icp/v-trace/app/pkg/contract/config"
	"gitee.com/qciip-icp/v-trace/pkg/constants"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var conf config.Config

var c0, c1 *VTraceContract

var addrs []evmutils.Address

var caClient v1.CAServiceClient

var _digest = []string{}

var base = math.MaxInt64

func Init() {
	conf := config.LoadConfig("./config")

	node := NewNode(
		conf.Nodes[0].Addr,
		conf.Nodes[0].ConnCnt,
		conf.Nodes[0].UseTls,
		conf.Nodes[0].CaPaths,
		conf.Nodes[0].TlsHostName,
	)

	client0, err := NewChainClient(
		node,
		conf.Client.ChainClientOrgId,
		conf.Client.ChainId,
		conf.Client.AuthType,
		conf.Endorsers[0],
	)
	if err != nil {
		panic(err)
	}

	client1, err := NewChainClient(
		node,
		conf.Client.ChainClientOrgId,
		conf.Client.ChainId,
		conf.Client.AuthType,
		conf.Endorsers[1],
	)
	if err != nil {
		panic(err)
	}

	contractData, err := NewContractData(conf.Contract.AbiPath, conf.Contract.BinPath, conf.Contract.Name)
	if err != nil {
		panic(err)
	}

	c0 = NewVTraceContract(client0, contractData, true, conf.Contract.Name) // addrs[0]
	c1 = NewVTraceContract(client1, contractData, true, conf.Contract.Name) // addrs[1]

	loadUser(conf.Endorsers)

	for i := 0; i < 10; i++ {
		_digest = append(_digest, uuid.NewString())
	}
}

func Test_name(t *testing.T) {
	Init()
	name, err := c0.Name()
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t, "VTrace", name)
}

func Test_transfer(t *testing.T) {
	Init()
	var tokenId int64 = int64(base - 140)

	if err := c0.SafeMint(addrs[0], big.NewInt(tokenId), _digest[0], _digest[0]); err != nil {
		t.Fatal(err)
	}

	owner, err := c0.OwnerOf(big.NewInt(tokenId))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t,
		fmt.Sprintf("0X%s", strings.ToUpper(addrs[0].String())),
		strings.ToUpper(owner.String()),
		"an error happened when try to mint token")

	if err := c0.TransferFrom(addrs[0], addrs[1], big.NewInt(tokenId), _digest[1], _digest[1]); err != nil {
		t.Log(err)
	}

	owner, err = c0.OwnerOf(big.NewInt(tokenId))
	if err != nil {
		t.Fatal(err)
	}
	require.Equal(t,
		fmt.Sprintf("0X%s", strings.ToUpper(addrs[1].String())),
		strings.ToUpper(owner.String()),
		"an error happened when try to mint token")
}

func Test_BatchOperation(t *testing.T) {
	Init()

	// 1. batch mint
	start := 20
	_tokenIds := []int{base - start, base - (start + 1), base - (start + 2), base - (start + 3)}
	tokenIds := make([]*big.Int, 0)
	for _, id := range _tokenIds {
		tokenIds = append(tokenIds, big.NewInt(int64(id)))
	}

	transIds := []string{_digest[2], _digest[3], _digest[4], _digest[5]}

	// 记录当前账户0的余额
	balance0_1, err := c0.BalanceOf(addrs[0])
	if err != nil {
		t.Fatal(err)
	}

	// 为账户0发币， 4个
	if err := c0.BatchMint(addrs[0], tokenIds, transIds, _digest[2]); err != nil {
		t.Fatal(err)
	}

	// 再次记录账户0的余额
	balance0_2, err := c0.BalanceOf(addrs[0])
	if err != nil {
		t.Fatal(err)
	}

	// 检查余额是否上升4个
	add := &big.Int{}
	add = add.Sub(balance0_2, balance0_1)
	require.Equal(t, int64(4), add.Int64())

	// 2. batch transfer

	// 记录账户1的余额
	balance1_1, err := c0.BalanceOf(addrs[1])
	if err != nil {
		t.Fatal(err)
	}

	// 账户0向账户1转账4个
	transIds2 := []string{_digest[6], _digest[7], _digest[8], _digest[9]}
	if err := c0.BatchTransform(addrs[0], addrs[1], tokenIds, transIds2, "batch_transform_digest"); err != nil {
		t.Fatal(err)
	}

	// 再次记录账户1的余额
	balance1_2, err := c0.BalanceOf(addrs[1])
	if err != nil {
		t.Fatal(err)
	}

	// 检查账户1余额是否上升4个
	add2 := &big.Int{}
	add2 = add2.Sub(balance1_2, balance1_1)
	require.Equal(t, int64(4), add2.Int64())

	// 查询账户0余额
	balance0_3, err := c0.BalanceOf(addrs[0])
	if err != nil {
		t.Fatal(err)
	}

	// 检查账户0余额是否减少4个
	sub1 := &big.Int{}
	sub1 = sub1.Sub(balance0_2, balance0_3)
	require.Equal(t, int64(4), sub1.Int64())
}

func loadUser(users []*config.User) {
	for _, user := range users {
		addrStr, err := cmSdk.GetEVMAddressFromCertPath(user.SignCrtPath)
		if err != nil {
			panic(err)
		}
		addr, err := evmutils.HexToAddress(addrStr)
		if err != nil {
			panic(err)
		}
		addrs = append(addrs, addr)
	}
}

func getCert(username string) *v1.GetCertResponse {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "127.0.0.1:10105", []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}...)
	if err != nil {
		panic(err)
	}
	caClient := v1.NewCAServiceClient(conn)

	resp, err := caClient.GetCert(ctx, &v1.GetCertRequest{
		OrgId:     constants.OrgId,
		Username:  username,
		UserType:  v1.UserType_client,
		CertUsage: []v1.CertUsage{v1.CertUsage_sign, v1.CertUsage_tls},
	})
	if err != nil {
		panic(err)
	}

	return resp
}

// 从ca服务中取出证书
func getEvmAddr(username string) evmutils.Address {
	if username == "admin" {
		return addrs[0]
	}
	resp := getCert(username)

	addrStr, err := cmSdk.GetEVMAddressFromCertBytes([]byte(resp.GetCert().GetValue()))
	if err != nil {
		panic(err)
	}

	addr := evmutils.Address{}
	addr.SetBytes([]byte(addrStr))
	return addr
}
