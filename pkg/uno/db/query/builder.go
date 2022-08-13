package query

import (
	"errors"
	"strings"
)

var (
	// ErrTableEmpty table not set
	ErrTableEmpty = errors.New("table empty")
	// ErrInsertEmpty insert content not set
	ErrInsertEmpty = errors.New("insert content empty")
	// ErrUpdateEmpty update content not set
	ErrUpdateEmpty = errors.New("update content empty")
)

// Builder sql builder
type Builder struct {
	_table        string
	_select       string
	_insert       string
	_update       string
	_delete       string
	_join         string
	_where        string
	_groupBy      string
	_having       string
	_orderBy      string
	_limit        string
	_insertParams []interface{}
	_updateParams []interface{}
	_whereParams  []interface{}
	_havingParams []interface{}
	_limitParams  []interface{}
	_joinParams   []interface{}
}

// NewBuilder init sql builder
func NewBuilder() *Builder {
	return &Builder{}
}

// GetQuery get sql
func (sb *Builder) GetQuery() (string, error) {
	if sb._table == "" {
		return "", ErrTableEmpty
	}
	var buf strings.Builder

	buf.WriteString("SELECT ")
	if sb._select != "" {
		buf.WriteString(sb._select)
	} else {
		buf.WriteString("*")
	}
	buf.WriteString(" FROM ")
	buf.WriteString(sb._table)
	if sb._join != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._join)
	}
	if sb._where != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._where)
	}
	if sb._groupBy != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._groupBy)
	}
	if sb._having != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._having)
	}
	if sb._orderBy != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._orderBy)
	}
	if sb._limit != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._limit)
	}

	return buf.String(), nil
}

// GetInsert get sql
func (sb *Builder) GetInsert() (string, error) {
	if sb._table == "" {
		return "", ErrTableEmpty
	}
	if sb._insert == "" {
		return "", ErrInsertEmpty
	}

	var buf strings.Builder

	buf.WriteString("INSERT INTO ")
	buf.WriteString(sb._table)
	buf.WriteString(" ")
	buf.WriteString(sb._insert)

	return buf.String(), nil
}

// GetUpdate get sql
func (sb *Builder) GetUpdate() (string, error) {
	if sb._table == "" {
		return "", ErrTableEmpty
	}

	if sb._update == "" {
		return "", ErrUpdateEmpty
	}

	var buf strings.Builder

	buf.WriteString("UPDATE ")
	buf.WriteString(sb._table)
	buf.WriteString(" ")
	buf.WriteString(sb._update)
	if sb._where != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._where)
	}

	return buf.String(), nil
}

// GetDelete get sql
func (sb *Builder) GetDelete() (string, error) {
	if sb._table == "" {
		return "", ErrTableEmpty
	}

	var buf strings.Builder

	buf.WriteString("DELETE FROM ")
	buf.WriteString(sb._table)
	if sb._where != "" {
		buf.WriteString(" ")
		buf.WriteString(sb._where)
	}

	return buf.String(), nil
}

// GetQueryParams get params
func (sb *Builder) GetQueryParams() []interface{} {
	params := []interface{}{}
	params = append(params, sb._joinParams...)
	params = append(params, sb._whereParams...)
	params = append(params, sb._havingParams...)
	params = append(params, sb._limitParams...)
	return params
}

// GetInsertParams get params
func (sb *Builder) GetInsertParams() []interface{} {
	params := []interface{}{}
	params = append(params, sb._insertParams...)
	return params
}

// GetUpdateParams get params
func (sb *Builder) GetUpdateParams() []interface{} {
	params := []interface{}{}
	params = append(params, sb._updateParams...)
	params = append(params, sb._whereParams...)
	return params
}

// GetDeleteParams get params
func (sb *Builder) GetDeleteParams() []interface{} {
	params := []interface{}{}
	params = append(params, sb._whereParams...)
	return params
}

// Table set table
func (sb *Builder) Table(table string) *Builder {

	sb._table = table

	return sb
}

// Select set select cols
func (sb *Builder) Select(cols ...string) *Builder {
	var buf strings.Builder

	for k, col := range cols {

		buf.WriteString(col)

		if k != len(cols)-1 {
			buf.WriteString(",")
		}
	}

	sb._select = buf.String()

	return sb
}

