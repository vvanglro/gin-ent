package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(), // Positive number
		field.String("name").
			Unique().Comment("用户名"),
		field.String("password").Sensitive().Comment("密码"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
