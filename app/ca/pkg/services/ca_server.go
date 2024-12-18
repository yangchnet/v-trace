/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package services

import (
	"crypto/rand"
	"crypto/x509/pkix"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/common/v2/crypto/x509"
	"gitee.com/qciip-icp/v-trace/app/ca/pkg/models"
	"gitee.com/qciip-icp/v-trace/app/ca/pkg/models/db"
	"gitee.com/qciip-icp/v-trace/app/ca/pkg/utils"
	"go.uber.org/zap"
)

// Generate cert by csr.
func GenCertByCsr(genCertByCsrReq *GenCertByCsrReq) (string, error) {
	var empty string
	err := CheckCert(genCertByCsrReq.OrgId, genCertByCsrReq.UserId, genCertByCsrReq.UserType, genCertByCsrReq.CertUsage)
	if err != nil {
		logger.Error("generate cert by csr failed", zap.Error(err))
		return empty, err
	}
	logger.Info("generate cert by csr", zap.Any("req", genCertByCsrReq))
	hashType, err := checkHashType(hashTypeFromConfig())
	if err != nil {
		logger.Error("generate cert by csr failed", zap.Error(err))
		return empty, err
	}

	issuerPrivateKey, issuerCertBytes, err := searchIssuerCa(genCertByCsrReq.OrgId,
		genCertByCsrReq.UserType, genCertByCsrReq.CertUsage)
	if err != nil {
		logger.Error("generate cert by csr failed", zap.Error(err))
		return empty, err
	}

	certRequestConfig := &CertRequestConfig{
		HashType:         hashType,
		IssuerPrivateKey: issuerPrivateKey,
		CsrBytes:         genCertByCsrReq.CsrBytes,
		IssuerCertBytes:  issuerCertBytes,
		ExpireYear:       int32(expireYearFromConfig()),
		CertUsage:        genCertByCsrReq.CertUsage,
		UserType:         genCertByCsrReq.UserType,
	}
	certContent, err := IssueCertificate(certRequestConfig)
	if err != nil {
		logger.Error("generate cert by csr failed", zap.Error(err))
		return empty, err
	}

	certConditions := &CertConditions{
		UserType:  genCertByCsrReq.UserType,
		CertUsage: genCertByCsrReq.CertUsage,
		UserId:    genCertByCsrReq.UserId,
		OrgId:     genCertByCsrReq.OrgId,
	}
	certInfo, err := CreateCertInfo(certContent, "", certConditions)
	if err != nil {
		logger.Error("generate cert by csr failed", zap.Error(err))
		return empty, err
	}
	logger.Info("generate cert by csr", zap.Any("cert info", certInfo))
	err = models.CreateCertAndInfoTransaction(certContent, certInfo)
	if err != nil {
		logger.Error("generate cert by csr failed", zap.Error(err))
		return empty, err
	}
	logger.Info("generate cert by csr successfully", zap.String("cert", certContent.Content))
	return certContent.Content, nil
}

