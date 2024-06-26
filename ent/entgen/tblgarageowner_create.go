// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"
	"rr-backend/ent/entgen/tbladdress"
	"rr-backend/ent/entgen/tbldocument"
	"rr-backend/ent/entgen/tblenum"
	"rr-backend/ent/entgen/tblgarageowner"
	"rr-backend/ent/entgen/tblusers"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TblGarageOwnerCreate is the builder for creating a TblGarageOwner entity.
type TblGarageOwnerCreate struct {
	config
	mutation *TblGarageOwnerMutation
	hooks    []Hook
}

// SetCreatedBy sets the "CreatedBy" field.
func (tgoc *TblGarageOwnerCreate) SetCreatedBy(s string) *TblGarageOwnerCreate {
	tgoc.mutation.SetCreatedBy(s)
	return tgoc
}

// SetNillableCreatedBy sets the "CreatedBy" field if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableCreatedBy(s *string) *TblGarageOwnerCreate {
	if s != nil {
		tgoc.SetCreatedBy(*s)
	}
	return tgoc
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (tgoc *TblGarageOwnerCreate) SetUpdatedBy(s string) *TblGarageOwnerCreate {
	tgoc.mutation.SetUpdatedBy(s)
	return tgoc
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableUpdatedBy(s *string) *TblGarageOwnerCreate {
	if s != nil {
		tgoc.SetUpdatedBy(*s)
	}
	return tgoc
}

// SetDeletedBy sets the "DeletedBy" field.
func (tgoc *TblGarageOwnerCreate) SetDeletedBy(s string) *TblGarageOwnerCreate {
	tgoc.mutation.SetDeletedBy(s)
	return tgoc
}

// SetNillableDeletedBy sets the "DeletedBy" field if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableDeletedBy(s *string) *TblGarageOwnerCreate {
	if s != nil {
		tgoc.SetDeletedBy(*s)
	}
	return tgoc
}

// SetIP sets the "IP" field.
func (tgoc *TblGarageOwnerCreate) SetIP(s string) *TblGarageOwnerCreate {
	tgoc.mutation.SetIP(s)
	return tgoc
}

// SetNillableIP sets the "IP" field if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableIP(s *string) *TblGarageOwnerCreate {
	if s != nil {
		tgoc.SetIP(*s)
	}
	return tgoc
}

// SetUserAgent sets the "UserAgent" field.
func (tgoc *TblGarageOwnerCreate) SetUserAgent(s string) *TblGarageOwnerCreate {
	tgoc.mutation.SetUserAgent(s)
	return tgoc
}

// SetNillableUserAgent sets the "UserAgent" field if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableUserAgent(s *string) *TblGarageOwnerCreate {
	if s != nil {
		tgoc.SetUserAgent(*s)
	}
	return tgoc
}

// SetCreatedAt sets the "CreatedAt" field.
func (tgoc *TblGarageOwnerCreate) SetCreatedAt(t time.Time) *TblGarageOwnerCreate {
	tgoc.mutation.SetCreatedAt(t)
	return tgoc
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableCreatedAt(t *time.Time) *TblGarageOwnerCreate {
	if t != nil {
		tgoc.SetCreatedAt(*t)
	}
	return tgoc
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (tgoc *TblGarageOwnerCreate) SetUpdatedAt(t time.Time) *TblGarageOwnerCreate {
	tgoc.mutation.SetUpdatedAt(t)
	return tgoc
}

// SetNillableUpdatedAt sets the "UpdatedAt" field if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableUpdatedAt(t *time.Time) *TblGarageOwnerCreate {
	if t != nil {
		tgoc.SetUpdatedAt(*t)
	}
	return tgoc
}

// SetDeletedAt sets the "DeletedAt" field.
func (tgoc *TblGarageOwnerCreate) SetDeletedAt(t time.Time) *TblGarageOwnerCreate {
	tgoc.mutation.SetDeletedAt(t)
	return tgoc
}

// SetNillableDeletedAt sets the "DeletedAt" field if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableDeletedAt(t *time.Time) *TblGarageOwnerCreate {
	if t != nil {
		tgoc.SetDeletedAt(*t)
	}
	return tgoc
}

// SetUserIdUlid sets the "UserId_ulid" field.
func (tgoc *TblGarageOwnerCreate) SetUserIdUlid(s string) *TblGarageOwnerCreate {
	tgoc.mutation.SetUserIdUlid(s)
	return tgoc
}