// Insert set Insert
func (sb *Builder) Insert(cols []string, values ...interface{}) *Builder {
	var buf strings.Builder

	buf.WriteString("(")
	for k, col := range cols {

		buf.WriteString(col)

		if k != len(cols)-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString(") VALUES (")

	for k := range cols {
		buf.WriteString("?")
		if k != len(cols)-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString(")")

	sb._insert = buf.String()

	for _, value := range values {
		sb._insertParams = append(sb._insertParams, value)
	}

	return sb
}

// Update set update
func (sb *Builder) Update(cols []string, values ...interface{}) *Builder {
	var buf strings.Builder

	buf.WriteString("SET ")

	for k, col := range cols {

		buf.WriteString(col)

		buf.WriteString(" = ?")
		if k != len(cols)-1 {
			buf.WriteString(",")
		}
	}

	sb._update = buf.String()

	for _, value := range values {
		sb._updateParams = append(sb._updateParams, value)
	}

	return sb
}

// WhereRaw set where raw string
func (sb *Builder) WhereRaw(s string, values ...interface{}) *Builder {
	return sb.whereRaw("AND", s, values)
}

// OrWhereRaw set where raw string
func (sb *Builder) OrWhereRaw(s string, values ...interface{}) *Builder {
	return sb.whereRaw("OR", s, values)
}

func (sb *Builder) whereRaw(operator string, s string, values []interface{}) *Builder {
	var buf strings.Builder

	buf.WriteString(sb._where) // append

	if buf.Len() == 0 {
		buf.WriteString("WHERE ")
	} else {
		buf.WriteString(" ")
		buf.WriteString(operator)
		buf.WriteString(" ")
	}

	buf.WriteString(s)
	sb._where = buf.String()
	sb._whereParams = append(sb._whereParams, values...)

	return sb
}

// Where set where cond
func (sb *Builder) Where(field string, condition string, value interface{}) *Builder {
	return sb.where("AND", condition, field, value)
}

// OrWhere set or where cond
func (sb *Builder) OrWhere(field string, condition string, value interface{}) *Builder {
	return sb.where("OR", condition, field, value)
}

func (sb *Builder) where(operator string, condition string, field string, value interface{}) *Builder {
	var buf strings.Builder

	buf.WriteString(sb._where) // append

	if buf.Len() == 0 {
		buf.WriteString("WHERE ")
	} else {
		buf.WriteString(" ")
		buf.WriteString(operator)
		buf.WriteString(" ")
	}

	buf.WriteString(field)

	buf.WriteString(" ")
	buf.WriteString(condition)
	buf.WriteString(" ")
	buf.WriteString("?")

	sb._where = buf.String()

	sb._whereParams = append(sb._whereParams, value)

	return sb
}

// WhereIn set where in cond
func (sb *Builder) WhereIn(field string, values ...interface{}) *Builder {
	return sb.whereIn("AND", "IN", field, values)
}

// OrWhereIn set or where in cond
func (sb *Builder) OrWhereIn(field string, values ...interface{}) *Builder {
	return sb.whereIn("OR", "IN", field, values)
}

// WhereNotIn set where not in cond
func (sb *Builder) WhereNotIn(field string, values ...interface{}) *Builder {
	return sb.whereIn("AND", "NOT IN", field, values)
}

// OrWhereNotIn set or where not in cond
func (sb *Builder) OrWhereNotIn(field string, values ...interface{}) *Builder {
	return sb.whereIn("OR", "NOT IN", field, values)
}

func (sb *Builder) whereIn(operator string, condition string, field string, values []interface{}) *Builder {
	var buf strings.Builder

	buf.WriteString(sb._where) // append

	if buf.Len() == 0 {
		buf.WriteString("WHERE ")
	} else {
		buf.WriteString(" ")
		buf.WriteString(operator)
		buf.WriteString(" ")
	}

	buf.WriteString(field)

	plhs := GenPlaceholders(len(values))
	buf.WriteString(" ")
	buf.WriteString(condition)
	buf.WriteString(" ")
	buf.WriteString("(")
	buf.WriteString(plhs)
	buf.WriteString(")")

	sb._where = buf.String()

	sb._whereParams = append(sb._whereParams, values...)

	return sb
}

// GroupBy set group by fields
func (sb *Builder) GroupBy(fields ...string) *Builder {
	var buf strings.Builder

	buf.WriteString("GROUP BY ")

	for k, field := range fields {

		buf.WriteString(field)

		if k != len(fields)-1 {
			buf.WriteString(",")
		}
	}

	sb._groupBy = buf.String()

	return sb
}

// HavingRaw set having raw string
func (sb *Builder) HavingRaw(s string, values ...interface{}) *Builder {
	return sb.havingRaw("AND", s, values)
}

// OrHavingRaw set having raw string
func (sb *Builder) OrHavingRaw(s string, values ...interface{}) *Builder {
	return sb.havingRaw("OR", s, values)
}

func (sb *Builder) havingRaw(operator string, s string, values []interface{}) *Builder {
	var buf strings.Builder

	buf.WriteString(sb._having) // append

	if buf.Len() == 0 {
		buf.WriteString("HAVING ")
	} else {
		buf.WriteString(" ")
		buf.WriteString(operator)
		buf.WriteString(" ")
	}

	buf.WriteString(s)
	sb._having = buf.String()

	sb._havingParams = append(sb._havingParams, values...)

	return sb
}

// Having set having cond
func (sb *Builder) Having(field string, condition string, value interface{}) *Builder {
	return sb.having("AND", condition, field, value)
}

// OrHaving set or having cond
func (sb *Builder) OrHaving(field string, condition string, value interface{}) *Builder {
	return sb.having("OR", condition, field, value)
}

func (sb *Builder) having(operator string, condition string, field string, value interface{}) *Builder {
	if sb._groupBy == "" { // group by not set
		return sb
	}

	var buf strings.Builder

	buf.WriteString(sb._having) // append

	if buf.Len() == 0 {
		buf.WriteString("HAVING ")
	} else {
		buf.WriteString(" ")
		buf.WriteString(operator)
		buf.WriteString(" ")
	}

	buf.WriteString(field)

	buf.WriteString(" ")
	buf.WriteString(condition)
	buf.WriteString(" ")
	buf.WriteString("?")

	sb._having = buf.String()

	sb._havingParams = append(sb._havingParams, value)

	return sb
}

// OrderBy set order by fields
func (sb *Builder) OrderBy(operator string, fields ...string) *Builder {
	var buf strings.Builder

	buf.WriteString("ORDER BY ")

	for k, field := range fields {

		buf.WriteString(field)

		if k != len(fields)-1 {
			buf.WriteString(",")
		}
	}

	buf.WriteString(" ")
	buf.WriteString(operator)

	sb._orderBy = buf.String()

	return sb
}

// Limit set limit
func (sb *Builder) Limit(offset, num interface{}) *Builder {
	var buf strings.Builder

	buf.WriteString("LIMIT ? OFFSET ?")

	sb._limit = buf.String()

	sb._limitParams = append(sb._limitParams, num, offset)

	return sb
}

// JoinRaw join with raw sql
func (sb *Builder) JoinRaw(join string, values ...interface{}) *Builder {
	var buf strings.Builder

	buf.WriteString(sb._join)
	if buf.Len() != 0 {
		buf.WriteString(" ")
	}
	buf.WriteString(join)

	sb._join = buf.String()
	sb._joinParams = append(sb._joinParams, values...)

	return sb
}

// GenPlaceholders generate placeholders
func GenPlaceholders(n int) string {
	var buf strings.Builder

	for i := 0; i < n-1; i++ {
		buf.WriteString("?,")
	}

	if n > 0 {
		buf.WriteString("?")
	}

	return buf.String()
}

/*type Builder interface {

	// Create saves a new model and return the instance.
	Create(fields core.Fields) any

	// Distinct forces the query to only return distinct results.
	Distinct() Builder

	// FirstOrCreate gets the first record matching the attributes or create it.
	FirstOrCreate(values ...core.Fields) any

	// From sets the table which the query is targeting.
	From(table string, as ...string) Builder

	// Get executes the query as a "select" statement.
	Get() core.Collection

	// GetBindings gets the current query value bindings in a flattened array.
	GetBindings() (results []any)

	// GroupBy adds a "group by" clause to the query.
	GroupBy(columns ...string) Builder

	// Having adds a "having" clause to the query.
	Having(column string, args ...any) Builder

	// Insert inserts new records into the database.
	Insert(values ...core.Fields) bool

	// Limit sets the "limit" value of the query.
	Limit(num int64) Builder

	// Offset sets the "offset" value of the query.
	Offset(offset int64) Builder

	// OrderBy adds an "order by" clause to the query.
	OrderBy(column string, columnOrderType ...string) Builder

	// OrHaving adds an "or having" clause to the query.
	OrHaving(column string, args ...any) Builder

	// OrWhere adds an "or where" clause to the query.
	OrWhere(column string, args ...any) Builder

	// Select sets the columns to be selected.
	Select(columns ...string) Builder

	// SelectSql gets the complete  string formed by the current specifications of this query builder.
	SelectSql() (string, []any)

	// Skip is alias to set the "offset" value of the query.
	Skip(offset int64) Builder

	// Take is alias to set the "limit" value of the query.
	Take(num int64) Builder

	// ToSql gets the  representation of the query.
	ToSql() string

	// Update updates records in the database.
	Update(fields core.Fields) int64

	// UpdateOrCreate creates or update a record matching the attributes, and fill it with values.
	UpdateOrCreate(attributes core.Fields, values ...core.Fields) any

	// UpdateOrInsert inserts or update a record matching the attributes, and fill it with values.
	UpdateOrInsert(attributes core.Fields, values ...core.Fields) bool

	// Where adds a basic where clause to the query.
	Where(column string, args ...any) Builder

	// WhereIn adds a "where in" clause to the query.
	WhereIn(column string, args any, whereType ...string) Builder
}

// QueryBuilder
type QueryBuilder struct {
	Builder
	limit    int64
	offset   int64
	distinct bool
	table    string
	fields   []string
	wheres   *Wheres
	orderBy  OrderBys
	groupBy  GroupBy
	joins    Joins
	unions   Unions
	having   *Wheres
	bindings map[string][]interface{}
}

func NewQueryBuilder() Builder {
	return &QueryBuilder{}
}*/
