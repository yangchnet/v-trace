INSERT INTO
    model (`id`, `name`, `version`, `status`, `des`, `metadata`)
VALUES
    (1, "model1", 1, "", "", "{}"),
    (2, "model2", 1, "", "", "{}");

-- model1（反射式）：['中国荷斯坦牛' '丹顶鹤' '云南驴' '利木赞牛' '南阳牛' '夏洛来牛' '天猫' '安格斯牛' '延边黄牛' '徐州牛' '德国黄牛' '德州驴' '摩拉水牛' '柴犬' '梅花鹿' '水牛' '渤海黑牛' '熊猫' '犏牛' '猪' '田园犬' '短角牛' '腾冲马' '英短' '荷斯坦奶牛' '蒙古牛' '袋鼠' '西藏牛' '西藏马' '西藏驴' '西门塔尔牛' '金丝猴' '青海黄牛' '马' '鸡' '鸭' '鸽子' '鹅' '黄牛' '黑熊']
-- model2（透射式）：['中国荷斯坦牛' '云南驴' '利木赞牛' '南阳牛' '原鸡' '夏洛来牛' '安格斯牛' '山鸡' '延边黄牛' '徐州牛' '德国黄牛' '德州驴' '摩拉水牛' '渤海黑牛' '犏牛' '猪' '瑞士褐牛' '白冠长尾雉' '白枕鹤' '白鹳' '短角牛' '红腹锦鸡' '腾冲马' '荷斯坦奶牛' '蒙古牛' '蓑羽鹤' '蓝黄鹦鹉' '西藏牛' '西藏马' '西藏驴' '西门塔尔牛' '辽白牛' '金丝猴' '金钱豹' '长臂猿' '青海黄牛' '马' '鸡' '鸭' '鸽子' '鹅' '黑熊']

INSERT INTO
    material (`id`, `name`, `alias`, `des`)
VALUES
    (1, "中国荷斯坦牛",  "", ""),
    (2, "丹顶鹤",  "", ""),
    (3, "云南驴",  "", "");

INSERT INTO
    relation (`id`, `model_id`, `material_id`, `index`)
VALUES
    (1, 1, 1, 0),
    (2, 1, 2, 1),
    (3, 1, 3, 2),
    (4, 2, 1, 0),
    (5, 2, 3, 1);