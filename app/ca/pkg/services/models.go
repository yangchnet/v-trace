/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package services

import (
	"gitee.com/qciip-icp/v-trace/app/ca/pkg/models/db"
)

type GenCertByCsrReq struct {
	OrgId     string
	UserId    string
	UserType  db.UserType
	CertUsage db.CertUsage
	CsrBytes  []byte
}

type GenCertReq struct {
	OrgId         string
	UserId        string
	UserType      db.UserType
	CertUsage     db.CertUsage
	PrivateKeyPwd string
	Country       string
	Locality      string
	Province      string
}

type QueryCertsReq struct {
	OrgId     string
	UserId    string
	UserType  string
	CertUsage string
}

type RenewCertReq struct {
	CertSn int64
}

type RevokeCertReq struct {
	RevokedCertSn int64
	IssuerCertSn  int64
	Reason        string
}

type GenCrlReq struct {
	IssuerCertSn int64
}

type GenCsrReq struct {
	OrgId         string
	UserId        string
	UserType      db.UserType
	PrivateKeyPwd string
	Country       string
	Locality      string
	Province      string
}

type CertAndPrivateKey struct {
	Cert       string `json:"cert"`
	PrivateKey string `json:"privateKey"`
}

type CertInfos struct {
	UserId         string `json:"userId"`
	OrgId          string `json:"orgId"`
	UserType       string `json:"userType"`
	CertUsage      string `json:"certUsage"`
	CertSn         int64  `json:"certSn"`
	IssuerSn       int64  `json:"issuerSn"`
	CertContent    string `json:"certContent"`
	ExpirationDate int64  `json:"expirationDate"`
	PrivateKey     string `json:"privateKey"`
}
