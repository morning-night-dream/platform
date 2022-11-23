// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/app/db/ent/auth"
	"github.com/morning-night-dream/platform/app/db/ent/predicate"
	"github.com/morning-night-dream/platform/app/db/ent/user"
)

// AuthQuery is the builder for querying Auth entities.
type AuthQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Auth
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthQuery builder.
func (aq *AuthQuery) Where(ps ...predicate.Auth) *AuthQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit adds a limit step to the query.
func (aq *AuthQuery) Limit(limit int) *AuthQuery {
	aq.limit = &limit
	return aq
}

// Offset adds an offset step to the query.
func (aq *AuthQuery) Offset(offset int) *AuthQuery {
	aq.offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AuthQuery) Unique(unique bool) *AuthQuery {
	aq.unique = &unique
	return aq
}

// Order adds an order step to the query.
func (aq *AuthQuery) Order(o ...OrderFunc) *AuthQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryUser chains the current query on the "user" edge.
func (aq *AuthQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: aq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(auth.Table, auth.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, auth.UserTable, auth.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Auth entity from the query.
// Returns a *NotFoundError when no Auth was found.
func (aq *AuthQuery) First(ctx context.Context) (*Auth, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{auth.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AuthQuery) FirstX(ctx context.Context) *Auth {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Auth ID from the query.
// Returns a *NotFoundError when no Auth ID was found.
func (aq *AuthQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{auth.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AuthQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Auth entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Auth entity is found.
// Returns a *NotFoundError when no Auth entities are found.
func (aq *AuthQuery) Only(ctx context.Context) (*Auth, error) {
	nodes, err := aq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{auth.Label}
	default:
		return nil, &NotSingularError{auth.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AuthQuery) OnlyX(ctx context.Context) *Auth {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Auth ID in the query.
// Returns a *NotSingularError when more than one Auth ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AuthQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{auth.Label}
	default:
		err = &NotSingularError{auth.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AuthQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Auths.
func (aq *AuthQuery) All(ctx context.Context) ([]*Auth, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aq *AuthQuery) AllX(ctx context.Context) []*Auth {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Auth IDs.
func (aq *AuthQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := aq.Select(auth.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AuthQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AuthQuery) Count(ctx context.Context) (int, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AuthQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AuthQuery) Exist(ctx context.Context) (bool, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AuthQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AuthQuery) Clone() *AuthQuery {
	if aq == nil {
		return nil
	}
	return &AuthQuery{
		config:     aq.config,
		limit:      aq.limit,
		offset:     aq.offset,
		order:      append([]OrderFunc{}, aq.order...),
		predicates: append([]predicate.Auth{}, aq.predicates...),
		withUser:   aq.withUser.Clone(),
		// clone intermediate query.
		sql:    aq.sql.Clone(),
		path:   aq.path,
		unique: aq.unique,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AuthQuery) WithUser(opts ...func(*UserQuery)) *AuthQuery {
	query := &UserQuery{config: aq.config}
	for _, opt := range opts {
		opt(query)
	}
	aq.withUser = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Auth.Query().
//		GroupBy(auth.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AuthQuery) GroupBy(field string, fields ...string) *AuthGroupBy {
	grbuild := &AuthGroupBy{config: aq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aq.sqlQuery(ctx), nil
	}
	grbuild.label = auth.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//	}
//
//	client.Auth.Query().
//		Select(auth.FieldUserID).
//		Scan(ctx, &v)
func (aq *AuthQuery) Select(fields ...string) *AuthSelect {
	aq.fields = append(aq.fields, fields...)
	selbuild := &AuthSelect{AuthQuery: aq}
	selbuild.label = auth.Label
	selbuild.flds, selbuild.scan = &aq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a AuthSelect configured with the given aggregations.
func (aq *AuthQuery) Aggregate(fns ...AggregateFunc) *AuthSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AuthQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aq.fields {
		if !auth.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AuthQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Auth, error) {
	var (
		nodes       = []*Auth{}
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Auth).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Auth{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withUser; query != nil {
		if err := aq.loadUser(ctx, query, nodes, nil,
			func(n *Auth, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AuthQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Auth, init func(*Auth), assign func(*Auth, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Auth)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aq *AuthQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.fields
	if len(aq.fields) > 0 {
		_spec.Unique = aq.unique != nil && *aq.unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AuthQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (aq *AuthQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   auth.Table,
			Columns: auth.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: auth.FieldID,
			},
		},
		From:   aq.sql,
		Unique: true,
	}
	if unique := aq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, auth.FieldID)
		for i := range fields {
			if fields[i] != auth.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AuthQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(auth.Table)
	columns := aq.fields
	if len(columns) == 0 {
		columns = auth.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.unique != nil && *aq.unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthGroupBy is the group-by builder for Auth entities.
type AuthGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AuthGroupBy) Aggregate(fns ...AggregateFunc) *AuthGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the group-by query and scans the result into the given value.
func (agb *AuthGroupBy) Scan(ctx context.Context, v any) error {
	query, err := agb.path(ctx)
	if err != nil {
		return err
	}
	agb.sql = query
	return agb.sqlScan(ctx, v)
}

func (agb *AuthGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range agb.fields {
		if !auth.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agb *AuthGroupBy) sqlQuery() *sql.Selector {
	selector := agb.sql.Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agb.fields)+len(agb.fns))
		for _, f := range agb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agb.fields...)...)
}

// AuthSelect is the builder for selecting fields of Auth entities.
type AuthSelect struct {
	*AuthQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AuthSelect) Aggregate(fns ...AggregateFunc) *AuthSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AuthSelect) Scan(ctx context.Context, v any) error {
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	as.sql = as.AuthQuery.sqlQuery(ctx)
	return as.sqlScan(ctx, v)
}

func (as *AuthSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(as.sql))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		as.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		as.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := as.sql.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
