INSERT INTO `menus` (`label`) VALUES ('Admin');
INSERT INTO `menus` (`label`) VALUES ('User');

INSERT INTO `menu_items` (`menu_id`, `label`, `path`, `ordering`) VALUES (1, 'Home', 'home', 1);
INSERT INTO `menu_items` (`menu_id`, `label`, `path`, `ordering`) VALUES (1, 'Users', 'users', 2);
INSERT INTO `menu_items` (`menu_id`, `label`, `path`, `ordering`) VALUES (1, 'Contacts', 'contacts', 3);
INSERT INTO `menu_items` (`menu_id`, `label`, `path`, `ordering`) VALUES (1, 'Accounts', 'accounts', 4);
INSERT INTO `menu_items` (`menu_id`, `label`, `path`, `ordering`) VALUES (1, 'Settings', 'settings', 5);