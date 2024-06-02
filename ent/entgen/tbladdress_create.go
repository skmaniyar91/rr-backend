// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"
	"rr-backend/ent/entgen/tbladdress"
	"rr-backend/ent/entgen/tblgarageowner"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TblAddressCreate is the builder for creating a TblAddress entity.
type TblAddressCreate struct {
	config
	mutation *TblAddressMutation
	hooks    []Hook
}

// SetCreatedBy sets the "CreatedBy" field.
func (tac *TblAddressCreate) SetCreatedBy(s string) *TblAddressCreate {
	tac.mutation.SetCreatedBy(s)
	return tac
}

// SetNillableCreatedBy sets the "CreatedBy" field if the given value is not nil.
func (tac *TblAddressCreate) SetNillableCreatedBy(s *string) *TblAddressCreate {
	if s != nil {
		tac.SetCreatedBy(*s)
	}
	return tac
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (tac *TblAddressCreate) SetUpdatedBy(s string) *TblAddressCreate {
	tac.mutation.SetUpdatedBy(s)
	return tac
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (tac *TblAddressCreate) SetNillableUpdatedBy(s *string) *TblAddressCreate {
	if s != nil {
		tac.SetUpdatedBy(*s)
	}
	return tac
}

// SetDeletedBy sets the "DeletedBy" field.
func (tac *TblAddressCreate) SetDeletedBy(s string) *TblAddressCreate {
	tac.mutation.SetDeletedBy(s)
	return tac
}

// SetNillableDeletedBy sets the "DeletedBy" field if the given value is not nil.
func (tac *TblAddressCreate) SetNillableDeletedBy(s *string) *TblAddressCreate {
	if s != nil {
		tac.SetDeletedBy(*s)
	}
	return tac
}

// SetIP sets the "IP" field.
func (tac *TblAddressCreate) SetIP(s string) *TblAddressCreate {
	tac.mutation.SetIP(s)
	return tac
}

// SetNillableIP sets the "IP" field if the given value is not nil.
func (tac *TblAddressCreate) SetNillableIP(s *string) *TblAddressCreate {
	if s != nil {
		tac.SetIP(*s)
	}
	return tac
}

// SetUserAgent sets the "UserAgent" field.
func (tac *TblAddressCreate) SetUserAgent(s string) *TblAddressCreate {
	tac.mutation.SetUserAgent(s)
	return tac
}

// SetNillableUserAgent sets the "UserAgent" field if the given value is not nil.
func (tac *TblAddressCreate) SetNillableUserAgent(s *string) *TblAddressCreate {
	if s != nil {
		tac.SetUserAgent(*s)
	}
	return tac
}

// SetCreatedAt sets the "CreatedAt" field.
func (tac *TblAddressCreate) SetCreatedAt(t time.Time) *TblAddressCreate {
	tac.mutation.SetCreatedAt(t)
	return tac
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (tac *TblAddressCreate) SetNillableCreatedAt(t *time.Time) *TblAddressCreate {
	if t != nil {
		tac.SetCreatedAt(*t)
	}
	return tac
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (tac *TblAddressCreate) SetUpdatedAt(t time.Time) *TblAddressCreate {
	tac.mutation.SetUpdatedAt(t)
	return tac
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (tac *TblAddressCreate) SetNillableUpdatedAt(t *time.Time) *TblAddressCreate {
	if t != nil {
		tac.SetUpdatedAt(*t)
	}
	return tac
}

// SetDeletedAt sets the "DeletedAt" field.
func (tac *TblAddressCreate) SetDeletedAt(t time.Time) *TblAddressCreate {
	tac.mutation.SetDeletedAt(t)
	return tac
}

// SetNillableDeletedAt sets the "DeletedAt" field if the given value is not nil.
func (tac *TblAddressCreate) SetNillableDeletedAt(t *time.Time) *TblAddressCreate {
	if t != nil {
		tac.SetDeletedAt(*t)
	}
	return tac
}

// SetLine1 sets the "Line1" field.
func (tac *TblAddressCreate) SetLine1(s string) *TblAddressCreate {
	tac.mutation.SetLine1(s)
	return tac
}

// SetLine2 sets the "Line2" field.
func (tac *TblAddressCreate) SetLine2(s string) *TblAddressCreate {
	tac.mutation.SetLine2(s)
	return tac
}

// SetNillableLine2 sets the "Line2" field if the given value is not nil.
func (tac *TblAddressCreate) SetNillableLine2(s *string) *TblAddressCreate {
	if s != nil {
		tac.SetLine2(*s)
	}
	return tac
}

// SetLine3 sets the "Line3" field.
func (tac *TblAddressCreate) SetLine3(s string) *TblAddressCreate {
	tac.mutation.SetLine3(s)
	return tac
}

// SetNillableLine3 sets the "Line3" field if the given value is not nil.
func (tac *TblAddressCreate) SetNillableLine3(s *string) *TblAddressCreate {
	if s != nil {
		tac.SetLine3(*s)
	}
	return tac
}

// SetCity sets the "City" field.
func (tac *TblAddressCreate) SetCity(s string) *TblAddressCreate {
	tac.mutation.SetCity(s)
	return tac
}

// SetDistrict sets the "District" field.
func (tac *TblAddressCreate) SetDistrict(s string) *TblAddressCreate {
	tac.mutation.SetDistrict(s)
	return tac
}

// SetNillableDistrict sets the "District" field if the given value is not nil.
func (tac *TblAddressCreate) SetNillableDistrict(s *string) *TblAddressCreate {
	if s != nil {
		tac.SetDistrict(*s)
	}
	return tac
}

// SetSubDistrict sets the "SubDistrict" field.
func (tac *TblAddressCreate) SetSubDistrict(s string) *TblAddressCreate {
	tac.mutation.SetSubDistrict(s)
	return tac
}

// SetNillableSubDistrict sets the "SubDistrict" field if the given value is not nil.
func (tac *TblAddressCreate) SetNillableSubDistrict(s *string) *TblAddressCreate {
	if s != nil {
		tac.SetSubDistrict(*s)
	}
	return tac
}

// SetState sets the "State" field.
func (tac *TblAddressCreate) SetState(s string) *TblAddressCreate {
	tac.mutation.SetState(s)
	return tac
}

// SetCountry sets the "Country" field.
func (tac *TblAddressCreate) SetCountry(s string) *TblAddressCreate {
	tac.mutation.SetCountry(s)
	return tac
}

// SetPostalCode sets the "PostalCode" field.
func (tac *TblAddressCreate) SetPostalCode(s string) *TblAddressCreate {
	tac.mutation.SetPostalCode(s)
	return tac
}

// SetNillablePostalCode sets the "PostalCode" field if the given value is not nil.
func (tac *TblAddressCreate) SetNillablePostalCode(s *string) *TblAddressCreate {
	if s != nil {
		tac.SetPostalCode(*s)
	}
	return tac
}

// SetID sets the "id" field.
func (tac *TblAddressCreate) SetID(s string) *TblAddressCreate {
	tac.mutation.SetID(s)
	return tac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tac *TblAddressCreate) SetNillableID(s *string) *TblAddressCreate {
	if s != nil {
		tac.SetID(*s)
	}
	return tac
}

// SetOwnerAddressID sets the "OwnerAddress" edge to the TblGarageOwner entity by ID.
func (tac *TblAddressCreate) SetOwnerAddressID(id string) *TblAddressCreate {
	tac.mutation.SetOwnerAddressID(id)
	return tac
}

// SetNillableOwnerAddressID sets the "OwnerAddress" edge to the TblGarageOwner entity by ID if the given value is not nil.
func (tac *TblAddressCreate) SetNillableOwnerAddressID(id *string) *TblAddressCreate {
	if id != nil {
		tac = tac.SetOwnerAddressID(*id)
	}
	return tac
}

// SetOwnerAddress sets the "OwnerAddress" edge to the TblGarageOwner entity.
func (tac *TblAddressCreate) SetOwnerAddress(t *TblGarageOwner) *TblAddressCreate {
	return tac.SetOwnerAddressID(t.ID)
}

// Mutation returns the TblAddressMutation object of the builder.
func (tac *TblAddressCreate) Mutation() *TblAddressMutation {
	return tac.mutation
}

// Save creates the TblAddress in the database.
func (tac *TblAddressCreate) Save(ctx context.Context) (*TblAddress, error) {
	if err := tac.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tac.sqlSave, tac.mutation, tac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tac *TblAddressCreate) SaveX(ctx context.Context) *TblAddress {
	v, err := tac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tac *TblAddressCreate) Exec(ctx context.Context) error {
	_, err := tac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tac *TblAddressCreate) ExecX(ctx context.Context) {
	if err := tac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tac *TblAddressCreate) defaults() error {
	if _, ok := tac.mutation.CreatedAt(); !ok {
		if tbladdress.DefaultCreatedAt == nil {
			return fmt.Errorf("entgen: uninitialized tbladdress.DefaultCreatedAt (forgotten import entgen/runtime?)")
		}
		v := tbladdress.DefaultCreatedAt()
		tac.mutation.SetCreatedAt(v)
	}
	if _, ok := tac.mutation.UpdatedAt(); !ok {
		if tbladdress.DefaultUpdatedAt == nil {
			return fmt.Errorf("entgen: uninitialized tbladdress.DefaultUpdatedAt (forgotten import entgen/runtime?)")
		}
		v := tbladdress.DefaultUpdatedAt()
		tac.mutation.SetUpdatedAt(v)
	}
	if _, ok := tac.mutation.ID(); !ok {
		if tbladdress.DefaultID == nil {
			return fmt.Errorf("entgen: uninitialized tbladdress.DefaultID (forgotten import entgen/runtime?)")
		}
		v := tbladdress.DefaultID()
		tac.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tac *TblAddressCreate) check() error {
	if v, ok := tac.mutation.CreatedBy(); ok {
		if err := tbladdress.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "CreatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblAddress.CreatedBy": %w`, err)}
		}
	}
	if v, ok := tac.mutation.UpdatedBy(); ok {
		if err := tbladdress.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "UpdatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblAddress.UpdatedBy": %w`, err)}
		}
	}
	if v, ok := tac.mutation.DeletedBy(); ok {
		if err := tbladdress.DeletedByValidator(v); err != nil {
			return &ValidationError{Name: "DeletedBy", err: fmt.Errorf(`entgen: validator failed for field "TblAddress.DeletedBy": %w`, err)}
		}
	}
	if _, ok := tac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`entgen: missing required field "TblAddress.CreatedAt"`)}
	}
	if _, ok := tac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`entgen: missing required field "TblAddress.UpdatedAt"`)}
	}
	if _, ok := tac.mutation.Line1(); !ok {
		return &ValidationError{Name: "Line1", err: errors.New(`entgen: missing required field "TblAddress.Line1"`)}
	}
	if v, ok := tac.mutation.Line1(); ok {
		if err := tbladdress.Line1Validator(v); err != nil {
			return &ValidationError{Name: "Line1", err: fmt.Errorf(`entgen: validator failed for field "TblAddress.Line1": %w`, err)}
		}
	}
	if v, ok := tac.mutation.Line2(); ok {
		if err := tbladdress.Line2Validator(v); err != nil {
			return &ValidationError{Name: "Line2", err: fmt.Errorf(`entgen: validator failed for field "TblAddress.Line2": %w`, err)}
		}
	}
	if v, ok := tac.mutation.Line3(); ok {
		if err := tbladdress.Line3Validator(v); err != nil {
			return &ValidationError{Name: "Line3", err: fmt.Errorf(`entgen: validator failed for field "TblAddress.Line3": %w`, err)}
		}
	}
	if _, ok := tac.mutation.City(); !ok {
		return &ValidationError{Name: "City", err: errors.New(`entgen: missing required field "TblAddress.City"`)}
	}
	if v, ok := tac.mutation.City(); ok {
		if err := tbladdress.CityValidator(v); err != nil {
			return &ValidationError{Name: "City", err: fmt.Errorf(`entgen: validator failed for field "TblAddress.City": %w`, err)}
		}
	}
	if v, ok := tac.mutation.District(); ok {
		if err := tbladdress.DistrictValidator(v); err != nil {
			return &ValidationError{Name: "District", err: fmt.Errorf(`entgen: validator failed for field "TblAddress.District": %w`, err)}
		}
	}
	if v, ok := tac.mutation.SubDistrict(); ok {
		if err := tbladdress.SubDistrictValidator(v); err != nil {
			return &ValidationError{Name: "SubDistrict", err: fmt.Errorf(`entgen: validator failed for field "TblAddress.SubDistrict": %w`, err)}
		}
	}
	if _, ok := tac.mutation.State(); !ok {
		return &ValidationError{Name: "State", err: errors.New(`entgen: missing required field "TblAddress.State"`)}
	}
	if v, ok := tac.mutation.State(); ok {
		if err := tbladdress.StateValidator(v); err != nil {
			return &ValidationError{Name: "State", err: fmt.Errorf(`entgen: validator failed for field "TblAddress.State": %w`, err)}
		}
	}
	if _, ok := tac.mutation.Country(); !ok {
		return &ValidationError{Name: "Country", err: errors.New(`entgen: missing required field "TblAddress.Country"`)}
	}
	if v, ok := tac.mutation.Country(); ok {
		if err := tbladdress.CountryValidator(v); err != nil {
			return &ValidationError{Name: "Country", err: fmt.Errorf(`entgen: validator failed for field "TblAddress.Country": %w`, err)}
		}
	}
	if v, ok := tac.mutation.PostalCode(); ok {
		if err := tbladdress.PostalCodeValidator(v); err != nil {
			return &ValidationError{Name: "PostalCode", err: fmt.Errorf(`entgen: validator failed for field "TblAddress.PostalCode": %w`, err)}
		}
	}
	if v, ok := tac.mutation.ID(); ok {
		if err := tbladdress.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`entgen: validator failed for field "TblAddress.id": %w`, err)}
		}
	}
	return nil
}

