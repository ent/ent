-- Create a functional (unique) index on the lowercased name column.
CREATE UNIQUE INDEX unique_name ON "users" ((lower("name")));