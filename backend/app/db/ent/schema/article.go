package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
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
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entproto.Field(1)),
		field.String("title").Annotations(entproto.Field(2)),
		field.String("url").Unique().Annotations(entproto.Field(3)),
		field.String("description").Annotations(entproto.Field(4)),
		field.String("thumbnail").Annotations(entproto.Field(5)),
		field.Time("created_at").Default(time.Now().UTC).Annotations(entproto.Field(6)),
		field.Time("updated_at").Default(time.Now().UTC).UpdateDefault(time.Now().UTC).Annotations(entproto.Field(7)),
		field.Time("deleted_at").Optional().Nillable().Annotations(entproto.Field(8)),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tags", ArticleTag.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}).
			Annotations(entproto.Field(9)),
		edge.To("read_articles", ReadArticle.Type).
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}).
			Annotations(entproto.Field(10)),
	}
}

// Indexes of the Article.
func (Article) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("title"),
	}
}

// Annotations of the Article.
func (Article) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
