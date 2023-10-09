// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateBookInput represents a mutation input for creating books.
type CreateBookInput struct {
	Title     string
	Body      *string
	CreatedAt *time.Time
}

// Mutate applies the CreateBookInput on the BookMutation builder.
func (i *CreateBookInput) Mutate(m *BookMutation) {
	m.SetTitle(i.Title)
	if v := i.Body; v != nil {
		m.SetBody(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
}

// SetInput applies the change-set in the CreateBookInput on the BookCreate builder.
func (c *BookCreate) SetInput(i CreateBookInput) *BookCreate {
	i.Mutate(c.Mutation())
	return c
}
