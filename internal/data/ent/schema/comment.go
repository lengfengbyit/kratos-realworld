package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id").Optional(),
		field.Int64("article_id"),
		field.String("body").NotEmpty(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).Ref("comments").Field("user_id").Unique(),
	}
}

func (Comment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("article_id"),
		index.Fields("user_id"),
	}
}