// Generate cert.
func GenCert(genCertReq *GenCertReq) (*CertAndPrivateKey, error) {
	err := CheckCert(genCertReq.OrgId, genCertReq.UserId, genCertReq.UserType, genCertReq.CertUsage)
	if err != nil {
		logger.Error("generate cert failed", zap.Error(err))
		return nil, err
	}
	logger.Info("generate cert", zap.Any("req", genCertReq))
	privateKeyTypeStr := keyTypeFromConfig()
	hashTypeStr := hashTypeFromConfig()
	privateKeyPwd := genCertReq.PrivateKeyPwd
	privateKey, keyPair, err := CreateKeyPair(privateKeyTypeStr, hashTypeStr, privateKeyPwd)
	if err != nil {
		logger.Error("generate cert failed", zap.Error(err))
		return nil, err
	}
	csrRequest := &CSRRequest{
		OrgId:      genCertReq.OrgId,
		UserId:     genCertReq.UserId,
		UserType:   genCertReq.UserType,
		Country:    genCertReq.Country,
		Locality:   genCertReq.Locality,
		Province:   genCertReq.Province,
		PrivateKey: privateKey,
	}
	csrRequestConf := BuildCSRReqConf(csrRequest)
	csrByte, err := createCSR(csrRequestConf)
	if err != nil {
		logger.Error("generate cert failed", zap.Error(err))
		return nil, err
	}

	issuerPrivateKey, issuerCertBytes, err := searchIssuerCa(genCertReq.OrgId, genCertReq.UserType, genCertReq.CertUsage)
	if err != nil {
		logger.Error("generate cert failed", zap.Error(err))
		return nil, err
	}
	hashType, err := checkHashType(hashTypeFromConfig())
	if err != nil {
		logger.Error("generate cert failed", zap.Error(err))
		return nil, err
	}
	certRequestConfig := &CertRequestConfig{
		HashType:         hashType,
		CsrBytes:         csrByte,
		ExpireYear:       int32(expireYearFromConfig()),
		CertUsage:        genCertReq.CertUsage,
		UserType:         genCertReq.UserType,
		IssuerPrivateKey: issuerPrivateKey,
		IssuerCertBytes:  issuerCertBytes,
	}
	certContent, err := IssueCertificate(certRequestConfig)
	if err != nil {
		logger.Error("generate cert failed", zap.Error(err))
		return nil, err
	}
	certConditions := &CertConditions{
		UserType:  genCertReq.UserType,
		CertUsage: genCertReq.CertUsage,
		UserId:    genCertReq.UserId,
		OrgId:     genCertReq.OrgId,
	}
	certInfo, err := CreateCertInfo(certContent, keyPair.Ski, certConditions)
	if err != nil {
		logger.Error("generate cert failed", zap.Error(err))
		return nil, err
	}
	logger.Info("generare cert", zap.Any("cert info", certInfo))
	err = models.CreateCertTransaction(certContent, certInfo, keyPair)
	if err != nil {
		logger.Error("generate cert failed", zap.Error(err))
		return nil, err
	}
	logger.Info("generate cert", zap.String("cert", certContent.Content))
	logger.Info("generate cert", zap.String("private key", keyPair.PrivateKey))
	return &CertAndPrivateKey{
		Cert:       certContent.Content,
		PrivateKey: keyPair.PrivateKey,
	}, nil
}

// Query certs.
func QueryCerts(req *QueryCertsReq) ([]*CertInfos, error) {
	var (
		userType  db.UserType
		certUsage db.CertUsage
		err       error
	)
	userType, err = CheckParametersUserType(req.UserType)
	if err != nil {
		userType = 0
	}
	certUsage, err = checkParametersCertUsage(req.CertUsage)
	if err != nil {
		certUsage = 0
	}
	logger.Info("query certs", zap.Any("req", req))
	certInfoList, err := models.FindCertInfos(req.UserId, req.OrgId, certUsage, userType)
	if err != nil {
		logger.Error("query cert by status failed", zap.Error(err))
		return nil, err
	}
	var res []*CertInfos
	for _, certInfo := range certInfoList {
		certContent, err := models.FindCertContentBySn(certInfo.SerialNumber)
		if err != nil {
			logger.Error("query cert by status failed", zap.Error(err))
			return nil, err
		}
		pk, err := models.FindKeyPairBySki(certInfo.PrivateKeyId)
		if err != nil {
			logger.Error("query pk by ski failed", zap.Error(err))
			return nil, err
		}
		res = append(res, &CertInfos{
			UserId:         certInfo.UserId,
			OrgId:          certInfo.OrgId,
			UserType:       db.UserType2NameMap[certInfo.UserType],
			CertUsage:      db.CertUsage2NameMap[certInfo.CertUsage],
			CertSn:         certInfo.SerialNumber,
			IssuerSn:       certInfo.IssuerSn,
			CertContent:    certContent.Content,
			ExpirationDate: certContent.ExpirationDate,
			PrivateKey:     pk.PrivateKey,
		})
	}
	logger.Info("query certs", zap.Any("resp", res))
	return res, nil
}

// renew the cert expiration date.
func RenewCert(renewCertReq *RenewCertReq) (string, error) {
	var empty string
	logger.Info("renew cert", zap.Int64("cert sn", renewCertReq.CertSn))
	certInfo, err := models.FindCertInfoBySn(renewCertReq.CertSn)
	if err != nil {
		logger.Error("renew cert failed", zap.Error(err))
		return empty, err
	}
	certContent, err := models.FindCertContentBySn(renewCertReq.CertSn)
	if err != nil {
		logger.Error("renew cert failed", zap.Error(err))
		return empty, err
	}
	issuerPrivateKey, issuerCertBytes, err := searchIssuerCa(certInfo.OrgId, certInfo.UserType, certInfo.CertUsage)
	if err != nil {
		logger.Error("renew cert failed", zap.Error(err))
		return empty, err
	}
	csrBytes := []byte(certContent.CsrContent)
	oldCert, err := ParseCertificate([]byte(certContent.Content))
	if err != nil {
		logger.Error("renew cert failed", zap.Error(err))
		return empty, err
	}
	logger.Info("renew cert", zap.Any("old cert expiration date", oldCert.NotAfter.UTC()))
	// renew invalid date
	oldCert.NotAfter = oldCert.NotAfter.Add(time.Duration(expireYearFromConfig()) * 365 * 24 * time.Hour).UTC()

	logger.Info("renew cert", zap.Any("new cert expiration date", oldCert.NotAfter.UTC()))

	newCertContent, err := UpdateCert(&UpdateCertConfig{
		OldCert:         oldCert,
		OldCsrBytes:     csrBytes,
		IssuerCertBytes: issuerCertBytes,
		IssuerKey:       issuerPrivateKey,
	})
	if err != nil {
		logger.Error("renew cert failed", zap.Error(err))
		return empty, err
	}
	err = models.UpdateCertContent(certContent, newCertContent)
	if err != nil {
		logger.Error("renew cert failed", zap.Error(err))
		return empty, err
	}
	logger.Info("renew cert", zap.String("cert", newCertContent.Content))
	return newCertContent.Content, nil
}

