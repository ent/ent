CREATE DOMAIN us_postal_code AS TEXT
CHECK(
   VALUE ~ '^\d{5}$'
   OR VALUE ~ '^\d{5}-\d{4}$'
);