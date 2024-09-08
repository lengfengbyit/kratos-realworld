package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Follow holds the schema definition for the Follow entity.
type Follow struct {
	ent.Schema
}

// Fields of the Follow.
func (Follow) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.Int64("user_id").Default(0),    // 关注用户 ID
		field.Int64("be_user_id").Default(0), // 被关注用户 ID
		field.Time("created_at").Immutable().Default(time.Now),
	}
}

func (Follow) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "be_user_id").Unique(),
		index.Fields("be_user_id"),
	}
}

// Edges of the Follow.
func (Follow) Edges() []ent.Edge {
	return nil
}
