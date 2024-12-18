/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package models

import "gitee.com/qciip-icp/v-trace/app/ca/pkg/models/db"

//Find certcontent which certstatus is active by conditions
func FindCertContent(userId, orgId string, usage db.CertUsage, userType db.UserType) (*db.CertContent, error) {
	certInfo, err := FindCertInfo(userId, orgId, usage, userType)
	if err != nil {
		return nil, err
	}
	certSn := certInfo.SerialNumber
	certContent, err := FindCertContentBySn(certSn)
	if err != nil {
		return nil, err
	}
	return certContent, nil
}

//Find certcontent which certstatus is active by conditions
func FindKeyPair(userId, orgId string, usage db.CertUsage, userType db.UserType) (*db.KeyPair, error) {
	certInfo, err := FindCertInfo(userId, orgId, usage, userType)
	if err != nil {
		return nil, err
	}
	keyPairSki := certInfo.PrivateKeyId
	keyPair, err := FindKeyPairBySki(keyPairSki)
	if err != nil {
		return nil, err
	}
	return keyPair, nil
}

//Find certcontent by conditions
func FindCertContents(userId, orgId string, usage db.CertUsage, userType db.UserType) ([]*db.CertContent, error) {
	certInfoList, err := FindCertInfos(userId, orgId, usage, userType)
	if err != nil {
		return nil, err
	}
	var res []*db.CertContent
	for _, value := range certInfoList {
		tmp, err := FindCertContentBySn(value.SerialNumber)
		if err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}
	return res, nil
}
