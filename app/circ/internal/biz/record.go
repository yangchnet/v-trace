package biz

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	v1 "gitee.com/qciip-icp/v-trace/api/circ/v1"
	"gitee.com/qciip-icp/v-trace/app/circ/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	"gitee.com/qciip-icp/v-trace/pkg/verr"
)

func (c *CircCase) CreateRecord(
	ctx context.Context,
	transId string,
	circType string,
	operator string,
	from string,
	to string,
	formValue []byte,
) (*db.CircRecord, error) {
	times, goodsId, _ := parseTransId(transId)

	if err := c.valid(ctx, circType, goodsId, transId, from); err != nil {
		return nil, err
	}

	recordId, err := c.repo.CreateRecord(ctx, &db.CreateRecordParams{
		Transid:  transId,
		Objectid: int32(goodsId),
		Circtype: sql.NullString{
			String: circType,
			Valid:  true,
		},
		Operator: sql.NullString{
			String: operator,
			Valid:  true,
		},
		From: sql.NullString{
			String: from,
			Valid:  true,
		},
		To: sql.NullString{
			String: to,
			Valid:  true,
		},
		Formvalue: formValue,
		Times:     int32(times),
		Status:    v1.RecordStatus_pending.String(),
	})
	if err != nil {
		logger.Error(err)
		if verr.IsDuplicate(err) {
			return nil, v1.ErrorCircAlreadyExist("创建流转记录失败: %s", err)
		}
	}

	record, err := c.repo.GetRecordById(ctx, int32(recordId))
	if err != nil {
		logger.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, v1.ErrorCircNotFound("不存在的流转记录ID:%d", recordId)
		}

		return nil, err
	}

	return record, nil
}

func (c *CircCase) BatchRecord(ctx context.Context, transIds []string, operator, from, to string, circType string, formValue []byte) ([]int, error) {
	ids := make([]int, 0)
	var reterr error = nil
	c.repo.ExecTx(ctx, func(q *db.Queries) error {
		n := len(transIds)
		for i := 0; i < n; i++ {
			times, goodsId, _ := parseTransId(transIds[i])

			if err := c.valid(ctx, circType, goodsId, transIds[i], from); err != nil {
				return err
			}

			id, err := c.repo.CreateRecord(ctx, &db.CreateRecordParams{
				Transid:  transIds[i],
				Objectid: int32(goodsId),
				Circtype: sql.NullString{
					String: circType,
					Valid:  true,
				},
				Operator: sql.NullString{
					String: operator,
					Valid:  true,
				},
				From: sql.NullString{
					String: from,
					Valid:  true,
				},
				To: sql.NullString{
					String: to,
					Valid:  true,
				},
				Formvalue: formValue,
				Times:     int32(times),
			})
			if err != nil {
				logger.Error(err)
				if verr.IsDuplicate(err) {
					reterr = err
				}

				return err
			}

			ids = append(ids, int(id))
		}

		return nil
	})

	if reterr != nil {
		return nil, reterr
	}

	return ids, nil
}

func (c *CircCase) GetRecordByTransId(ctx context.Context, transId string) (*db.CircRecord, error) {
	record, err := c.repo.GetRecordByTransId(ctx, transId)
	if err != nil {
		logger.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, v1.ErrorCircNotFound("不存在的流转记录TransID:%s", transId)
		}

		return nil, err
	}

	return record, nil
}

func (c *CircCase) GetRecordByID(ctx context.Context, circId int) (*db.CircRecord, error) {
	record, err := c.repo.GetRecordById(ctx, int32(circId))
	if err != nil {
		logger.Error(err)

		return nil, v1.ErrorCircNotFound("不存在的流转记录ID:%d", circId)
	}

	return record, nil
}

func (c *CircCase) GetRecordByGoodsId(ctx context.Context, goodsId int) ([]*db.CircRecord, error) {
	record, err := c.repo.GetRecordByObjIdDesc(ctx, int32(goodsId))
	if err != nil {
		logger.Error(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, v1.ErrorCircNotFound("不存在的流转记录, 产品id[%d]不存在: %v", goodsId, err)
		}

		return nil, err
	}

	return record, nil
}

func (c *CircCase) TransId(ctx context.Context, goodsId int32) (string, error) {
	records, err := c.repo.GetRecordByObjIdDesc(ctx, int32(goodsId))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return buildTransId(0, goodsId), nil
		}

		return "", nil
	}
	if len(records) <= 0 {
		return buildTransId(0, goodsId), nil
	}

	latestTransId := records[0].Transid
	preTimes, _, _ := parseTransId(latestTransId)
	return buildTransId(preTimes+1, goodsId), nil
}

func (c *CircCase) BatchTransId(ctx context.Context, goodsIds []int32) (map[int32]string, error) {
	transIds := make(map[int32]string)
	for _, goodsId := range goodsIds {
		records, err := c.repo.GetRecordByObjIdDesc(ctx, int32(goodsId))
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				transIds[goodsId] = buildTransId(0, goodsId)
			}
			continue
		}
		if len(records) <= 0 {
			transIds[goodsId] = buildTransId(0, goodsId)
		}

		latestTransId := records[0].Transid
		preTimes, _, _ := parseTransId(latestTransId)
		transIds[goodsId] = buildTransId(preTimes+1, goodsId)
	}

	return transIds, nil
}

func (c *CircCase) UpdateCircStatus(ctx context.Context, transId, status string) error {
	err := c.repo.UpdateStatus(ctx, &db.UpdateStatusParams{
		Status:  status,
		Transid: transId,
	})
	if err != nil {
		logger.Error(err)

		return err
	}

	return nil
}

func (c *CircCase) valid(ctx context.Context, circType string, goodsId int, transId string, from string) error {
	switch circType {
	case v1.CircType_produce.String():
		return nil
	case v1.CircType_process.String(), v1.CircType_transfer.String(), v1.CircType_exam.String():
		// TODO: 5s后仍有一定概率前一个步骤未完成
		currentOwner, err := c.repo.GetObjOwner(ctx, int32(goodsId))
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			logger.Error(err)
			return err
		}

		if !currentOwner.Valid || currentOwner.String != from {
			return v1.ErrorNotOwner("当前产品所有者为[%s]， 非[%s]", currentOwner.String, from)
		}

		return nil
	default:
		return errors.New("invalid circ type")
	}
}

func buildTransId(times int, goodsId int32) string {
	return fmt.Sprintf("%d-%d-%d", times, goodsId, time.Now().Unix())
}

func parseTransId(transId string) (times, goodsId, timestamp int) {
	r := strings.SplitN(transId, "-", 3)
	times, _ = strconv.Atoi(r[0])
	goodsId, _ = strconv.Atoi(r[1])
	timestamp, _ = strconv.Atoi(r[2])

	return
}
