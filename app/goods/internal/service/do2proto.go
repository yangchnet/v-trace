package service

import (
	"database/sql"

	v1 "gitee.com/qciip-icp/v-trace/api/goods/v1"
	"gitee.com/qciip-icp/v-trace/app/goods/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
)

func Goods2Proto(goods *db.Good, serial *db.Serial, class *db.Class) *v1.Goods {
	pbGoods := &v1.Goods{
		ID:      pbtools.ToProtoInt64(goods.ID),
		State:   pbtools.ToProtoString(goods.Status.String),
		Creator: pbtools.ToProtoString(goods.Creator.String),
	}

	if serial != nil {
		pbGoods.Serial = Serial2Proto(serial, class)
	}

	return pbGoods
}

func Serial2Proto(serial *db.Serial, class *db.Class) *v1.Serial {
	pbSerial := &v1.Serial{
		ID:        pbtools.ToProtoInt64(serial.ID),
		Timestamp: pbtools.ToProtoTimestamp(serial.ProductTime.Time),
		State:     pbtools.ToProtoString(serial.Status.String),
		Creator:   pbtools.ToProtoString(serial.Creator.String),
	}

	if class != nil {
		pbSerial.Class = Class2Proto(class)
	}

	return pbSerial
}

func Class2Proto(class *db.Class) *v1.Class {
	return &v1.Class{
		ID:         pbtools.ToProtoInt64(class.ID),
		Name:       pbtools.ToProtoString(class.Name.String),
		Des:        pbtools.ToProtoBytes(class.Des),
		State:      pbtools.ToProtoString(class.Status.String),
		Creator:    pbtools.ToProtoString(class.Creator.String),
		OrgId:      pbtools.ToProtoInt64(class.OrgID),
		Tm:         pbtools.ToProtoString(class.Tm.String),
		MaterialId: pbtools.ToProtoInt32(class.MaterialID),
	}
}

func Proto2Class(class *v1.Class) *db.Class {
	value := class.Des.Value
	return &db.Class{
		ID:         int32(class.ID.Value),
		Name:       sql.NullString{String: class.GetName().GetValue(), Valid: true},
		Des:        value,
		Status:     sql.NullString{String: class.GetState().GetValue(), Valid: true},
		Creator:    sql.NullString{String: class.GetCreator().GetValue(), Valid: true},
		MaterialID: class.GetMaterialId().GetValue(),
	}
}

func Proto2Goods(goods *v1.Goods) *db.Good {
	return &db.Good{
		ID:       int32(goods.ID.Value),
		Status:   sql.NullString{String: goods.GetState().GetValue(), Valid: true},
		Creator:  sql.NullString{String: goods.GetCreator().GetValue(), Valid: true},
		SerialID: sql.NullInt32{Int32: int32(goods.GetSerial().GetID().GetValue()), Valid: true},
	}
}

func Proto2Serial(serial *v1.Serial) *db.Serial {
	return &db.Serial{
		ID:          int32(serial.ID.Value),
		ProductTime: sql.NullTime{Time: serial.GetTimestamp().AsTime(), Valid: true},
		Status:      sql.NullString{String: serial.GetState().GetValue(), Valid: true},
		Creator:     sql.NullString{String: serial.GetCreator().GetValue(), Valid: true},
		ClassID:     sql.NullInt32{Int32: int32(serial.GetClass().GetID().GetValue()), Valid: true},
	}
}
