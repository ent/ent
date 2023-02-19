-- create "cars" table
CREATE TABLE `cars` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `model` text NOT NULL, `registered_at` datetime NOT NULL, `user_cars` integer NULL, CONSTRAINT `cars_users_cars` FOREIGN KEY (`user_cars`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- create "groups" table
CREATE TABLE `groups` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL);
-- create "users" table
CREATE TABLE `users` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `age` integer NOT NULL, `name` text NOT NULL DEFAULT 'unknown');
-- create "group_users" table
CREATE TABLE `group_users` (`group_id` integer NOT NULL, `user_id` integer NOT NULL, PRIMARY KEY (`group_id`, `user_id`), CONSTRAINT `group_users_group_id` FOREIGN KEY (`group_id`) REFERENCES `groups` (`id`) ON DELETE CASCADE, CONSTRAINT `group_users_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
