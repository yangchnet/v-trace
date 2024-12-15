package service

import (
	v1 "gitee.com/qciip-icp/v-trace/api/trans/v1"
	"gitee.com/qciip-icp/v-trace/app/trans/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
)

func Record2Proto(r *db.TransRecord) *v1.TransRecord {
	pbr := &v1.TransRecord{
		ID:           pbtools.ToProtoInt64(r.ID),
		TransId:      pbtools.ToProtoString(r.Transid),
		Sender:       pbtools.ToProtoString(r.Sender),
		Contract:     pbtools.ToProtoString(r.Contract),
		Method:       pbtools.ToProtoString(r.Method),
		Params:       pbtools.ToProtoBytes([]byte(r.Params.String)),
		TxParamsHash: pbtools.ToProtoString(r.TxParamsHash),
		TxHash:       pbtools.ToProtoString(r.Txhash),
		CreatedAt:    pbtools.ToProtoTimestamp(r.CreatedAt.Time),
		UpdatedAt:    pbtools.ToProtoTimestamp(r.UpdatedAt.Time),
	}

	switch r.Status.String {
	case v1.TransStatus_waiting.String():
		pbr.Status = pbtools.ToProtoString(v1.TransStatus_waiting.String())
	case v1.TransStatus_failed.String():
		pbr.Status = pbtools.ToProtoString(v1.TransStatus_failed.String())
	case v1.TransStatus_success.String():
		pbr.Status = pbtools.ToProtoString(v1.TransStatus_success.String())
	}

	return pbr
}
