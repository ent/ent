data "composite_schema" "app" {
  # Load the ent schema first with all tables.
  schema "public" {
    url = "ent://ent/schema"
  }
  # Then, load the views defined in the schema.sql file.
  schema "public" {
    url = "file://schema.sql"
  }
}

env "local" {
  src = data.composite_schema.app.url
  dev = "docker://postgres/15/dev?search_path=public"
}