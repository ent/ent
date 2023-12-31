-- Modify "cards" table
ALTER TABLE `cards` ADD CONSTRAINT `number_length` CHECK (length(`number`) = 16), ADD COLUMN `number` varchar(255) NOT NULL;
