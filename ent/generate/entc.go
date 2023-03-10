//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithConfigPath("../../gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("../../graph/schemas/schema.graphqls"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	options := []entc.Option{
		entc.Extensions(ex),
	}

	entc.Generate("../schema", &gen.Config{}, options...)
}
