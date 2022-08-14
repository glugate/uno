CREATE TABLE IF NOT EXISTS users
	(
		`id`          INTEGER PRIMARY KEY AUTO_INCREMENT,
		`first_name`  VARCHAR(255) NOT NULL,
		`last_name`   VARCHAR(255) NOT NULL,
		`email`       VARCHAR(255) NOT NULL,
		`phone`       VARCHAR(255) NOT NULL,
		`birth_date`  DATE,
		`is_active`   BOOLEAN DEFAULT true,
		`status`      VARCHAR(255) NOT NULL,
		`title`       VARCHAR(255) NOT NULL,
		`password`    VARCHAR(255) NOT NULL,
		`created_at`  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		`updated_at`  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)
	