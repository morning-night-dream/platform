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

// ArticleTag holds the schema definition for the ArticleTag entity.
type ArticleTag struct {
	ent.Schema
}

// Fields of the ArticleTag.
func (ArticleTag) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entproto.Field(1)),
		field.String("tag").Annotations(entproto.Field(2)),
		field.UUID("article_id", uuid.UUID{}).Annotations(entproto.Field(3)),
		field.Time("created_at").Default(time.Now().UTC).Annotations(entproto.Field(4)),
		field.Time("updated_at").Default(time.Now().UTC).UpdateDefault(time.Now().UTC).Annotations(entproto.Field(5)),
	}
}

// Edges of the ArticleTag.
func (ArticleTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("article", Article.Type).
			Ref("tags").
			Unique().
			Field("article_id").
			// https://github.com/ent/ent/issues/1561
			Required().
			Annotations(entproto.Field(6)),
	}
}

// Indexes of the ArticleTag.
func (ArticleTag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tag", "article_id").Unique(),
		index.Fields("tag"),
		index.Fields("article_id"),
	}
}

// Annotations of the ArticleTag.
func (ArticleTag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
