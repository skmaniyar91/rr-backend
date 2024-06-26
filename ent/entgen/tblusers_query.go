// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"rr-backend/ent/entgen/predicate"
	"rr-backend/ent/entgen/tblgarageowner"
	"rr-backend/ent/entgen/tblusers"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TblUSersQuery is the builder for querying TblUSers entities.
type TblUSersQuery struct {
	config
	ctx        *QueryContext
	order      []tblusers.OrderOption
	inters     []Interceptor
	predicates []predicate.TblUSers
	withOwner  *TblGarageOwnerQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TblUSersQuery builder.
func (tuq *TblUSersQuery) Where(ps ...predicate.TblUSers) *TblUSersQuery {
	tuq.predicates = append(tuq.predicates, ps...)
	return tuq
}

// Limit the number of records to be returned by this query.
func (tuq *TblUSersQuery) Limit(limit int) *TblUSersQuery {
	tuq.ctx.Limit = &limit
	return tuq
}

// Offset to start from.
func (tuq *TblUSersQuery) Offset(offset int) *TblUSersQuery {
	tuq.ctx.Offset = &offset
	return tuq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tuq *TblUSersQuery) Unique(unique bool) *TblUSersQuery {
	tuq.ctx.Unique = &unique
	return tuq
}

// Order specifies how the records should be ordered.
func (tuq *TblUSersQuery) Order(o ...tblusers.OrderOption) *TblUSersQuery {
	tuq.order = append(tuq.order, o...)
	return tuq
}

// QueryOwner chains the current query on the "Owner" edge.
func (tuq *TblUSersQuery) QueryOwner() *TblGarageOwnerQuery {
	query := (&TblGarageOwnerClient{config: tuq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tblusers.Table, tblusers.FieldID, selector),
			sqlgraph.To(tblgarageowner.Table, tblgarageowner.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, tblusers.OwnerTable, tblusers.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(tuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TblUSers entity from the query.
// Returns a *NotFoundError when no TblUSers was found.
func (tuq *TblUSersQuery) First(ctx context.Context) (*TblUSers, error) {
	nodes, err := tuq.Limit(1).All(setContextOp(ctx, tuq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tblusers.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tuq *TblUSersQuery) FirstX(ctx context.Context) *TblUSers {
	node, err := tuq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TblUSers ID from the query.
// Returns a *NotFoundError when no TblUSers ID was found.
func (tuq *TblUSersQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tuq.Limit(1).IDs(setContextOp(ctx, tuq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tblusers.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tuq *TblUSersQuery) FirstIDX(ctx context.Context) string {
	id, err := tuq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TblUSers entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TblUSers entity is found.
// Returns a *NotFoundError when no TblUSers entities are found.
func (tuq *TblUSersQuery) Only(ctx context.Context) (*TblUSers, error) {
	nodes, err := tuq.Limit(2).All(setContextOp(ctx, tuq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tblusers.Label}
	default:
		return nil, &NotSingularError{tblusers.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tuq *TblUSersQuery) OnlyX(ctx context.Context) *TblUSers {
	node, err := tuq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TblUSers ID in the query.
// Returns a *NotSingularError when more than one TblUSers ID is found.
// Returns a *NotFoundError when no entities are found.
func (tuq *TblUSersQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tuq.Limit(2).IDs(setContextOp(ctx, tuq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tblusers.Label}
	default:
		err = &NotSingularError{tblusers.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tuq *TblUSersQuery) OnlyIDX(ctx context.Context) string {
	id, err := tuq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TblUSersSlice.
func (tuq *TblUSersQuery) All(ctx context.Context) ([]*TblUSers, error) {
	ctx = setContextOp(ctx, tuq.ctx, "All")
	if err := tuq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TblUSers, *TblUSersQuery]()
	return withInterceptors[[]*TblUSers](ctx, tuq, qr, tuq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tuq *TblUSersQuery) AllX(ctx context.Context) []*TblUSers {
	nodes, err := tuq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TblUSers IDs.
func (tuq *TblUSersQuery) IDs(ctx context.Context) (ids []string, err error) {
	if tuq.ctx.Unique == nil && tuq.path != nil {
		tuq.Unique(true)
	}
	ctx = setContextOp(ctx, tuq.ctx, "IDs")
	if err = tuq.Select(tblusers.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tuq *TblUSersQuery) IDsX(ctx context.Context) []string {
	ids, err := tuq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tuq *TblUSersQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tuq.ctx, "Count")
	if err := tuq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tuq, querierCount[*TblUSersQuery](), tuq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tuq *TblUSersQuery) CountX(ctx context.Context) int {
	count, err := tuq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tuq *TblUSersQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tuq.ctx, "Exist")
	switch _, err := tuq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entgen: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tuq *TblUSersQuery) ExistX(ctx context.Context) bool {
	exist, err := tuq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TblUSersQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tuq *TblUSersQuery) Clone() *TblUSersQuery {
	if tuq == nil {
		return nil
	}
	return &TblUSersQuery{
		config:     tuq.config,
		ctx:        tuq.ctx.Clone(),
		order:      append([]tblusers.OrderOption{}, tuq.order...),
		inters:     append([]Interceptor{}, tuq.inters...),
		predicates: append([]predicate.TblUSers{}, tuq.predicates...),
		withOwner:  tuq.withOwner.Clone(),
		// clone intermediate query.
		sql:  tuq.sql.Clone(),
		path: tuq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "Owner" edge. The optional arguments are used to configure the query builder of the edge.
func (tuq *TblUSersQuery) WithOwner(opts ...func(*TblGarageOwnerQuery)) *TblUSersQuery {
	query := (&TblGarageOwnerClient{config: tuq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tuq.withOwner = query
	return tuq
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
//	client.TblUSers.Query().
//		GroupBy(tblusers.FieldCreatedBy).
//		Aggregate(entgen.Count()).
//		Scan(ctx, &v)
func (tuq *TblUSersQuery) GroupBy(field string, fields ...string) *TblUSersGroupBy {
	tuq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TblUSersGroupBy{build: tuq}
	grbuild.flds = &tuq.ctx.Fields
	grbuild.label = tblusers.Label
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
//	client.TblUSers.Query().
//		Select(tblusers.FieldCreatedBy).
//		Scan(ctx, &v)
func (tuq *TblUSersQuery) Select(fields ...string) *TblUSersSelect {
	tuq.ctx.Fields = append(tuq.ctx.Fields, fields...)
	sbuild := &TblUSersSelect{TblUSersQuery: tuq}
	sbuild.label = tblusers.Label
	sbuild.flds, sbuild.scan = &tuq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TblUSersSelect configured with the given aggregations.
func (tuq *TblUSersQuery) Aggregate(fns ...AggregateFunc) *TblUSersSelect {
	return tuq.Select().Aggregate(fns...)
}

func (tuq *TblUSersQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tuq.inters {
		if inter == nil {
			return fmt.Errorf("entgen: uninitialized interceptor (forgotten import entgen/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tuq); err != nil {
				return err
			}
		}
	}
	for _, f := range tuq.ctx.Fields {
		if !tblusers.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
		}
	}
	if tuq.path != nil {
		prev, err := tuq.path(ctx)
		if err != nil {
			return err
		}
		tuq.sql = prev
	}
	return nil
}

func (tuq *TblUSersQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TblUSers, error) {
	var (
		nodes       = []*TblUSers{}
		_spec       = tuq.querySpec()
		loadedTypes = [1]bool{
			tuq.withOwner != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TblUSers).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TblUSers{config: tuq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tuq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tuq.withOwner; query != nil {
		if err := tuq.loadOwner(ctx, query, nodes, nil,
			func(n *TblUSers, e *TblGarageOwner) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tuq *TblUSersQuery) loadOwner(ctx context.Context, query *TblGarageOwnerQuery, nodes []*TblUSers, init func(*TblUSers), assign func(*TblUSers, *TblGarageOwner)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*TblUSers)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(tblgarageowner.FieldUserIdUlid)
	}
	query.Where(predicate.TblGarageOwner(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(tblusers.OwnerColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserIdUlid
		if fk == nil {
			return fmt.Errorf(`foreign-key "UserId_ulid" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "UserId_ulid" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tuq *TblUSersQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tuq.querySpec()
	_spec.Node.Columns = tuq.ctx.Fields
	if len(tuq.ctx.Fields) > 0 {
		_spec.Unique = tuq.ctx.Unique != nil && *tuq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tuq.driver, _spec)
}

func (tuq *TblUSersQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tblusers.Table, tblusers.Columns, sqlgraph.NewFieldSpec(tblusers.FieldID, field.TypeString))
	_spec.From = tuq.sql
	if unique := tuq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tuq.path != nil {
		_spec.Unique = true
	}
	if fields := tuq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tblusers.FieldID)
		for i := range fields {
			if fields[i] != tblusers.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tuq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tuq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tuq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tuq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tuq *TblUSersQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tuq.driver.Dialect())
	t1 := builder.Table(tblusers.Table)
	columns := tuq.ctx.Fields
	if len(columns) == 0 {
		columns = tblusers.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tuq.sql != nil {
		selector = tuq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tuq.ctx.Unique != nil && *tuq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tuq.predicates {
		p(selector)
	}
	for _, p := range tuq.order {
		p(selector)
	}
	if offset := tuq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tuq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TblUSersGroupBy is the group-by builder for TblUSers entities.
type TblUSersGroupBy struct {
	selector
	build *TblUSersQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tugb *TblUSersGroupBy) Aggregate(fns ...AggregateFunc) *TblUSersGroupBy {
	tugb.fns = append(tugb.fns, fns...)
	return tugb
}

// Scan applies the selector query and scans the result into the given value.
func (tugb *TblUSersGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tugb.build.ctx, "GroupBy")
	if err := tugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TblUSersQuery, *TblUSersGroupBy](ctx, tugb.build, tugb, tugb.build.inters, v)
}

func (tugb *TblUSersGroupBy) sqlScan(ctx context.Context, root *TblUSersQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tugb.fns))
	for _, fn := range tugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tugb.flds)+len(tugb.fns))
		for _, f := range *tugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TblUSersSelect is the builder for selecting fields of TblUSers entities.
type TblUSersSelect struct {
	*TblUSersQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tus *TblUSersSelect) Aggregate(fns ...AggregateFunc) *TblUSersSelect {
	tus.fns = append(tus.fns, fns...)
	return tus
}

// Scan applies the selector query and scans the result into the given value.
func (tus *TblUSersSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tus.ctx, "Select")
	if err := tus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TblUSersQuery, *TblUSersSelect](ctx, tus.TblUSersQuery, tus, tus.inters, v)
}

func (tus *TblUSersSelect) sqlScan(ctx context.Context, root *TblUSersQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tus.fns))
	for _, fn := range tus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
