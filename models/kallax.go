// IMPORTANT! This is auto generated code by https://github.com/src-d/go-kallax
// Please, do not touch the code below, and if you do, do it under your own
// risk. Take into account that all the code you write here will be completely
// erased from earth the next time you generate the kallax models.
package models

import (
	"database/sql"
	"fmt"
	"time"

	"gopkg.in/src-d/go-kallax.v1"
	"gopkg.in/src-d/go-kallax.v1/types"
)

var _ types.SQLType
var _ fmt.Formatter

type modelSaveFunc func(*kallax.Store) error

// NewCharacter returns a new instance of Character.
func NewCharacter() (record *Character) {
	return new(Character)
}

// GetID returns the primary key of the model.
func (r *Character) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Character) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "player_id":
		return types.Nullable(kallax.VirtualColumn("player_id", r, new(kallax.NumericID))), nil
	case "name":
		return &r.Name, nil
	case "job":
		return &r.Job, nil
	case "map_id":
		return types.Nullable(kallax.VirtualColumn("map_id", r, new(kallax.NumericID))), nil
	case "level":
		return &r.Level, nil
	case "race":
		return &r.Race, nil
	case "enabled":
		return &r.Enabled, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Character: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Character) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil
	case "player_id":
		v := r.Model.VirtualColumn(col)
		if v == nil {
			return nil, kallax.ErrEmptyVirtualColumn
		}
		return v, nil
	case "name":
		return r.Name, nil
	case "job":
		return r.Job, nil
	case "map_id":
		v := r.Model.VirtualColumn(col)
		if v == nil {
			return nil, kallax.ErrEmptyVirtualColumn
		}
		return v, nil
	case "level":
		return r.Level, nil
	case "race":
		return r.Race, nil
	case "enabled":
		return r.Enabled, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Character: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Character) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "Player":
		return new(Player), nil
	case "Map":
		return new(Map), nil

	}
	return nil, fmt.Errorf("kallax: model Character has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Character) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "Player":
		val, ok := rel.(*Player)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Player", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Player = val
		}

		return nil
	case "Map":
		val, ok := rel.(*Map)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Map", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Map = val
		}

		return nil

	}
	return fmt.Errorf("kallax: model Character has no relationship %s", field)
}

// CharacterStore is the entity to access the records of the type Character
// in the database.
type CharacterStore struct {
	*kallax.Store
}

// NewCharacterStore creates a new instance of CharacterStore
// using a SQL database.
func NewCharacterStore(db *sql.DB) *CharacterStore {
	return &CharacterStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *CharacterStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *CharacterStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *CharacterStore) Debug() *CharacterStore {
	return &CharacterStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *CharacterStore) DebugWith(logger kallax.LoggerFunc) *CharacterStore {
	return &CharacterStore{s.Store.DebugWith(logger)}
}

func (s *CharacterStore) inverseRecords(record *Character) []modelSaveFunc {
	var result []modelSaveFunc

	if record.Player != nil && !record.Player.IsSaving() {
		record.AddVirtualColumn("player_id", record.Player.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&PlayerStore{store}).Save(record.Player)
			return err
		})
	}

	if record.Map != nil && !record.Map.IsSaving() {
		record.AddVirtualColumn("map_id", record.Map.GetID())
		result = append(result, func(store *kallax.Store) error {
			_, err := (&MapStore{store}).Save(record.Map)
			return err
		})
	}

	return result
}

