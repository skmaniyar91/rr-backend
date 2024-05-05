// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"errors"
	"fmt"
	"rr-backend/ent/entgen/predicate"
	"rr-backend/ent/entgen/tblsuperadmin"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TblSuperAdminUpdate is the builder for updating TblSuperAdmin entities.
type TblSuperAdminUpdate struct {
	config
	hooks    []Hook
	mutation *TblSuperAdminMutation
}

// Where appends a list predicates to the TblSuperAdminUpdate builder.
func (tsau *TblSuperAdminUpdate) Where(ps ...predicate.TblSuperAdmin) *TblSuperAdminUpdate {
	tsau.mutation.Where(ps...)
	return tsau
}

// SetCreatedBy sets the "CreatedBy" field.
func (tsau *TblSuperAdminUpdate) SetCreatedBy(s string) *TblSuperAdminUpdate {
	tsau.mutation.SetCreatedBy(s)
	return tsau
}

// SetNillableCreatedBy sets the "CreatedBy" field if the given value is not nil.
func (tsau *TblSuperAdminUpdate) SetNillableCreatedBy(s *string) *TblSuperAdminUpdate {
	if s != nil {
		tsau.SetCreatedBy(*s)
	}
	return tsau
}

// ClearCreatedBy clears the value of the "CreatedBy" field.
func (tsau *TblSuperAdminUpdate) ClearCreatedBy() *TblSuperAdminUpdate {
	tsau.mutation.ClearCreatedBy()
	return tsau
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (tsau *TblSuperAdminUpdate) SetUpdatedBy(s string) *TblSuperAdminUpdate {
	tsau.mutation.SetUpdatedBy(s)
	return tsau
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (tsau *TblSuperAdminUpdate) SetNillableUpdatedBy(s *string) *TblSuperAdminUpdate {
	if s != nil {
		tsau.SetUpdatedBy(*s)
	}
	return tsau
}

// ClearUpdatedBy clears the value of the "UpdatedBy" field.
func (tsau *TblSuperAdminUpdate) ClearUpdatedBy() *TblSuperAdminUpdate {
	tsau.mutation.ClearUpdatedBy()
	return tsau
}

// SetDeletedBy sets the "DeletedBy" field.
func (tsau *TblSuperAdminUpdate) SetDeletedBy(s string) *TblSuperAdminUpdate {
	tsau.mutation.SetDeletedBy(s)
	return tsau
}

// SetNillableDeletedBy sets the "DeletedBy" field if the given value is not nil.
func (tsau *TblSuperAdminUpdate) SetNillableDeletedBy(s *string) *TblSuperAdminUpdate {
	if s != nil {
		tsau.SetDeletedBy(*s)
	}
	return tsau
}

// ClearDeletedBy clears the value of the "DeletedBy" field.
func (tsau *TblSuperAdminUpdate) ClearDeletedBy() *TblSuperAdminUpdate {
	tsau.mutation.ClearDeletedBy()
	return tsau
}

// SetIP sets the "IP" field.
func (tsau *TblSuperAdminUpdate) SetIP(s string) *TblSuperAdminUpdate {
	tsau.mutation.SetIP(s)
	return tsau
}

// SetNillableIP sets the "IP" field if the given value is not nil.
func (tsau *TblSuperAdminUpdate) SetNillableIP(s *string) *TblSuperAdminUpdate {
	if s != nil {
		tsau.SetIP(*s)
	}
	return tsau
}

// ClearIP clears the value of the "IP" field.
func (tsau *TblSuperAdminUpdate) ClearIP() *TblSuperAdminUpdate {
	tsau.mutation.ClearIP()
	return tsau
}

// SetUserAgent sets the "UserAgent" field.
func (tsau *TblSuperAdminUpdate) SetUserAgent(s string) *TblSuperAdminUpdate {
	tsau.mutation.SetUserAgent(s)
	return tsau
}

// SetNillableUserAgent sets the "UserAgent" field if the given value is not nil.
func (tsau *TblSuperAdminUpdate) SetNillableUserAgent(s *string) *TblSuperAdminUpdate {
	if s != nil {
		tsau.SetUserAgent(*s)
	}
	return tsau
}

// ClearUserAgent clears the value of the "UserAgent" field.
func (tsau *TblSuperAdminUpdate) ClearUserAgent() *TblSuperAdminUpdate {
	tsau.mutation.ClearUserAgent()
	return tsau
}

// SetCreatedAt sets the "CreatedAt" field.
func (tsau *TblSuperAdminUpdate) SetCreatedAt(t time.Time) *TblSuperAdminUpdate {
	tsau.mutation.SetCreatedAt(t)
	return tsau
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (tsau *TblSuperAdminUpdate) SetNillableCreatedAt(t *time.Time) *TblSuperAdminUpdate {
	if t != nil {
		tsau.SetCreatedAt(*t)
	}
	return tsau
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (tsau *TblSuperAdminUpdate) SetUpdatedAt(t time.Time) *TblSuperAdminUpdate {
	tsau.mutation.SetUpdatedAt(t)
	return tsau
}

// SetDeletedAt sets the "DeletedAt" field.
func (tsau *TblSuperAdminUpdate) SetDeletedAt(t time.Time) *TblSuperAdminUpdate {
	tsau.mutation.SetDeletedAt(t)
	return tsau
}

// SetNillableDeletedAt sets the "DeletedAt" field if the given value is not nil.
func (tsau *TblSuperAdminUpdate) SetNillableDeletedAt(t *time.Time) *TblSuperAdminUpdate {
	if t != nil {
		tsau.SetDeletedAt(*t)
	}
	return tsau
}

// ClearDeletedAt clears the value of the "DeletedAt" field.
func (tsau *TblSuperAdminUpdate) ClearDeletedAt() *TblSuperAdminUpdate {
	tsau.mutation.ClearDeletedAt()
	return tsau
}

// SetUserName sets the "UserName" field.
func (tsau *TblSuperAdminUpdate) SetUserName(s string) *TblSuperAdminUpdate {
	tsau.mutation.SetUserName(s)
	return tsau
}

// SetNillableUserName sets the "UserName" field if the given value is not nil.
func (tsau *TblSuperAdminUpdate) SetNillableUserName(s *string) *TblSuperAdminUpdate {
	if s != nil {
		tsau.SetUserName(*s)
	}
	return tsau
}

// SetPassWord sets the "PassWord" field.
func (tsau *TblSuperAdminUpdate) SetPassWord(s string) *TblSuperAdminUpdate {
	tsau.mutation.SetPassWord(s)
	return tsau
}

// SetNillablePassWord sets the "PassWord" field if the given value is not nil.
func (tsau *TblSuperAdminUpdate) SetNillablePassWord(s *string) *TblSuperAdminUpdate {
	if s != nil {
		tsau.SetPassWord(*s)
	}
	return tsau
}

// Mutation returns the TblSuperAdminMutation object of the builder.
func (tsau *TblSuperAdminUpdate) Mutation() *TblSuperAdminMutation {
	return tsau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tsau *TblSuperAdminUpdate) Save(ctx context.Context) (int, error) {
	if err := tsau.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, tsau.sqlSave, tsau.mutation, tsau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tsau *TblSuperAdminUpdate) SaveX(ctx context.Context) int {
	affected, err := tsau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tsau *TblSuperAdminUpdate) Exec(ctx context.Context) error {
	_, err := tsau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsau *TblSuperAdminUpdate) ExecX(ctx context.Context) {
	if err := tsau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tsau *TblSuperAdminUpdate) defaults() error {
	if _, ok := tsau.mutation.UpdatedAt(); !ok {
		if tblsuperadmin.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("entgen: uninitialized tblsuperadmin.UpdateDefaultUpdatedAt (forgotten import entgen/runtime?)")
		}
		v := tblsuperadmin.UpdateDefaultUpdatedAt()
		tsau.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tsau *TblSuperAdminUpdate) check() error {
	if v, ok := tsau.mutation.CreatedBy(); ok {
		if err := tblsuperadmin.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "CreatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.CreatedBy": %w`, err)}
		}
	}
	if v, ok := tsau.mutation.UpdatedBy(); ok {
		if err := tblsuperadmin.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "UpdatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.UpdatedBy": %w`, err)}
		}
	}
	if v, ok := tsau.mutation.DeletedBy(); ok {
		if err := tblsuperadmin.DeletedByValidator(v); err != nil {
			return &ValidationError{Name: "DeletedBy", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.DeletedBy": %w`, err)}
		}
	}
	if v, ok := tsau.mutation.UserName(); ok {
		if err := tblsuperadmin.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "UserName", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.UserName": %w`, err)}
		}
	}
	if v, ok := tsau.mutation.PassWord(); ok {
		if err := tblsuperadmin.PassWordValidator(v); err != nil {
			return &ValidationError{Name: "PassWord", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.PassWord": %w`, err)}
		}
	}
	return nil
}

func (tsau *TblSuperAdminUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tsau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tblsuperadmin.Table, tblsuperadmin.Columns, sqlgraph.NewFieldSpec(tblsuperadmin.FieldID, field.TypeString))
	if ps := tsau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tsau.mutation.CreatedBy(); ok {
		_spec.SetField(tblsuperadmin.FieldCreatedBy, field.TypeString, value)
	}
	if tsau.mutation.CreatedByCleared() {
		_spec.ClearField(tblsuperadmin.FieldCreatedBy, field.TypeString)
	}
	if value, ok := tsau.mutation.UpdatedBy(); ok {
		_spec.SetField(tblsuperadmin.FieldUpdatedBy, field.TypeString, value)
	}
	if tsau.mutation.UpdatedByCleared() {
		_spec.ClearField(tblsuperadmin.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := tsau.mutation.DeletedBy(); ok {
		_spec.SetField(tblsuperadmin.FieldDeletedBy, field.TypeString, value)
	}
	if tsau.mutation.DeletedByCleared() {
		_spec.ClearField(tblsuperadmin.FieldDeletedBy, field.TypeString)
	}
	if value, ok := tsau.mutation.IP(); ok {
		_spec.SetField(tblsuperadmin.FieldIP, field.TypeString, value)
	}
	if tsau.mutation.IPCleared() {
		_spec.ClearField(tblsuperadmin.FieldIP, field.TypeString)
	}
	if value, ok := tsau.mutation.UserAgent(); ok {
		_spec.SetField(tblsuperadmin.FieldUserAgent, field.TypeString, value)
	}
	if tsau.mutation.UserAgentCleared() {
		_spec.ClearField(tblsuperadmin.FieldUserAgent, field.TypeString)
	}
	if value, ok := tsau.mutation.CreatedAt(); ok {
		_spec.SetField(tblsuperadmin.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tsau.mutation.UpdatedAt(); ok {
		_spec.SetField(tblsuperadmin.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tsau.mutation.DeletedAt(); ok {
		_spec.SetField(tblsuperadmin.FieldDeletedAt, field.TypeTime, value)
	}
	if tsau.mutation.DeletedAtCleared() {
		_spec.ClearField(tblsuperadmin.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tsau.mutation.UserName(); ok {
		_spec.SetField(tblsuperadmin.FieldUserName, field.TypeString, value)
	}
	if value, ok := tsau.mutation.PassWord(); ok {
		_spec.SetField(tblsuperadmin.FieldPassWord, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tsau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tblsuperadmin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tsau.mutation.done = true
	return n, nil
}

// TblSuperAdminUpdateOne is the builder for updating a single TblSuperAdmin entity.
type TblSuperAdminUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TblSuperAdminMutation
}

// SetCreatedBy sets the "CreatedBy" field.
func (tsauo *TblSuperAdminUpdateOne) SetCreatedBy(s string) *TblSuperAdminUpdateOne {
	tsauo.mutation.SetCreatedBy(s)
	return tsauo
}

// SetNillableCreatedBy sets the "CreatedBy" field if the given value is not nil.
func (tsauo *TblSuperAdminUpdateOne) SetNillableCreatedBy(s *string) *TblSuperAdminUpdateOne {
	if s != nil {
		tsauo.SetCreatedBy(*s)
	}
	return tsauo
}

// ClearCreatedBy clears the value of the "CreatedBy" field.
func (tsauo *TblSuperAdminUpdateOne) ClearCreatedBy() *TblSuperAdminUpdateOne {
	tsauo.mutation.ClearCreatedBy()
	return tsauo
}

// SetUpdatedBy sets the "UpdatedBy" field.
func (tsauo *TblSuperAdminUpdateOne) SetUpdatedBy(s string) *TblSuperAdminUpdateOne {
	tsauo.mutation.SetUpdatedBy(s)
	return tsauo
}

// SetNillableUpdatedBy sets the "UpdatedBy" field if the given value is not nil.
func (tsauo *TblSuperAdminUpdateOne) SetNillableUpdatedBy(s *string) *TblSuperAdminUpdateOne {
	if s != nil {
		tsauo.SetUpdatedBy(*s)
	}
	return tsauo
}

// ClearUpdatedBy clears the value of the "UpdatedBy" field.
func (tsauo *TblSuperAdminUpdateOne) ClearUpdatedBy() *TblSuperAdminUpdateOne {
	tsauo.mutation.ClearUpdatedBy()
	return tsauo
}

// SetDeletedBy sets the "DeletedBy" field.
func (tsauo *TblSuperAdminUpdateOne) SetDeletedBy(s string) *TblSuperAdminUpdateOne {
	tsauo.mutation.SetDeletedBy(s)
	return tsauo
}

// SetNillableDeletedBy sets the "DeletedBy" field if the given value is not nil.
func (tsauo *TblSuperAdminUpdateOne) SetNillableDeletedBy(s *string) *TblSuperAdminUpdateOne {
	if s != nil {
		tsauo.SetDeletedBy(*s)
	}
	return tsauo
}

// ClearDeletedBy clears the value of the "DeletedBy" field.
func (tsauo *TblSuperAdminUpdateOne) ClearDeletedBy() *TblSuperAdminUpdateOne {
	tsauo.mutation.ClearDeletedBy()
	return tsauo
}

// SetIP sets the "IP" field.
func (tsauo *TblSuperAdminUpdateOne) SetIP(s string) *TblSuperAdminUpdateOne {
	tsauo.mutation.SetIP(s)
	return tsauo
}

// SetNillableIP sets the "IP" field if the given value is not nil.
func (tsauo *TblSuperAdminUpdateOne) SetNillableIP(s *string) *TblSuperAdminUpdateOne {
	if s != nil {
		tsauo.SetIP(*s)
	}
	return tsauo
}

// ClearIP clears the value of the "IP" field.
func (tsauo *TblSuperAdminUpdateOne) ClearIP() *TblSuperAdminUpdateOne {
	tsauo.mutation.ClearIP()
	return tsauo
}

// SetUserAgent sets the "UserAgent" field.
func (tsauo *TblSuperAdminUpdateOne) SetUserAgent(s string) *TblSuperAdminUpdateOne {
	tsauo.mutation.SetUserAgent(s)
	return tsauo
}

// SetNillableUserAgent sets the "UserAgent" field if the given value is not nil.
func (tsauo *TblSuperAdminUpdateOne) SetNillableUserAgent(s *string) *TblSuperAdminUpdateOne {
	if s != nil {
		tsauo.SetUserAgent(*s)
	}
	return tsauo
}

// ClearUserAgent clears the value of the "UserAgent" field.
func (tsauo *TblSuperAdminUpdateOne) ClearUserAgent() *TblSuperAdminUpdateOne {
	tsauo.mutation.ClearUserAgent()
	return tsauo
}

// SetCreatedAt sets the "CreatedAt" field.
func (tsauo *TblSuperAdminUpdateOne) SetCreatedAt(t time.Time) *TblSuperAdminUpdateOne {
	tsauo.mutation.SetCreatedAt(t)
	return tsauo
}

// SetNillableCreatedAt sets the "CreatedAt" field if the given value is not nil.
func (tsauo *TblSuperAdminUpdateOne) SetNillableCreatedAt(t *time.Time) *TblSuperAdminUpdateOne {
	if t != nil {
		tsauo.SetCreatedAt(*t)
	}
	return tsauo
}

// SetUpdatedAt sets the "UpdatedAt" field.
func (tsauo *TblSuperAdminUpdateOne) SetUpdatedAt(t time.Time) *TblSuperAdminUpdateOne {
	tsauo.mutation.SetUpdatedAt(t)
	return tsauo
}

// SetDeletedAt sets the "DeletedAt" field.
func (tsauo *TblSuperAdminUpdateOne) SetDeletedAt(t time.Time) *TblSuperAdminUpdateOne {
	tsauo.mutation.SetDeletedAt(t)
	return tsauo
}

// SetNillableDeletedAt sets the "DeletedAt" field if the given value is not nil.
func (tsauo *TblSuperAdminUpdateOne) SetNillableDeletedAt(t *time.Time) *TblSuperAdminUpdateOne {
	if t != nil {
		tsauo.SetDeletedAt(*t)
	}
	return tsauo
}

// ClearDeletedAt clears the value of the "DeletedAt" field.
func (tsauo *TblSuperAdminUpdateOne) ClearDeletedAt() *TblSuperAdminUpdateOne {
	tsauo.mutation.ClearDeletedAt()
	return tsauo
}

// SetUserName sets the "UserName" field.
func (tsauo *TblSuperAdminUpdateOne) SetUserName(s string) *TblSuperAdminUpdateOne {
	tsauo.mutation.SetUserName(s)
	return tsauo
}

// SetNillableUserName sets the "UserName" field if the given value is not nil.
func (tsauo *TblSuperAdminUpdateOne) SetNillableUserName(s *string) *TblSuperAdminUpdateOne {
	if s != nil {
		tsauo.SetUserName(*s)
	}
	return tsauo
}

// SetPassWord sets the "PassWord" field.
func (tsauo *TblSuperAdminUpdateOne) SetPassWord(s string) *TblSuperAdminUpdateOne {
	tsauo.mutation.SetPassWord(s)
	return tsauo
}

// SetNillablePassWord sets the "PassWord" field if the given value is not nil.
func (tsauo *TblSuperAdminUpdateOne) SetNillablePassWord(s *string) *TblSuperAdminUpdateOne {
	if s != nil {
		tsauo.SetPassWord(*s)
	}
	return tsauo
}

// Mutation returns the TblSuperAdminMutation object of the builder.
func (tsauo *TblSuperAdminUpdateOne) Mutation() *TblSuperAdminMutation {
	return tsauo.mutation
}

// Where appends a list predicates to the TblSuperAdminUpdate builder.
func (tsauo *TblSuperAdminUpdateOne) Where(ps ...predicate.TblSuperAdmin) *TblSuperAdminUpdateOne {
	tsauo.mutation.Where(ps...)
	return tsauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tsauo *TblSuperAdminUpdateOne) Select(field string, fields ...string) *TblSuperAdminUpdateOne {
	tsauo.fields = append([]string{field}, fields...)
	return tsauo
}

// Save executes the query and returns the updated TblSuperAdmin entity.
func (tsauo *TblSuperAdminUpdateOne) Save(ctx context.Context) (*TblSuperAdmin, error) {
	if err := tsauo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tsauo.sqlSave, tsauo.mutation, tsauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tsauo *TblSuperAdminUpdateOne) SaveX(ctx context.Context) *TblSuperAdmin {
	node, err := tsauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tsauo *TblSuperAdminUpdateOne) Exec(ctx context.Context) error {
	_, err := tsauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tsauo *TblSuperAdminUpdateOne) ExecX(ctx context.Context) {
	if err := tsauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tsauo *TblSuperAdminUpdateOne) defaults() error {
	if _, ok := tsauo.mutation.UpdatedAt(); !ok {
		if tblsuperadmin.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("entgen: uninitialized tblsuperadmin.UpdateDefaultUpdatedAt (forgotten import entgen/runtime?)")
		}
		v := tblsuperadmin.UpdateDefaultUpdatedAt()
		tsauo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tsauo *TblSuperAdminUpdateOne) check() error {
	if v, ok := tsauo.mutation.CreatedBy(); ok {
		if err := tblsuperadmin.CreatedByValidator(v); err != nil {
			return &ValidationError{Name: "CreatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.CreatedBy": %w`, err)}
		}
	}
	if v, ok := tsauo.mutation.UpdatedBy(); ok {
		if err := tblsuperadmin.UpdatedByValidator(v); err != nil {
			return &ValidationError{Name: "UpdatedBy", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.UpdatedBy": %w`, err)}
		}
	}
	if v, ok := tsauo.mutation.DeletedBy(); ok {
		if err := tblsuperadmin.DeletedByValidator(v); err != nil {
			return &ValidationError{Name: "DeletedBy", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.DeletedBy": %w`, err)}
		}
	}
	if v, ok := tsauo.mutation.UserName(); ok {
		if err := tblsuperadmin.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "UserName", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.UserName": %w`, err)}
		}
	}
	if v, ok := tsauo.mutation.PassWord(); ok {
		if err := tblsuperadmin.PassWordValidator(v); err != nil {
			return &ValidationError{Name: "PassWord", err: fmt.Errorf(`entgen: validator failed for field "TblSuperAdmin.PassWord": %w`, err)}
		}
	}
	return nil
}

func (tsauo *TblSuperAdminUpdateOne) sqlSave(ctx context.Context) (_node *TblSuperAdmin, err error) {
	if err := tsauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tblsuperadmin.Table, tblsuperadmin.Columns, sqlgraph.NewFieldSpec(tblsuperadmin.FieldID, field.TypeString))
	id, ok := tsauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entgen: missing "TblSuperAdmin.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tsauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tblsuperadmin.FieldID)
		for _, f := range fields {
			if !tblsuperadmin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
			}
			if f != tblsuperadmin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tsauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tsauo.mutation.CreatedBy(); ok {
		_spec.SetField(tblsuperadmin.FieldCreatedBy, field.TypeString, value)
	}
	if tsauo.mutation.CreatedByCleared() {
		_spec.ClearField(tblsuperadmin.FieldCreatedBy, field.TypeString)
	}
	if value, ok := tsauo.mutation.UpdatedBy(); ok {
		_spec.SetField(tblsuperadmin.FieldUpdatedBy, field.TypeString, value)
	}
	if tsauo.mutation.UpdatedByCleared() {
		_spec.ClearField(tblsuperadmin.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := tsauo.mutation.DeletedBy(); ok {
		_spec.SetField(tblsuperadmin.FieldDeletedBy, field.TypeString, value)
	}
	if tsauo.mutation.DeletedByCleared() {
		_spec.ClearField(tblsuperadmin.FieldDeletedBy, field.TypeString)
	}
	if value, ok := tsauo.mutation.IP(); ok {
		_spec.SetField(tblsuperadmin.FieldIP, field.TypeString, value)
	}
	if tsauo.mutation.IPCleared() {
		_spec.ClearField(tblsuperadmin.FieldIP, field.TypeString)
	}
	if value, ok := tsauo.mutation.UserAgent(); ok {
		_spec.SetField(tblsuperadmin.FieldUserAgent, field.TypeString, value)
	}
	if tsauo.mutation.UserAgentCleared() {
		_spec.ClearField(tblsuperadmin.FieldUserAgent, field.TypeString)
	}
	if value, ok := tsauo.mutation.CreatedAt(); ok {
		_spec.SetField(tblsuperadmin.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tsauo.mutation.UpdatedAt(); ok {
		_spec.SetField(tblsuperadmin.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tsauo.mutation.DeletedAt(); ok {
		_spec.SetField(tblsuperadmin.FieldDeletedAt, field.TypeTime, value)
	}
	if tsauo.mutation.DeletedAtCleared() {
		_spec.ClearField(tblsuperadmin.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := tsauo.mutation.UserName(); ok {
		_spec.SetField(tblsuperadmin.FieldUserName, field.TypeString, value)
	}
	if value, ok := tsauo.mutation.PassWord(); ok {
		_spec.SetField(tblsuperadmin.FieldPassWord, field.TypeString, value)
	}
	_node = &TblSuperAdmin{config: tsauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tsauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tblsuperadmin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tsauo.mutation.done = true
	return _node, nil
}
