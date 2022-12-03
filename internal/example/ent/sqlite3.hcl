table "users" {
  schema = schema.main
  column "id" {
    null           = false
    type           = integer
    auto_increment = true
  }
  primary_key {
    columns = [column.id]
  }
}
schema "main" {
}
