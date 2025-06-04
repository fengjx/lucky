package entx

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// TimeMixin 用于为所有表添加 ctime 和 utime 字段
type TimeMixin struct {
	// 可选：嵌入 ent.Mixin 以便扩展
	ent.Mixin
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("ctime").
			Default(time.Now).
			Immutable().
			Comment("创建时间").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
		field.Time("utime").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间").
			SchemaType(map[string]string{
				dialect.MySQL: "datetime",
			}),
	}
}

func (TimeMixin) Indexes() []ent.Index {
	return []ent.Index{}
}

func (TimeMixin) Edges() []ent.Edge {
	return nil
}

func (TimeMixin) Hooks() []ent.Hook {
	return []ent.Hook{}
}

func (TimeMixin) Interceptors() []ent.Interceptor {
	return nil
}

func (TimeMixin) Policy() ent.Policy {
	return nil
}

func (TimeMixin) Annotations() []schema.Annotation {
	return nil
}
