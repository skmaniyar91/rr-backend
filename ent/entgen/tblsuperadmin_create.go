// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"
	"rr-backend/ent/entgen/tblsuperadmin"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TblSuperAdminCreate is the builder for creating a TblSuperAdmin entity.
type TblSuperAdminCreate struct {
	config
	mutation *TblSuperAdminMutation
	hooks    []Hook
}

// SetCreatedBy sets the "CreatedBy" field.
func (tsac *TblSuperAdminCreate) SetCreatedBy(s string) *TblSuperAdminCreate {
	tsac.mutation.SetCreatedBy(s)
	return tsac
}

// SetNillableCreatedBy sets the "CreatedBy" field if the given value is not nil.
func (tsac *TblSuperAdminCreate) SetNillableCreatedBy(s *string) *TblSuperAdminCreate {
	if s != nil {
		tsac.SetCreatedBy(*s)
	}
	return tsac
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (tsac *TblSuperAdminCreate) SetUpdatedBy(s string) *TblSuperAdminCreate {
	tsac.mutation.SetUpdatedBy(s)
	return tsac
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (tsac *TblSuperAdminCreate) SetNillableUpdatedBy(s *string) *TblSuperAdminCreate {
	if s != nil {
		tsac.SetUpdatedBy(*s)
	}
	return tsac
}

// SetDeletedBy sets the "DeletedBy" field.
func (tsac *TblSuperAdminCreate) SetDeletedBy(s string) *TblSuperAdminCreate {
	tsac.mutation.SetDeletedBy(s)
	return tsac
}

// SetNillableDeletedBy sets the "DeletedBy" field if the given value is not nil.
func (tsac *TblSuperAdminCreate) SetNillableDeletedBy(s *string) *TblSuperAdminCreate {
	if s != nil {
		tsac.SetDeletedBy(*s)
	}
	return tsac
}

// SetIP sets the "Ip" field.
func (tsac *TblSuperAdminCreate) SetIP(s string) *TblSuperAdminCreate {
	tsac.mutation.SetIP(s)
	return tsac
}

// SetNillableIP sets the "Ip" field if the given value is not nil.
func (tsac *TblSuperAdminCreate) SetNillableIP(s *string) *TblSuperAdminCreate {
	if s != nil {
		tsac.SetIP(*s)
	}
	return tsac
}

// SetUserAgent sets the "UserAgent" field.
func (tsac *TblSuperAdminCreate) SetUserAgent(s string) *TblSuperAdminCreate {
	tsac.mutation.SetUserAgent(s)
	return tsac
}

// SetNillableUserAgent sets the "UserAgent" field if the given value is not nil.
func (tsac *TblSuperAdminCreate) SetNillableUserAgent(s *string) *TblSuperAdminCreate {
	if s != nil {
		tsac.SetUserAgent(*s)
	}
	return tsac
}

// SetCreatedAt sets the "CreatedAt" field.
func (tsac *TblSuperAdminCreate) SetCreatedAt(t time.Time) *TblSuperAdminCreate {
	tsac.mutation.SetCreatedAt(t)
	return tsac
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (tsac *TblSuperAdminCreate) SetNillableCreatedAt(t *time.Time) *TblSuperAdminCreate {
	if t != nil {
		tsac.SetCreatedAt(*t)
	}
	return tsac
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (tsac *TblSuperAdminCreate) SetUpdatedAt(t time.Time) *TblSuperAdminCreate {
	tsac.mutation.SetUpdatedAt(t)
	return tsac
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (tsac *TblSuperAdminCreate) SetNillableUpdatedAt(t *time.Time) *TblSuperAdminCreate {
	if t != nil {
		tsac.SetUpdatedAt(*t)
	}
	return tsac
}

// SetDeletedAt sets the "DeletedAt" field.
func (tsac *TblSuperAdminCreate) SetDeletedAt(t time.Time) *TblSuperAdminCreate {
	tsac.mutation.SetDeletedAt(t)
	return tsac
}

// SetNillableDeletedAt sets the "DeletedAt" field if the given value is not nil.
func (tsac *TblSuperAdminCreate) SetNillableDeletedAt(t *time.Time) *TblSuperAdminCreate {
	if t != nil {
		tsac.SetDeletedAt(*t)
	}
	return tsac
}

// SetUserName sets the "UserName" field.
func (tsac *TblSuperAdminCreate) SetUserName(s string) *TblSuperAdminCreate {
	tsac.mutation.SetUserName(s)
	return tsac
}

// SetPassWord sets the "PassWord" field.
func (tsac *TblSuperAdminCreate) SetPassWord(s string) *TblSuperAdminCreate {
	tsac.mutation.SetPassWord(s)
	return tsac
}

// SetID sets the "id" field.
func (tsac *TblSuperAdminCreate) SetID(s string) *TblSuperAdminCreate {
	tsac.mutation.SetID(s)
	return tsac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tsac *TblSuperAdminCreate) SetNillableID(s *string) *TblSuperAdminCreate {
	if s != nil {
		tsac.SetID(*s)
	}
	return tsac
}

// Mutation returns the TblSuperAdminMutation object of the builder.
func (tsac *TblSuperAdminCreate) Mutation() *TblSuperAdminMutation {
	return tsac.mutation
}

// Save creates the TblSuperAdmin in the database.
func (tsac *TblSuperAdminCreate) Save(ctx context.Context) (*TblSuperAdmin, error) {
	if err := tsac.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tsac.sqlSave, tsac.mutation, tsac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tsac *TblSuperAdminCreate) SaveX(ctx context.Context) *TblSuperAdmin {
	v, err := tsac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tsac *TblSuperAdminCreate) Exec(ctx context.Context) error {
	_, err := tsac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsac *TblSuperAdminCreate) ExecX(ctx context.Context) {
	if err := tsac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tsac *TblSuperAdminCreate) defaults() error {
	if _, ok := tsac.mutation.CreatedAt(); !ok {
		if tblsuperadmin.DefaultCreatedAt == nil {
			return fmt.Errorf("entgen: uninitialized tblsuperadmin.DefaultCreatedAt (forgotten import entgen/runtime?)")
		}
		v := tblsuperadmin.DefaultCreatedAt()
		tsac.mutation.SetCreatedAt(v)
	}
	if _, ok := tsac.mutation.UpdatedAt(); !ok {
		v := tblsuperadmin.DefaultUpdatedAt
		tsac.mutation.SetUpdatedAt(v)
	}
	if _, ok := tsac.mutation.ID(); !ok {
		if tblsuperadmin.DefaultID == nil {
			return fmt.Errorf("entgen: uninitialized tblsuperadmin.DefaultID (forgotten import entgen/runtime?)")
		}
		v := tblsuperadmin.DefaultID()
		tsac.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tsac *TblSuperAdminCreate) check() error {
	if v, ok := tsac.mutation.CreatedBy(); ok {
		if err := tblsuperadmin.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "CreatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.CreatedBy": %w`, err)}
		}
	}
	if v, ok := tsac.mutation.UpdatedBy(); ok {
		if err := tblsuperadmin.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "UpdatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.UpdatedBy": %w`, err)}
		}
	}
	if v, ok := tsac.mutation.DeletedBy(); ok {
		if err := tblsuperadmin.DeletedByValidator(v); err != nil {
			return &ValidationError{Name: "DeletedBy", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.DeletedBy": %w`, err)}
		}
	}
	if _, ok := tsac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`entgen: missing required field "TblSuperAdmin.CreatedAt"`)}
	}
	if _, ok := tsac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`entgen: missing required field "TblSuperAdmin.UpdatedAt"`)}
	}
	if _, ok := tsac.mutation.UserName(); !ok {
		return &ValidationError{Name: "UserName", err: errors.New(`entgen: missing required field "TblSuperAdmin.UserName"`)}
	}
	if v, ok := tsac.mutation.UserName(); ok {
		if err := tblsuperadmin.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "UserName", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.UserName": %w`, err)}
		}
	}
	if _, ok := tsac.mutation.PassWord(); !ok {
		return &ValidationError{Name: "PassWord", err: errors.New(`entgen: missing required field "TblSuperAdmin.PassWord"`)}
	}
	if v, ok := tsac.mutation.PassWord(); ok {
		if err := tblsuperadmin.PassWordValidator(v); err != nil {
			return &ValidationError{Name: "PassWord", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.PassWord": %w`, err)}
		}
	}
	if v, ok := tsac.mutation.ID(); ok {
		if err := tblsuperadmin.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.id": %w`, err)}
		}
	}
	return nil
}

