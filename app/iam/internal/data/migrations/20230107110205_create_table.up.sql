CREATE TABLE `user` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) UNIQUE NOT NULL,
  `nickname` varchar(255),
  `passwd` varchar(255) NOT NULL,
  `phone` varchar(255) NOT NULL,
  `email` varchar(255) UNIQUE,
  `created_at` timestamp,
  `realname` varchar(255),
  `idcard` varchar(255) UNIQUE
);

CREATE INDEX `usernameIdx` ON `user` (`username`);

CREATE TABLE `role` (`rolename` varchar(255) PRIMARY KEY);

CREATE TABLE `user_role` (
  `username` varchar(255),
  `rolename` varchar(255),
  `created_at` timestamp,
  CONSTRAINT relation PRIMARY KEY(`username`, `rolename`)
);

CREATE TABLE `member` (
  `username` varchar(255),
  `org_id` int,
  `created_at` timestamp,
  CONSTRAINT relation PRIMARY KEY(`username`, `org_id`)
);

CREATE TABLE `org` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `org_name` varchar(255),
  `org_code` varchar(255) UNIQUE,
  `legal_person_name` varchar(255),
  `legal_person_phone` varchar(255),
  `created_at` timestamp,
  `owner` varchar(255) NOT NULL
);