// Insert inserts a Character in the database. A non-persisted object is
// required for this operation.
func (s *CharacterStore) Insert(record *Character) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	if err := record.BeforeSave(); err != nil {
		return err
	}

	inverseRecords := s.inverseRecords(record)

	if len(inverseRecords) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {
			for _, r := range inverseRecords {
				if err := r(s); err != nil {
					return err
				}
			}

			if err := s.Insert(Schema.Character.BaseSchema, record); err != nil {
				return err
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.Character.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *CharacterStore) Update(record *Character, cols ...kallax.SchemaField) (updated int64, err error) {
	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	record.SetSaving(true)
	defer record.SetSaving(false)

	if err := record.BeforeSave(); err != nil {
		return 0, err
	}

	inverseRecords := s.inverseRecords(record)

	if len(inverseRecords) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {
			for _, r := range inverseRecords {
				if err := r(s); err != nil {
					return err
				}
			}

			updated, err = s.Update(Schema.Character.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

	return s.Store.Update(Schema.Character.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *CharacterStore) Save(record *Character) (updated bool, err error) {
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
func (s *CharacterStore) Delete(record *Character) error {
	return s.Store.Delete(Schema.Character.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *CharacterStore) Find(q *CharacterQuery) (*CharacterResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewCharacterResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *CharacterStore) MustFind(q *CharacterQuery) *CharacterResultSet {
	return NewCharacterResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *CharacterStore) Count(q *CharacterQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *CharacterStore) MustCount(q *CharacterQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *CharacterStore) FindOne(q *CharacterQuery) (*Character, error) {
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
func (s *CharacterStore) FindAll(q *CharacterQuery) ([]*Character, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *CharacterStore) MustFindOne(q *CharacterQuery) *Character {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Character with the data in the database and
// makes it writable.
func (s *CharacterStore) Reload(record *Character) error {
	return s.Store.Reload(Schema.Character.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *CharacterStore) Transaction(callback func(*CharacterStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&CharacterStore{store})
	})
}

// CharacterQuery is the object used to create queries for the Character
// entity.
type CharacterQuery struct {
	*kallax.BaseQuery
}

// NewCharacterQuery returns a new instance of CharacterQuery.
func NewCharacterQuery() *CharacterQuery {
	return &CharacterQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Character.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *CharacterQuery) Select(columns ...kallax.SchemaField) *CharacterQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *CharacterQuery) SelectNot(columns ...kallax.SchemaField) *CharacterQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *CharacterQuery) Copy() *CharacterQuery {
	return &CharacterQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *CharacterQuery) Order(cols ...kallax.ColumnOrder) *CharacterQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *CharacterQuery) BatchSize(size uint64) *CharacterQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *CharacterQuery) Limit(n uint64) *CharacterQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *CharacterQuery) Offset(n uint64) *CharacterQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *CharacterQuery) Where(cond kallax.Condition) *CharacterQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *CharacterQuery) WithPlayer() *CharacterQuery {
	q.AddRelation(Schema.Player.BaseSchema, "Player", kallax.OneToOne, nil)
	return q
}

