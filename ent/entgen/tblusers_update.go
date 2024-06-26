// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"
	"rr-backend/ent/entgen/predicate"
	"rr-backend/ent/entgen/tblgarageowner"
	"rr-backend/ent/entgen/tblusers"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TblUSersUpdate is the builder for updating TblUSers entities.
type TblUSersUpdate struct {
	config
	hooks    []Hook
	mutation *TblUSersMutation
}

// Where appends a list predicates to the TblUSersUpdate builder.
func (tuu *TblUSersUpdate) Where(ps ...predicate.TblUSers) *TblUSersUpdate {
	tuu.mutation.Where(ps...)
	return tuu
}

// SetCreatedBy sets the "CreatedBy" field.
func (tuu *TblUSersUpdate) SetCreatedBy(s string) *TblUSersUpdate {
	tuu.mutation.SetCreatedBy(s)
	return tuu
}

// SetNillableCreatedBy sets the "CreatedBy" field if the given value is not nil.
func (tuu *TblUSersUpdate) SetNillableCreatedBy(s *string) *TblUSersUpdate {
	if s != nil {
		tuu.SetCreatedBy(*s)
	}
	return tuu
}

// ClearCreatedBy clears the value of the "CreatedBy" field.
func (tuu *TblUSersUpdate) ClearCreatedBy() *TblUSersUpdate {
	tuu.mutation.ClearCreatedBy()
	return tuu
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (tuu *TblUSersUpdate) SetUpdatedBy(s string) *TblUSersUpdate {
	tuu.mutation.SetUpdatedBy(s)
	return tuu
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (tuu *TblUSersUpdate) SetNillableUpdatedBy(s *string) *TblUSersUpdate {
	if s != nil {
		tuu.SetUpdatedBy(*s)
	}
	return tuu
}

// ClearUpdatedBy clears the value of the "UpdatedBy" field.
func (tuu *TblUSersUpdate) ClearUpdatedBy() *TblUSersUpdate {
	tuu.mutation.ClearUpdatedBy()
	return tuu
}

// SetDeletedBy sets the "DeletedBy" field.
func (tuu *TblUSersUpdate) SetDeletedBy(s string) *TblUSersUpdate {
	tuu.mutation.SetDeletedBy(s)
	return tuu
}

// SetNillableDeletedBy sets the "DeletedBy" field if the given value is not nil.
func (tuu *TblUSersUpdate) SetNillableDeletedBy(s *string) *TblUSersUpdate {
	if s != nil {
		tuu.SetDeletedBy(*s)
	}
	return tuu
}

// ClearDeletedBy clears the value of the "DeletedBy" field.
func (tuu *TblUSersUpdate) ClearDeletedBy() *TblUSersUpdate {
	tuu.mutation.ClearDeletedBy()
	return tuu
}

// SetIP sets the "IP" field.
func (tuu *TblUSersUpdate) SetIP(s string) *TblUSersUpdate {
	tuu.mutation.SetIP(s)
	return tuu
}

// SetNillableIP sets the "IP" field if the given value is not nil.
func (tuu *TblUSersUpdate) SetNillableIP(s *string) *TblUSersUpdate {
	if s != nil {
		tuu.SetIP(*s)
	}
	return tuu
}

// ClearIP clears the value of the "IP" field.
func (tuu *TblUSersUpdate) ClearIP() *TblUSersUpdate {
	tuu.mutation.ClearIP()
	return tuu
}

// SetUserAgent sets the "UserAgent" field.
func (tuu *TblUSersUpdate) SetUserAgent(s string) *TblUSersUpdate {
	tuu.mutation.SetUserAgent(s)
	return tuu
}

// SetNillableUserAgent sets the "UserAgent" field if the given value is not nil.
func (tuu *TblUSersUpdate) SetNillableUserAgent(s *string) *TblUSersUpdate {
	if s != nil {
		tuu.SetUserAgent(*s)
	}
	return tuu
}

// ClearUserAgent clears the value of the "UserAgent" field.
func (tuu *TblUSersUpdate) ClearUserAgent() *TblUSersUpdate {
	tuu.mutation.ClearUserAgent()
	return tuu
}

// SetCreatedAt sets the "CreatedAt" field.
func (tuu *TblUSersUpdate) SetCreatedAt(t time.Time) *TblUSersUpdate {
	tuu.mutation.SetCreatedAt(t)
	return tuu
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (tuu *TblUSersUpdate) SetNillableCreatedAt(t *time.Time) *TblUSersUpdate {
	if t != nil {
		tuu.SetCreatedAt(*t)
	}
	return tuu
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (tuu *TblUSersUpdate) SetUpdatedAt(t time.Time) *TblUSersUpdate {
	tuu.mutation.SetUpdatedAt(t)
	return tuu
}

// SetDeletedAt sets the "DeletedAt" field.
func (tuu *TblUSersUpdate) SetDeletedAt(t time.Time) *TblUSersUpdate {
	tuu.mutation.SetDeletedAt(t)
	return tuu
}

// SetNillableDeletedAt sets the "DeletedAt" field if the given value is not nil.
func (tuu *TblUSersUpdate) SetNillableDeletedAt(t *time.Time) *TblUSersUpdate {
	if t != nil {
		tuu.SetDeletedAt(*t)
	}
	return tuu
}

// ClearDeletedAt clears the value of the "DeletedAt" field.
func (tuu *TblUSersUpdate) ClearDeletedAt() *TblUSersUpdate {
	tuu.mutation.ClearDeletedAt()
	return tuu
}

// SetUserName sets the "UserName" field.
func (tuu *TblUSersUpdate) SetUserName(s string) *TblUSersUpdate {
	tuu.mutation.SetUserName(s)
	return tuu
}

// SetNillableUserName sets the "UserName" field if the given value is not nil.
func (tuu *TblUSersUpdate) SetNillableUserName(s *string) *TblUSersUpdate {
	if s != nil {
		tuu.SetUserName(*s)
	}
	return tuu
}

// SetPassword sets the "Password" field.
func (tuu *TblUSersUpdate) SetPassword(s string) *TblUSersUpdate {
	tuu.mutation.SetPassword(s)
	return tuu
}

// SetNillablePassword sets the "Password" field if the given value is not nil.
func (tuu *TblUSersUpdate) SetNillablePassword(s *string) *TblUSersUpdate {
	if s != nil {
		tuu.SetPassword(*s)
	}
	return tuu
}

// SetEmail sets the "Email" field.
func (tuu *TblUSersUpdate) SetEmail(s string) *TblUSersUpdate {
	tuu.mutation.SetEmail(s)
	return tuu
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (tuu *TblUSersUpdate) SetNillableEmail(s *string) *TblUSersUpdate {
	if s != nil {
		tuu.SetEmail(*s)
	}
	return tuu
}

// ClearEmail clears the value of the "Email" field.
func (tuu *TblUSersUpdate) ClearEmail() *TblUSersUpdate {
	tuu.mutation.ClearEmail()
	return tuu
}

// SetOwnerID sets the "Owner" edge to the TblGarageOwner entity by ID.
func (tuu *TblUSersUpdate) SetOwnerID(id string) *TblUSersUpdate {
	tuu.mutation.SetOwnerID(id)
	return tuu
}

// SetNillableOwnerID sets the "Owner" edge to the TblGarageOwner entity by ID if the given value is not nil.
func (tuu *TblUSersUpdate) SetNillableOwnerID(id *string) *TblUSersUpdate {
	if id != nil {
		tuu = tuu.SetOwnerID(*id)
	}
	return tuu
}

// SetOwner sets the "Owner" edge to the TblGarageOwner entity.
func (tuu *TblUSersUpdate) SetOwner(t *TblGarageOwner) *TblUSersUpdate {
	return tuu.SetOwnerID(t.ID)
}

// Mutation returns the TblUSersMutation object of the builder.
func (tuu *TblUSersUpdate) Mutation() *TblUSersMutation {
	return tuu.mutation
}

// ClearOwner clears the "Owner" edge to the TblGarageOwner entity.
func (tuu *TblUSersUpdate) ClearOwner() *TblUSersUpdate {
	tuu.mutation.ClearOwner()
	return tuu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tuu *TblUSersUpdate) Save(ctx context.Context) (int, error) {
	if err := tuu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, tuu.sqlSave, tuu.mutation, tuu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuu *TblUSersUpdate) SaveX(ctx context.Context) int {
	affected, err := tuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tuu *TblUSersUpdate) Exec(ctx context.Context) error {
	_, err := tuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuu *TblUSersUpdate) ExecX(ctx context.Context) {
	if err := tuu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuu *TblUSersUpdate) defaults() error {
	if _, ok := tuu.mutation.UpdatedAt(); !ok {
		if tblusers.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("entgen: uninitialized tblusers.UpdateDefaultUpdatedAt (forgotten import entgen/runtime?)")
		}
		v := tblusers.UpdateDefaultUpdatedAt()
		tuu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tuu *TblUSersUpdate) check() error {
	if v, ok := tuu.mutation.CreatedBy(); ok {
		if err := tblusers.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "CreatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblUSers.CreatedBy": %w`, err)}
		}
	}
	if v, ok := tuu.mutation.UpdatedBy(); ok {
		if err := tblusers.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "UpdatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblUSers.UpdatedBy": %w`, err)}
		}
	}
	if v, ok := tuu.mutation.DeletedBy(); ok {
		if err := tblusers.DeletedByValidator(v); err != nil {
			return &ValidationError{Name: "DeletedBy", err: fmt.Errorf(`entgen: validator failed for field "TblUSers.DeletedBy": %w`, err)}
		}
	}
	if v, ok := tuu.mutation.UserName(); ok {
		if err := tblusers.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "UserName", err: fmt.Errorf(`entgen: validator failed for field "TblUSers.UserName": %w`, err)}
		}
	}
	if v, ok := tuu.mutation.Password(); ok {
		if err := tblusers.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "Password", err: fmt.Errorf(`entgen: validator failed for field "TblUSers.Password": %w`, err)}
		}
	}
	if v, ok := tuu.mutation.Email(); ok {
		if err := tblusers.EmailValidator(v); err != nil {
			return &ValidationError{Name: "Email", err: fmt.Errorf(`entgen: validator failed for field "TblUSers.Email": %w`, err)}
		}
	}
	return nil
}

