CREATE TABLE `trans_record` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `transId` varchar(255) UNIQUE NOT NULL,
  `sender` varchar(255) NOT NULL,
  `contract` varchar(255) NOT NULL,
  `method` varchar(255) NOT NULL,
  `params` varchar(1000),
  `status` varchar(255),
  `txHash` varchar(100),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE INDEX `transID_Idx` ON `trans_record` (`transId`);