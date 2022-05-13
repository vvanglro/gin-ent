package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

type TimeMixin struct {
	mixin.Schema
}

func NowOfUTC() time.Time {
	l, _ := time.LoadLocation("Asia/Shanghai")
	return time.Now().In(l)
}

// 自定义字段类型
var timeSchema = map[string]string{
	dialect.Postgres: "timestamp",
	dialect.SQLite:   "timestamp",
	dialect.MySQL:    "datetime(3)",
}

// Fields of the AuditMixin.
func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").SchemaType(timeSchema).
			Immutable().
			Default(NowOfUTC), // 这里time.now 也可以因为使用了postgres的timestamp类型 不在保存为utc时间
		field.Time("updated_at").SchemaType(timeSchema).
			Optional().
			UpdateDefault(NowOfUTC),
		field.Time("deleted_at").SchemaType(timeSchema).
			StructTag(`json:"-"`).
			Optional(),
	}
}
