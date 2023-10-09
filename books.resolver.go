package bookers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"bookers/ent"
)

type Resolver struct {
	Client *ent.Client
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Books(ctx context.Context) (*ent.Book, error) {
	return r.Client.Book.Query().First(ctx)
}
