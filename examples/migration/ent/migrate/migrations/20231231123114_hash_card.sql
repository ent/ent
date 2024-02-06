-- Modify "cards" table
ALTER TABLE `cards` DROP CONSTRAINT `number_length`, ADD CONSTRAINT `number_hash_length` CHECK (length(`number_hash`) = 16), RENAME COLUMN `number` TO `number_hash`, ADD COLUMN `cvv_hash` varchar(255) NOT NULL, ADD COLUMN `expires_at` timestamp NULL;
-- Modify "pets" table
ALTER TABLE `pets` ADD COLUMN `age` double NOT NULL AFTER `id`, ADD COLUMN `weight` double NOT NULL AFTER `id`;
