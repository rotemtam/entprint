table "users" {
  schema = schema.main
  column "id" {
    null           = false
    type           = integer
    auto_increment = true
  }
  column "name" {
    null    = false
    type    = text
    comment = "Name of the user"
  }
  primary_key {
    columns = [column.id]
  }
}
schema "main" {
}
