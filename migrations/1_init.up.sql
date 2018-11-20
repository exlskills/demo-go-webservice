CREATE TABLE `gopher` (
	`id` INT UNSIGNED AUTO_INCREMENT,
    `full_name` VARCHAR(255) NOT NULL,
    `headline` VARCHAR(255) NOT NULL,
    `avatar_url` VARCHAR(255) NOT NULL,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

