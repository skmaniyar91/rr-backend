// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"
	"rr-backend/ent/entgen/tblenum"
	"rr-backend/ent/entgen/tblgarageowner"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TblEnumCreate is the builder for creating a TblEnum entity.
type TblEnumCreate struct {
	config
	mutation *TblEnumMutation
	hooks    []Hook
}

// SetCreatedBy sets the "CreatedBy" field.
func (tec *TblEnumCreate) SetCreatedBy(s string) *TblEnumCreate {
	tec.mutation.SetCreatedBy(s)
	return tec
}

// SetNillableCreatedBy sets the "CreatedBy" field if the given value is not nil.
func (tec *TblEnumCreate) SetNillableCreatedBy(s *string) *TblEnumCreate {
	if s != nil {
		tec.SetCreatedBy(*s)
	}
	return tec
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (tec *TblEnumCreate) SetUpdatedBy(s string) *TblEnumCreate {
	tec.mutation.SetUpdatedBy(s)
	return tec
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (tec *TblEnumCreate) SetNillableUpdatedBy(s *string) *TblEnumCreate {
	if s != nil {
		tec.SetUpdatedBy(*s)
	}
	return tec
}

// SetDeletedBy sets the "DeletedBy" field.
func (tec *TblEnumCreate) SetDeletedBy(s string) *TblEnumCreate {
	tec.mutation.SetDeletedBy(s)
	return tec
}

// SetNillableDeletedBy sets the "DeletedBy" field if the given value is not nil.
func (tec *TblEnumCreate) SetNillableDeletedBy(s *string) *TblEnumCreate {
	if s != nil {
		tec.SetDeletedBy(*s)
	}
	return tec
}

// SetIP sets the "IP" field.
func (tec *TblEnumCreate) SetIP(s string) *TblEnumCreate {
	tec.mutation.SetIP(s)
	return tec
}

// SetNillableIP sets the "IP" field if the given value is not nil.
func (tec *TblEnumCreate) SetNillableIP(s *string) *TblEnumCreate {
	if s != nil {
		tec.SetIP(*s)
	}
	return tec
}

// SetUserAgent sets the "UserAgent" field.
func (tec *TblEnumCreate) SetUserAgent(s string) *TblEnumCreate {
	tec.mutation.SetUserAgent(s)
	return tec
}

// SetNillableUserAgent sets the "UserAgent" field if the given value is not nil.
func (tec *TblEnumCreate) SetNillableUserAgent(s *string) *TblEnumCreate {
	if s != nil {
		tec.SetUserAgent(*s)
	}
	return tec
}

// SetCreatedAt sets the "CreatedAt" field.
func (tec *TblEnumCreate) SetCreatedAt(t time.Time) *TblEnumCreate {
	tec.mutation.SetCreatedAt(t)
	return tec
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (tec *TblEnumCreate) SetNillableCreatedAt(t *time.Time) *TblEnumCreate {
	if t != nil {
		tec.SetCreatedAt(*t)
	}
	return tec
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (tec *TblEnumCreate) SetUpdatedAt(t time.Time) *TblEnumCreate {
	tec.mutation.SetUpdatedAt(t)
	return tec
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (tec *TblEnumCreate) SetNillableUpdatedAt(t *time.Time) *TblEnumCreate {
	if t != nil {
		tec.SetUpdatedAt(*t)
	}
	return tec
}

// SetDeletedAt sets the "DeletedAt" field.
func (tec *TblEnumCreate) SetDeletedAt(t time.Time) *TblEnumCreate {
	tec.mutation.SetDeletedAt(t)
	return tec
}

// SetNillableDeletedAt sets the "DeletedAt" field if the given value is not nil.
func (tec *TblEnumCreate) SetNillableDeletedAt(t *time.Time) *TblEnumCreate {
	if t != nil {
		tec.SetDeletedAt(*t)
	}
	return tec
}

// SetCode sets the "Code" field.
func (tec *TblEnumCreate) SetCode(s string) *TblEnumCreate {
	tec.mutation.SetCode(s)
	return tec
}

// SetCodeType sets the "CodeType" field.
func (tec *TblEnumCreate) SetCodeType(s string) *TblEnumCreate {
	tec.mutation.SetCodeType(s)
	return tec
}

// SetDisplayText sets the "DisplayText" field.
func (tec *TblEnumCreate) SetDisplayText(s string) *TblEnumCreate {
	tec.mutation.SetDisplayText(s)
	return tec
}

// SetID sets the "id" field.
func (tec *TblEnumCreate) SetID(i int) *TblEnumCreate {
	tec.mutation.SetID(i)
	return tec
}

// SetInitialEnumID sets the "InitialEnum" edge to the TblGarageOwner entity by ID.
func (tec *TblEnumCreate) SetInitialEnumID(id string) *TblEnumCreate {
	tec.mutation.SetInitialEnumID(id)
	return tec
}

// SetNillableInitialEnumID sets the "InitialEnum" edge to the TblGarageOwner entity by ID if the given value is not nil.
func (tec *TblEnumCreate) SetNillableInitialEnumID(id *string) *TblEnumCreate {
	if id != nil {
		tec = tec.SetInitialEnumID(*id)
	}
	return tec
}

// SetInitialEnum sets the "InitialEnum" edge to the TblGarageOwner entity.
func (tec *TblEnumCreate) SetInitialEnum(t *TblGarageOwner) *TblEnumCreate {
	return tec.SetInitialEnumID(t.ID)
}

// Mutation returns the TblEnumMutation object of the builder.
func (tec *TblEnumCreate) Mutation() *TblEnumMutation {
	return tec.mutation
}

// Save creates the TblEnum in the database.
func (tec *TblEnumCreate) Save(ctx context.Context) (*TblEnum, error) {
	if err := tec.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tec.sqlSave, tec.mutation, tec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tec *TblEnumCreate) SaveX(ctx context.Context) *TblEnum {
	v, err := tec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tec *TblEnumCreate) Exec(ctx context.Context) error {
	_, err := tec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tec *TblEnumCreate) ExecX(ctx context.Context) {
	if err := tec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tec *TblEnumCreate) defaults() error {
	if _, ok := tec.mutation.CreatedAt(); !ok {
		if tblenum.DefaultCreatedAt == nil {
			return fmt.Errorf("entgen: uninitialized tblenum.DefaultCreatedAt (forgotten import entgen/runtime?)")
		}
		v := tblenum.DefaultCreatedAt()
		tec.mutation.SetCreatedAt(v)
	}
	if _, ok := tec.mutation.UpdatedAt(); !ok {
		if tblenum.DefaultUpdatedAt == nil {
			return fmt.Errorf("entgen: uninitialized tblenum.DefaultUpdatedAt (forgotten import entgen/runtime?)")
		}
		v := tblenum.DefaultUpdatedAt()
		tec.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tec *TblEnumCreate) check() error {
	if v, ok := tec.mutation.CreatedBy(); ok {
		if err := tblenum.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "CreatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblEnum.CreatedBy": %w`, err)}
		}
	}
	if v, ok := tec.mutation.UpdatedBy(); ok {
		if err := tblenum.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "UpdatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblEnum.UpdatedBy": %w`, err)}
		}
	}
	if v, ok := tec.mutation.DeletedBy(); ok {
		if err := tblenum.DeletedByValidator(v); err != nil {
			return &ValidationError{Name: "DeletedBy", err: fmt.Errorf(`entgen: validator failed for field "TblEnum.DeletedBy": %w`, err)}
		}
	}
	if _, ok := tec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`entgen: missing required field "TblEnum.CreatedAt"`)}
	}
	if _, ok := tec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`entgen: missing required field "TblEnum.UpdatedAt"`)}
	}
	if _, ok := tec.mutation.Code(); !ok {
		return &ValidationError{Name: "Code", err: errors.New(`entgen: missing required field "TblEnum.Code"`)}
	}
	if v, ok := tec.mutation.Code(); ok {
		if err := tblenum.CodeValidator(v); err != nil {
			return &ValidationError{Name: "Code", err: fmt.Errorf(`entgen: validator failed for field "TblEnum.Code": %w`, err)}
		}
	}
	if _, ok := tec.mutation.CodeType(); !ok {
		return &ValidationError{Name: "CodeType", err: errors.New(`entgen: missing required field "TblEnum.CodeType"`)}
	}
	if v, ok := tec.mutation.CodeType(); ok {
		if err := tblenum.CodeTypeValidator(v); err != nil {
			return &ValidationError{Name: "CodeType", err: fmt.Errorf(`entgen: validator failed for field "TblEnum.CodeType": %w`, err)}
		}
	}
	if _, ok := tec.mutation.DisplayText(); !ok {
		return &ValidationError{Name: "DisplayText", err: errors.New(`entgen: missing required field "TblEnum.DisplayText"`)}
	}
	if v, ok := tec.mutation.DisplayText(); ok {
		if err := tblenum.DisplayTextValidator(v); err != nil {
			return &ValidationError{Name: "DisplayText", err: fmt.Errorf(`entgen: validator failed for field "TblEnum.DisplayText": %w`, err)}
		}
	}
	return nil
}

