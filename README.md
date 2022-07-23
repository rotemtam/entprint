# entprint

a tool to print [ent](https://github.com/ent/ent) schemas as Atlas HCL documents

### Docs

From within your Ent project:

```text
go run -mod=mod github.com/rotemtam/entprint --help
```

Output:

```text
Usage: entprint --dir=STRING --dev=STRING

a tool to print ent schemas as Atlas HCL documents

Flags:
  -h, --help                      Show context-sensitive help.
      --dir=STRING                ent schema directory
      --dev=STRING                dev db url
      --with-global-unique-ids    set global unique ids to true
```

### Example usage

Suppose your Ent schema is located in `./ent/schema`, and you have an
empty MySQL db running at `mysql://root:pass@localhost:3306/test`:

```text
go run github.com/rotemtam/entprint  --dir ./internal/example/ent/schema --dev mysql://root:pass@localhost:3306/test
```

Output:

```hcl
table "users" {
  schema  = schema.test
  collate = "utf8mb4_bin"
  column "id" {
    null           = false
    type           = bigint
    auto_increment = true
  }
  primary_key {
    columns = [column.id]
  }
}
schema "test" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
```

### Support

Join us on our [Discord server](https://discord.gg/qZmPgTE6RX).