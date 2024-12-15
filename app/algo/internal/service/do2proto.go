package service

import (
	"encoding/json"

	v1 "gitee.com/qciip-icp/v-trace/api/algo/v1"
	"gitee.com/qciip-icp/v-trace/app/algo/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/tools/pbtools"
	"google.golang.org/protobuf/types/known/structpb"
)

func Material2Proto(m *db.Material, models []*db.Model) *v1.Material {
	pbModels := make([]*v1.Model, 0)
	if len(models) > 0 {
		for _, model := range models {
			pbModels = append(pbModels, Model2Proto(model))
		}
	}
	return &v1.Material{
		Id:              pbtools.ToProtoInt64(m.ID),
		Name:            pbtools.ToProtoString(m.Name.String),
		Alias:           pbtools.ToProtoString(m.Alias.String),
		Des:             pbtools.ToProtoString(m.Des.String),
		AvailableModels: pbModels,
	}
}

func Model2Proto(model *db.Model) *v1.Model {
	pbModel := &v1.Model{
		ID:      pbtools.ToProtoInt64(model.ID),
		Name:    pbtools.ToProtoString(model.Name.String),
		Version: pbtools.ToProtoInt64(model.Version.Int32),
		State:   pbtools.ToProtoString(model.Status.String),
		Des:     pbtools.ToProtoString(model.Des.String),
	}

	md := make(map[string]interface{})
	data := model.Metadata
	if err := json.Unmarshal(data, md); err != nil {
		logger.Error(err)

		return pbModel
	}

	pbMd := make(map[string]*structpb.Value)
	for k, v := range md {
		value, err := structpb.NewValue(v)
		if err != nil {
			logger.Error(err)
			continue
		}
		pbMd[k] = value
	}

	pbModel.Metadata = pbMd

	return pbModel
}