func (tec *TblEnumCreate) sqlSave(ctx context.Context) (*TblEnum, error) {
	if err := tec.check(); err != nil {
		return nil, err
	}
	_node, _spec := tec.createSpec()
	if err := sqlgraph.CreateNode(ctx, tec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	tec.mutation.id = &_node.ID
	tec.mutation.done = true
	return _node, nil
}

func (tec *TblEnumCreate) createSpec() (*TblEnum, *sqlgraph.CreateSpec) {
	var (
		_node = &TblEnum{config: tec.config}
		_spec = sqlgraph.NewCreateSpec(tblenum.Table, sqlgraph.NewFieldSpec(tblenum.FieldID, field.TypeInt))
	)
	if id, ok := tec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tec.mutation.CreatedBy(); ok {
		_spec.SetField(tblenum.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = &value
	}
	if value, ok := tec.mutation.UpdatedBy(); ok {
		_spec.SetField(tblenum.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = &value
	}
	if value, ok := tec.mutation.DeletedBy(); ok {
		_spec.SetField(tblenum.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = &value
	}
	if value, ok := tec.mutation.IP(); ok {
		_spec.SetField(tblenum.FieldIP, field.TypeString, value)
		_node.IP = &value
	}
	if value, ok := tec.mutation.UserAgent(); ok {
		_spec.SetField(tblenum.FieldUserAgent, field.TypeString, value)
		_node.UserAgent = &value
	}
	if value, ok := tec.mutation.CreatedAt(); ok {
		_spec.SetField(tblenum.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tec.mutation.UpdatedAt(); ok {
		_spec.SetField(tblenum.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tec.mutation.DeletedAt(); ok {
		_spec.SetField(tblenum.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := tec.mutation.Code(); ok {
		_spec.SetField(tblenum.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := tec.mutation.CodeType(); ok {
		_spec.SetField(tblenum.FieldCodeType, field.TypeString, value)
		_node.CodeType = value
	}
	if value, ok := tec.mutation.DisplayText(); ok {
		_spec.SetField(tblenum.FieldDisplayText, field.TypeString, value)
		_node.DisplayText = value
	}
	if nodes := tec.mutation.InitialEnumIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   tblenum.InitialEnumTable,
			Columns: []string{tblenum.InitialEnumColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tblgarageowner.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TblEnumCreateBulk is the builder for creating many TblEnum entities in bulk.
type TblEnumCreateBulk struct {
	config
	err      error
	builders []*TblEnumCreate
}

// Save creates the TblEnum entities in the database.
func (tecb *TblEnumCreateBulk) Save(ctx context.Context) ([]*TblEnum, error) {
	if tecb.err != nil {
		return nil, tecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tecb.builders))
	nodes := make([]*TblEnum, len(tecb.builders))
	mutators := make([]Mutator, len(tecb.builders))
	for i := range tecb.builders {
		func(i int, root context.Context) {
			builder := tecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TblEnumMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tecb *TblEnumCreateBulk) SaveX(ctx context.Context) []*TblEnum {
	v, err := tecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tecb *TblEnumCreateBulk) Exec(ctx context.Context) error {
	_, err := tecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tecb *TblEnumCreateBulk) ExecX(ctx context.Context) {
	if err := tecb.Exec(ctx); err != nil {
		panic(err)
	}
}
