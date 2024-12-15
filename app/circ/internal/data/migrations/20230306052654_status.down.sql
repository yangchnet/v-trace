CREATE TABLE `object_status`(
    `objectId`      int          NOT NULL UNIQUE,
    `current_owner` varchar(255) NOT NULL
);

ALTER TABLE `circ_record` DROP `status`;