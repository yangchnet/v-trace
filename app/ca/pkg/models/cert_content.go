/*
Copyright (C) BABEC. All rights reserved.

SPDX-License-Identifier: Apache-2.0
*/

package models

import (
	"fmt"

	"gitee.com/qciip-icp/v-trace/app/ca/pkg/models/db"
)

//Find certcontent by Sn
func FindCertContentBySn(sn int64) (*db.CertContent, error) {
	var certContent db.CertContent
	if err := db.DB.Where("serial_number=?", sn).First(&certContent).Error; err != nil {
		return nil, fmt.Errorf("[DB] find cert content by sn error: %s, sn: %d", err.Error(), sn)
	}
	return &certContent, nil
}

//Update cert content
func UpdateCertContent(oldCertContent, newCertContent *db.CertContent) error {
	if err := db.DB.Model(oldCertContent).
		Select("content", "cert_raw", "key_usage", "ext_key_usage", "is_ca", "issue_date", "expiration_date").
		Updates(newCertContent).Error; err != nil {
		return fmt.Errorf("[DB] update cert content failed: %s", err.Error())
	}
	return nil
}
