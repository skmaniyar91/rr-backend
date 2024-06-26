// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"fmt"
	"math"
	"rr-backend/ent/entgen/predicate"
	"rr-backend/ent/entgen/tblsuperadmin"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TblSuperAdminQuery is the builder for querying TblSuperAdmin entities.
type TblSuperAdminQuery struct {
	config
	ctx        *QueryContext
	order      []tblsuperadmin.OrderOption
	inters     []Interceptor
	predicates []predicate.TblSuperAdmin
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TblSuperAdminQuery builder.
func (tsaq *TblSuperAdminQuery) Where(ps ...predicate.TblSuperAdmin) *TblSuperAdminQuery {
	tsaq.predicates = append(tsaq.predicates, ps...)
	return tsaq
}

// Limit the number of records to be returned by this query.
func (tsaq *TblSuperAdminQuery) Limit(limit int) *TblSuperAdminQuery {
	tsaq.ctx.Limit = &limit
	return tsaq
}

// Offset to start from.
func (tsaq *TblSuperAdminQuery) Offset(offset int) *TblSuperAdminQuery {
	tsaq.ctx.Offset = &offset
	return tsaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tsaq *TblSuperAdminQuery) Unique(unique bool) *TblSuperAdminQuery {
	tsaq.ctx.Unique = &unique
	return tsaq
}

// Order specifies how the records should be ordered.
func (tsaq *TblSuperAdminQuery) Order(o ...tblsuperadmin.OrderOption) *TblSuperAdminQuery {
	tsaq.order = append(tsaq.order, o...)
	return tsaq
}

// First returns the first TblSuperAdmin entity from the query.
// Returns a *NotFoundError when no TblSuperAdmin was found.
func (tsaq *TblSuperAdminQuery) First(ctx context.Context) (*TblSuperAdmin, error) {
	nodes, err := tsaq.Limit(1).All(setContextOp(ctx, tsaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tblsuperadmin.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tsaq *TblSuperAdminQuery) FirstX(ctx context.Context) *TblSuperAdmin {
	node, err := tsaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TblSuperAdmin ID from the query.
// Returns a *NotFoundError when no TblSuperAdmin ID was found.
func (tsaq *TblSuperAdminQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tsaq.Limit(1).IDs(setContextOp(ctx, tsaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tblsuperadmin.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tsaq *TblSuperAdminQuery) FirstIDX(ctx context.Context) string {
	id, err := tsaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TblSuperAdmin entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TblSuperAdmin entity is found.
// Returns a *NotFoundError when no TblSuperAdmin entities are found.
func (tsaq *TblSuperAdminQuery) Only(ctx context.Context) (*TblSuperAdmin, error) {
	nodes, err := tsaq.Limit(2).All(setContextOp(ctx, tsaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tblsuperadmin.Label}
	default:
		return nil, &NotSingularError{tblsuperadmin.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tsaq *TblSuperAdminQuery) OnlyX(ctx context.Context) *TblSuperAdmin {
	node, err := tsaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TblSuperAdmin ID in the query.
// Returns a *NotSingularError when more than one TblSuperAdmin ID is found.
// Returns a *NotFoundError when no entities are found.
func (tsaq *TblSuperAdminQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tsaq.Limit(2).IDs(setContextOp(ctx, tsaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tblsuperadmin.Label}
	default:
		err = &NotSingularError{tblsuperadmin.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tsaq *TblSuperAdminQuery) OnlyIDX(ctx context.Context) string {
	id, err := tsaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TblSuperAdmins.
func (tsaq *TblSuperAdminQuery) All(ctx context.Context) ([]*TblSuperAdmin, error) {
	ctx = setContextOp(ctx, tsaq.ctx, "All")
	if err := tsaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TblSuperAdmin, *TblSuperAdminQuery]()
	return withInterceptors[[]*TblSuperAdmin](ctx, tsaq, qr, tsaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tsaq *TblSuperAdminQuery) AllX(ctx context.Context) []*TblSuperAdmin {
	nodes, err := tsaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TblSuperAdmin IDs.
func (tsaq *TblSuperAdminQuery) IDs(ctx context.Context) (ids []string, err error) {
	if tsaq.ctx.Unique == nil && tsaq.path != nil {
		tsaq.Unique(true)
	}
	ctx = setContextOp(ctx, tsaq.ctx, "IDs")
	if err = tsaq.Select(tblsuperadmin.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tsaq *TblSuperAdminQuery) IDsX(ctx context.Context) []string {
	ids, err := tsaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tsaq *TblSuperAdminQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tsaq.ctx, "Count")
	if err := tsaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tsaq, querierCount[*TblSuperAdminQuery](), tsaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tsaq *TblSuperAdminQuery) CountX(ctx context.Context) int {
	count, err := tsaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tsaq *TblSuperAdminQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tsaq.ctx, "Exist")
	switch _, err := tsaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entgen: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tsaq *TblSuperAdminQuery) ExistX(ctx context.Context) bool {
	exist, err := tsaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TblSuperAdminQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tsaq *TblSuperAdminQuery) Clone() *TblSuperAdminQuery {
	if tsaq == nil {
		return nil
	}
	return &TblSuperAdminQuery{
		config:     tsaq.config,
		ctx:        tsaq.ctx.Clone(),
		order:      append([]tblsuperadmin.OrderOption{}, tsaq.order...),
		inters:     append([]Interceptor{}, tsaq.inters...),
		predicates: append([]predicate.TblSuperAdmin{}, tsaq.predicates...),
		// clone intermediate query.
		sql:  tsaq.sql.Clone(),
		path: tsaq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy string `json:"CreatedBy,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TblSuperAdmin.Query().
//		GroupBy(tblsuperadmin.FieldCreatedBy).
//		Aggregate(entgen.Count()).
//		Scan(ctx, &v)
func (tsaq *TblSuperAdminQuery) GroupBy(field string, fields ...string) *TblSuperAdminGroupBy {
	tsaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TblSuperAdminGroupBy{build: tsaq}
	grbuild.flds = &tsaq.ctx.Fields
	grbuild.label = tblsuperadmin.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy string `json:"CreatedBy,omitempty"`
//	}
//
//	client.TblSuperAdmin.Query().
//		Select(tblsuperadmin.FieldCreatedBy).
//		Scan(ctx, &v)
func (tsaq *TblSuperAdminQuery) Select(fields ...string) *TblSuperAdminSelect {
	tsaq.ctx.Fields = append(tsaq.ctx.Fields, fields...)
	sbuild := &TblSuperAdminSelect{TblSuperAdminQuery: tsaq}
	sbuild.label = tblsuperadmin.Label
	sbuild.flds, sbuild.scan = &tsaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TblSuperAdminSelect configured with the given aggregations.
func (tsaq *TblSuperAdminQuery) Aggregate(fns ...AggregateFunc) *TblSuperAdminSelect {
	return tsaq.Select().Aggregate(fns...)
}

func (tsaq *TblSuperAdminQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tsaq.inters {
		if inter == nil {
			return fmt.Errorf("entgen: uninitialized interceptor (forgotten import entgen/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tsaq); err != nil {
				return err
			}
		}
	}
	for _, f := range tsaq.ctx.Fields {
		if !tblsuperadmin.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
		}
	}
	if tsaq.path != nil {
		prev, err := tsaq.path(ctx)
		if err != nil {
			return err
		}
		tsaq.sql = prev
	}
	return nil
}

func (tsaq *TblSuperAdminQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TblSuperAdmin, error) {
	var (
		nodes = []*TblSuperAdmin{}
		_spec = tsaq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TblSuperAdmin).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TblSuperAdmin{config: tsaq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tsaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tsaq *TblSuperAdminQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tsaq.querySpec()
	_spec.Node.Columns = tsaq.ctx.Fields
	if len(tsaq.ctx.Fields) > 0 {
		_spec.Unique = tsaq.ctx.Unique != nil && *tsaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tsaq.driver, _spec)
}

func (tsaq *TblSuperAdminQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tblsuperadmin.Table, tblsuperadmin.Columns, sqlgraph.NewFieldSpec(tblsuperadmin.FieldID, field.TypeString))
	_spec.From = tsaq.sql
	if unique := tsaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tsaq.path != nil {
		_spec.Unique = true
	}
	if fields := tsaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tblsuperadmin.FieldID)
		for i := range fields {
			if fields[i] != tblsuperadmin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tsaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tsaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tsaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tsaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tsaq *TblSuperAdminQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tsaq.driver.Dialect())
	t1 := builder.Table(tblsuperadmin.Table)
	columns := tsaq.ctx.Fields
	if len(columns) == 0 {
		columns = tblsuperadmin.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tsaq.sql != nil {
		selector = tsaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tsaq.ctx.Unique != nil && *tsaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tsaq.predicates {
		p(selector)
	}
	for _, p := range tsaq.order {
		p(selector)
	}
	if offset := tsaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tsaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TblSuperAdminGroupBy is the group-by builder for TblSuperAdmin entities.
type TblSuperAdminGroupBy struct {
	selector
	build *TblSuperAdminQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tsagb *TblSuperAdminGroupBy) Aggregate(fns ...AggregateFunc) *TblSuperAdminGroupBy {
	tsagb.fns = append(tsagb.fns, fns...)
	return tsagb
}

// Scan applies the selector query and scans the result into the given value.
func (tsagb *TblSuperAdminGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tsagb.build.ctx, "GroupBy")
	if err := tsagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TblSuperAdminQuery, *TblSuperAdminGroupBy](ctx, tsagb.build, tsagb, tsagb.build.inters, v)
}

func (tsagb *TblSuperAdminGroupBy) sqlScan(ctx context.Context, root *TblSuperAdminQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tsagb.fns))
	for _, fn := range tsagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tsagb.flds)+len(tsagb.fns))
		for _, f := range *tsagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tsagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tsagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TblSuperAdminSelect is the builder for selecting fields of TblSuperAdmin entities.
type TblSuperAdminSelect struct {
	*TblSuperAdminQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tsas *TblSuperAdminSelect) Aggregate(fns ...AggregateFunc) *TblSuperAdminSelect {
	tsas.fns = append(tsas.fns, fns...)
	return tsas
}

// Scan applies the selector query and scans the result into the given value.
func (tsas *TblSuperAdminSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tsas.ctx, "Select")
	if err := tsas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TblSuperAdminQuery, *TblSuperAdminSelect](ctx, tsas.TblSuperAdminQuery, tsas, tsas.inters, v)
}

func (tsas *TblSuperAdminSelect) sqlScan(ctx context.Context, root *TblSuperAdminQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tsas.fns))
	for _, fn := range tsas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tsas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tsas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
