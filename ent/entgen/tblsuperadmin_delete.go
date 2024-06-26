// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"rr-backend/ent/entgen/predicate"
	"rr-backend/ent/entgen/tblsuperadmin"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TblSuperAdminDelete is the builder for deleting a TblSuperAdmin entity.
type TblSuperAdminDelete struct {
	config
	hooks    []Hook
	mutation *TblSuperAdminMutation
}

// Where appends a list predicates to the TblSuperAdminDelete builder.
func (tsad *TblSuperAdminDelete) Where(ps ...predicate.TblSuperAdmin) *TblSuperAdminDelete {
	tsad.mutation.Where(ps...)
	return tsad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tsad *TblSuperAdminDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, tsad.sqlExec, tsad.mutation, tsad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (tsad *TblSuperAdminDelete) ExecX(ctx context.Context) int {
	n, err := tsad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tsad *TblSuperAdminDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(tblsuperadmin.Table, sqlgraph.NewFieldSpec(tblsuperadmin.FieldID, field.TypeString))
	if ps := tsad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tsad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	tsad.mutation.done = true
	return affected, err
}

// TblSuperAdminDeleteOne is the builder for deleting a single TblSuperAdmin entity.
type TblSuperAdminDeleteOne struct {
	tsad *TblSuperAdminDelete
}

// Where appends a list predicates to the TblSuperAdminDelete builder.
func (tsado *TblSuperAdminDeleteOne) Where(ps ...predicate.TblSuperAdmin) *TblSuperAdminDeleteOne {
	tsado.tsad.mutation.Where(ps...)
	return tsado
}

// Exec executes the deletion query.
func (tsado *TblSuperAdminDeleteOne) Exec(ctx context.Context) error {
	n, err := tsado.tsad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tblsuperadmin.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tsado *TblSuperAdminDeleteOne) ExecX(ctx context.Context) {
	if err := tsado.Exec(ctx); err != nil {
		panic(err)
	}
}
