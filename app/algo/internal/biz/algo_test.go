package biz

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"gitee.com/qciip-icp/v-trace/app/algo/internal/data/db"
	"gitee.com/qciip-icp/v-trace/app/pkg/algo"
	gomock "github.com/golang/mock/gomock"
)

func Test_predict(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := NewMockAlgoRepo(ctrl)

	repo.EXPECT().GetModelByName(gomock.Any(), gomock.Any()).Return(&db.Model{
		ID: 1,
		Name: sql.NullString{
			String: "model1",
			Valid:  true,
		},
		Version: sql.NullInt32{
			Int32: 1,
			Valid: true,
		},
		Status: sql.NullString{
			String: "",
			Valid:  true,
		},
		Des: sql.NullString{
			String: "",
			Valid:  true,
		},
		Metadata: []byte{},
	}, nil)

	repo.EXPECT().GetMaterialID(gomock.Any(), &db.GetMaterialIDParams{
		ModelID: sql.NullInt32{
			Int32: 1,
			Valid: true,
		},
		Index: sql.NullInt32{
			Int32: 0,
			Valid: true,
		},
	}).Return(sql.NullInt32{
		Int32: 1,
		Valid: true,
	}, nil)

	repo.EXPECT().GetMaterialByID(gomock.Any(), int32(1)).Return(&db.Material{
		ID: 1,
		Name: sql.NullString{
			String: "中国荷斯坦牛",
			Valid:  true,
		},
		Alias: sql.NullString{
			String: "",
			Valid:  true,
		},
		Des: sql.NullString{
			String: "",
			Valid:  true,
		},
	}, nil)

	algoClient := algo.NewAlgoClient(ctx, &algo.AlgoConfig{
		Address: "vtrace_tf_serving:8500",
	})

	cas := NewAlgoCase(ctx, repo, algoClient)

	data, err := os.ReadFile("./testdata/Reflection_data/中国荷斯坦牛/0bd15d59-8540-4e19-ae4a-616c31954918.dx")
	if err != nil {
		t.Fatal(err)
	}

	_, err = cas.Predict(ctx, "model1", data)
	if err != nil {
		t.Fatal(err)
	}
}
