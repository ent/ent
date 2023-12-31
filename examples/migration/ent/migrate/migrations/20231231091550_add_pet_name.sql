-- Modify "pets" table
ALTER TABLE `pets` ADD COLUMN `name` varchar(255) NOT NULL, ADD UNIQUE INDEX `pet_name_owner_id` (`name`, `owner_id`);
