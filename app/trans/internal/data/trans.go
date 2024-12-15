package data

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	v1 "gitee.com/qciip-icp/v-trace/api/trans/v1"
	"gitee.com/qciip-icp/v-trace/app/trans/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
	"golang.org/x/net/context"
)

func (d *Data) CreateTransWithParams(ctx context.Context, transId, contractName, sender, method string, argsMap map[string]any) error {
	var params string
	paramsBytes, err := json.Marshal(argsMap)
	if err != nil {
		logger.Error(err)

		params = fmt.Sprint(argsMap)
	}
	params = string(paramsBytes)

	hs := hashParams(contractName, sender, method, params)

	if _, err := d.Store.CreateTrans(ctx, &db.CreateTransParams{
		Transid:  transId,
		Sender:   sender,
		Contract: contractName,
		Method:   method,
		Params: sql.NullString{
			String: params,
			Valid:  true,
		},
		Status: sql.NullString{
			String: v1.TransStatus_waiting.String(),
			Valid:  true,
		},
		Txhash:       fmt.Sprintf("fake-%s", hs),
		TxParamsHash: hs,
	}); err != nil {
		logger.Error(err)
		if verr.IsDuplicate(err) {
			if strings.Contains(err.Error(), "trans_record.txHash") || strings.Contains(err.Error(), "trans_record.tx_params_hash") {
				return v1.ErrorTransAlreadyExist("交易已存在，不能重复创建")
			}

			return v1.ErrorDuplicateErr("重复创建交易: %v", err)
		}

		return err
	}

	return nil
}

// hashParams 根据参数构造一个假的交易hash用于占位.
func hashParams(contractName, sender, method, params string) string {
	buf := bytes.Buffer{}
	buf.WriteString(contractName)
	buf.WriteString(sender)
	buf.WriteString(method)
	buf.WriteString(params)
	h := sha256.New()
	h.Write(buf.Bytes())
	hs := h.Sum(nil)

	return fmt.Sprintf("%x", hs)
}
