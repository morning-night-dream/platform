// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/morning-night-dream/platform/pkg/ent/article"
	"github.com/morning-night-dream/platform/pkg/ent/articletag"
	"github.com/morning-night-dream/platform/pkg/ent/predicate"
)

// ArticleTagUpdate is the builder for updating ArticleTag entities.
type ArticleTagUpdate struct {
	config
	hooks    []Hook
	mutation *ArticleTagMutation
}

// Where appends a list predicates to the ArticleTagUpdate builder.
func (atu *ArticleTagUpdate) Where(ps ...predicate.ArticleTag) *ArticleTagUpdate {
	atu.mutation.Where(ps...)
	return atu
}

// SetTag sets the "tag" field.
func (atu *ArticleTagUpdate) SetTag(s string) *ArticleTagUpdate {
	atu.mutation.SetTag(s)
	return atu
}

// SetArticleID sets the "article_id" field.
func (atu *ArticleTagUpdate) SetArticleID(u uuid.UUID) *ArticleTagUpdate {
	atu.mutation.SetArticleID(u)
	return atu
}

// SetCreatedAt sets the "created_at" field.
func (atu *ArticleTagUpdate) SetCreatedAt(t time.Time) *ArticleTagUpdate {
	atu.mutation.SetCreatedAt(t)
	return atu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (atu *ArticleTagUpdate) SetNillableCreatedAt(t *time.Time) *ArticleTagUpdate {
	if t != nil {
		atu.SetCreatedAt(*t)
	}
	return atu
}

// SetUpdatedAt sets the "updated_at" field.
func (atu *ArticleTagUpdate) SetUpdatedAt(t time.Time) *ArticleTagUpdate {
	atu.mutation.SetUpdatedAt(t)
	return atu
}

// SetArticle sets the "article" edge to the Article entity.
func (atu *ArticleTagUpdate) SetArticle(a *Article) *ArticleTagUpdate {
	return atu.SetArticleID(a.ID)
}

// Mutation returns the ArticleTagMutation object of the builder.
func (atu *ArticleTagUpdate) Mutation() *ArticleTagMutation {
	return atu.mutation
}

// ClearArticle clears the "article" edge to the Article entity.
func (atu *ArticleTagUpdate) ClearArticle() *ArticleTagUpdate {
	atu.mutation.ClearArticle()
	return atu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (atu *ArticleTagUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	atu.defaults()
	if len(atu.hooks) == 0 {
		if err = atu.check(); err != nil {
			return 0, err
		}
		affected, err = atu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArticleTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atu.check(); err != nil {
				return 0, err
			}
			atu.mutation = mutation
			affected, err = atu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(atu.hooks) - 1; i >= 0; i-- {
			if atu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (atu *ArticleTagUpdate) SaveX(ctx context.Context) int {
	affected, err := atu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (atu *ArticleTagUpdate) Exec(ctx context.Context) error {
	_, err := atu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atu *ArticleTagUpdate) ExecX(ctx context.Context) {
	if err := atu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atu *ArticleTagUpdate) defaults() {
	if _, ok := atu.mutation.UpdatedAt(); !ok {
		v := articletag.UpdateDefaultUpdatedAt()
		atu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atu *ArticleTagUpdate) check() error {
	if _, ok := atu.mutation.ArticleID(); atu.mutation.ArticleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArticleTag.article"`)
	}
	return nil
}

func (atu *ArticleTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   articletag.Table,
			Columns: articletag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: articletag.FieldID,
			},
		},
	}
	if ps := atu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atu.mutation.Tag(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: articletag.FieldTag,
		})
	}
	if value, ok := atu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: articletag.FieldCreatedAt,
		})
	}
	if value, ok := atu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: articletag.FieldUpdatedAt,
		})
	}
	if atu.mutation.ArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   articletag.ArticleTable,
			Columns: []string{articletag.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: article.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atu.mutation.ArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   articletag.ArticleTable,
			Columns: []string{articletag.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, atu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{articletag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ArticleTagUpdateOne is the builder for updating a single ArticleTag entity.
type ArticleTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArticleTagMutation
}

// SetTag sets the "tag" field.
func (atuo *ArticleTagUpdateOne) SetTag(s string) *ArticleTagUpdateOne {
	atuo.mutation.SetTag(s)
	return atuo
}

// SetArticleID sets the "article_id" field.
func (atuo *ArticleTagUpdateOne) SetArticleID(u uuid.UUID) *ArticleTagUpdateOne {
	atuo.mutation.SetArticleID(u)
	return atuo
}

// SetCreatedAt sets the "created_at" field.
func (atuo *ArticleTagUpdateOne) SetCreatedAt(t time.Time) *ArticleTagUpdateOne {
	atuo.mutation.SetCreatedAt(t)
	return atuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (atuo *ArticleTagUpdateOne) SetNillableCreatedAt(t *time.Time) *ArticleTagUpdateOne {
	if t != nil {
		atuo.SetCreatedAt(*t)
	}
	return atuo
}

// SetUpdatedAt sets the "updated_at" field.
func (atuo *ArticleTagUpdateOne) SetUpdatedAt(t time.Time) *ArticleTagUpdateOne {
	atuo.mutation.SetUpdatedAt(t)
	return atuo
}

// SetArticle sets the "article" edge to the Article entity.
func (atuo *ArticleTagUpdateOne) SetArticle(a *Article) *ArticleTagUpdateOne {
	return atuo.SetArticleID(a.ID)
}

// Mutation returns the ArticleTagMutation object of the builder.
func (atuo *ArticleTagUpdateOne) Mutation() *ArticleTagMutation {
	return atuo.mutation
}

// ClearArticle clears the "article" edge to the Article entity.
func (atuo *ArticleTagUpdateOne) ClearArticle() *ArticleTagUpdateOne {
	atuo.mutation.ClearArticle()
	return atuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (atuo *ArticleTagUpdateOne) Select(field string, fields ...string) *ArticleTagUpdateOne {
	atuo.fields = append([]string{field}, fields...)
	return atuo
}

// Save executes the query and returns the updated ArticleTag entity.
func (atuo *ArticleTagUpdateOne) Save(ctx context.Context) (*ArticleTag, error) {
	var (
		err  error
		node *ArticleTag
	)
	atuo.defaults()
	if len(atuo.hooks) == 0 {
		if err = atuo.check(); err != nil {
			return nil, err
		}
		node, err = atuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ArticleTagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atuo.check(); err != nil {
				return nil, err
			}
			atuo.mutation = mutation
			node, err = atuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(atuo.hooks) - 1; i >= 0; i-- {
			if atuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, atuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ArticleTag)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ArticleTagMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (atuo *ArticleTagUpdateOne) SaveX(ctx context.Context) *ArticleTag {
	node, err := atuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (atuo *ArticleTagUpdateOne) Exec(ctx context.Context) error {
	_, err := atuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atuo *ArticleTagUpdateOne) ExecX(ctx context.Context) {
	if err := atuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atuo *ArticleTagUpdateOne) defaults() {
	if _, ok := atuo.mutation.UpdatedAt(); !ok {
		v := articletag.UpdateDefaultUpdatedAt()
		atuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atuo *ArticleTagUpdateOne) check() error {
	if _, ok := atuo.mutation.ArticleID(); atuo.mutation.ArticleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ArticleTag.article"`)
	}
	return nil
}

func (atuo *ArticleTagUpdateOne) sqlSave(ctx context.Context) (_node *ArticleTag, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   articletag.Table,
			Columns: articletag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: articletag.FieldID,
			},
		},
	}
	id, ok := atuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ArticleTag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := atuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, articletag.FieldID)
		for _, f := range fields {
			if !articletag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != articletag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := atuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := atuo.mutation.Tag(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: articletag.FieldTag,
		})
	}
	if value, ok := atuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: articletag.FieldCreatedAt,
		})
	}
	if value, ok := atuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: articletag.FieldUpdatedAt,
		})
	}
	if atuo.mutation.ArticleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   articletag.ArticleTable,
			Columns: []string{articletag.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: article.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := atuo.mutation.ArticleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   articletag.ArticleTable,
			Columns: []string{articletag.ArticleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: article.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ArticleTag{config: atuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, atuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{articletag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
