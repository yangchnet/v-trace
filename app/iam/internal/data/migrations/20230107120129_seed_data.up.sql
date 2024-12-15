-- 插入角色
INSERT INTO
    `role` (`rolename`)
VALUES
    ("admin"),
    ("producer"),
    ("transporter"),
    ("examiner"),
    ("normal");

-- 插入用户
INSERT INTO
    `user` (`id`, `username`, `nickname`, `passwd`, `phone`, `email`, `created_at`, `realname`, `idcard`)
VALUES
    (1, "admin", "admin", "$2a$10$0xQZY5rCsEZY0tvQhX.Xg.BR6quh6nf8SYuSoMO4hiXuolHM4LcMK", "17354277557", "lich.mailer@gmail.com", now(), "admin", "admin"),
    (2, "normalUser", "normalUser", "$2a$10$ObYRXowTLvX4hfU1HsJ4oOK2AYYshM8oavgibevs9RbmNrzpl5U6e", "11111111111", "normal@vtrace.com", now(), "", "fake-normalUser"),
    (3, "transporterUser", "transporterUser", "$2a$10$bNsR0hJMttTulQNWC0.BLuMDUUEHYdfrLp19N3s5GNHir2/rNnXne", "22222222222", "transporter@vtrace.com", now(), "transporterUser", "transporterUser"),
    (4, "examinerUser", "examinerUser", "$2a$10$hGEyH4mX2TUAbsehkOeXFe0hKZ0UhVUiHV7gl5vQ8v/acWgDgor/6", "33333333333", "examiner@vtrace.com", now(), "", "examinerUser"),
    (5, "producerUser", "producerUser", "$2a$10$SiyhTjLT2vNbDXNeLajQtutB2J9OD246nFGSp7tF0BjOYOu1SgbK.", "44444444444", "producer@vtrace.com", now(), "producerUser", "producerUser");

-- 赋予角色
INSERT INTO
    `user_role` (`username`, `rolename`, `created_at`)
VALUES
    ("admin", "admin", now()),
    ("normalUser", "normal", now()),
    ("examinerUser", "examiner", now()),
    ("transporterUser", "transporter", now()),
    ("producerUser", "producer", now());

-- 插入企业
INSERT INTO
    `org` (`id`, `org_name`, `org_code`, `legal_person_name`, `legal_person_phone`, `created_at`, `owner`)
VALUES
    (1, "高值产品溯源平台", "xxxxxxxxxxx", "admin", "11122223333", now(), "admin");

-- 插入企业成员
INSERT INTO
    `member` (`username`, `org_id`, `created_at`)
VALUES
    ("admin", 1, now()),
    ("producerUser", 1, now());