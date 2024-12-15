package service

import (
	"context"
	"strings"

	v1 "gitee.com/qciip-icp/v-trace/api/ca/v1"
	"gitee.com/qciip-icp/v-trace/app/ca/pkg/loggers"
	"gitee.com/qciip-icp/v-trace/app/ca/pkg/services"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
)

func (s *CaService) GenCert(ctx context.Context, req *v1.GenCertRequest) (*v1.GenCertResponse, error) {
	orgId := req.GetOrgId()
	username := req.GetUsername()
	userType := v1.UserType_name[int32(req.GetUserType())]
	certUsageList := req.GetCertUsage()
	privateKeyPwd := req.GetPrivateKeyPwd()
	country := req.GetCountry()
	locality := req.GetLocality()
	province := req.GetProvince()

	userId := strings.Join([]string{username, orgId}, ".")

	var signCrt, signKey, tlsCrt, tlsKey string
	for _, usage := range certUsageList {
		certUsage := v1.CertUsage_name[int32(usage)]
		curUserType, curCertUsage, err := services.CheckParameters(orgId, userId,
			userType, certUsage)
		if err != nil {
			loggers.GetSugaLogger().Errorf("params check error: %+v", err)

			return nil, verr.Error(s, v1.ErrorInvalidParams("参数校验失败:%v", err))
		}

		cert, err := services.GenCert(&services.GenCertReq{
			OrgId:         orgId,
			UserId:        userId,
			UserType:      curUserType,
			CertUsage:     curCertUsage,
			PrivateKeyPwd: privateKeyPwd,
			Country:       country,
			Locality:      locality,
			Province:      province,
		})
		if err != nil {
			loggers.GetSugaLogger().Errorf("generate cert failed: %+v", err)

			return nil, verr.Error(s, v1.ErrorGenCertErr("证书生成失败:%v", err))
		}

		if usage == v1.CertUsage_sign {
			signCrt = cert.Cert
			signKey = cert.PrivateKey
		} else if usage == v1.CertUsage_tls {
			tlsCrt = cert.Cert
			tlsKey = cert.PrivateKey
		}
	}

	return &v1.GenCertResponse{
		Cert:       pbtools.ToProtoString(signCrt),
		PrivateKey: pbtools.ToProtoString(signKey),
		Username:   pbtools.ToProtoString(username),
		TlsCert:    pbtools.ToProtoString(tlsCrt),
		TlsKey:     pbtools.ToProtoString(tlsKey),
	}, nil
}

func (s *CaService) GetCert(ctx context.Context, req *v1.GetCertRequest) (*v1.GetCertResponse, error) {
	orgId := req.GetOrgId()
	username := req.GetUsername()
	userType := v1.UserType_name[int32(req.GetUserType())]
	certUsageList := req.GetCertUsage()

	userId := strings.Join([]string{username, orgId}, ".")

	var signCrt, signKey, tlsCrt, tlsKey string
	for _, usage := range certUsageList {
		certUsage := v1.CertUsage_name[int32(usage)]
		cert, err := services.QueryCerts(&services.QueryCertsReq{
			OrgId:     orgId,
			UserId:    userId,
			UserType:  userType,
			CertUsage: certUsage,
		})
		if err != nil {
			return nil, verr.Error(s, err)
		}

		if len(cert) <= 0 {
			return nil, v1.ErrorCertNotFound("证书不存在")
		}

		if usage == v1.CertUsage_sign {
			signCrt = cert[0].CertContent
			signKey = cert[0].PrivateKey
		} else if usage == v1.CertUsage_tls {
			tlsCrt = cert[0].CertContent
			tlsKey = cert[0].PrivateKey
		}
	}

	return &v1.GetCertResponse{
		Cert:       pbtools.ToProtoString(signCrt),
		PrivateKey: pbtools.ToProtoString(signKey),
		Username:   pbtools.ToProtoString(username),
		TlsCert:    pbtools.ToProtoString(tlsCrt),
		TlsKey:     pbtools.ToProtoString(tlsKey),
	}, nil
}
