-- Create "clean_users" view
CREATE VIEW "clean_users" ("id", "name", "public_info") AS SELECT id,
    name,
    public_info
   FROM users;

-- Create "pet_user_names" view
CREATE VIEW "pet_user_names" ("name") AS SELECT DISTINCT name
   FROM ( SELECT users.name
           FROM users
        UNION
         SELECT pets.name
           FROM pets) all_names;