func (tsac *TblSuperAdminCreate) sqlSave(ctx context.Context) (*TblSuperAdmin, error) {
	if err := tsac.check(); err != nil {
		return nil, err
	}
	_node, _spec := tsac.createSpec()
	if err := sqlgraph.CreateNode(ctx, tsac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected TblSuperAdmin.ID type: %T", _spec.ID.Value)
		}
	}
	tsac.mutation.id = &_node.ID
	tsac.mutation.done = true
	return _node, nil
}

func (tsac *TblSuperAdminCreate) createSpec() (*TblSuperAdmin, *sqlgraph.CreateSpec) {
	var (
		_node = &TblSuperAdmin{config: tsac.config}
		_spec = sqlgraph.NewCreateSpec(tblsuperadmin.Table, sqlgraph.NewFieldSpec(tblsuperadmin.FieldID, field.TypeString))
	)
	if id, ok := tsac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tsac.mutation.CreatedBy(); ok {
		_spec.SetField(tblsuperadmin.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = &value
	}
	if value, ok := tsac.mutation.UpdatedBy(); ok {
		_spec.SetField(tblsuperadmin.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = &value
	}
	if value, ok := tsac.mutation.DeletedBy(); ok {
		_spec.SetField(tblsuperadmin.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = &value
	}
	if value, ok := tsac.mutation.IP(); ok {
		_spec.SetField(tblsuperadmin.FieldIP, field.TypeString, value)
		_node.IP = &value
	}
	if value, ok := tsac.mutation.UserAgent(); ok {
		_spec.SetField(tblsuperadmin.FieldUserAgent, field.TypeString, value)
		_node.UserAgent = &value
	}
	if value, ok := tsac.mutation.CreatedAt(); ok {
		_spec.SetField(tblsuperadmin.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tsac.mutation.UpdatedAt(); ok {
		_spec.SetField(tblsuperadmin.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tsac.mutation.DeletedAt(); ok {
		_spec.SetField(tblsuperadmin.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := tsac.mutation.UserName(); ok {
		_spec.SetField(tblsuperadmin.FieldUserName, field.TypeString, value)
		_node.UserName = value
	}
	if value, ok := tsac.mutation.PassWord(); ok {
		_spec.SetField(tblsuperadmin.FieldPassWord, field.TypeString, value)
		_node.PassWord = value
	}
	return _node, _spec
}

// TblSuperAdminCreateBulk is the builder for creating many TblSuperAdmin entities in bulk.
type TblSuperAdminCreateBulk struct {
	config
	err      error
	builders []*TblSuperAdminCreate
}

// Save creates the TblSuperAdmin entities in the database.
func (tsacb *TblSuperAdminCreateBulk) Save(ctx context.Context) ([]*TblSuperAdmin, error) {
	if tsacb.err != nil {
		return nil, tsacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tsacb.builders))
	nodes := make([]*TblSuperAdmin, len(tsacb.builders))
	mutators := make([]Mutator, len(tsacb.builders))
	for i := range tsacb.builders {
		func(i int, root context.Context) {
			builder := tsacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TblSuperAdminMutation)
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
					_, err = mutators[i+1].Mutate(root, tsacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tsacb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, tsacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tsacb *TblSuperAdminCreateBulk) SaveX(ctx context.Context) []*TblSuperAdmin {
	v, err := tsacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tsacb *TblSuperAdminCreateBulk) Exec(ctx context.Context) error {
	_, err := tsacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsacb *TblSuperAdminCreateBulk) ExecX(ctx context.Context) {
	if err := tsacb.Exec(ctx); err != nil {
		panic(err)
	}
}