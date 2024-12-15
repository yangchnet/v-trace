CREATE TABLE `class` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255),
  `des` json,
  `status` varchar(255),
  `created_at` timestamp,
  `creator` varchar(255),
  `material_id` int NOT NULL
);

CREATE TABLE `serial` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `product_time` date,
  `status` varchar(255),
  `created_at` timestamp,
  `creator` varchar(255),
  `class_id` int
);

CREATE TABLE `goods` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `status` varchar(255),
  `created_at` timestamp,
  `creator` varchar(255),
  `serial_id` int
);