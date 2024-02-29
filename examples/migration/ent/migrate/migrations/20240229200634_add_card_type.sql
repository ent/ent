-- Modify "cards" table
ALTER TABLE `cards` ADD COLUMN `type` varchar(255) NOT NULL DEFAULT "unknown";
