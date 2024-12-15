package service

import (
	v1 "gitee.com/qciip-icp/v-trace/api/circ/v1"
	"gitee.com/qciip-icp/v-trace/app/circ/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
)

func Circ2Proto(record *db.CircRecord) *v1.CircRecord {
	return &v1.CircRecord{
		Id:        pbtools.ToProtoInt64(record.ID),
		ObjectId:  pbtools.ToProtoInt64(record.Objectid),
		CircType:  v1.CircType(v1.CircType_value[record.Circtype.String]),
		Operator:  pbtools.ToProtoString(record.Operator.String),
		From:      pbtools.ToProtoString(record.From.String),
		To:        pbtools.ToProtoString(record.To.String),
		FromValue: pbtools.ToProtoString(string(record.Formvalue)),
		TransId:   pbtools.ToProtoString(record.Transid),
		Times:     pbtools.ToProtoInt64(record.Times),
		Status:    pbtools.ToProtoString(record.Status),
	}
}