// SetNillableUserIdUlid sets the "UserId_ulid" field if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableUserIdUlid(s *string) *TblGarageOwnerCreate {
	if s != nil {
		tgoc.SetUserIdUlid(*s)
	}
	return tgoc
}

// SetInitial sets the "Initial" field.
func (tgoc *TblGarageOwnerCreate) SetInitial(i int) *TblGarageOwnerCreate {
	tgoc.mutation.SetInitial(i)
	return tgoc
}

// SetNillableInitial sets the "Initial" field if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableInitial(i *int) *TblGarageOwnerCreate {
	if i != nil {
		tgoc.SetInitial(*i)
	}
	return tgoc
}

// SetFirstName sets the "FirstName" field.
func (tgoc *TblGarageOwnerCreate) SetFirstName(s string) *TblGarageOwnerCreate {
	tgoc.mutation.SetFirstName(s)
	return tgoc
}

// SetMiddleName sets the "MiddleName" field.
func (tgoc *TblGarageOwnerCreate) SetMiddleName(s string) *TblGarageOwnerCreate {
	tgoc.mutation.SetMiddleName(s)
	return tgoc
}

// SetNillableMiddleName sets the "MiddleName" field if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableMiddleName(s *string) *TblGarageOwnerCreate {
	if s != nil {
		tgoc.SetMiddleName(*s)
	}
	return tgoc
}

// SetLastName sets the "LastName" field.
func (tgoc *TblGarageOwnerCreate) SetLastName(s string) *TblGarageOwnerCreate {
	tgoc.mutation.SetLastName(s)
	return tgoc
}

// SetContactNumber sets the "ContactNumber" field.
func (tgoc *TblGarageOwnerCreate) SetContactNumber(s string) *TblGarageOwnerCreate {
	tgoc.mutation.SetContactNumber(s)
	return tgoc
}

// SetEmail sets the "Email" field.
func (tgoc *TblGarageOwnerCreate) SetEmail(s string) *TblGarageOwnerCreate {
	tgoc.mutation.SetEmail(s)
	return tgoc
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableEmail(s *string) *TblGarageOwnerCreate {
	if s != nil {
		tgoc.SetEmail(*s)
	}
	return tgoc
}

// SetAge sets the "Age" field.
func (tgoc *TblGarageOwnerCreate) SetAge(i int) *TblGarageOwnerCreate {
	tgoc.mutation.SetAge(i)
	return tgoc
}

// SetNillableAge sets the "Age" field if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableAge(i *int) *TblGarageOwnerCreate {
	if i != nil {
		tgoc.SetAge(*i)
	}
	return tgoc
}

// SetPhotoIdUlid sets the "PhotoId_ulid" field.
func (tgoc *TblGarageOwnerCreate) SetPhotoIdUlid(s string) *TblGarageOwnerCreate {
	tgoc.mutation.SetPhotoIdUlid(s)
	return tgoc
}

// SetNillablePhotoIdUlid sets the "PhotoId_ulid" field if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillablePhotoIdUlid(s *string) *TblGarageOwnerCreate {
	if s != nil {
		tgoc.SetPhotoIdUlid(*s)
	}
	return tgoc
}

// SetAddressIdUlid sets the "AddressId_ulid" field.
func (tgoc *TblGarageOwnerCreate) SetAddressIdUlid(s string) *TblGarageOwnerCreate {
	tgoc.mutation.SetAddressIdUlid(s)
	return tgoc
}

// SetID sets the "id" field.
func (tgoc *TblGarageOwnerCreate) SetID(s string) *TblGarageOwnerCreate {
	tgoc.mutation.SetID(s)
	return tgoc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableID(s *string) *TblGarageOwnerCreate {
	if s != nil {
		tgoc.SetID(*s)
	}
	return tgoc
}

// SetUserID sets the "User" edge to the TblUSers entity by ID.
func (tgoc *TblGarageOwnerCreate) SetUserID(id string) *TblGarageOwnerCreate {
	tgoc.mutation.SetUserID(id)
	return tgoc
}

// SetNillableUserID sets the "User" edge to the TblUSers entity by ID if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableUserID(id *string) *TblGarageOwnerCreate {
	if id != nil {
		tgoc = tgoc.SetUserID(*id)
	}
	return tgoc
}

// SetUser sets the "User" edge to the TblUSers entity.
func (tgoc *TblGarageOwnerCreate) SetUser(t *TblUSers) *TblGarageOwnerCreate {
	return tgoc.SetUserID(t.ID)
}

// SetNameInitialID sets the "NameInitial" edge to the TblEnum entity by ID.
func (tgoc *TblGarageOwnerCreate) SetNameInitialID(id int) *TblGarageOwnerCreate {
	tgoc.mutation.SetNameInitialID(id)
	return tgoc
}

