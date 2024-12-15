CREATE TABLE `material` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(255),
    `alias` varchar(255),
    `des` text
);

CREATE TABLE `model` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `name` varchar(255) UNIQUE,
    `version` int,
    `status` varchar(255),
    `des` text,
    `metadata` json
);

CREATE TABLE `relation` (
    `id` int PRIMARY KEY AUTO_INCREMENT,
    `model_id` int,
    `material_id` int,
    `index` int
);