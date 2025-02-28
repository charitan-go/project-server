package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/charitan-go/project-server/internal/project/dto"
	"github.com/google/uuid"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("readable_id", uuid.UUID{}).
			Default(uuid.New).
			StructTag(`json:"readableId"`),
		field.String("name").
			NotEmpty(),
		field.String("description").
			Default("A charity project"),
		field.Float("goal").
			Positive(),
		field.Time("start_date").
			StructTag(`json:"startDate"`),
		field.Time("end_date").
			StructTag(`json:"endDate"`),
		field.Enum("category").
			GoType(dto.CategoryEnum("")),
		field.String("countryCode").
			NotEmpty(),
		field.Enum("status").
			GoType(dto.StatusEnum("")),
		field.String("owner_charity_readable_id").
			StructTag(`json:"ownerCharityReadableId"`),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return nil
}