// SetNillableNameInitialID sets the "NameInitial" edge to the TblEnum entity by ID if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableNameInitialID(id *int) *TblGarageOwnerCreate {
	if id != nil {
		tgoc = tgoc.SetNameInitialID(*id)
	}
	return tgoc
}

// SetNameInitial sets the "NameInitial" edge to the TblEnum entity.
func (tgoc *TblGarageOwnerCreate) SetNameInitial(t *TblEnum) *TblGarageOwnerCreate {
	return tgoc.SetNameInitialID(t.ID)
}

// SetOwnerPhotoID sets the "OwnerPhoto" edge to the TblDocument entity by ID.
func (tgoc *TblGarageOwnerCreate) SetOwnerPhotoID(id string) *TblGarageOwnerCreate {
	tgoc.mutation.SetOwnerPhotoID(id)
	return tgoc
}

// SetNillableOwnerPhotoID sets the "OwnerPhoto" edge to the TblDocument entity by ID if the given value is not nil.
func (tgoc *TblGarageOwnerCreate) SetNillableOwnerPhotoID(id *string) *TblGarageOwnerCreate {
	if id != nil {
		tgoc = tgoc.SetOwnerPhotoID(*id)
	}
	return tgoc
}

// SetOwnerPhoto sets the "OwnerPhoto" edge to the TblDocument entity.
func (tgoc *TblGarageOwnerCreate) SetOwnerPhoto(t *TblDocument) *TblGarageOwnerCreate {
	return tgoc.SetOwnerPhotoID(t.ID)
}

// SetAddressID sets the "Address" edge to the TblAddress entity by ID.
func (tgoc *TblGarageOwnerCreate) SetAddressID(id string) *TblGarageOwnerCreate {
	tgoc.mutation.SetAddressID(id)
	return tgoc
}

// SetAddress sets the "Address" edge to the TblAddress entity.
func (tgoc *TblGarageOwnerCreate) SetAddress(t *TblAddress) *TblGarageOwnerCreate {
	return tgoc.SetAddressID(t.ID)
}

// Mutation returns the TblGarageOwnerMutation object of the builder.
func (tgoc *TblGarageOwnerCreate) Mutation() *TblGarageOwnerMutation {
	return tgoc.mutation
}

