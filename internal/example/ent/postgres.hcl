table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
    }
  }
  column "name" {
    null    = false
    type    = character_varying
    comment = "Name of the user"
  }
  primary_key {
    columns = [column.id]
  }
}
schema "public" {
}
