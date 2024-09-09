package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"), // 会自动设置为主键
		field.Int64("author_id"),
		field.UUID("slug", uuid.UUID{}).Default(uuid.New), // uuid 用于接口搜索
		field.String("title").NotEmpty(),
		field.String("description").Default(""),
		field.String("body"),
		field.String("tags"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

func (Article) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("slug").Unique(),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("articles").Unique(),
	}
}
