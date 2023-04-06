-- Backfill NULL or null tags with a default value.
UPDATE `users` SET `tags` = '["foo","bar"]' WHERE `tags` IS NULL OR JSON_CONTAINS(`tags`, 'null', '$');
-- Append the 'org' tag for organization accounts in case they don't have it.
UPDATE `users` SET `tags` = CASE WHEN (JSON_TYPE(JSON_EXTRACT(`tags`, '$')) IS NULL OR JSON_TYPE(JSON_EXTRACT(`tags`, '$')) = 'NULL') THEN JSON_ARRAY('org') ELSE JSON_ARRAY_APPEND(`tags`, '$', 'org') END WHERE (`users`.`name` LIKE 'org-%' OR `users`.`name` LIKE '%-org') AND (NOT (JSON_CONTAINS(`tags`, '"org"', '$') = 1));
