ALTER TABLE `circ_record` MODIFY COLUMN `transId` varchar(255);
DROP INDEX `obj_times_idx` ON `circ_record`;
ALTER TABLE `circ_record` DROP `times`;
