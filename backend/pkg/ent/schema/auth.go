package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Article holds the schema definition for the Article entity.
type Auth struct {
	ent.Schema
}

// Fields of the Auth.
func (Auth) Fields() []ent.Field {
	return []ent.Field{
		// ユーザーID
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("login_id").Unique(),
		field.String("email").Unique(),
		field.String("password"),
	}
}

// Edges of the Auth.
func (Auth) Edges() []ent.Edge {
	return nil
}

// Indexes of the Auth.
func (Auth) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("login_id").Unique().StorageKey("login_id_index"),
	}
}
