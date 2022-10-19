package schema

import (
	"time"

	"entgo.io/ent"
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
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.UUID("user_id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.String("login_id").Unique(),
		field.String("email").Unique(),
		field.String("password"),
		field.Time("created_at").Default(time.Now().UTC),
		field.Time("updated_at").Default(time.Now().UTC).UpdateDefault(time.Now().UTC),
	}
}

// Edges of the Auth.
func (Auth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("auths").
			Field("user_id").
			Required().
			Unique(),
	}
}

// Indexes of the Auth.
func (Auth) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("login_id").Unique().StorageKey("login_id_index"),
	}
}