// Save creates the TblGarageOwner in the database.
func (tgoc *TblGarageOwnerCreate) Save(ctx context.Context) (*TblGarageOwner, error) {
	if err := tgoc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tgoc.sqlSave, tgoc.mutation, tgoc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tgoc *TblGarageOwnerCreate) SaveX(ctx context.Context) *TblGarageOwner {
	v, err := tgoc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tgoc *TblGarageOwnerCreate) Exec(ctx context.Context) error {
	_, err := tgoc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tgoc *TblGarageOwnerCreate) ExecX(ctx context.Context) {
	if err := tgoc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tgoc *TblGarageOwnerCreate) defaults() error {
	if _, ok := tgoc.mutation.CreatedAt(); !ok {
		if tblgarageowner.DefaultCreatedAt == nil {
			return fmt.Errorf("entgen: uninitialized tblgarageowner.DefaultCreatedAt (forgotten import entgen/runtime?)")
		}
		v := tblgarageowner.DefaultCreatedAt()
		tgoc.mutation.SetCreatedAt(v)
	}
	if _, ok := tgoc.mutation.UpdatedAt(); !ok {
		if tblgarageowner.DefaultUpdatedAt == nil {
			return fmt.Errorf("entgen: uninitialized tblgarageowner.DefaultUpdatedAt (forgotten import entgen/runtime?)")
		}
		v := tblgarageowner.DefaultUpdatedAt()
		tgoc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tgoc.mutation.ID(); !ok {
		if tblgarageowner.DefaultID == nil {
			return fmt.Errorf("entgen: uninitialized tblgarageowner.DefaultID (forgotten import entgen/runtime?)")
		}
		v := tblgarageowner.DefaultID()
		tgoc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tgoc *TblGarageOwnerCreate) check() error {
	if v, ok := tgoc.mutation.CreatedBy(); ok {
		if err := tblgarageowner.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "CreatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblGarageOwner.CreatedBy": %w`, err)}
		}
	}
	if v, ok := tgoc.mutation.UpdatedBy(); ok {
		if err := tblgarageowner.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "UpdatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblGarageOwner.UpdatedBy": %w`, err)}
		}
	}
	if v, ok := tgoc.mutation.DeletedBy(); ok {
		if err := tblgarageowner.DeletedByValidator(v); err != nil {
			return &ValidationError{Name: "DeletedBy", err: fmt.Errorf(`entgen: validator failed for field "TblGarageOwner.DeletedBy": %w`, err)}
		}
	}
	if _, ok := tgoc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "CreatedAt", err: errors.New(`entgen: missing required field "TblGarageOwner.CreatedAt"`)}
	}
	if _, ok := tgoc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "UpdatedAt", err: errors.New(`entgen: missing required field "TblGarageOwner.UpdatedAt"`)}
	}
	if v, ok := tgoc.mutation.UserIdUlid(); ok {
		if err := tblgarageowner.UserIdUlidValidator(v); err != nil {
			return &ValidationError{Name: "UserId_ulid", err: fmt.Errorf(`entgen: validator failed for field "TblGarageOwner.UserId_ulid": %w`, err)}
		}
	}
	if _, ok := tgoc.mutation.FirstName(); !ok {
		return &ValidationError{Name: "FirstName", err: errors.New(`entgen: missing required field "TblGarageOwner.FirstName"`)}
	}
	if v, ok := tgoc.mutation.FirstName(); ok {
		if err := tblgarageowner.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "FirstName", err: fmt.Errorf(`entgen: validator failed for field "TblGarageOwner.FirstName": %w`, err)}
		}
	}
	if v, ok := tgoc.mutation.MiddleName(); ok {
		if err := tblgarageowner.MiddleNameValidator(v); err != nil {
			return &ValidationError{Name: "MiddleName", err: fmt.Errorf(`entgen: validator failed for field "TblGarageOwner.MiddleName": %w`, err)}
		}
	}
	if _, ok := tgoc.mutation.LastName(); !ok {
		return &ValidationError{Name: "LastName", err: errors.New(`entgen: missing required field "TblGarageOwner.LastName"`)}
	}
	if v, ok := tgoc.mutation.LastName(); ok {
		if err := tblgarageowner.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "LastName", err: fmt.Errorf(`entgen: validator failed for field "TblGarageOwner.LastName": %w`, err)}
		}
	}
	if _, ok := tgoc.mutation.ContactNumber(); !ok {
		return &ValidationError{Name: "ContactNumber", err: errors.New(`entgen: missing required field "TblGarageOwner.ContactNumber"`)}
	}
	if v, ok := tgoc.mutation.ContactNumber(); ok {
		if err := tblgarageowner.ContactNumberValidator(v); err != nil {
			return &ValidationError{Name: "ContactNumber", err: fmt.Errorf(`entgen: validator failed for field "TblGarageOwner.ContactNumber": %w`, err)}
		}
	}
	if v, ok := tgoc.mutation.Email(); ok {
		if err := tblgarageowner.EmailValidator(v); err != nil {
			return &ValidationError{Name: "Email", err: fmt.Errorf(`entgen: validator failed for field "TblGarageOwner.Email": %w`, err)}
		}
	}
	if v, ok := tgoc.mutation.PhotoIdUlid(); ok {
		if err := tblgarageowner.PhotoIdUlidValidator(v); err != nil {
			return &ValidationError{Name: "PhotoId_ulid", err: fmt.Errorf(`entgen: validator failed for field "TblGarageOwner.PhotoId_ulid": %w`, err)}
		}
	}
	if _, ok := tgoc.mutation.AddressIdUlid(); !ok {
		return &ValidationError{Name: "AddressId_ulid", err: errors.New(`entgen: missing required field "TblGarageOwner.AddressId_ulid"`)}
	}
	if v, ok := tgoc.mutation.AddressIdUlid(); ok {
		if err := tblgarageowner.AddressIdUlidValidator(v); err != nil {
			return &ValidationError{Name: "AddressId_ulid", err: fmt.Errorf(`entgen: validator failed for field "TblGarageOwner.AddressId_ulid": %w`, err)}
		}
	}
	if v, ok := tgoc.mutation.ID(); ok {
		if err := tblgarageowner.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`entgen: validator failed for field "TblGarageOwner.id": %w`, err)}
		}
	}
	if _, ok := tgoc.mutation.AddressID(); !ok {
		return &ValidationError{Name: "Address", err: errors.New(`entgen: missing required edge "TblGarageOwner.Address"`)}
	}
	return nil
}

func (tgoc *TblGarageOwnerCreate) sqlSave(ctx context.Context) (*TblGarageOwner, error) {
	if err := tgoc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tgoc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tgoc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected TblGarageOwner.ID type: %T", _spec.ID.Value)
		}
	}
	tgoc.mutation.id = &_node.ID
	tgoc.mutation.done = true
	return _node, nil
}

