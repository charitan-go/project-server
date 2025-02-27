// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/charitan-go/project-server/ent/predicate"
	"github.com/charitan-go/project-server/ent/project"
	"github.com/charitan-go/project-server/internal/project/dto"
	"github.com/google/uuid"
)

// ProjectUpdate is the builder for updating Project entities.
type ProjectUpdate struct {
	config
	hooks    []Hook
	mutation *ProjectMutation
}

// Where appends a list predicates to the ProjectUpdate builder.
func (pu *ProjectUpdate) Where(ps ...predicate.Project) *ProjectUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetReadableID sets the "readable_id" field.
func (pu *ProjectUpdate) SetReadableID(u uuid.UUID) *ProjectUpdate {
	pu.mutation.SetReadableID(u)
	return pu
}

// SetNillableReadableID sets the "readable_id" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableReadableID(u *uuid.UUID) *ProjectUpdate {
	if u != nil {
		pu.SetReadableID(*u)
	}
	return pu
}

// SetName sets the "name" field.
func (pu *ProjectUpdate) SetName(s string) *ProjectUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableName(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetDescription sets the "description" field.
func (pu *ProjectUpdate) SetDescription(s string) *ProjectUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableDescription(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// SetGoal sets the "goal" field.
func (pu *ProjectUpdate) SetGoal(f float64) *ProjectUpdate {
	pu.mutation.ResetGoal()
	pu.mutation.SetGoal(f)
	return pu
}

// SetNillableGoal sets the "goal" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableGoal(f *float64) *ProjectUpdate {
	if f != nil {
		pu.SetGoal(*f)
	}
	return pu
}

// AddGoal adds f to the "goal" field.
func (pu *ProjectUpdate) AddGoal(f float64) *ProjectUpdate {
	pu.mutation.AddGoal(f)
	return pu
}

// SetStartDate sets the "start_date" field.
func (pu *ProjectUpdate) SetStartDate(t time.Time) *ProjectUpdate {
	pu.mutation.SetStartDate(t)
	return pu
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableStartDate(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetStartDate(*t)
	}
	return pu
}

// SetEndDate sets the "end_date" field.
func (pu *ProjectUpdate) SetEndDate(t time.Time) *ProjectUpdate {
	pu.mutation.SetEndDate(t)
	return pu
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableEndDate(t *time.Time) *ProjectUpdate {
	if t != nil {
		pu.SetEndDate(*t)
	}
	return pu
}

// SetCategory sets the "category" field.
func (pu *ProjectUpdate) SetCategory(de dto.CategoryEnum) *ProjectUpdate {
	pu.mutation.SetCategory(de)
	return pu
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableCategory(de *dto.CategoryEnum) *ProjectUpdate {
	if de != nil {
		pu.SetCategory(*de)
	}
	return pu
}

// SetCountryCode sets the "countryCode" field.
func (pu *ProjectUpdate) SetCountryCode(s string) *ProjectUpdate {
	pu.mutation.SetCountryCode(s)
	return pu
}

// SetNillableCountryCode sets the "countryCode" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableCountryCode(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetCountryCode(*s)
	}
	return pu
}

// SetStatus sets the "status" field.
func (pu *ProjectUpdate) SetStatus(de dto.StatusEnum) *ProjectUpdate {
	pu.mutation.SetStatus(de)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableStatus(de *dto.StatusEnum) *ProjectUpdate {
	if de != nil {
		pu.SetStatus(*de)
	}
	return pu
}

// SetOwnerCharityReadableID sets the "owner_charity_readable_id" field.
func (pu *ProjectUpdate) SetOwnerCharityReadableID(s string) *ProjectUpdate {
	pu.mutation.SetOwnerCharityReadableID(s)
	return pu
}

// SetNillableOwnerCharityReadableID sets the "owner_charity_readable_id" field if the given value is not nil.
func (pu *ProjectUpdate) SetNillableOwnerCharityReadableID(s *string) *ProjectUpdate {
	if s != nil {
		pu.SetOwnerCharityReadableID(*s)
	}
	return pu
}

// Mutation returns the ProjectMutation object of the builder.
func (pu *ProjectUpdate) Mutation() *ProjectMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProjectUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProjectUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProjectUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProjectUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProjectUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := project.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Project.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Goal(); ok {
		if err := project.GoalValidator(v); err != nil {
			return &ValidationError{Name: "goal", err: fmt.Errorf(`ent: validator failed for field "Project.goal": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Category(); ok {
		if err := project.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "Project.category": %w`, err)}
		}
	}
	if v, ok := pu.mutation.CountryCode(); ok {
		if err := project.CountryCodeValidator(v); err != nil {
			return &ValidationError{Name: "countryCode", err: fmt.Errorf(`ent: validator failed for field "Project.countryCode": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Status(); ok {
		if err := project.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Project.status": %w`, err)}
		}
	}
	return nil
}

func (pu *ProjectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.ReadableID(); ok {
		_spec.SetField(project.FieldReadableID, field.TypeUUID, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.Goal(); ok {
		_spec.SetField(project.FieldGoal, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.AddedGoal(); ok {
		_spec.AddField(project.FieldGoal, field.TypeFloat64, value)
	}
	if value, ok := pu.mutation.StartDate(); ok {
		_spec.SetField(project.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := pu.mutation.EndDate(); ok {
		_spec.SetField(project.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Category(); ok {
		_spec.SetField(project.FieldCategory, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.CountryCode(); ok {
		_spec.SetField(project.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(project.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.OwnerCharityReadableID(); ok {
		_spec.SetField(project.FieldOwnerCharityReadableID, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProjectUpdateOne is the builder for updating a single Project entity.
type ProjectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProjectMutation
}

// SetReadableID sets the "readable_id" field.
func (puo *ProjectUpdateOne) SetReadableID(u uuid.UUID) *ProjectUpdateOne {
	puo.mutation.SetReadableID(u)
	return puo
}

// SetNillableReadableID sets the "readable_id" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableReadableID(u *uuid.UUID) *ProjectUpdateOne {
	if u != nil {
		puo.SetReadableID(*u)
	}
	return puo
}

// SetName sets the "name" field.
func (puo *ProjectUpdateOne) SetName(s string) *ProjectUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableName(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetDescription sets the "description" field.
func (puo *ProjectUpdateOne) SetDescription(s string) *ProjectUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableDescription(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// SetGoal sets the "goal" field.
func (puo *ProjectUpdateOne) SetGoal(f float64) *ProjectUpdateOne {
	puo.mutation.ResetGoal()
	puo.mutation.SetGoal(f)
	return puo
}

// SetNillableGoal sets the "goal" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableGoal(f *float64) *ProjectUpdateOne {
	if f != nil {
		puo.SetGoal(*f)
	}
	return puo
}

// AddGoal adds f to the "goal" field.
func (puo *ProjectUpdateOne) AddGoal(f float64) *ProjectUpdateOne {
	puo.mutation.AddGoal(f)
	return puo
}

// SetStartDate sets the "start_date" field.
func (puo *ProjectUpdateOne) SetStartDate(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetStartDate(t)
	return puo
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableStartDate(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetStartDate(*t)
	}
	return puo
}

// SetEndDate sets the "end_date" field.
func (puo *ProjectUpdateOne) SetEndDate(t time.Time) *ProjectUpdateOne {
	puo.mutation.SetEndDate(t)
	return puo
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableEndDate(t *time.Time) *ProjectUpdateOne {
	if t != nil {
		puo.SetEndDate(*t)
	}
	return puo
}

// SetCategory sets the "category" field.
func (puo *ProjectUpdateOne) SetCategory(de dto.CategoryEnum) *ProjectUpdateOne {
	puo.mutation.SetCategory(de)
	return puo
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableCategory(de *dto.CategoryEnum) *ProjectUpdateOne {
	if de != nil {
		puo.SetCategory(*de)
	}
	return puo
}

// SetCountryCode sets the "countryCode" field.
func (puo *ProjectUpdateOne) SetCountryCode(s string) *ProjectUpdateOne {
	puo.mutation.SetCountryCode(s)
	return puo
}

// SetNillableCountryCode sets the "countryCode" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableCountryCode(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetCountryCode(*s)
	}
	return puo
}

// SetStatus sets the "status" field.
func (puo *ProjectUpdateOne) SetStatus(de dto.StatusEnum) *ProjectUpdateOne {
	puo.mutation.SetStatus(de)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableStatus(de *dto.StatusEnum) *ProjectUpdateOne {
	if de != nil {
		puo.SetStatus(*de)
	}
	return puo
}

// SetOwnerCharityReadableID sets the "owner_charity_readable_id" field.
func (puo *ProjectUpdateOne) SetOwnerCharityReadableID(s string) *ProjectUpdateOne {
	puo.mutation.SetOwnerCharityReadableID(s)
	return puo
}

// SetNillableOwnerCharityReadableID sets the "owner_charity_readable_id" field if the given value is not nil.
func (puo *ProjectUpdateOne) SetNillableOwnerCharityReadableID(s *string) *ProjectUpdateOne {
	if s != nil {
		puo.SetOwnerCharityReadableID(*s)
	}
	return puo
}

// Mutation returns the ProjectMutation object of the builder.
func (puo *ProjectUpdateOne) Mutation() *ProjectMutation {
	return puo.mutation
}

// Where appends a list predicates to the ProjectUpdate builder.
func (puo *ProjectUpdateOne) Where(ps ...predicate.Project) *ProjectUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProjectUpdateOne) Select(field string, fields ...string) *ProjectUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Project entity.
func (puo *ProjectUpdateOne) Save(ctx context.Context) (*Project, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProjectUpdateOne) SaveX(ctx context.Context) *Project {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProjectUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProjectUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProjectUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := project.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Project.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Goal(); ok {
		if err := project.GoalValidator(v); err != nil {
			return &ValidationError{Name: "goal", err: fmt.Errorf(`ent: validator failed for field "Project.goal": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Category(); ok {
		if err := project.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "Project.category": %w`, err)}
		}
	}
	if v, ok := puo.mutation.CountryCode(); ok {
		if err := project.CountryCodeValidator(v); err != nil {
			return &ValidationError{Name: "countryCode", err: fmt.Errorf(`ent: validator failed for field "Project.countryCode": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Status(); ok {
		if err := project.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Project.status": %w`, err)}
		}
	}
	return nil
}

func (puo *ProjectUpdateOne) sqlSave(ctx context.Context) (_node *Project, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(project.Table, project.Columns, sqlgraph.NewFieldSpec(project.FieldID, field.TypeUUID))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Project.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, project.FieldID)
		for _, f := range fields {
			if !project.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != project.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.ReadableID(); ok {
		_spec.SetField(project.FieldReadableID, field.TypeUUID, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.Goal(); ok {
		_spec.SetField(project.FieldGoal, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.AddedGoal(); ok {
		_spec.AddField(project.FieldGoal, field.TypeFloat64, value)
	}
	if value, ok := puo.mutation.StartDate(); ok {
		_spec.SetField(project.FieldStartDate, field.TypeTime, value)
	}
	if value, ok := puo.mutation.EndDate(); ok {
		_spec.SetField(project.FieldEndDate, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Category(); ok {
		_spec.SetField(project.FieldCategory, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.CountryCode(); ok {
		_spec.SetField(project.FieldCountryCode, field.TypeString, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(project.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.OwnerCharityReadableID(); ok {
		_spec.SetField(project.FieldOwnerCharityReadableID, field.TypeString, value)
	}
	_node = &Project{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{project.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
