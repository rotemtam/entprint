package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "ariga.io/atlas/sql/mysql"
	_ "ariga.io/atlas/sql/postgres"
	atlas "ariga.io/atlas/sql/schema"
	"ariga.io/atlas/sql/sqlclient"
	_ "ariga.io/atlas/sql/sqlite"

	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/alecthomas/kong"

	_ "github.com/rotemtam/entprint/internal/docker"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

var CLI struct {
	Dir                 string `help:"ent schema directory" required:""`
	Dev                 string `help:"dev db url" required:""`
	WithGlobalUniqueIDs bool   `name:"with-global-unique-ids" help:"set global unique ids to true"`
}

var skip = errors.New("skip")

func main() {
	kong.Parse(&CLI,
		kong.Name("entprint"),
		kong.Description("a tool to print ent schemas as Atlas HCL documents"))
	graph, err := entc.LoadGraph(CLI.Dir, &gen.Config{})
	if err != nil {
		log.Fatalf("loading schema: %v", err)
	}
	var sch *atlas.Schema
	opts := []schema.MigrateOption{
		schema.WithGlobalUniqueID(CLI.WithGlobalUniqueIDs),
		schema.WithDiffHook(func(differ schema.Differ) schema.Differ {
			return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
				sch = desired
				return nil, skip
			})
		}),
	}

	parts, err := url.Parse(CLI.Dev)
	if err != nil {
		log.Fatalf("parsing dev url: %v", err)
	}
	if parts.Scheme == "docker" {
		opts = append(opts, schema.WithDialect(parts.Host))
	}
	mig, err := schema.NewMigrateURL(CLI.Dev, opts...)
	if err != nil {
		log.Fatalf("connecting: %v", err)
	}
	tbl, err := graph.Tables()
	if err != nil {
		log.Fatalf("reading tables: %v", err)
	}
	if err := mig.Create(context.Background(), tbl...); err != nil && !errors.Is(err, skip) {
		log.Fatalf("failed: %v", err)
	}
	drv, err := sqlclient.Open(context.Background(), CLI.Dev)
	if err != nil {
		log.Fatalf("connecting: %v", err)
	}
	norm, ok := drv.Driver.(atlas.Normalizer)
	if !ok {
		log.Fatalf("driver %T does not impl Normalizer", drv.Driver)
	}
	sch, err = norm.NormalizeSchema(context.Background(), sch)
	if err != nil {
		log.Fatalf("normalzing schema: %v", err)
	}
	spec, err := drv.MarshalSpec(sch)
	if err != nil {
		log.Fatalf("marshaling schema: %v", err)
	}
	fmt.Fprint(os.Stdout, string(spec))
}