func (tgoc *TblGarageOwnerCreate) createSpec() (*TblGarageOwner, *sqlgraph.CreateSpec) {
	var (
		_node = &TblGarageOwner{config: tgoc.config}
		_spec = sqlgraph.NewCreateSpec(tblgarageowner.Table, sqlgraph.NewFieldSpec(tblgarageowner.FieldID, field.TypeString))
	)
	if id, ok := tgoc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tgoc.mutation.CreatedBy(); ok {
		_spec.SetField(tblgarageowner.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = &value
	}
	if value, ok := tgoc.mutation.UpdatedBy(); ok {
		_spec.SetField(tblgarageowner.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = &value
	}
	if value, ok := tgoc.mutation.DeletedBy(); ok {
		_spec.SetField(tblgarageowner.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = &value
	}
	if value, ok := tgoc.mutation.IP(); ok {
		_spec.SetField(tblgarageowner.FieldIP, field.TypeString, value)
		_node.IP = &value
	}
	if value, ok := tgoc.mutation.UserAgent(); ok {
		_spec.SetField(tblgarageowner.FieldUserAgent, field.TypeString, value)
		_node.UserAgent = &value
	}
	if value, ok := tgoc.mutation.CreatedAt(); ok {
		_spec.SetField(tblgarageowner.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tgoc.mutation.UpdatedAt(); ok {
		_spec.SetField(tblgarageowner.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tgoc.mutation.DeletedAt(); ok {
		_spec.SetField(tblgarageowner.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = &value
	}
	if value, ok := tgoc.mutation.FirstName(); ok {
		_spec.SetField(tblgarageowner.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := tgoc.mutation.MiddleName(); ok {
		_spec.SetField(tblgarageowner.FieldMiddleName, field.TypeString, value)
		_node.MiddleName = &value
	}
	if value, ok := tgoc.mutation.LastName(); ok {
		_spec.SetField(tblgarageowner.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := tgoc.mutation.ContactNumber(); ok {
		_spec.SetField(tblgarageowner.FieldContactNumber, field.TypeString, value)
		_node.ContactNumber = value
	}
	if value, ok := tgoc.mutation.Email(); ok {
		_spec.SetField(tblgarageowner.FieldEmail, field.TypeString, value)
		_node.Email = &value
	}
	if value, ok := tgoc.mutation.Age(); ok {
		_spec.SetField(tblgarageowner.FieldAge, field.TypeInt, value)
		_node.Age = &value
	}
	if nodes := tgoc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   tblgarageowner.UserTable,
			Columns: []string{tblgarageowner.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tblusers.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserIdUlid = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tgoc.mutation.NameInitialIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   tblgarageowner.NameInitialTable,
			Columns: []string{tblgarageowner.NameInitialColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tblenum.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.Initial = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tgoc.mutation.OwnerPhotoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   tblgarageowner.OwnerPhotoTable,
			Columns: []string{tblgarageowner.OwnerPhotoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tbldocument.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PhotoIdUlid = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tgoc.mutation.AddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   tblgarageowner.AddressTable,
			Columns: []string{tblgarageowner.AddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tbladdress.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AddressIdUlid = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TblGarageOwnerCreateBulk is the builder for creating many TblGarageOwner entities in bulk.
type TblGarageOwnerCreateBulk struct {
	config
	err      error
	builders []*TblGarageOwnerCreate
}

// Save creates the TblGarageOwner entities in the database.
func (tgocb *TblGarageOwnerCreateBulk) Save(ctx context.Context) ([]*TblGarageOwner, error) {
	if tgocb.err != nil {
		return nil, tgocb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tgocb.builders))
	nodes := make([]*TblGarageOwner, len(tgocb.builders))
	mutators := make([]Mutator, len(tgocb.builders))
	for i := range tgocb.builders {
		func(i int, root context.Context) {
			builder := tgocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TblGarageOwnerMutation)
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
					_, err = mutators[i+1].Mutate(root, tgocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tgocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tgocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tgocb *TblGarageOwnerCreateBulk) SaveX(ctx context.Context) []*TblGarageOwner {
	v, err := tgocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tgocb *TblGarageOwnerCreateBulk) Exec(ctx context.Context) error {
	_, err := tgocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tgocb *TblGarageOwnerCreateBulk) ExecX(ctx context.Context) {
	if err := tgocb.Exec(ctx); err != nil {
		panic(err)
	}
}
