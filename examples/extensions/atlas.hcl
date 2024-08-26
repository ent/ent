data "composite_schema" "app" {
  # Install extensions first (PostGIS).
  schema "public" {
    url = "file://schema.sql"
  }
  # Then, load the Ent schema.
  schema "public" {
    url = "ent://ent/schema"
  }
}

env "local" {
  src = data.composite_schema.app.url
  dev = "docker://postgis/latest/dev"
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}