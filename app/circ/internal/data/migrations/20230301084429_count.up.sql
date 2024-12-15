ALTER TABLE `circ_record` ADD `times` int NOT NULL ; -- 商品的第几次流转
ALTER TABLE `circ_record` MODIFY COLUMN `transId` varchar(255) NOT NULL; -- transId不空

CREATE UNIQUE INDEX `obj_times_idx` ON `circ_record` (`objectId`, `times`); -- 防止流转记录重复