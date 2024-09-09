// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kratos-realworld/internal/data/ent/article"
	"kratos-realworld/internal/data/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ArticleCreate is the builder for creating a Article entity.
type ArticleCreate struct {
	config
	mutation *ArticleMutation
	hooks    []Hook
}

// SetAuthorID sets the "author_id" field.
func (ac *ArticleCreate) SetAuthorID(i int64) *ArticleCreate {
	ac.mutation.SetAuthorID(i)
	return ac
}

// SetSlug sets the "slug" field.
func (ac *ArticleCreate) SetSlug(u uuid.UUID) *ArticleCreate {
	ac.mutation.SetSlug(u)
	return ac
}

// SetNillableSlug sets the "slug" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableSlug(u *uuid.UUID) *ArticleCreate {
	if u != nil {
		ac.SetSlug(*u)
	}
	return ac
}

// SetTitle sets the "title" field.
func (ac *ArticleCreate) SetTitle(s string) *ArticleCreate {
	ac.mutation.SetTitle(s)
	return ac
}

// SetDescription sets the "description" field.
func (ac *ArticleCreate) SetDescription(s string) *ArticleCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableDescription(s *string) *ArticleCreate {
	if s != nil {
		ac.SetDescription(*s)
	}
	return ac
}

// SetBody sets the "body" field.
func (ac *ArticleCreate) SetBody(s string) *ArticleCreate {
	ac.mutation.SetBody(s)
	return ac
}

// SetTags sets the "tags" field.
func (ac *ArticleCreate) SetTags(s string) *ArticleCreate {
	ac.mutation.SetTags(s)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *ArticleCreate) SetCreatedAt(t time.Time) *ArticleCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableCreatedAt(t *time.Time) *ArticleCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *ArticleCreate) SetUpdatedAt(t time.Time) *ArticleCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *ArticleCreate) SetNillableUpdatedAt(t *time.Time) *ArticleCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *ArticleCreate) SetID(i int64) *ArticleCreate {
	ac.mutation.SetID(i)
	return ac
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (ac *ArticleCreate) SetOwnerID(id int64) *ArticleCreate {
	ac.mutation.SetOwnerID(id)
	return ac
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (ac *ArticleCreate) SetNillableOwnerID(id *int64) *ArticleCreate {
	if id != nil {
		ac = ac.SetOwnerID(*id)
	}
	return ac
}

// SetOwner sets the "owner" edge to the User entity.
func (ac *ArticleCreate) SetOwner(u *User) *ArticleCreate {
	return ac.SetOwnerID(u.ID)
}

// Mutation returns the ArticleMutation object of the builder.
func (ac *ArticleCreate) Mutation() *ArticleMutation {
	return ac.mutation
}

// Save creates the Article in the database.
func (ac *ArticleCreate) Save(ctx context.Context) (*Article, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ArticleCreate) SaveX(ctx context.Context) *Article {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ArticleCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ArticleCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ArticleCreate) defaults() {
	if _, ok := ac.mutation.Slug(); !ok {
		v := article.DefaultSlug()
		ac.mutation.SetSlug(v)
	}
	if _, ok := ac.mutation.Description(); !ok {
		v := article.DefaultDescription
		ac.mutation.SetDescription(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := article.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := article.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ArticleCreate) check() error {
	if _, ok := ac.mutation.AuthorID(); !ok {
		return &ValidationError{Name: "author_id", err: errors.New(`ent: missing required field "Article.author_id"`)}
	}
	if _, ok := ac.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Article.slug"`)}
	}
	if _, ok := ac.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Article.title"`)}
	}
	if v, ok := ac.mutation.Title(); ok {
		if err := article.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Article.title": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Article.description"`)}
	}
	if _, ok := ac.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "Article.body"`)}
	}
	if _, ok := ac.mutation.Tags(); !ok {
		return &ValidationError{Name: "tags", err: errors.New(`ent: missing required field "Article.tags"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Article.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Article.updated_at"`)}
	}
	return nil
}

func (ac *ArticleCreate) sqlSave(ctx context.Context) (*Article, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ArticleCreate) createSpec() (*Article, *sqlgraph.CreateSpec) {
	var (
		_node = &Article{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(article.Table, sqlgraph.NewFieldSpec(article.FieldID, field.TypeInt64))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.AuthorID(); ok {
		_spec.SetField(article.FieldAuthorID, field.TypeInt64, value)
		_node.AuthorID = value
	}
	if value, ok := ac.mutation.Slug(); ok {
		_spec.SetField(article.FieldSlug, field.TypeUUID, value)
		_node.Slug = value
	}
	if value, ok := ac.mutation.Title(); ok {
		_spec.SetField(article.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.SetField(article.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ac.mutation.Body(); ok {
		_spec.SetField(article.FieldBody, field.TypeString, value)
		_node.Body = value
	}
	if value, ok := ac.mutation.Tags(); ok {
		_spec.SetField(article.FieldTags, field.TypeString, value)
		_node.Tags = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(article.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(article.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := ac.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   article.OwnerTable,
			Columns: []string{article.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_articles = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ArticleCreateBulk is the builder for creating many Article entities in bulk.
type ArticleCreateBulk struct {
	config
	err      error
	builders []*ArticleCreate
}

// Save creates the Article entities in the database.
func (acb *ArticleCreateBulk) Save(ctx context.Context) ([]*Article, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Article, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ArticleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ArticleCreateBulk) SaveX(ctx context.Context) []*Article {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ArticleCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ArticleCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
