data "composite_schema" "app" {
  # Load first custom types first.
  schema "public" {
    url = "file://schema.sql"
  }
  # Second, load the Ent schema.
  schema "public" {
    url = "ent://ent/schema"
  }
}

env "local" {
  src = data.composite_schema.app.url
  dev = "docker://postgres/15/dev?search_path=public"
}