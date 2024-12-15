package service

import (
	"context"

	v1 "gitee.com/qciip-icp/v-trace/api/ca/v1"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewCaService)

type CaService struct {
	v1.UnimplementedCAServiceServer
}

func NewCaService() *CaService {
	s := &CaService{}
	s.seedCert("normalUser")
	s.seedCert("transporterUser")
	s.seedCert("examinerUser")
	s.seedCert("producerUser")
	return s
}

func (s *CaService) GetDomain() string {
	return "CaService"
}

func (s *CaService) seedCert(username string) {
	_, err := s.GetCert(context.Background(), &v1.GetCertRequest{
		OrgId:     "org1",
		Username:  username,
		UserType:  v1.UserType_client,
		CertUsage: []v1.CertUsage{v1.CertUsage_sign, v1.CertUsage_tls},
	})

	if !v1.IsCertNotFound(err) {
		return
	}

	s.GenCert(context.Background(), &v1.GenCertRequest{
		OrgId:         "org1",
		Username:      username,
		UserType:      v1.UserType(v1.UserType_value["client"]),
		CertUsage:     []v1.CertUsage{v1.CertUsage_sign, v1.CertUsage_tls},
		PrivateKeyPwd: "",
		Country:       "CN",
		Locality:      "ShangHai",
		Province:      "ShangHai",
	})
}
