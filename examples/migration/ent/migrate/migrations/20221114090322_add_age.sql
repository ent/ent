-- modify "users" table
ALTER TABLE `users` ADD CHECK (age > 0), ADD COLUMN `age` double NOT NULL;
