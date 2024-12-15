CREATE TABLE `circ_record` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `transId` varchar(255),
  `objectId` int NOT NULL,
  `circType` varchar(255),
  `operator` varchar(255),
  `from` varchar(255),
  `to` varchar(255),
  `created_at` timestamp,
  `formValue` json
)
