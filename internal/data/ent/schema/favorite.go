package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Favorite struct {
	ent.Schema
}

func (Favorite) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("user_id"),
		field.Int64("article_id"),
		field.Time("created_at").Default(time.Now),
	}
}

func (Favorite) Edges() []ent.Edge {
	return nil
}