// Revoke  certificate.
func RevokeCert(revokeCertReq *RevokeCertReq) ([]byte, error) {
	logger.Info("revoke cert", zap.Any("req", revokeCertReq))
	_, err := models.QueryRevokedCertByRevokedSn(revokeCertReq.RevokedCertSn)
	if err == nil { // find it and is already revoked
		err = fmt.Errorf("this cert had already been revoked")
		logger.Error("revoked cert failed", zap.Error(err))
		return nil, err
	}

	_, err = models.QueryRevokedCertByRevokedSn(revokeCertReq.IssuerCertSn)
	if err == nil { // find it and is already revoked
		err = fmt.Errorf("issuer cert had already been revoked")
		logger.Error("revoked cert failed", zap.Error(err))
		return nil, err
	}

	ok, err := searchCertChain(revokeCertReq.RevokedCertSn, revokeCertReq.IssuerCertSn)
	if err != nil {
		logger.Error("revoked cert failed", zap.Error(err))
		return nil, err
	}
	if !ok {
		err = fmt.Errorf("issue cert is not in revoked cert chain")
		logger.Error("revoked cert failed", zap.Error(err))
		return nil, err
	}
	issueCertInfo, err := models.FindCertInfoBySn(revokeCertReq.IssuerCertSn)
	if err != nil {
		logger.Error("revoked cert failed", zap.Error(err))
		return nil, err
	}
	logger.Info("issuer cert info", zap.Any("issuer cert info", issueCertInfo))
	revokedCertContent, err := models.FindCertContentBySn(revokeCertReq.RevokedCertSn)
	if err != nil {
		logger.Error("revoked cert failed", zap.Error(err))
		return nil, err
	}
	revokedCert := &db.RevokedCert{
		OrgId:            issueCertInfo.OrgId,
		RevokedCertSN:    revokeCertReq.RevokedCertSn,
		Reason:           revokeCertReq.Reason,
		RevokedStartTime: time.Now().Unix(),
		RevokedEndTime:   revokedCertContent.ExpirationDate,
		RevokedBy:        revokeCertReq.IssuerCertSn,
	}
	err = models.InsertRevokedCert(revokedCert)
	if err != nil {
		logger.Error("revoked cert failed", zap.Error(err))
		return nil, err
	}
	// create crl
	genCrlReq := &GenCrlReq{
		IssuerCertSn: revokeCertReq.IssuerCertSn,
	}
	crlBytes, err := GenCrl(genCrlReq)
	if err != nil {
		logger.Error("revoked cert failed", zap.Error(err))
		return nil, err
	}
	logger.Info("revoke cert", zap.String("crl", string(crlBytes)))
	return crlBytes, nil
}

