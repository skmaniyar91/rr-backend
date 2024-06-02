// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"context"
	"fmt"
	"math"
	"rr-backend/ent/entgen/predicate"
	"rr-backend/ent/entgen/tbladdress"
	"rr-backend/ent/entgen/tbldocument"
	"rr-backend/ent/entgen/tblenum"
	"rr-backend/ent/entgen/tblgarageowner"
	"rr-backend/ent/entgen/tblusers"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TblGarageOwnerQuery is the builder for querying TblGarageOwner entities.
type TblGarageOwnerQuery struct {
	config
	ctx             *QueryContext
	order           []tblgarageowner.OrderOption
	inters          []Interceptor
	predicates      []predicate.TblGarageOwner
	withUser        *TblUSersQuery
	withNameInitial *TblEnumQuery
	withOwnerPhoto  *TblDocumentQuery
	withAddress     *TblAddressQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TblGarageOwnerQuery builder.
func (tgoq *TblGarageOwnerQuery) Where(ps ...predicate.TblGarageOwner) *TblGarageOwnerQuery {
	tgoq.predicates = append(tgoq.predicates, ps...)
	return tgoq
}

// Limit the number of records to be returned by this query.
func (tgoq *TblGarageOwnerQuery) Limit(limit int) *TblGarageOwnerQuery {
	tgoq.ctx.Limit = &limit
	return tgoq
}

// Offset to start from.
func (tgoq *TblGarageOwnerQuery) Offset(offset int) *TblGarageOwnerQuery {
	tgoq.ctx.Offset = &offset
	return tgoq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tgoq *TblGarageOwnerQuery) Unique(unique bool) *TblGarageOwnerQuery {
	tgoq.ctx.Unique = &unique
	return tgoq
}

// Order specifies how the records should be ordered.
func (tgoq *TblGarageOwnerQuery) Order(o ...tblgarageowner.OrderOption) *TblGarageOwnerQuery {
	tgoq.order = append(tgoq.order, o...)
	return tgoq
}

// QueryUser chains the current query on the "User" edge.
func (tgoq *TblGarageOwnerQuery) QueryUser() *TblUSersQuery {
	query := (&TblUSersClient{config: tgoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tgoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tgoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tblgarageowner.Table, tblgarageowner.FieldID, selector),
			sqlgraph.To(tblusers.Table, tblusers.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, tblgarageowner.UserTable, tblgarageowner.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(tgoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNameInitial chains the current query on the "NameInitial" edge.
func (tgoq *TblGarageOwnerQuery) QueryNameInitial() *TblEnumQuery {
	query := (&TblEnumClient{config: tgoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tgoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tgoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tblgarageowner.Table, tblgarageowner.FieldID, selector),
			sqlgraph.To(tblenum.Table, tblenum.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, tblgarageowner.NameInitialTable, tblgarageowner.NameInitialColumn),
		)
		fromU = sqlgraph.SetNeighbors(tgoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOwnerPhoto chains the current query on the "OwnerPhoto" edge.
func (tgoq *TblGarageOwnerQuery) QueryOwnerPhoto() *TblDocumentQuery {
	query := (&TblDocumentClient{config: tgoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tgoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tgoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tblgarageowner.Table, tblgarageowner.FieldID, selector),
			sqlgraph.To(tbldocument.Table, tbldocument.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, tblgarageowner.OwnerPhotoTable, tblgarageowner.OwnerPhotoColumn),
		)
		fromU = sqlgraph.SetNeighbors(tgoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAddress chains the current query on the "Address" edge.
func (tgoq *TblGarageOwnerQuery) QueryAddress() *TblAddressQuery {
	query := (&TblAddressClient{config: tgoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tgoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tgoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tblgarageowner.Table, tblgarageowner.FieldID, selector),
			sqlgraph.To(tbladdress.Table, tbladdress.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, tblgarageowner.AddressTable, tblgarageowner.AddressColumn),
		)
		fromU = sqlgraph.SetNeighbors(tgoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TblGarageOwner entity from the query.
// Returns a *NotFoundError when no TblGarageOwner was found.
func (tgoq *TblGarageOwnerQuery) First(ctx context.Context) (*TblGarageOwner, error) {
	nodes, err := tgoq.Limit(1).All(setContextOp(ctx, tgoq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tblgarageowner.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tgoq *TblGarageOwnerQuery) FirstX(ctx context.Context) *TblGarageOwner {
	node, err := tgoq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TblGarageOwner ID from the query.
// Returns a *NotFoundError when no TblGarageOwner ID was found.
func (tgoq *TblGarageOwnerQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tgoq.Limit(1).IDs(setContextOp(ctx, tgoq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tblgarageowner.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tgoq *TblGarageOwnerQuery) FirstIDX(ctx context.Context) string {
	id, err := tgoq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TblGarageOwner entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TblGarageOwner entity is found.
// Returns a *NotFoundError when no TblGarageOwner entities are found.
func (tgoq *TblGarageOwnerQuery) Only(ctx context.Context) (*TblGarageOwner, error) {
	nodes, err := tgoq.Limit(2).All(setContextOp(ctx, tgoq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tblgarageowner.Label}
	default:
		return nil, &NotSingularError{tblgarageowner.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tgoq *TblGarageOwnerQuery) OnlyX(ctx context.Context) *TblGarageOwner {
	node, err := tgoq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TblGarageOwner ID in the query.
// Returns a *NotSingularError when more than one TblGarageOwner ID is found.
// Returns a *NotFoundError when no entities are found.
func (tgoq *TblGarageOwnerQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tgoq.Limit(2).IDs(setContextOp(ctx, tgoq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tblgarageowner.Label}
	default:
		err = &NotSingularError{tblgarageowner.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tgoq *TblGarageOwnerQuery) OnlyIDX(ctx context.Context) string {
	id, err := tgoq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TblGarageOwners.
func (tgoq *TblGarageOwnerQuery) All(ctx context.Context) ([]*TblGarageOwner, error) {
	ctx = setContextOp(ctx, tgoq.ctx, "All")
	if err := tgoq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TblGarageOwner, *TblGarageOwnerQuery]()
	return withInterceptors[[]*TblGarageOwner](ctx, tgoq, qr, tgoq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tgoq *TblGarageOwnerQuery) AllX(ctx context.Context) []*TblGarageOwner {
	nodes, err := tgoq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TblGarageOwner IDs.
func (tgoq *TblGarageOwnerQuery) IDs(ctx context.Context) (ids []string, err error) {
	if tgoq.ctx.Unique == nil && tgoq.path != nil {
		tgoq.Unique(true)
	}
	ctx = setContextOp(ctx, tgoq.ctx, "IDs")
	if err = tgoq.Select(tblgarageowner.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tgoq *TblGarageOwnerQuery) IDsX(ctx context.Context) []string {
	ids, err := tgoq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tgoq *TblGarageOwnerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tgoq.ctx, "Count")
	if err := tgoq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tgoq, querierCount[*TblGarageOwnerQuery](), tgoq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tgoq *TblGarageOwnerQuery) CountX(ctx context.Context) int {
	count, err := tgoq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tgoq *TblGarageOwnerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tgoq.ctx, "Exist")
	switch _, err := tgoq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("entgen: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tgoq *TblGarageOwnerQuery) ExistX(ctx context.Context) bool {
	exist, err := tgoq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TblGarageOwnerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tgoq *TblGarageOwnerQuery) Clone() *TblGarageOwnerQuery {
	if tgoq == nil {
		return nil
	}
	return &TblGarageOwnerQuery{
		config:          tgoq.config,
		ctx:             tgoq.ctx.Clone(),
		order:           append([]tblgarageowner.OrderOption{}, tgoq.order...),
		inters:          append([]Interceptor{}, tgoq.inters...),
		predicates:      append([]predicate.TblGarageOwner{}, tgoq.predicates...),
		withUser:        tgoq.withUser.Clone(),
		withNameInitial: tgoq.withNameInitial.Clone(),
		withOwnerPhoto:  tgoq.withOwnerPhoto.Clone(),
		withAddress:     tgoq.withAddress.Clone(),
		// clone intermediate query.
		sql:  tgoq.sql.Clone(),
		path: tgoq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "User" edge. The optional arguments are used to configure the query builder of the edge.
func (tgoq *TblGarageOwnerQuery) WithUser(opts ...func(*TblUSersQuery)) *TblGarageOwnerQuery {
	query := (&TblUSersClient{config: tgoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tgoq.withUser = query
	return tgoq
}

// WithNameInitial tells the query-builder to eager-load the nodes that are connected to
// the "NameInitial" edge. The optional arguments are used to configure the query builder of the edge.
func (tgoq *TblGarageOwnerQuery) WithNameInitial(opts ...func(*TblEnumQuery)) *TblGarageOwnerQuery {
	query := (&TblEnumClient{config: tgoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tgoq.withNameInitial = query
	return tgoq
}

// WithOwnerPhoto tells the query-builder to eager-load the nodes that are connected to
// the "OwnerPhoto" edge. The optional arguments are used to configure the query builder of the edge.
func (tgoq *TblGarageOwnerQuery) WithOwnerPhoto(opts ...func(*TblDocumentQuery)) *TblGarageOwnerQuery {
	query := (&TblDocumentClient{config: tgoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tgoq.withOwnerPhoto = query
	return tgoq
}

// WithAddress tells the query-builder to eager-load the nodes that are connected to
// the "Address" edge. The optional arguments are used to configure the query builder of the edge.
func (tgoq *TblGarageOwnerQuery) WithAddress(opts ...func(*TblAddressQuery)) *TblGarageOwnerQuery {
	query := (&TblAddressClient{config: tgoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tgoq.withAddress = query
	return tgoq
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
//	client.TblGarageOwner.Query().
//		GroupBy(tblgarageowner.FieldCreatedBy).
//		Aggregate(entgen.Count()).
//		Scan(ctx, &v)
func (tgoq *TblGarageOwnerQuery) GroupBy(field string, fields ...string) *TblGarageOwnerGroupBy {
	tgoq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TblGarageOwnerGroupBy{build: tgoq}
	grbuild.flds = &tgoq.ctx.Fields
	grbuild.label = tblgarageowner.Label
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
//	client.TblGarageOwner.Query().
//		Select(tblgarageowner.FieldCreatedBy).
//		Scan(ctx, &v)
func (tgoq *TblGarageOwnerQuery) Select(fields ...string) *TblGarageOwnerSelect {
	tgoq.ctx.Fields = append(tgoq.ctx.Fields, fields...)
	sbuild := &TblGarageOwnerSelect{TblGarageOwnerQuery: tgoq}
	sbuild.label = tblgarageowner.Label
	sbuild.flds, sbuild.scan = &tgoq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TblGarageOwnerSelect configured with the given aggregations.
func (tgoq *TblGarageOwnerQuery) Aggregate(fns ...AggregateFunc) *TblGarageOwnerSelect {
	return tgoq.Select().Aggregate(fns...)
}

func (tgoq *TblGarageOwnerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tgoq.inters {
		if inter == nil {
			return fmt.Errorf("entgen: uninitialized interceptor (forgotten import entgen/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tgoq); err != nil {
				return err
			}
		}
	}
	for _, f := range tgoq.ctx.Fields {
		if !tblgarageowner.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("entgen: invalid field %q for query", f)}
		}
	}
	if tgoq.path != nil {
		prev, err := tgoq.path(ctx)
		if err != nil {
			return err
		}
		tgoq.sql = prev
	}
	return nil
}

func (tgoq *TblGarageOwnerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TblGarageOwner, error) {
	var (
		nodes       = []*TblGarageOwner{}
		_spec       = tgoq.querySpec()
		loadedTypes = [4]bool{
			tgoq.withUser != nil,
			tgoq.withNameInitial != nil,
			tgoq.withOwnerPhoto != nil,
			tgoq.withAddress != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TblGarageOwner).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TblGarageOwner{config: tgoq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tgoq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tgoq.withUser; query != nil {
		if err := tgoq.loadUser(ctx, query, nodes, nil,
			func(n *TblGarageOwner, e *TblUSers) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := tgoq.withNameInitial; query != nil {
		if err := tgoq.loadNameInitial(ctx, query, nodes, nil,
			func(n *TblGarageOwner, e *TblEnum) { n.Edges.NameInitial = e }); err != nil {
			return nil, err
		}
	}
	if query := tgoq.withOwnerPhoto; query != nil {
		if err := tgoq.loadOwnerPhoto(ctx, query, nodes, nil,
			func(n *TblGarageOwner, e *TblDocument) { n.Edges.OwnerPhoto = e }); err != nil {
			return nil, err
		}
	}
	if query := tgoq.withAddress; query != nil {
		if err := tgoq.loadAddress(ctx, query, nodes, nil,
			func(n *TblGarageOwner, e *TblAddress) { n.Edges.Address = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tgoq *TblGarageOwnerQuery) loadUser(ctx context.Context, query *TblUSersQuery, nodes []*TblGarageOwner, init func(*TblGarageOwner), assign func(*TblGarageOwner, *TblUSers)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*TblGarageOwner)
	for i := range nodes {
		if nodes[i].UserIdUlid == nil {
			continue
		}
		fk := *nodes[i].UserIdUlid
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(tblusers.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "UserId_ulid" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tgoq *TblGarageOwnerQuery) loadNameInitial(ctx context.Context, query *TblEnumQuery, nodes []*TblGarageOwner, init func(*TblGarageOwner), assign func(*TblGarageOwner, *TblEnum)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TblGarageOwner)
	for i := range nodes {
		if nodes[i].Initial == nil {
			continue
		}
		fk := *nodes[i].Initial
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(tblenum.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "Initial" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tgoq *TblGarageOwnerQuery) loadOwnerPhoto(ctx context.Context, query *TblDocumentQuery, nodes []*TblGarageOwner, init func(*TblGarageOwner), assign func(*TblGarageOwner, *TblDocument)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*TblGarageOwner)
	for i := range nodes {
		if nodes[i].PhotoIdUlid == nil {
			continue
		}
		fk := *nodes[i].PhotoIdUlid
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(tbldocument.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "PhotoId_ulid" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tgoq *TblGarageOwnerQuery) loadAddress(ctx context.Context, query *TblAddressQuery, nodes []*TblGarageOwner, init func(*TblGarageOwner), assign func(*TblGarageOwner, *TblAddress)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*TblGarageOwner)
	for i := range nodes {
		fk := nodes[i].AddressIdUlid
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(tbladdress.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "AddressId_ulid" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tgoq *TblGarageOwnerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tgoq.querySpec()
	_spec.Node.Columns = tgoq.ctx.Fields
	if len(tgoq.ctx.Fields) > 0 {
		_spec.Unique = tgoq.ctx.Unique != nil && *tgoq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tgoq.driver, _spec)
}

func (tgoq *TblGarageOwnerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tblgarageowner.Table, tblgarageowner.Columns, sqlgraph.NewFieldSpec(tblgarageowner.FieldID, field.TypeString))
	_spec.From = tgoq.sql
	if unique := tgoq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tgoq.path != nil {
		_spec.Unique = true
	}
	if fields := tgoq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tblgarageowner.FieldID)
		for i := range fields {
			if fields[i] != tblgarageowner.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if tgoq.withUser != nil {
			_spec.Node.AddColumnOnce(tblgarageowner.FieldUserIdUlid)
		}
		if tgoq.withNameInitial != nil {
			_spec.Node.AddColumnOnce(tblgarageowner.FieldInitial)
		}
		if tgoq.withOwnerPhoto != nil {
			_spec.Node.AddColumnOnce(tblgarageowner.FieldPhotoIdUlid)
		}
		if tgoq.withAddress != nil {
			_spec.Node.AddColumnOnce(tblgarageowner.FieldAddressIdUlid)
		}
	}
	if ps := tgoq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tgoq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tgoq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tgoq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tgoq *TblGarageOwnerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tgoq.driver.Dialect())
	t1 := builder.Table(tblgarageowner.Table)
	columns := tgoq.ctx.Fields
	if len(columns) == 0 {
		columns = tblgarageowner.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tgoq.sql != nil {
		selector = tgoq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tgoq.ctx.Unique != nil && *tgoq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tgoq.predicates {
		p(selector)
	}
	for _, p := range tgoq.order {
		p(selector)
	}
	if offset := tgoq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tgoq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TblGarageOwnerGroupBy is the group-by builder for TblGarageOwner entities.
type TblGarageOwnerGroupBy struct {
	selector
	build *TblGarageOwnerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgogb *TblGarageOwnerGroupBy) Aggregate(fns ...AggregateFunc) *TblGarageOwnerGroupBy {
	tgogb.fns = append(tgogb.fns, fns...)
	return tgogb
}

// Scan applies the selector query and scans the result into the given value.
func (tgogb *TblGarageOwnerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgogb.build.ctx, "GroupBy")
	if err := tgogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TblGarageOwnerQuery, *TblGarageOwnerGroupBy](ctx, tgogb.build, tgogb, tgogb.build.inters, v)
}

func (tgogb *TblGarageOwnerGroupBy) sqlScan(ctx context.Context, root *TblGarageOwnerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgogb.fns))
	for _, fn := range tgogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgogb.flds)+len(tgogb.fns))
		for _, f := range *tgogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TblGarageOwnerSelect is the builder for selecting fields of TblGarageOwner entities.
type TblGarageOwnerSelect struct {
	*TblGarageOwnerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tgos *TblGarageOwnerSelect) Aggregate(fns ...AggregateFunc) *TblGarageOwnerSelect {
	tgos.fns = append(tgos.fns, fns...)
	return tgos
}

// Scan applies the selector query and scans the result into the given value.
func (tgos *TblGarageOwnerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgos.ctx, "Select")
	if err := tgos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TblGarageOwnerQuery, *TblGarageOwnerSelect](ctx, tgos.TblGarageOwnerQuery, tgos, tgos.inters, v)
}

func (tgos *TblGarageOwnerSelect) sqlScan(ctx context.Context, root *TblGarageOwnerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tgos.fns))
	for _, fn := range tgos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tgos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}