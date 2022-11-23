package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entproto.Field(1)),
		field.UUID("article_id", uuid.UUID{}).Annotations(entproto.Field(2)),
		field.UUID("user_id", uuid.UUID{}).Annotations(entproto.Field(3)),
		field.Time("read_at").Default(time.Now().UTC).Annotations(entproto.Field(4)),
	}
}

// Edges of the Article.
func (ReadArticle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("article", Article.Type).
			Ref("read_articles").
			Field("article_id").
			Required().
			Unique().
			Annotations(entproto.Field(5)),
	}
}

// Indexes of the ReadArticle.
func (ReadArticle) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "article_id").Unique(),
		index.Fields("user_id"),
	}
}

// Annotations of the ReadArticle.
func (ReadArticle) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