// Get the latest crl.
func GenCrl(genCrlReq *GenCrlReq) ([]byte, error) {
	logger.Info("generate crl", zap.Any("req", genCrlReq))
	issueCertUse, err := GetX509Certificate(genCrlReq.IssuerCertSn)
	if err != nil {
		logger.Error("crl list get failed", zap.Error(err))
		return nil, err
	}
	issueCertInfo, err := models.FindCertInfoBySn(genCrlReq.IssuerCertSn)
	if err != nil {
		logger.Error("crl list get failed", zap.Error(err))
		return nil, err
	}
	logger.Info("generate crl", zap.Any("issuer cert info", issueCertInfo))
	var issuePrivateKey crypto.PrivateKey
	if issueCertInfo.UserType == db.ROOT_CA {
		issuePrivateKey, err = GetRootPrivateKey(issueCertInfo.CertUsage)
		if err != nil {
			logger.Error("crl list get failed", zap.Error(err))
			return nil, err
		}
	} else {
		var issueKeyPair *db.KeyPair
		issueKeyPair, err = models.FindKeyPairBySki(issueCertInfo.PrivateKeyId)
		if err != nil {
			logger.Error("crl list get failed", zap.Error(err))
			return nil, err
		}
		issuePrivateKeyByte := []byte(issueKeyPair.PrivateKey)
		var hashPwd []byte
		hashPwd, err = hex.DecodeString(issueKeyPair.PrivateKeyPwd)
		if err != nil {
			logger.Error("crl list get failed", zap.Error(err))
			return nil, err
		}

		issuePrivateKey, err = KeyBytesToPrivateKey(issuePrivateKeyByte, string(hashPwd))
		if err != nil {
			logger.Error("crl list get failed", zap.Error(err))
			return nil, err
		}
	}
	revokedCertsList, err := models.QueryRevokedCertByIssueSn(genCrlReq.IssuerCertSn)
	if err != nil {
		logger.Error("crl list get failed", zap.Error(err))
		return nil, err
	}
	var revokedCerts []pkix.RevokedCertificate
	for _, value := range revokedCertsList {
		revoked := pkix.RevokedCertificate{
			SerialNumber:   big.NewInt(value.RevokedCertSN),
			RevocationTime: time.Unix(value.RevokedEndTime, 0),
		}
		revokedCerts = append(revokedCerts, revoked)
	}
	now := time.Now()
	next := now.Add(utils.DefaultCRLNextTime)
	crlBytes, err := x509.CreateCRL(rand.Reader, issueCertUse, issuePrivateKey.ToStandardKey(), revokedCerts, now, next)
	if err != nil {
		logger.Error("crl list get failed", zap.Error(err))
		return nil, err
	}
	logger.Info("generate crl", zap.String("crl", string(crlBytes)))
	return crlBytes, nil
}

// Generate csr.
func GenCsr(genCsrReq *GenCsrReq) ([]byte, error) {
	logger.Info("generate csr", zap.Any("req", genCsrReq))
	privateKeyTypeStr := keyTypeFromConfig()
	hashTypeStr := hashTypeFromConfig()
	privateKeyPwd := genCsrReq.PrivateKeyPwd
	privateKey, _, err := CreateKeyPair(privateKeyTypeStr, hashTypeStr, privateKeyPwd)
	if err != nil {
		logger.Error("generate cert failed", zap.Error(err))
		return nil, err
	}
	csrRequest := &CSRRequest{
		OrgId:      genCsrReq.OrgId,
		UserId:     genCsrReq.UserId,
		UserType:   genCsrReq.UserType,
		Country:    genCsrReq.Country,
		Locality:   genCsrReq.Locality,
		Province:   genCsrReq.Province,
		PrivateKey: privateKey,
	}
	csrRequestConf := BuildCSRReqConf(csrRequest)
	csrByte, err := createCSR(csrRequestConf)
	if err != nil {
		logger.Error("generate cert failed", zap.Error(err))
		return nil, err
	}
	logger.Info("generate csr", zap.String("csr", string(csrByte)))
	return csrByte, nil
}

// check orgId userId usertype certusage and determine whether to provIde certificate service.
func CheckParameters(orgId, userId, userTypeStr, certUsageStr string) (userType db.UserType,
	certUsage db.CertUsage, err error,
) {
	userType, err = CheckParametersUserType(userTypeStr)
	if err != nil {
		return
	}
	certUsage, err = checkParametersCertUsage(certUsageStr)
	if err != nil {
		return
	}

	if err = checkParamsOfCertReq(orgId, userId, userType, certUsage); err != nil {
		return
	}
	return
}

// check the string parametes if empty.
func CheckParametersEmpty(parameters ...string) error {
	for _, parameter := range parameters {
		if len(parameter) == 0 {
			err := fmt.Errorf("check parameters failed: required parameters cannot be empty")
			return err
		}
	}
	return nil
}

// Check if the certificate already exists.
func CheckCert(orgId string, userId string, userType db.UserType, certUsage db.CertUsage) error {
	if userType == db.INTERMRDIARY_CA {
		_, err := models.FindCertInfo("", orgId, certUsage, db.INTERMRDIARY_CA)
		if err == nil {
			return fmt.Errorf("the ca cert has already existed")
		}
		return nil
	}
	_, err := models.FindCertInfo(userId, orgId, certUsage, userType)
	if err == nil {
		return fmt.Errorf("the cert has already existed")
	}
	return nil
}
