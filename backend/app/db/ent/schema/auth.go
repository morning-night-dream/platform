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

// Auth holds the schema definition for the Auth entity.
type Auth struct {
	ent.Schema
}

// Fields of the Auth.
func (Auth) Fields() []ent.Field {
	return []ent.Field{
		// ユーザーID
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Annotations(entproto.Field(1)),
		field.UUID("user_id", uuid.UUID{}).Default(uuid.New).Unique().Annotations(entproto.Field(2)),
		field.String("login_id").Unique().Annotations(entproto.Field(3)),
		field.String("email").Unique().Annotations(entproto.Field(4)),
		field.String("password").Annotations(entproto.Field(5)),
		field.Time("created_at").Default(time.Now().UTC).Annotations(entproto.Field(6)),
		field.Time("updated_at").Default(time.Now().UTC).UpdateDefault(time.Now().UTC).Annotations(entproto.Field(7)),
	}
}

// Edges of the Auth.
func (Auth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("auths").
			Field("user_id").
			Required().
			Unique().
			Annotations(entproto.Field(8)),
	}
}

// Indexes of the Auth.
func (Auth) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("login_id").Unique().StorageKey("login_id_index"),
	}
}

// Annotations of the Auth.
func (Auth) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(),
	}
}
