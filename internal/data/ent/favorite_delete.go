// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"kratos-realworld/internal/data/ent/favorite"
	"kratos-realworld/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FavoriteDelete is the builder for deleting a Favorite entity.
type FavoriteDelete struct {
	config
	hooks    []Hook
	mutation *FavoriteMutation
}

// Where appends a list predicates to the FavoriteDelete builder.
func (fd *FavoriteDelete) Where(ps ...predicate.Favorite) *FavoriteDelete {
	fd.mutation.Where(ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FavoriteDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fd.sqlExec, fd.mutation, fd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FavoriteDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FavoriteDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(favorite.Table, sqlgraph.NewFieldSpec(favorite.FieldID, field.TypeInt))
	if ps := fd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fd.mutation.done = true
	return affected, err
}

// FavoriteDeleteOne is the builder for deleting a single Favorite entity.
type FavoriteDeleteOne struct {
	fd *FavoriteDelete
}

// Where appends a list predicates to the FavoriteDelete builder.
func (fdo *FavoriteDeleteOne) Where(ps ...predicate.Favorite) *FavoriteDeleteOne {
	fdo.fd.mutation.Where(ps...)
	return fdo
}

// Exec executes the deletion query.
func (fdo *FavoriteDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{favorite.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FavoriteDeleteOne) ExecX(ctx context.Context) {
	if err := fdo.Exec(ctx); err != nil {
		panic(err)
	}
}
