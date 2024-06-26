// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"rr-backend/ent/entgen/predicate"
	"rr-backend/ent/entgen/tblenum"
	"rr-backend/ent/entgen/tblgarageowner"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TblEnumQuery is the builder for querying TblEnum entities.
type TblEnumQuery struct {
	config
	ctx             *QueryContext
	order           []tblenum.OrderOption
	inters          []Interceptor
	predicates      []predicate.TblEnum
	withInitialEnum *TblGarageOwnerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TblEnumQuery builder.
func (teq *TblEnumQuery) Where(ps ...predicate.TblEnum) *TblEnumQuery {
	teq.predicates = append(teq.predicates, ps...)
	return teq
}

// Limit the number of records to be returned by this query.
func (teq *TblEnumQuery) Limit(limit int) *TblEnumQuery {
	teq.ctx.Limit = &limit
	return teq
}

// Offset to start from.
func (teq *TblEnumQuery) Offset(offset int) *TblEnumQuery {
	teq.ctx.Offset = &offset
	return teq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (teq *TblEnumQuery) Unique(unique bool) *TblEnumQuery {
	teq.ctx.Unique = &unique
	return teq
}

// Order specifies how the records should be ordered.
func (teq *TblEnumQuery) Order(o ...tblenum.OrderOption) *TblEnumQuery {
	teq.order = append(teq.order, o...)
	return teq
}

// QueryInitialEnum chains the current query on the "InitialEnum" edge.
func (teq *TblEnumQuery) QueryInitialEnum() *TblGarageOwnerQuery {
	query := (&TblGarageOwnerClient{config: teq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := teq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := teq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tblenum.Table, tblenum.FieldID, selector),
			sqlgraph.To(tblgarageowner.Table, tblgarageowner.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, tblenum.InitialEnumTable, tblenum.InitialEnumColumn),
		)
		fromU = sqlgraph.SetNeighbors(teq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TblEnum entity from the query.
// Returns a *NotFoundError when no TblEnum was found.
func (teq *TblEnumQuery) First(ctx context.Context) (*TblEnum, error) {
	nodes, err := teq.Limit(1).All(setContextOp(ctx, teq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tblenum.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (teq *TblEnumQuery) FirstX(ctx context.Context) *TblEnum {
	node, err := teq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TblEnum ID from the query.
// Returns a *NotFoundError when no TblEnum ID was found.
func (teq *TblEnumQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = teq.Limit(1).IDs(setContextOp(ctx, teq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tblenum.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (teq *TblEnumQuery) FirstIDX(ctx context.Context) int {
	id, err := teq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TblEnum entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TblEnum entity is found.
// Returns a *NotFoundError when no TblEnum entities are found.
func (teq *TblEnumQuery) Only(ctx context.Context) (*TblEnum, error) {
	nodes, err := teq.Limit(2).All(setContextOp(ctx, teq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tblenum.Label}
	default:
		return nil, &NotSingularError{tblenum.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (teq *TblEnumQuery) OnlyX(ctx context.Context) *TblEnum {
	node, err := teq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TblEnum ID in the query.
// Returns a *NotSingularError when more than one TblEnum ID is found.
// Returns a *NotFoundError when no entities are found.
func (teq *TblEnumQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = teq.Limit(2).IDs(setContextOp(ctx, teq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tblenum.Label}
	default:
		err = &NotSingularError{tblenum.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (teq *TblEnumQuery) OnlyIDX(ctx context.Context) int {
	id, err := teq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TblEnums.
func (teq *TblEnumQuery) All(ctx context.Context) ([]*TblEnum, error) {
	ctx = setContextOp(ctx, teq.ctx, "All")
	if err := teq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TblEnum, *TblEnumQuery]()
	return withInterceptors[[]*TblEnum](ctx, teq, qr, teq.inters)
}

// AllX is like All, but panics if an error occurs.
func (teq *TblEnumQuery) AllX(ctx context.Context) []*TblEnum {
	nodes, err := teq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TblEnum IDs.
func (teq *TblEnumQuery) IDs(ctx context.Context) (ids []int, err error) {
	if teq.ctx.Unique == nil && teq.path != nil {
		teq.Unique(true)
	}
	ctx = setContextOp(ctx, teq.ctx, "IDs")
	if err = teq.Select(tblenum.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (teq *TblEnumQuery) IDsX(ctx context.Context) []int {
	ids, err := teq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (teq *TblEnumQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, teq.ctx, "Count")
	if err := teq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, teq, querierCount[*TblEnumQuery](), teq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (teq *TblEnumQuery) CountX(ctx context.Context) int {
	count, err := teq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (teq *TblEnumQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, teq.ctx, "Exist")
	switch _, err := teq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entgen: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (teq *TblEnumQuery) ExistX(ctx context.Context) bool {
	exist, err := teq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TblEnumQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (teq *TblEnumQuery) Clone() *TblEnumQuery {
	if teq == nil {
		return nil
	}
	return &TblEnumQuery{
		config:          teq.config,
		ctx:             teq.ctx.Clone(),
		order:           append([]tblenum.OrderOption{}, teq.order...),
		inters:          append([]Interceptor{}, teq.inters...),
		predicates:      append([]predicate.TblEnum{}, teq.predicates...),
		withInitialEnum: teq.withInitialEnum.Clone(),
		// clone intermediate query.
		sql:  teq.sql.Clone(),
		path: teq.path,
	}
}

// WithInitialEnum tells the query-builder to eager-load the nodes that are connected to
// the "InitialEnum" edge. The optional arguments are used to configure the query builder of the edge.
func (teq *TblEnumQuery) WithInitialEnum(opts ...func(*TblGarageOwnerQuery)) *TblEnumQuery {
	query := (&TblGarageOwnerClient{config: teq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	teq.withInitialEnum = query
	return teq
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
//	client.TblEnum.Query().
//		GroupBy(tblenum.FieldCreatedBy).
//		Aggregate(entgen.Count()).
//		Scan(ctx, &v)
func (teq *TblEnumQuery) GroupBy(field string, fields ...string) *TblEnumGroupBy {
	teq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TblEnumGroupBy{build: teq}
	grbuild.flds = &teq.ctx.Fields
	grbuild.label = tblenum.Label
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
//	client.TblEnum.Query().
//		Select(tblenum.FieldCreatedBy).
//		Scan(ctx, &v)
func (teq *TblEnumQuery) Select(fields ...string) *TblEnumSelect {
	teq.ctx.Fields = append(teq.ctx.Fields, fields...)
	sbuild := &TblEnumSelect{TblEnumQuery: teq}
	sbuild.label = tblenum.Label
	sbuild.flds, sbuild.scan = &teq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TblEnumSelect configured with the given aggregations.
func (teq *TblEnumQuery) Aggregate(fns ...AggregateFunc) *TblEnumSelect {
	return teq.Select().Aggregate(fns...)
}

func (teq *TblEnumQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range teq.inters {
		if inter == nil {
			return fmt.Errorf("entgen: uninitialized interceptor (forgotten import entgen/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, teq); err != nil {
				return err
			}
		}
	}
	for _, f := range teq.ctx.Fields {
		if !tblenum.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
		}
	}
	if teq.path != nil {
		prev, err := teq.path(ctx)
		if err != nil {
			return err
		}
		teq.sql = prev
	}
	return nil
}

func (teq *TblEnumQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TblEnum, error) {
	var (
		nodes       = []*TblEnum{}
		_spec       = teq.querySpec()
		loadedTypes = [1]bool{
			teq.withInitialEnum != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TblEnum).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TblEnum{config: teq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, teq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := teq.withInitialEnum; query != nil {
		if err := teq.loadInitialEnum(ctx, query, nodes, nil,
			func(n *TblEnum, e *TblGarageOwner) { n.Edges.InitialEnum = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (teq *TblEnumQuery) loadInitialEnum(ctx context.Context, query *TblGarageOwnerQuery, nodes []*TblEnum, init func(*TblEnum), assign func(*TblEnum, *TblGarageOwner)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*TblEnum)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(tblgarageowner.FieldInitial)
	}
	query.Where(predicate.TblGarageOwner(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(tblenum.InitialEnumColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.Initial
		if fk == nil {
			return fmt.Errorf(`foreign-key "Initial" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "Initial" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (teq *TblEnumQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := teq.querySpec()
	_spec.Node.Columns = teq.ctx.Fields
	if len(teq.ctx.Fields) > 0 {
		_spec.Unique = teq.ctx.Unique != nil && *teq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, teq.driver, _spec)
}

func (teq *TblEnumQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tblenum.Table, tblenum.Columns, sqlgraph.NewFieldSpec(tblenum.FieldID, field.TypeInt))
	_spec.From = teq.sql
	if unique := teq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if teq.path != nil {
		_spec.Unique = true
	}
	if fields := teq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tblenum.FieldID)
		for i := range fields {
			if fields[i] != tblenum.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := teq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := teq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := teq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := teq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (teq *TblEnumQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(teq.driver.Dialect())
	t1 := builder.Table(tblenum.Table)
	columns := teq.ctx.Fields
	if len(columns) == 0 {
		columns = tblenum.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if teq.sql != nil {
		selector = teq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if teq.ctx.Unique != nil && *teq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range teq.predicates {
		p(selector)
	}
	for _, p := range teq.order {
		p(selector)
	}
	if offset := teq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := teq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TblEnumGroupBy is the group-by builder for TblEnum entities.
type TblEnumGroupBy struct {
	selector
	build *TblEnumQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tegb *TblEnumGroupBy) Aggregate(fns ...AggregateFunc) *TblEnumGroupBy {
	tegb.fns = append(tegb.fns, fns...)
	return tegb
}

// Scan applies the selector query and scans the result into the given value.
func (tegb *TblEnumGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tegb.build.ctx, "GroupBy")
	if err := tegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TblEnumQuery, *TblEnumGroupBy](ctx, tegb.build, tegb, tegb.build.inters, v)
}

func (tegb *TblEnumGroupBy) sqlScan(ctx context.Context, root *TblEnumQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tegb.fns))
	for _, fn := range tegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tegb.flds)+len(tegb.fns))
		for _, f := range *tegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TblEnumSelect is the builder for selecting fields of TblEnum entities.
type TblEnumSelect struct {
	*TblEnumQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tes *TblEnumSelect) Aggregate(fns ...AggregateFunc) *TblEnumSelect {
	tes.fns = append(tes.fns, fns...)
	return tes
}

// Scan applies the selector query and scans the result into the given value.
func (tes *TblEnumSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tes.ctx, "Select")
	if err := tes.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TblEnumQuery, *TblEnumSelect](ctx, tes.TblEnumQuery, tes, tes.inters, v)
}

func (tes *TblEnumSelect) sqlScan(ctx context.Context, root *TblEnumQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tes.fns))
	for _, fn := range tes.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