func (q *CharacterQuery) WithMap() *CharacterQuery {
	q.AddRelation(Schema.Map.BaseSchema, "Map", kallax.OneToOne, nil)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *CharacterQuery) FindByID(v ...int64) *CharacterQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Character.ID, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *CharacterQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *CharacterQuery {
	return q.Where(cond(Schema.Character.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *CharacterQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *CharacterQuery {
	return q.Where(cond(Schema.Character.UpdatedAt, v))
}

// FindByPlayer adds a new filter to the query that will require that
// the foreign key of Player is equal to the passed value.
func (q *CharacterQuery) FindByPlayer(v int64) *CharacterQuery {
	return q.Where(kallax.Eq(Schema.Character.PlayerFK, v))
}

// FindByName adds a new filter to the query that will require that
// the Name property is equal to the passed value.
func (q *CharacterQuery) FindByName(v string) *CharacterQuery {
	return q.Where(kallax.Eq(Schema.Character.Name, v))
}

// FindByJob adds a new filter to the query that will require that
// the Job property is equal to the passed value.
func (q *CharacterQuery) FindByJob(v string) *CharacterQuery {
	return q.Where(kallax.Eq(Schema.Character.Job, v))
}

// FindByMap adds a new filter to the query that will require that
// the foreign key of Map is equal to the passed value.
func (q *CharacterQuery) FindByMap(v int64) *CharacterQuery {
	return q.Where(kallax.Eq(Schema.Character.MapFK, v))
}

// FindByLevel adds a new filter to the query that will require that
// the Level property is equal to the passed value.
func (q *CharacterQuery) FindByLevel(cond kallax.ScalarCond, v uint16) *CharacterQuery {
	return q.Where(cond(Schema.Character.Level, v))
}

// FindByRace adds a new filter to the query that will require that
// the Race property is equal to the passed value.
func (q *CharacterQuery) FindByRace(cond kallax.ScalarCond, v uint16) *CharacterQuery {
	return q.Where(cond(Schema.Character.Race, v))
}

// FindByEnabled adds a new filter to the query that will require that
// the Enabled property is equal to the passed value.
func (q *CharacterQuery) FindByEnabled(v bool) *CharacterQuery {
	return q.Where(kallax.Eq(Schema.Character.Enabled, v))
}

// CharacterResultSet is the set of results returned by a query to the
// database.
type CharacterResultSet struct {
	ResultSet kallax.ResultSet
	last      *Character
	lastErr   error
}

// NewCharacterResultSet creates a new result set for rows of the type
// Character.
func NewCharacterResultSet(rs kallax.ResultSet) *CharacterResultSet {
	return &CharacterResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *CharacterResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Character.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Character)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Character")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *CharacterResultSet) Get() (*Character, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *CharacterResultSet) ForEach(fn func(*Character) error) error {
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
func (rs *CharacterResultSet) All() ([]*Character, error) {
	var result []*Character
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
func (rs *CharacterResultSet) One() (*Character, error) {
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
func (rs *CharacterResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *CharacterResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewMap returns a new instance of Map.
func NewMap() (record *Map) {
	return new(Map)
}

// GetID returns the primary key of the model.
func (r *Map) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Map) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "name":
		return types.Slice(&r.Name), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Map: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Map) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil
	case "name":
		return types.Slice(r.Name), nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Map: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Map) NewRelationshipRecord(field string) (kallax.Record, error) {
	return nil, fmt.Errorf("kallax: model Map has no relationships")
}

// SetRelationship sets the given relationship in the given field.
func (r *Map) SetRelationship(field string, rel interface{}) error {
	return fmt.Errorf("kallax: model Map has no relationships")
}

// MapStore is the entity to access the records of the type Map
// in the database.
type MapStore struct {
	*kallax.Store
}

// NewMapStore creates a new instance of MapStore
// using a SQL database.
func NewMapStore(db *sql.DB) *MapStore {
	return &MapStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *MapStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *MapStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *MapStore) Debug() *MapStore {
	return &MapStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *MapStore) DebugWith(logger kallax.LoggerFunc) *MapStore {
	return &MapStore{s.Store.DebugWith(logger)}
}

// Insert inserts a Map in the database. A non-persisted object is
// required for this operation.
func (s *MapStore) Insert(record *Map) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	if err := record.BeforeSave(); err != nil {
		return err
	}

	return s.Store.Insert(Schema.Map.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *MapStore) Update(record *Map, cols ...kallax.SchemaField) (updated int64, err error) {
	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	record.SetSaving(true)
	defer record.SetSaving(false)

	if err := record.BeforeSave(); err != nil {
		return 0, err
	}

	return s.Store.Update(Schema.Map.BaseSchema, record, cols...)
}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *MapStore) Save(record *Map) (updated bool, err error) {
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
func (s *MapStore) Delete(record *Map) error {
	return s.Store.Delete(Schema.Map.BaseSchema, record)
}

// Find returns the set of results for the given query.
func (s *MapStore) Find(q *MapQuery) (*MapResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewMapResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *MapStore) MustFind(q *MapQuery) *MapResultSet {
	return NewMapResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *MapStore) Count(q *MapQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *MapStore) MustCount(q *MapQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *MapStore) FindOne(q *MapQuery) (*Map, error) {
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
func (s *MapStore) FindAll(q *MapQuery) ([]*Map, error) {
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	return rs.All()
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *MapStore) MustFindOne(q *MapQuery) *Map {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Map with the data in the database and
// makes it writable.
func (s *MapStore) Reload(record *Map) error {
	return s.Store.Reload(Schema.Map.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *MapStore) Transaction(callback func(*MapStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&MapStore{store})
	})
}

// MapQuery is the object used to create queries for the Map
// entity.
type MapQuery struct {
	*kallax.BaseQuery
}

// NewMapQuery returns a new instance of MapQuery.
func NewMapQuery() *MapQuery {
	return &MapQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Map.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *MapQuery) Select(columns ...kallax.SchemaField) *MapQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *MapQuery) SelectNot(columns ...kallax.SchemaField) *MapQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *MapQuery) Copy() *MapQuery {
	return &MapQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *MapQuery) Order(cols ...kallax.ColumnOrder) *MapQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *MapQuery) BatchSize(size uint64) *MapQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *MapQuery) Limit(n uint64) *MapQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *MapQuery) Offset(n uint64) *MapQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *MapQuery) Where(cond kallax.Condition) *MapQuery {
	q.BaseQuery.Where(cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *MapQuery) FindByID(v ...int64) *MapQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Map.ID, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *MapQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *MapQuery {
	return q.Where(cond(Schema.Map.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *MapQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *MapQuery {
	return q.Where(cond(Schema.Map.UpdatedAt, v))
}

// FindByName adds a new filter to the query that will require that
// the Name property contains all the passed values; if no passed values,
// it will do nothing.
func (q *MapQuery) FindByName(v ...string) *MapQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.ArrayContains(Schema.Map.Name, values...))
}

// MapResultSet is the set of results returned by a query to the
// database.
type MapResultSet struct {
	ResultSet kallax.ResultSet
	last      *Map
	lastErr   error
}

// NewMapResultSet creates a new result set for rows of the type
// Map.
func NewMapResultSet(rs kallax.ResultSet) *MapResultSet {
	return &MapResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *MapResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Map.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Map)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Map")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *MapResultSet) Get() (*Map, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *MapResultSet) ForEach(fn func(*Map) error) error {
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
func (rs *MapResultSet) All() ([]*Map, error) {
	var result []*Map
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
func (rs *MapResultSet) One() (*Map, error) {
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
func (rs *MapResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *MapResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewPlayer returns a new instance of Player.
func NewPlayer() (record *Player) {
	return new(Player)
}

// GetID returns the primary key of the model.
func (r *Player) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Player) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "username":
		return &r.Username, nil
	case "email":
		return &r.Email, nil
	case "password":
		return &r.Password, nil
	case "pincode":
		return types.Nullable(&r.Pincode), nil
	case "is_active":
		return &r.IsActive, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Player: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Player) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil
	case "username":
		return r.Username, nil
	case "email":
		return r.Email, nil
	case "password":
		return r.Password, nil
	case "pincode":
		if r.Pincode == (*string)(nil) {
			return nil, nil
		}
		return r.Pincode, nil
	case "is_active":
		return r.IsActive, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Player: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Player) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "Characters":
		return new(Character), nil

	}
	return nil, fmt.Errorf("kallax: model Player has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Player) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "Characters":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Characters = make([]*Character, len(records))
		for i, record := range records {
			rel, ok := record.(*Character)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Characters[i] = rel
		}
		return nil

	}
	return fmt.Errorf("kallax: model Player has no relationship %s", field)
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

func (s *PlayerStore) relationshipRecords(record *Player) []modelSaveFunc {
	var result []modelSaveFunc

	for i := range record.Characters {
		r := record.Characters[i]
		if !r.IsSaving() {
			r.AddVirtualColumn("player_id", record.GetID())
			result = append(result, func(store *kallax.Store) error {
				_, err := (&CharacterStore{store}).Save(r)
				return err
			})
		}
	}

	return result
}

// Insert inserts a Player in the database. A non-persisted object is
// required for this operation.
func (s *PlayerStore) Insert(record *Player) error {
	record.SetSaving(true)
	defer record.SetSaving(false)

	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	if err := record.BeforeSave(); err != nil {
		return err
	}

	records := s.relationshipRecords(record)

	if len(records) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {
			if err := s.Insert(Schema.Player.BaseSchema, record); err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.Player.BaseSchema, record)
}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *PlayerStore) Update(record *Player, cols ...kallax.SchemaField) (updated int64, err error) {
	record.CreatedAt = record.CreatedAt.Truncate(time.Microsecond)
	record.UpdatedAt = record.UpdatedAt.Truncate(time.Microsecond)

	record.SetSaving(true)
	defer record.SetSaving(false)

	if err := record.BeforeSave(); err != nil {
		return 0, err
	}

	records := s.relationshipRecords(record)

	if len(records) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {
			updated, err = s.Update(Schema.Player.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			for _, r := range records {
				if err := r(s); err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

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

// RemoveCharacters removes the given items of the Characters field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
// Note that is required that `Characters` is not empty. This method clears the
// the elements of Characters in a model, it does not retrieve them to know
// what relationships the model has.
func (s *PlayerStore) RemoveCharacters(record *Player, deleted ...*Character) error {
	var updated []*Character
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Characters
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Character.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Characters = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Character.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Character.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Characters {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Characters = updated
	return nil
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

func (q *PlayerQuery) WithCharacters(cond kallax.Condition) *PlayerQuery {
	q.AddRelation(Schema.Character.BaseSchema, "Characters", kallax.OneToMany, cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *PlayerQuery) FindByID(v ...int64) *PlayerQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Player.ID, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *PlayerQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *PlayerQuery {
	return q.Where(cond(Schema.Player.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *PlayerQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *PlayerQuery {
	return q.Where(cond(Schema.Player.UpdatedAt, v))
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

// FindByIsActive adds a new filter to the query that will require that
// the IsActive property is equal to the passed value.
func (q *PlayerQuery) FindByIsActive(v bool) *PlayerQuery {
	return q.Where(kallax.Eq(Schema.Player.IsActive, v))
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
	Character *schemaCharacter
	Map       *schemaMap
	Player    *schemaPlayer
}

type schemaCharacter struct {
	*kallax.BaseSchema
	ID        kallax.SchemaField
	CreatedAt kallax.SchemaField
	UpdatedAt kallax.SchemaField
	PlayerFK  kallax.SchemaField
	Name      kallax.SchemaField
	Job       kallax.SchemaField
	MapFK     kallax.SchemaField
	Level     kallax.SchemaField
	Race      kallax.SchemaField
	Enabled   kallax.SchemaField
}

type schemaMap struct {
	*kallax.BaseSchema
	ID        kallax.SchemaField
	CreatedAt kallax.SchemaField
	UpdatedAt kallax.SchemaField
	Name      kallax.SchemaField
}

type schemaPlayer struct {
	*kallax.BaseSchema
	ID        kallax.SchemaField
	CreatedAt kallax.SchemaField
	UpdatedAt kallax.SchemaField
	Username  kallax.SchemaField
	Email     kallax.SchemaField
	Password  kallax.SchemaField
	Pincode   kallax.SchemaField
	IsActive  kallax.SchemaField
}

var Schema = &schema{
	Character: &schemaCharacter{
		BaseSchema: kallax.NewBaseSchema(
			"characters",
			"__character",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"Player": kallax.NewForeignKey("player_id", true),
				"Map":    kallax.NewForeignKey("map_id", true),
			},
			func() kallax.Record {
				return new(Character)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("player_id"),
			kallax.NewSchemaField("name"),
			kallax.NewSchemaField("job"),
			kallax.NewSchemaField("map_id"),
			kallax.NewSchemaField("level"),
			kallax.NewSchemaField("race"),
			kallax.NewSchemaField("enabled"),
		),
		ID:        kallax.NewSchemaField("id"),
		CreatedAt: kallax.NewSchemaField("created_at"),
		UpdatedAt: kallax.NewSchemaField("updated_at"),
		PlayerFK:  kallax.NewSchemaField("player_id"),
		Name:      kallax.NewSchemaField("name"),
		Job:       kallax.NewSchemaField("job"),
		MapFK:     kallax.NewSchemaField("map_id"),
		Level:     kallax.NewSchemaField("level"),
		Race:      kallax.NewSchemaField("race"),
		Enabled:   kallax.NewSchemaField("enabled"),
	},
	Map: &schemaMap{
		BaseSchema: kallax.NewBaseSchema(
			"maps",
			"__map",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{},
			func() kallax.Record {
				return new(Map)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("name"),
		),
		ID:        kallax.NewSchemaField("id"),
		CreatedAt: kallax.NewSchemaField("created_at"),
		UpdatedAt: kallax.NewSchemaField("updated_at"),
		Name:      kallax.NewSchemaField("name"),
	},
	Player: &schemaPlayer{
		BaseSchema: kallax.NewBaseSchema(
			"players",
			"__player",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"Characters": kallax.NewForeignKey("player_id", false),
			},
			func() kallax.Record {
				return new(Player)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("username"),
			kallax.NewSchemaField("email"),
			kallax.NewSchemaField("password"),
			kallax.NewSchemaField("pincode"),
			kallax.NewSchemaField("is_active"),
		),
		ID:        kallax.NewSchemaField("id"),
		CreatedAt: kallax.NewSchemaField("created_at"),
		UpdatedAt: kallax.NewSchemaField("updated_at"),
		Username:  kallax.NewSchemaField("username"),
		Email:     kallax.NewSchemaField("email"),
		Password:  kallax.NewSchemaField("password"),
		Pincode:   kallax.NewSchemaField("pincode"),
		IsActive:  kallax.NewSchemaField("is_active"),
	},
}
