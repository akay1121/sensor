package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Sensor 定义传感器的 schema
type Sensor struct {
	ent.Schema
}

// Fields 定义传感器字段
func (Sensor) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Comment("sensor id"),
		field.Int64("type_id").
			Comment("sensor type"),
		field.String("desc").
			Optional().
			Default("").
			Comment("description on sensor"),
		field.Int64("rule_id").
			Optional().
			Default(-1).
			Comment("sensor rule iD"),
		field.Float("threshold").
			Optional().
			Default(0).
			Comment("sensor threshold"),
		field.Float("previous_value").
			Optional().
			Default(0).
			Comment("old value of sensor"),
		field.Bool("deleted").
			Default(false).
			Comment("whether is deleted"),
		field.Time("create_time").
			Default(time.Now).
			Immutable().
			Comment("create time"),
		field.Time("last_update").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("last update time"),
	}
}

// Indexes
func (Sensor) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type_id"),
	}
}
