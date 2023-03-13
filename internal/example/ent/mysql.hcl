table "users" {
  schema  = schema.test
  collate = "utf8mb4_bin"
  column "id" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  column "name" {
    null    = false
    type    = varchar(255)
    comment = "Name of the user"
  }
  primary_key {
    columns = [column.id]
  }
}
schema "test" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
