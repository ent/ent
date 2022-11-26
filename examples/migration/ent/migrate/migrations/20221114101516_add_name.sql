-- modify "users" table
ALTER TABLE `users` ADD CONSTRAINT `name_not_empty` CHECK (name <> ''), ADD COLUMN `name` varchar(255) NOT NULL;
