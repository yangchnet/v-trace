-- name: CreateTrans :execlastid
-- CreateTrans 创建交易记录
INSERT INTO
    `trans_record` (`transId`, `sender`, `contract`, `method`, `params`, `status`, `txHash`, `tx_params_hash`, created_at)
VALUES
    (?, ?, ?, ?, ?, ?, ?, ?, now());

-- name: UpdateTrans :execlastid
-- 根据transId更新交易状态
UPDATE
    `trans_record`
SET
    `txHash` = ?, `status` = ?, `updated_at` = now()
WHERE
    `transId` = ?;

-- name: GetTransByTransId :one
-- 根据transId获取记录
SELECT
    *
FROM
    `trans_record`
WHERE
    `transId` = ?;