func (tac *TblAddressCreate) sqlSave(ctx context.Context) (*TblAddress, error) {
	if err := tac.check(); err != nil {
		return nil, err
	}
	_node, _spec := tac.createSpec()
	if err := sqlgraph.CreateNode(ctx, tac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected TblAddress.ID type: %T", _spec.ID.Value)
		}
	}
	tac.mutation.id = &_node.ID
	tac.mutation.done = true
	return _node, nil
}

func (tac *TblAddressCreate) createSpec() (*TblAddress, *sqlgraph.CreateSpec) {
	var (
		_node = &TblAddress{config: tac.config}
		_spec = sqlgraph.NewCreateSpec(tbladdress.Table, sqlgraph.NewFieldSpec(tbladdress.FieldID, field.TypeString))
	)
	if id, ok := tac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tac.mutation.CreatedBy(); ok {
		_spec.SetField(tbladdress.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = &value
	}
	if value, ok := tac.mutation.UpdatedBy(); ok {
		_spec.SetField(tbladdress.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = &value
	}
	if value, ok := tac.mutation.DeletedBy(); ok {
		_spec.SetField(tbladdress.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = &value
	}
	if value, ok := tac.mutation.IP(); ok {
		_spec.SetField(tbladdress.FieldIP, field.TypeString, value)
		_node.IP = &value
	}
	if value, ok := tac.mutation.UserAgent(); ok {
		_spec.SetField(tbladdress.FieldUserAgent, field.TypeString, value)
		_node.UserAgent = &value
	}
	if value, ok := tac.mutation.CreatedAt(); ok {
		_spec.SetField(tbladdress.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tac.mutation.UpdatedAt(); ok {
		_spec.SetField(tbladdress.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tac.mutation.DeletedAt(); ok {
		_spec.SetField(tbladdress.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := tac.mutation.Line1(); ok {
		_spec.SetField(tbladdress.FieldLine1, field.TypeString, value)
		_node.Line1 = value
	}
	if value, ok := tac.mutation.Line2(); ok {
		_spec.SetField(tbladdress.FieldLine2, field.TypeString, value)
		_node.Line2 = &value
	}
	if value, ok := tac.mutation.Line3(); ok {
		_spec.SetField(tbladdress.FieldLine3, field.TypeString, value)
		_node.Line3 = &value
	}
	if value, ok := tac.mutation.City(); ok {
		_spec.SetField(tbladdress.FieldCity, field.TypeString, value)
		_node.City = value
	}
	if value, ok := tac.mutation.District(); ok {
		_spec.SetField(tbladdress.FieldDistrict, field.TypeString, value)
		_node.District = &value
	}
	if value, ok := tac.mutation.SubDistrict(); ok {
		_spec.SetField(tbladdress.FieldSubDistrict, field.TypeString, value)
		_node.SubDistrict = &value
	}
	if value, ok := tac.mutation.State(); ok {
		_spec.SetField(tbladdress.FieldState, field.TypeString, value)
		_node.State = value
	}
	if value, ok := tac.mutation.Country(); ok {
		_spec.SetField(tbladdress.FieldCountry, field.TypeString, value)
		_node.Country = value
	}
	if value, ok := tac.mutation.PostalCode(); ok {
		_spec.SetField(tbladdress.FieldPostalCode, field.TypeString, value)
		_node.PostalCode = &value
	}
	if nodes := tac.mutation.OwnerAddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   tbladdress.OwnerAddressTable,
			Columns: []string{tbladdress.OwnerAddressColumn},
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

// TblAddressCreateBulk is the builder for creating many TblAddress entities in bulk.
type TblAddressCreateBulk struct {
	config
	err      error
	builders []*TblAddressCreate
}

// Save creates the TblAddress entities in the database.
func (tacb *TblAddressCreateBulk) Save(ctx context.Context) ([]*TblAddress, error) {
	if tacb.err != nil {
		return nil, tacb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tacb.builders))
	nodes := make([]*TblAddress, len(tacb.builders))
	mutators := make([]Mutator, len(tacb.builders))
	for i := range tacb.builders {
		func(i int, root context.Context) {
			builder := tacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TblAddressMutation)
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
					_, err = mutators[i+1].Mutate(root, tacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tacb *TblAddressCreateBulk) SaveX(ctx context.Context) []*TblAddress {
	v, err := tacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tacb *TblAddressCreateBulk) Exec(ctx context.Context) error {
	_, err := tacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tacb *TblAddressCreateBulk) ExecX(ctx context.Context) {
	if err := tacb.Exec(ctx); err != nil {
		panic(err)
	}
}