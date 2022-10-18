package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// ReadArticle holds the schema definition for the Article entity.
type ReadArticle struct {
	ent.Schema
}

// Fields of the ReadArticle.
func (ReadArticle) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("article_id", uuid.UUID{}),
		field.UUID("user_id", uuid.UUID{}),
		field.Time("read_at").Default(time.Now().UTC),
	}
}

// Edges of the Article.
func (ReadArticle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("article", Article.Type).
			Ref("read_articles").
			Field("article_id").
			Required().
			Unique(),
	}
}

// Indexes of the ReadArticle.
func (ReadArticle) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "article_id").Unique(),
		index.Fields("user_id"),
	}
}
