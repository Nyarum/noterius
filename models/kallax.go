// IMPORTANT! This is auto generated code by https://github.com/src-d/go-kallax
// Please, do not touch the code below, and if you do, do it under your own
// risk. Take into account that all the code you write here will be completely
// erased from earth the next time you generate the kallax models.
package models

import (
	"database/sql"
	"fmt"

	"gopkg.in/src-d/go-kallax.v1"
	"gopkg.in/src-d/go-kallax.v1/types"
)

var _ types.SQLType
var _ fmt.Formatter

// NewPlayer returns a new instance of Player.
func NewPlayer() (record *Player) {
	return new(Player)
}

// GetID returns the primary key of the model.
func (r *Player) GetID() kallax.Identifier {
	return (*kallax.ULID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Player) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.ULID)(&r.ID), nil
	case "username":
		return &r.Username, nil
	case "email":
		return &r.Email, nil
	case "password":
		return &r.Password, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Player: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Player) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "username":
		return r.Username, nil
	case "email":
		return r.Email, nil
	case "password":
		return r.Password, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Player: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Player) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Player has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *Player) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Player has no relationships")
}

// PlayerStore is the entity to access the records of the type Player
// in the database.
type PlayerStore struct {
	*kallax.Store
}

// NewPlayerStore creates a new instance of PlayerStore
// using a SQL database.
func NewPlayerStore(db *sql.DB) *PlayerStore {
	return &PlayerStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *PlayerStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *PlayerStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *PlayerStore) Debug() *PlayerStore {
	return &PlayerStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *PlayerStore) DebugWith(logger kallax.LoggerFunc) *PlayerStore {
	return &PlayerStore{s.Store.DebugWith(logger)}
}

// Insert inserts a Player in the database. A non-persisted object is
// required for this operation.
func (s *PlayerStore) Insert(record *Player) error {
	return s.Store.Insert(Schema.Player.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *PlayerStore) Update(record *Player, cols ...kallax.SchemaField) (updated int64, err error) {
	return s.Store.Update(Schema.Player.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *PlayerStore) Save(record *Player) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *PlayerStore) Delete(record *Player) error {
	return s.Store.Delete(Schema.Player.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *PlayerStore) Find(q *PlayerQuery) (*PlayerResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewPlayerResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *PlayerStore) MustFind(q *PlayerQuery) *PlayerResultSet {
	return NewPlayerResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *PlayerStore) Count(q *PlayerQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *PlayerStore) MustCount(q *PlayerQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *PlayerStore) FindOne(q *PlayerQuery) (*Player, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// FindAll returns a list of all the rows returned by the given query.
func (s *PlayerStore) FindAll(q *PlayerQuery) ([]*Player, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *PlayerStore) MustFindOne(q *PlayerQuery) *Player {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Player with the data in the database and
// makes it writable.
func (s *PlayerStore) Reload(record *Player) error {
	return s.Store.Reload(Schema.Player.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *PlayerStore) Transaction(callback func(*PlayerStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&PlayerStore{store})
	})
}

// PlayerQuery is the object used to create queries for the Player
// entity.
type PlayerQuery struct {
	*kallax.BaseQuery
}

// NewPlayerQuery returns a new instance of PlayerQuery.
func NewPlayerQuery() *PlayerQuery {
	return &PlayerQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Player.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *PlayerQuery) Select(columns ...kallax.SchemaField) *PlayerQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *PlayerQuery) SelectNot(columns ...kallax.SchemaField) *PlayerQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *PlayerQuery) Copy() *PlayerQuery {
	return &PlayerQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *PlayerQuery) Order(cols ...kallax.ColumnOrder) *PlayerQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *PlayerQuery) BatchSize(size uint64) *PlayerQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *PlayerQuery) Limit(n uint64) *PlayerQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *PlayerQuery) Offset(n uint64) *PlayerQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *PlayerQuery) Where(cond kallax.Condition) *PlayerQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *PlayerQuery) FindByID(v ...kallax.ULID) *PlayerQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Player.ID, values...))
}

// FindByUsername adds a new filter to the query that will require that
// the Username property is equal to the passed value.
func (q *PlayerQuery) FindByUsername(v string) *PlayerQuery {
	return q.Where(kallax.Eq(Schema.Player.Username, v))
}

// FindByEmail adds a new filter to the query that will require that
// the Email property is equal to the passed value.
func (q *PlayerQuery) FindByEmail(v string) *PlayerQuery {
	return q.Where(kallax.Eq(Schema.Player.Email, v))
}

// FindByPassword adds a new filter to the query that will require that
// the Password property is equal to the passed value.
func (q *PlayerQuery) FindByPassword(v string) *PlayerQuery {
	return q.Where(kallax.Eq(Schema.Player.Password, v))
}

// PlayerResultSet is the set of results returned by a query to the
// database.
type PlayerResultSet struct {
	ResultSet kallax.ResultSet
	last      *Player
	lastErr   error
}

// NewPlayerResultSet creates a new result set for rows of the type
// Player.
func NewPlayerResultSet(rs kallax.ResultSet) *PlayerResultSet {
	return &PlayerResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *PlayerResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Player.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Player)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Player")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *PlayerResultSet) Get() (*Player, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *PlayerResultSet) ForEach(fn func(*Player) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *PlayerResultSet) All() ([]*Player, error) {
	var result []*Player
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *PlayerResultSet) One() (*Player, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *PlayerResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *PlayerResultSet) Close() error {
	return rs.ResultSet.Close()
}

type schema struct {
	Player *schemaPlayer
}

type schemaPlayer struct {
	*kallax.BaseSchema
	ID       kallax.SchemaField
	Username kallax.SchemaField
	Email    kallax.SchemaField
	Password kallax.SchemaField
}

var Schema = &schema{
	Player: &schemaPlayer{
		BaseSchema: kallax.NewBaseSchema(
			"users",
			"__player",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Player)
			},
			false,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("username"),
			kallax.NewSchemaField("email"),
			kallax.NewSchemaField("password"),
		),
		ID:       kallax.NewSchemaField("id"),
		Username: kallax.NewSchemaField("username"),
		Email:    kallax.NewSchemaField("email"),
		Password: kallax.NewSchemaField("password"),
	},
}
