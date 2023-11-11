-- Add new schema named "db1"
CREATE DATABASE IF NOT EXISTS`db1`;
-- Add new schema named "db2"
CREATE DATABASE IF NOT EXISTS `db2`;
-- Add new schema named "db3"
CREATE DATABASE IF NOT EXISTS `db3`;
-- Create "users" table
CREATE TABLE `db3`.`users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT "unknown",
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "friendships" table
CREATE TABLE `db1`.`friendships` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `weight` bigint NOT NULL DEFAULT 1,
  `created_at` timestamp NOT NULL,
  `user_id` bigint NOT NULL,
  `friend_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `friendship_created_at` (`created_at`),
  UNIQUE INDEX `friendship_user_id_friend_id` (`user_id`, `friend_id`),
  INDEX `friendships_users_friend` (`friend_id`),
  CONSTRAINT `friendships_users_friend` FOREIGN KEY (`friend_id`) REFERENCES `db3`.`users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT `friendships_users_user` FOREIGN KEY (`user_id`) REFERENCES `db3`.`users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "groups" table
CREATE TABLE `db1`.`groups` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT "unknown",
  PRIMARY KEY (`id`)
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "group_users" table
CREATE TABLE `db1`.`group_users` (
  `group_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  PRIMARY KEY (`group_id`, `user_id`),
  INDEX `group_users_user_id` (`user_id`),
  CONSTRAINT `group_users_group_id` FOREIGN KEY (`group_id`) REFERENCES `db1`.`groups` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT `group_users_user_id` FOREIGN KEY (`user_id`) REFERENCES `db3`.`users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "pets" table
CREATE TABLE `db2`.`pets` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT "unknown",
  `owner_id` bigint NULL,
  PRIMARY KEY (`id`),
  INDEX `pets_users_pets` (`owner_id`),
  CONSTRAINT `pets_users_pets` FOREIGN KEY (`owner_id`) REFERENCES `db3`.`users` (`id`) ON UPDATE NO ACTION ON DELETE SET NULL
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "user_following" table
CREATE TABLE `db3`.`user_following` (
  `user_id` bigint NOT NULL,
  `follower_id` bigint NOT NULL,
  PRIMARY KEY (`user_id`, `follower_id`),
  INDEX `user_following_follower_id` (`follower_id`),
  CONSTRAINT `user_following_follower_id` FOREIGN KEY (`follower_id`) REFERENCES `db3`.`users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE,
  CONSTRAINT `user_following_user_id` FOREIGN KEY (`user_id`) REFERENCES `db3`.`users` (`id`) ON UPDATE NO ACTION ON DELETE CASCADE
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
