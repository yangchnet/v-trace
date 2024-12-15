package data

import (
	"context"
	"io/ioutil"

	caV1 "gitee.com/qciip-icp/v-trace/api/ca/v1"
	transV1 "gitee.com/qciip-icp/v-trace/api/trans/v1"
	"gitee.com/qciip-icp/v-trace/pkg/constants"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
)

func (r *Data) CreateCert(ctx context.Context, username string) (*transV1.Identity, error) {
	resp, err := r.Ca().GenCert(ctx, &caV1.GenCertRequest{
		OrgId:         constants.OrgId,
		Username:      username,
		UserType:      caV1.UserType_client,
		CertUsage:     []caV1.CertUsage{caV1.CertUsage_sign, caV1.CertUsage_tls},
		PrivateKeyPwd: "",
		Country:       "CN",
		Locality:      "ShangHai",
		Province:      "ShangHai",
	})
	if err != nil {
		logger.Errorf("generate cert failed: %+v", err)

		return nil, err
	}

	return &transV1.Identity{
		Cert:     []byte(resp.GetCert().GetValue()),
		Key:      []byte(resp.GetPrivateKey().GetValue()),
		Username: pbtools.ToProtoString(resp.GetUsername().GetValue()),
		TlsCert:  []byte(resp.GetTlsCert().GetValue()),
		TlsKey:   []byte(resp.GetTlsKey().GetValue()),
	}, nil
}

func (r *Data) GetCert(ctx context.Context, username string) (*transV1.Identity, error) {
	if username == "admin" {
		return getAdminCert(ctx)
	}

	resp, err := r.Ca().GetCert(ctx, &caV1.GetCertRequest{
		OrgId:     constants.OrgId,
		Username:  username,
		UserType:  caV1.UserType_client,
		CertUsage: []caV1.CertUsage{caV1.CertUsage_sign, caV1.CertUsage_tls},
	})
	if err != nil {
		logger.Errorf("get crt failed: %+v", err)

		return nil, err
	}

	return &transV1.Identity{
		Cert:     []byte(resp.GetCert().GetValue()),
		Key:      []byte(resp.GetPrivateKey().GetValue()),
		Username: pbtools.ToProtoString(resp.GetUsername().GetValue()),
		TlsCert:  []byte(resp.GetTlsCert().GetValue()),
		TlsKey:   []byte(resp.GetTlsKey().GetValue()),
	}, nil
}

func getAdminCert(ctx context.Context) (*transV1.Identity, error) {
	crtBytes, err := ioutil.ReadFile("build/chainmaker/config/org1/certs/user/admin1/admin1.sign.crt")
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	pkBytes, err := ioutil.ReadFile("build/chainmaker/config/org1/certs/user/admin1/admin1.sign.key")
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	tlsCrtBytes, err := ioutil.ReadFile("build/chainmaker/config/org1/certs/user/admin1/admin1.tls.crt")
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	tlsPkBytes, err := ioutil.ReadFile("build/chainmaker/config/org1/certs/user/admin1/admin1.tls.key")
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	return &transV1.Identity{
		Cert:     crtBytes,
		Key:      pkBytes,
		Username: pbtools.ToProtoString("admin"),
		TlsCert:  tlsCrtBytes,
		TlsKey:   tlsPkBytes,
	}, nil
}