func (tuu *TblUSersUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tuu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tblusers.Table, tblusers.Columns, sqlgraph.NewFieldSpec(tblusers.FieldID, field.TypeString))
	if ps := tuu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuu.mutation.CreatedBy(); ok {
		_spec.SetField(tblusers.FieldCreatedBy, field.TypeString, value)
	}
	if tuu.mutation.CreatedByCleared() {
		_spec.ClearField(tblusers.FieldCreatedBy, field.TypeString)
	}
	if value, ok := tuu.mutation.UpdatedBy(); ok {
		_spec.SetField(tblusers.FieldUpdatedBy, field.TypeString, value)
	}
	if tuu.mutation.UpdatedByCleared() {
		_spec.ClearField(tblusers.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := tuu.mutation.DeletedBy(); ok {
		_spec.SetField(tblusers.FieldDeletedBy, field.TypeString, value)
	}
	if tuu.mutation.DeletedByCleared() {
		_spec.ClearField(tblusers.FieldDeletedBy, field.TypeString)
	}
	if value, ok := tuu.mutation.IP(); ok {
		_spec.SetField(tblusers.FieldIP, field.TypeString, value)
	}
	if tuu.mutation.IPCleared() {
		_spec.ClearField(tblusers.FieldIP, field.TypeString)
	}
	if value, ok := tuu.mutation.UserAgent(); ok {
		_spec.SetField(tblusers.FieldUserAgent, field.TypeString, value)
	}
	if tuu.mutation.UserAgentCleared() {
		_spec.ClearField(tblusers.FieldUserAgent, field.TypeString)
	}
	if value, ok := tuu.mutation.CreatedAt(); ok {
		_spec.SetField(tblusers.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tuu.mutation.UpdatedAt(); ok {
		_spec.SetField(tblusers.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuu.mutation.DeletedAt(); ok {
		_spec.SetField(tblusers.FieldDeletedAt, field.TypeTime, value)
	}
	if tuu.mutation.DeletedAtCleared() {
		_spec.ClearField(tblusers.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tuu.mutation.UserName(); ok {
		_spec.SetField(tblusers.FieldUserName, field.TypeString, value)
	}
	if value, ok := tuu.mutation.Password(); ok {
		_spec.SetField(tblusers.FieldPassword, field.TypeString, value)
	}
	if value, ok := tuu.mutation.Email(); ok {
		_spec.SetField(tblusers.FieldEmail, field.TypeString, value)
	}
	if tuu.mutation.EmailCleared() {
		_spec.ClearField(tblusers.FieldEmail, field.TypeString)
	}
	if tuu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   tblusers.OwnerTable,
			Columns: []string{tblusers.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tblgarageowner.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   tblusers.OwnerTable,
			Columns: []string{tblusers.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tblgarageowner.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tuu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tblusers.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tuu.mutation.done = true
	return n, nil
}

// TblUSersUpdateOne is the builder for updating a single TblUSers entity.
type TblUSersUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TblUSersMutation
}

// SetCreatedBy sets the "CreatedBy" field.
func (tuuo *TblUSersUpdateOne) SetCreatedBy(s string) *TblUSersUpdateOne {
	tuuo.mutation.SetCreatedBy(s)
	return tuuo
}

// SetNillableCreatedBy sets the "CreatedBy" field if the given value is not nil.
func (tuuo *TblUSersUpdateOne) SetNillableCreatedBy(s *string) *TblUSersUpdateOne {
	if s != nil {
		tuuo.SetCreatedBy(*s)
	}
	return tuuo
}

// ClearCreatedBy clears the value of the "CreatedBy" field.
func (tuuo *TblUSersUpdateOne) ClearCreatedBy() *TblUSersUpdateOne {
	tuuo.mutation.ClearCreatedBy()
	return tuuo
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (tuuo *TblUSersUpdateOne) SetUpdatedBy(s string) *TblUSersUpdateOne {
	tuuo.mutation.SetUpdatedBy(s)
	return tuuo
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (tuuo *TblUSersUpdateOne) SetNillableUpdatedBy(s *string) *TblUSersUpdateOne {
	if s != nil {
		tuuo.SetUpdatedBy(*s)
	}
	return tuuo
}

// ClearUpdatedBy clears the value of the "UpdatedBy" field.
func (tuuo *TblUSersUpdateOne) ClearUpdatedBy() *TblUSersUpdateOne {
	tuuo.mutation.ClearUpdatedBy()
	return tuuo
}

// SetDeletedBy sets the "DeletedBy" field.
func (tuuo *TblUSersUpdateOne) SetDeletedBy(s string) *TblUSersUpdateOne {
	tuuo.mutation.SetDeletedBy(s)
	return tuuo
}

// SetNillableDeletedBy sets the "DeletedBy" field if the given value is not nil.
func (tuuo *TblUSersUpdateOne) SetNillableDeletedBy(s *string) *TblUSersUpdateOne {
	if s != nil {
		tuuo.SetDeletedBy(*s)
	}
	return tuuo
}

// ClearDeletedBy clears the value of the "DeletedBy" field.
func (tuuo *TblUSersUpdateOne) ClearDeletedBy() *TblUSersUpdateOne {
	tuuo.mutation.ClearDeletedBy()
	return tuuo
}

// SetIP sets the "IP" field.
func (tuuo *TblUSersUpdateOne) SetIP(s string) *TblUSersUpdateOne {
	tuuo.mutation.SetIP(s)
	return tuuo
}

// SetNillableIP sets the "IP" field if the given value is not nil.
func (tuuo *TblUSersUpdateOne) SetNillableIP(s *string) *TblUSersUpdateOne {
	if s != nil {
		tuuo.SetIP(*s)
	}
	return tuuo
}

// ClearIP clears the value of the "IP" field.
func (tuuo *TblUSersUpdateOne) ClearIP() *TblUSersUpdateOne {
	tuuo.mutation.ClearIP()
	return tuuo
}

// SetUserAgent sets the "UserAgent" field.
func (tuuo *TblUSersUpdateOne) SetUserAgent(s string) *TblUSersUpdateOne {
	tuuo.mutation.SetUserAgent(s)
	return tuuo
}

// SetNillableUserAgent sets the "UserAgent" field if the given value is not nil.
func (tuuo *TblUSersUpdateOne) SetNillableUserAgent(s *string) *TblUSersUpdateOne {
	if s != nil {
		tuuo.SetUserAgent(*s)
	}
	return tuuo
}

// ClearUserAgent clears the value of the "UserAgent" field.
func (tuuo *TblUSersUpdateOne) ClearUserAgent() *TblUSersUpdateOne {
	tuuo.mutation.ClearUserAgent()
	return tuuo
}

// SetCreatedAt sets the "CreatedAt" field.
func (tuuo *TblUSersUpdateOne) SetCreatedAt(t time.Time) *TblUSersUpdateOne {
	tuuo.mutation.SetCreatedAt(t)
	return tuuo
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (tuuo *TblUSersUpdateOne) SetNillableCreatedAt(t *time.Time) *TblUSersUpdateOne {
	if t != nil {
		tuuo.SetCreatedAt(*t)
	}
	return tuuo
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (tuuo *TblUSersUpdateOne) SetUpdatedAt(t time.Time) *TblUSersUpdateOne {
	tuuo.mutation.SetUpdatedAt(t)
	return tuuo
}

// SetDeletedAt sets the "DeletedAt" field.
func (tuuo *TblUSersUpdateOne) SetDeletedAt(t time.Time) *TblUSersUpdateOne {
	tuuo.mutation.SetDeletedAt(t)
	return tuuo
}

// SetNillableDeletedAt sets the "DeletedAt" field if the given value is not nil.
func (tuuo *TblUSersUpdateOne) SetNillableDeletedAt(t *time.Time) *TblUSersUpdateOne {
	if t != nil {
		tuuo.SetDeletedAt(*t)
	}
	return tuuo
}

// ClearDeletedAt clears the value of the "DeletedAt" field.
func (tuuo *TblUSersUpdateOne) ClearDeletedAt() *TblUSersUpdateOne {
	tuuo.mutation.ClearDeletedAt()
	return tuuo
}

// SetUserName sets the "UserName" field.
func (tuuo *TblUSersUpdateOne) SetUserName(s string) *TblUSersUpdateOne {
	tuuo.mutation.SetUserName(s)
	return tuuo
}

// SetNillableUserName sets the "UserName" field if the given value is not nil.
func (tuuo *TblUSersUpdateOne) SetNillableUserName(s *string) *TblUSersUpdateOne {
	if s != nil {
		tuuo.SetUserName(*s)
	}
	return tuuo
}

// SetPassword sets the "Password" field.
func (tuuo *TblUSersUpdateOne) SetPassword(s string) *TblUSersUpdateOne {
	tuuo.mutation.SetPassword(s)
	return tuuo
}

// SetNillablePassword sets the "Password" field if the given value is not nil.
func (tuuo *TblUSersUpdateOne) SetNillablePassword(s *string) *TblUSersUpdateOne {
	if s != nil {
		tuuo.SetPassword(*s)
	}
	return tuuo
}

// SetEmail sets the "Email" field.
func (tuuo *TblUSersUpdateOne) SetEmail(s string) *TblUSersUpdateOne {
	tuuo.mutation.SetEmail(s)
	return tuuo
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (tuuo *TblUSersUpdateOne) SetNillableEmail(s *string) *TblUSersUpdateOne {
	if s != nil {
		tuuo.SetEmail(*s)
	}
	return tuuo
}

// ClearEmail clears the value of the "Email" field.
func (tuuo *TblUSersUpdateOne) ClearEmail() *TblUSersUpdateOne {
	tuuo.mutation.ClearEmail()
	return tuuo
}

// SetOwnerID sets the "Owner" edge to the TblGarageOwner entity by ID.
func (tuuo *TblUSersUpdateOne) SetOwnerID(id string) *TblUSersUpdateOne {
	tuuo.mutation.SetOwnerID(id)
	return tuuo
}

// SetNillableOwnerID sets the "Owner" edge to the TblGarageOwner entity by ID if the given value is not nil.
func (tuuo *TblUSersUpdateOne) SetNillableOwnerID(id *string) *TblUSersUpdateOne {
	if id != nil {
		tuuo = tuuo.SetOwnerID(*id)
	}
	return tuuo
}

// SetOwner sets the "Owner" edge to the TblGarageOwner entity.
func (tuuo *TblUSersUpdateOne) SetOwner(t *TblGarageOwner) *TblUSersUpdateOne {
	return tuuo.SetOwnerID(t.ID)
}

// Mutation returns the TblUSersMutation object of the builder.
func (tuuo *TblUSersUpdateOne) Mutation() *TblUSersMutation {
	return tuuo.mutation
}

// ClearOwner clears the "Owner" edge to the TblGarageOwner entity.
func (tuuo *TblUSersUpdateOne) ClearOwner() *TblUSersUpdateOne {
	tuuo.mutation.ClearOwner()
	return tuuo
}

// Where appends a list predicates to the TblUSersUpdate builder.
func (tuuo *TblUSersUpdateOne) Where(ps ...predicate.TblUSers) *TblUSersUpdateOne {
	tuuo.mutation.Where(ps...)
	return tuuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuuo *TblUSersUpdateOne) Select(field string, fields ...string) *TblUSersUpdateOne {
	tuuo.fields = append([]string{field}, fields...)
	return tuuo
}

// Save executes the query and returns the updated TblUSers entity.
func (tuuo *TblUSersUpdateOne) Save(ctx context.Context) (*TblUSers, error) {
	if err := tuuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tuuo.sqlSave, tuuo.mutation, tuuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuuo *TblUSersUpdateOne) SaveX(ctx context.Context) *TblUSers {
	node, err := tuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuuo *TblUSersUpdateOne) Exec(ctx context.Context) error {
	_, err := tuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuuo *TblUSersUpdateOne) ExecX(ctx context.Context) {
	if err := tuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuuo *TblUSersUpdateOne) defaults() error {
	if _, ok := tuuo.mutation.UpdatedAt(); !ok {
		if tblusers.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("entgen: uninitialized tblusers.UpdateDefaultUpdatedAt (forgotten import entgen/runtime?)")
		}
		v := tblusers.UpdateDefaultUpdatedAt()
		tuuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tuuo *TblUSersUpdateOne) check() error {
	if v, ok := tuuo.mutation.CreatedBy(); ok {
		if err := tblusers.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "CreatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblUSers.CreatedBy": %w`, err)}
		}
	}
	if v, ok := tuuo.mutation.UpdatedBy(); ok {
		if err := tblusers.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "UpdatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblUSers.UpdatedBy": %w`, err)}
		}
	}
	if v, ok := tuuo.mutation.DeletedBy(); ok {
		if err := tblusers.DeletedByValidator(v); err != nil {
			return &ValidationError{Name: "DeletedBy", err: fmt.Errorf(`entgen: validator failed for field "TblUSers.DeletedBy": %w`, err)}
		}
	}
	if v, ok := tuuo.mutation.UserName(); ok {
		if err := tblusers.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "UserName", err: fmt.Errorf(`entgen: validator failed for field "TblUSers.UserName": %w`, err)}
		}
	}
	if v, ok := tuuo.mutation.Password(); ok {
		if err := tblusers.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "Password", err: fmt.Errorf(`entgen: validator failed for field "TblUSers.Password": %w`, err)}
		}
	}
	if v, ok := tuuo.mutation.Email(); ok {
		if err := tblusers.EmailValidator(v); err != nil {
			return &ValidationError{Name: "Email", err: fmt.Errorf(`entgen: validator failed for field "TblUSers.Email": %w`, err)}
		}
	}
	return nil
}

func (tuuo *TblUSersUpdateOne) sqlSave(ctx context.Context) (_node *TblUSers, err error) {
	if err := tuuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tblusers.Table, tblusers.Columns, sqlgraph.NewFieldSpec(tblusers.FieldID, field.TypeString))
	id, ok := tuuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entgen: missing "TblUSers.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tblusers.FieldID)
		for _, f := range fields {
			if !tblusers.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
			}
			if f != tblusers.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuuo.mutation.CreatedBy(); ok {
		_spec.SetField(tblusers.FieldCreatedBy, field.TypeString, value)
	}
	if tuuo.mutation.CreatedByCleared() {
		_spec.ClearField(tblusers.FieldCreatedBy, field.TypeString)
	}
	if value, ok := tuuo.mutation.UpdatedBy(); ok {
		_spec.SetField(tblusers.FieldUpdatedBy, field.TypeString, value)
	}
	if tuuo.mutation.UpdatedByCleared() {
		_spec.ClearField(tblusers.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := tuuo.mutation.DeletedBy(); ok {
		_spec.SetField(tblusers.FieldDeletedBy, field.TypeString, value)
	}
	if tuuo.mutation.DeletedByCleared() {
		_spec.ClearField(tblusers.FieldDeletedBy, field.TypeString)
	}
	if value, ok := tuuo.mutation.IP(); ok {
		_spec.SetField(tblusers.FieldIP, field.TypeString, value)
	}
	if tuuo.mutation.IPCleared() {
		_spec.ClearField(tblusers.FieldIP, field.TypeString)
	}
	if value, ok := tuuo.mutation.UserAgent(); ok {
		_spec.SetField(tblusers.FieldUserAgent, field.TypeString, value)
	}
	if tuuo.mutation.UserAgentCleared() {
		_spec.ClearField(tblusers.FieldUserAgent, field.TypeString)
	}
	if value, ok := tuuo.mutation.CreatedAt(); ok {
		_spec.SetField(tblusers.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tuuo.mutation.UpdatedAt(); ok {
		_spec.SetField(tblusers.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuuo.mutation.DeletedAt(); ok {
		_spec.SetField(tblusers.FieldDeletedAt, field.TypeTime, value)
	}
	if tuuo.mutation.DeletedAtCleared() {
		_spec.ClearField(tblusers.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tuuo.mutation.UserName(); ok {
		_spec.SetField(tblusers.FieldUserName, field.TypeString, value)
	}
	if value, ok := tuuo.mutation.Password(); ok {
		_spec.SetField(tblusers.FieldPassword, field.TypeString, value)
	}
	if value, ok := tuuo.mutation.Email(); ok {
		_spec.SetField(tblusers.FieldEmail, field.TypeString, value)
	}
	if tuuo.mutation.EmailCleared() {
		_spec.ClearField(tblusers.FieldEmail, field.TypeString)
	}
	if tuuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   tblusers.OwnerTable,
			Columns: []string{tblusers.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tblgarageowner.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   tblusers.OwnerTable,
			Columns: []string{tblusers.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tblgarageowner.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TblUSers{config: tuuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tblusers.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuuo.mutation.done = true
	return _node, nil
}
