CREATE TABLE menu_items
	(
		`id`         INTEGER PRIMARY KEY AUTO_INCREMENT,
		`menu_id`   INTEGER,
		`label`      TEXT NOT NULL,
		`path`       TEXT,
		`ordering`   INTEGER,
		`parent_id`  INTEGER,
		`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);