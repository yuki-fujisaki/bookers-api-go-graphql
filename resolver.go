package bookers

import (
	"bookers/ent"

	"github.com/99designs/gqlgen/graphql"
)

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{client},
	})
}
