package biz

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	v1 "gitee.com/qciip-icp/v-trace/api/algo/v1"
	"gitee.com/qciip-icp/v-trace/app/algo/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"github.com/yangchnet/tf-serving/tensorflow/core/framework"
	"github.com/yangchnet/tf-serving/tensorflow_serving/apis"
)

func (c *AlgoCase) ListModels(ctx context.Context) ([]*db.Model, error) {
	dbModels, err := c.repo.ListModels(ctx)
	if err != nil {
		return nil, err
	}

	for _, model := range dbModels {
		statusResp, err := c.client.GetModelStatus(ctx, &apis.GetModelStatusRequest{
			ModelSpec: &apis.ModelSpec{
				Name: model.Des.String,
			},
		})
		if err != nil {
			logger.Error(err)

			continue
		}
		state := statusResp.ModelVersionStatus[len(statusResp.ModelVersionStatus)].GetStatus().String()
		if model.Status.String != state {
			if err := c.repo.UpdateModelStatus(ctx, &db.UpdateModelStatusParams{
				Status: sql.NullString{
					String: state,
					Valid:  true,
				},
				Name: sql.NullString{
					String: model.Name.String,
					Valid:  true,
				},
			}); err != nil {
				logger.Error(ctx, err)
			}
		}

		mdResp, err := c.client.GetModelMetadata(ctx, &apis.GetModelMetadataRequest{
			ModelSpec: &apis.ModelSpec{
				Name: model.Name.String,
			},
		})

		md := make(map[string]interface{})
		for k, v := range mdResp.Metadata {
			md[k] = v.Value
		}
	}

	return dbModels, nil
}

func (c *AlgoCase) Predict(ctx context.Context, modelName string, data []byte) (*db.Material, error) {
	model, err := c.repo.GetModelByName(ctx, sql.NullString{
		String: modelName,
		Valid:  true,
	})
	if err != nil {
		logger.Error(err)

		if errors.Is(err, sql.ErrNoRows) {
			return nil, v1.ErrorModelNotFound("不存在的模型: %s", modelName)
		}

		return nil, err
	}

	b := bytes.NewReader(data)
	reader := csv.NewReader(b)
	float32Data, err := readColumn1(reader)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Predict(ctx, &apis.PredictRequest{
		ModelSpec: &apis.ModelSpec{
			Name:          modelName,
			VersionChoice: nil,
			SignatureName: "",
		},
		Inputs: map[string]*framework.TensorProto{
			"input_1": {
				Dtype: framework.DataType_DT_FLOAT,
				TensorShape: &framework.TensorShapeProto{
					Dim: []*framework.TensorShapeProto_Dim{
						{Size: 1},
						{Size: 1400},
						{Size: 1},
					},
				},
				FloatVal: float32Data,
			},
		},
		OutputFilter: []string{},
	})
	if err != nil {
		return nil, v1.ErrorModelPredictFailed("调用算法模型失败: %v", err)
	}

	idx := maxIndex(resp.Outputs["output_1"].FloatVal)

	materialId, err := c.repo.GetMaterialID(ctx, &db.GetMaterialIDParams{
		ModelID: sql.NullInt32{
			Int32: model.ID,
			Valid: true,
		},
		Index: sql.NullInt32{
			Int32: int32(idx),
			Valid: true,
		},
	})
	if err != nil {
		logger.Error(err)

		return nil, v1.ErrorPredictResultUnavailable("分类结果不在数据库中\n")
	}

	material, err := c.repo.GetMaterialByID(ctx, materialId.Int32)
	if err != nil {
		logger.Error(err)

		return nil, v1.ErrorMaterialNotFound("不存在的原料数据: %d", materialId.Int32)
	}

	return material, nil
}

func readColumn1(reader *csv.Reader) ([]float32, error) {
	ret := make([]float32, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		value, err := strconv.ParseFloat(strings.TrimSpace(row[1]), 32)
		if err != nil {
			logger.Error(err)

			return nil, fmt.Errorf("解析数据失败: %v", err)
		}

		ret = append(ret, float32(value))
	}

	return ret, nil
}

func maxIndex(nums []float32) uint {
	var maxIdx uint = 0
	for i := 1; i < len(nums); i++ {
		if nums[maxIdx] < nums[i] {
			maxIdx = uint(i)
		}
	}

	return maxIdx
}
