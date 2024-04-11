-- Modify "payments" table
ALTER TABLE `payments` MODIFY COLUMN `currency` enum('USD','EUR','VND','ILS') NOT NULL;
