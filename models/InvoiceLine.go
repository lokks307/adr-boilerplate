// Code generated by SQLBoiler 4.8.3 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// InvoiceLine is an object representing the database table.
type InvoiceLine struct {
	InvoiceLineId int64  `boil:"InvoiceLineId" json:"InvoiceLineId" toml:"InvoiceLineId" yaml:"InvoiceLineId"`
	InvoiceId     int64  `boil:"InvoiceId" json:"InvoiceId" toml:"InvoiceId" yaml:"InvoiceId"`
	TrackId       int64  `boil:"TrackId" json:"TrackId" toml:"TrackId" yaml:"TrackId"`
	UnitPrice     string `boil:"UnitPrice" json:"UnitPrice" toml:"UnitPrice" yaml:"UnitPrice"`
	Quantity      int64  `boil:"Quantity" json:"Quantity" toml:"Quantity" yaml:"Quantity"`

	R *invoiceLineR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L invoiceLineL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var InvoiceLineColumns = struct {
	InvoiceLineId string
	InvoiceId     string
	TrackId       string
	UnitPrice     string
	Quantity      string
}{
	InvoiceLineId: "InvoiceLineId",
	InvoiceId:     "InvoiceId",
	TrackId:       "TrackId",
	UnitPrice:     "UnitPrice",
	Quantity:      "Quantity",
}

var InvoiceLineTableColumns = struct {
	InvoiceLineId string
	InvoiceId     string
	TrackId       string
	UnitPrice     string
	Quantity      string
}{
	InvoiceLineId: "InvoiceLine.InvoiceLineId",
	InvoiceId:     "InvoiceLine.InvoiceId",
	TrackId:       "InvoiceLine.TrackId",
	UnitPrice:     "InvoiceLine.UnitPrice",
	Quantity:      "InvoiceLine.Quantity",
}

// Generated where

var InvoiceLineWhere = struct {
	InvoiceLineId whereHelperint64
	InvoiceId     whereHelperint64
	TrackId       whereHelperint64
	UnitPrice     whereHelperstring
	Quantity      whereHelperint64
}{
	InvoiceLineId: whereHelperint64{field: "\"InvoiceLine\".\"InvoiceLineId\""},
	InvoiceId:     whereHelperint64{field: "\"InvoiceLine\".\"InvoiceId\""},
	TrackId:       whereHelperint64{field: "\"InvoiceLine\".\"TrackId\""},
	UnitPrice:     whereHelperstring{field: "\"InvoiceLine\".\"UnitPrice\""},
	Quantity:      whereHelperint64{field: "\"InvoiceLine\".\"Quantity\""},
}

// InvoiceLineRels is where relationship names are stored.
var InvoiceLineRels = struct {
	TrackIdTrack     string
	InvoiceIdInvoice string
}{
	TrackIdTrack:     "TrackIdTrack",
	InvoiceIdInvoice: "InvoiceIdInvoice",
}

// invoiceLineR is where relationships are stored.
type invoiceLineR struct {
	TrackIdTrack     *Track   `boil:"TrackIdTrack" json:"TrackIdTrack" toml:"TrackIdTrack" yaml:"TrackIdTrack"`
	InvoiceIdInvoice *Invoice `boil:"InvoiceIdInvoice" json:"InvoiceIdInvoice" toml:"InvoiceIdInvoice" yaml:"InvoiceIdInvoice"`
}

// NewStruct creates a new relationship struct
func (*invoiceLineR) NewStruct() *invoiceLineR {
	return &invoiceLineR{}
}

// invoiceLineL is where Load methods for each relationship are stored.
type invoiceLineL struct{}

var (
	invoiceLineAllColumns            = []string{"InvoiceLineId", "InvoiceId", "TrackId", "UnitPrice", "Quantity"}
	invoiceLineColumnsWithoutDefault = []string{"InvoiceId", "TrackId", "UnitPrice", "Quantity"}
	invoiceLineColumnsWithDefault    = []string{"InvoiceLineId"}
	invoiceLinePrimaryKeyColumns     = []string{"InvoiceLineId"}
)

type (
	// InvoiceLineSlice is an alias for a slice of pointers to InvoiceLine.
	// This should almost always be used instead of []InvoiceLine.
	InvoiceLineSlice []*InvoiceLine
	// InvoiceLineHook is the signature for custom InvoiceLine hook methods
	InvoiceLineHook func(context.Context, boil.ContextExecutor, *InvoiceLine) error

	invoiceLineQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	invoiceLineType                 = reflect.TypeOf(&InvoiceLine{})
	invoiceLineMapping              = queries.MakeStructMapping(invoiceLineType)
	invoiceLinePrimaryKeyMapping, _ = queries.BindMapping(invoiceLineType, invoiceLineMapping, invoiceLinePrimaryKeyColumns)
	invoiceLineInsertCacheMut       sync.RWMutex
	invoiceLineInsertCache          = make(map[string]insertCache)
	invoiceLineUpdateCacheMut       sync.RWMutex
	invoiceLineUpdateCache          = make(map[string]updateCache)
	invoiceLineUpsertCacheMut       sync.RWMutex
	invoiceLineUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var invoiceLineBeforeInsertHooks []InvoiceLineHook
var invoiceLineBeforeUpdateHooks []InvoiceLineHook
var invoiceLineBeforeDeleteHooks []InvoiceLineHook
var invoiceLineBeforeUpsertHooks []InvoiceLineHook

var invoiceLineAfterInsertHooks []InvoiceLineHook
var invoiceLineAfterSelectHooks []InvoiceLineHook
var invoiceLineAfterUpdateHooks []InvoiceLineHook
var invoiceLineAfterDeleteHooks []InvoiceLineHook
var invoiceLineAfterUpsertHooks []InvoiceLineHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *InvoiceLine) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *InvoiceLine) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *InvoiceLine) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *InvoiceLine) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *InvoiceLine) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *InvoiceLine) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *InvoiceLine) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *InvoiceLine) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *InvoiceLine) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range invoiceLineAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddInvoiceLineHook registers your hook function for all future operations.
func AddInvoiceLineHook(hookPoint boil.HookPoint, invoiceLineHook InvoiceLineHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		invoiceLineBeforeInsertHooks = append(invoiceLineBeforeInsertHooks, invoiceLineHook)
	case boil.BeforeUpdateHook:
		invoiceLineBeforeUpdateHooks = append(invoiceLineBeforeUpdateHooks, invoiceLineHook)
	case boil.BeforeDeleteHook:
		invoiceLineBeforeDeleteHooks = append(invoiceLineBeforeDeleteHooks, invoiceLineHook)
	case boil.BeforeUpsertHook:
		invoiceLineBeforeUpsertHooks = append(invoiceLineBeforeUpsertHooks, invoiceLineHook)
	case boil.AfterInsertHook:
		invoiceLineAfterInsertHooks = append(invoiceLineAfterInsertHooks, invoiceLineHook)
	case boil.AfterSelectHook:
		invoiceLineAfterSelectHooks = append(invoiceLineAfterSelectHooks, invoiceLineHook)
	case boil.AfterUpdateHook:
		invoiceLineAfterUpdateHooks = append(invoiceLineAfterUpdateHooks, invoiceLineHook)
	case boil.AfterDeleteHook:
		invoiceLineAfterDeleteHooks = append(invoiceLineAfterDeleteHooks, invoiceLineHook)
	case boil.AfterUpsertHook:
		invoiceLineAfterUpsertHooks = append(invoiceLineAfterUpsertHooks, invoiceLineHook)
	}
}

// One returns a single invoiceLine record from the query.
func (q invoiceLineQuery) One(ctx context.Context, exec boil.ContextExecutor) (*InvoiceLine, error) {
	o := &InvoiceLine{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for InvoiceLine")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all InvoiceLine records from the query.
func (q invoiceLineQuery) All(ctx context.Context, exec boil.ContextExecutor) (InvoiceLineSlice, error) {
	var o []*InvoiceLine

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to InvoiceLine slice")
	}

	if len(invoiceLineAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all InvoiceLine records in the query.
func (q invoiceLineQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count InvoiceLine rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q invoiceLineQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if InvoiceLine exists")
	}

	return count > 0, nil
}

// TrackIdTrack pointed to by the foreign key.
func (o *InvoiceLine) TrackIdTrack(mods ...qm.QueryMod) trackQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"TrackId\" = ?", o.TrackId),
	}

	queryMods = append(queryMods, mods...)

	query := Tracks(queryMods...)
	queries.SetFrom(query.Query, "\"Track\"")

	return query
}

// InvoiceIdInvoice pointed to by the foreign key.
func (o *InvoiceLine) InvoiceIdInvoice(mods ...qm.QueryMod) invoiceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"InvoiceId\" = ?", o.InvoiceId),
	}

	queryMods = append(queryMods, mods...)

	query := Invoices(queryMods...)
	queries.SetFrom(query.Query, "\"Invoice\"")

	return query
}

// LoadTrackIdTrack allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (invoiceLineL) LoadTrackIdTrack(ctx context.Context, e boil.ContextExecutor, singular bool, maybeInvoiceLine interface{}, mods queries.Applicator) error {
	var slice []*InvoiceLine
	var object *InvoiceLine

	if singular {
		object = maybeInvoiceLine.(*InvoiceLine)
	} else {
		slice = *maybeInvoiceLine.(*[]*InvoiceLine)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &invoiceLineR{}
		}
		args = append(args, object.TrackId)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &invoiceLineR{}
			}

			for _, a := range args {
				if a == obj.TrackId {
					continue Outer
				}
			}

			args = append(args, obj.TrackId)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`Track`),
		qm.WhereIn(`Track.TrackId in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Track")
	}

	var resultSlice []*Track
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Track")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for Track")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Track")
	}

	if len(invoiceLineAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.TrackIdTrack = foreign
		if foreign.R == nil {
			foreign.R = &trackR{}
		}
		foreign.R.TrackIdInvoiceLines = append(foreign.R.TrackIdInvoiceLines, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TrackId == foreign.TrackId {
				local.R.TrackIdTrack = foreign
				if foreign.R == nil {
					foreign.R = &trackR{}
				}
				foreign.R.TrackIdInvoiceLines = append(foreign.R.TrackIdInvoiceLines, local)
				break
			}
		}
	}

	return nil
}

// LoadInvoiceIdInvoice allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (invoiceLineL) LoadInvoiceIdInvoice(ctx context.Context, e boil.ContextExecutor, singular bool, maybeInvoiceLine interface{}, mods queries.Applicator) error {
	var slice []*InvoiceLine
	var object *InvoiceLine

	if singular {
		object = maybeInvoiceLine.(*InvoiceLine)
	} else {
		slice = *maybeInvoiceLine.(*[]*InvoiceLine)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &invoiceLineR{}
		}
		args = append(args, object.InvoiceId)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &invoiceLineR{}
			}

			for _, a := range args {
				if a == obj.InvoiceId {
					continue Outer
				}
			}

			args = append(args, obj.InvoiceId)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`Invoice`),
		qm.WhereIn(`Invoice.InvoiceId in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Invoice")
	}

	var resultSlice []*Invoice
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Invoice")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for Invoice")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Invoice")
	}

	if len(invoiceLineAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.InvoiceIdInvoice = foreign
		if foreign.R == nil {
			foreign.R = &invoiceR{}
		}
		foreign.R.InvoiceIdInvoiceLines = append(foreign.R.InvoiceIdInvoiceLines, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.InvoiceId == foreign.InvoiceId {
				local.R.InvoiceIdInvoice = foreign
				if foreign.R == nil {
					foreign.R = &invoiceR{}
				}
				foreign.R.InvoiceIdInvoiceLines = append(foreign.R.InvoiceIdInvoiceLines, local)
				break
			}
		}
	}

	return nil
}

// SetTrackIdTrack of the invoiceLine to the related item.
// Sets o.R.TrackIdTrack to related.
// Adds o to related.R.TrackIdInvoiceLines.
func (o *InvoiceLine) SetTrackIdTrack(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Track) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"InvoiceLine\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"TrackId"}),
		strmangle.WhereClause("\"", "\"", 0, invoiceLinePrimaryKeyColumns),
	)
	values := []interface{}{related.TrackId, o.InvoiceLineId}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TrackId = related.TrackId
	if o.R == nil {
		o.R = &invoiceLineR{
			TrackIdTrack: related,
		}
	} else {
		o.R.TrackIdTrack = related
	}

	if related.R == nil {
		related.R = &trackR{
			TrackIdInvoiceLines: InvoiceLineSlice{o},
		}
	} else {
		related.R.TrackIdInvoiceLines = append(related.R.TrackIdInvoiceLines, o)
	}

	return nil
}

// SetInvoiceIdInvoice of the invoiceLine to the related item.
// Sets o.R.InvoiceIdInvoice to related.
// Adds o to related.R.InvoiceIdInvoiceLines.
func (o *InvoiceLine) SetInvoiceIdInvoice(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Invoice) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"InvoiceLine\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, []string{"InvoiceId"}),
		strmangle.WhereClause("\"", "\"", 0, invoiceLinePrimaryKeyColumns),
	)
	values := []interface{}{related.InvoiceId, o.InvoiceLineId}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.InvoiceId = related.InvoiceId
	if o.R == nil {
		o.R = &invoiceLineR{
			InvoiceIdInvoice: related,
		}
	} else {
		o.R.InvoiceIdInvoice = related
	}

	if related.R == nil {
		related.R = &invoiceR{
			InvoiceIdInvoiceLines: InvoiceLineSlice{o},
		}
	} else {
		related.R.InvoiceIdInvoiceLines = append(related.R.InvoiceIdInvoiceLines, o)
	}

	return nil
}

// InvoiceLines retrieves all the records using an executor.
func InvoiceLines(mods ...qm.QueryMod) invoiceLineQuery {
	mods = append(mods, qm.From("\"InvoiceLine\""))
	return invoiceLineQuery{NewQuery(mods...)}
}

// FindInvoiceLine retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindInvoiceLine(ctx context.Context, exec boil.ContextExecutor, invoiceLineId int64, selectCols ...string) (*InvoiceLine, error) {
	invoiceLineObj := &InvoiceLine{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"InvoiceLine\" where \"InvoiceLineId\"=?", sel,
	)

	q := queries.Raw(query, invoiceLineId)

	err := q.Bind(ctx, exec, invoiceLineObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from InvoiceLine")
	}

	if err = invoiceLineObj.doAfterSelectHooks(ctx, exec); err != nil {
		return invoiceLineObj, err
	}

	return invoiceLineObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *InvoiceLine) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no InvoiceLine provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(invoiceLineColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	invoiceLineInsertCacheMut.RLock()
	cache, cached := invoiceLineInsertCache[key]
	invoiceLineInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			invoiceLineAllColumns,
			invoiceLineColumnsWithDefault,
			invoiceLineColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(invoiceLineType, invoiceLineMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(invoiceLineType, invoiceLineMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"InvoiceLine\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"InvoiceLine\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into InvoiceLine")
	}

	if !cached {
		invoiceLineInsertCacheMut.Lock()
		invoiceLineInsertCache[key] = cache
		invoiceLineInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the InvoiceLine.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *InvoiceLine) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	invoiceLineUpdateCacheMut.RLock()
	cache, cached := invoiceLineUpdateCache[key]
	invoiceLineUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			invoiceLineAllColumns,
			invoiceLinePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update InvoiceLine, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"InvoiceLine\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, invoiceLinePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(invoiceLineType, invoiceLineMapping, append(wl, invoiceLinePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update InvoiceLine row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for InvoiceLine")
	}

	if !cached {
		invoiceLineUpdateCacheMut.Lock()
		invoiceLineUpdateCache[key] = cache
		invoiceLineUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q invoiceLineQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for InvoiceLine")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for InvoiceLine")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o InvoiceLineSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invoiceLinePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"InvoiceLine\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, invoiceLinePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in invoiceLine slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all invoiceLine")
	}
	return rowsAff, nil
}

// Delete deletes a single InvoiceLine record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *InvoiceLine) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no InvoiceLine provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), invoiceLinePrimaryKeyMapping)
	sql := "DELETE FROM \"InvoiceLine\" WHERE \"InvoiceLineId\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from InvoiceLine")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for InvoiceLine")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q invoiceLineQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no invoiceLineQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from InvoiceLine")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for InvoiceLine")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o InvoiceLineSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(invoiceLineBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invoiceLinePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"InvoiceLine\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, invoiceLinePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from invoiceLine slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for InvoiceLine")
	}

	if len(invoiceLineAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *InvoiceLine) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindInvoiceLine(ctx, exec, o.InvoiceLineId)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InvoiceLineSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := InvoiceLineSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), invoiceLinePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"InvoiceLine\".* FROM \"InvoiceLine\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, invoiceLinePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in InvoiceLineSlice")
	}

	*o = slice

	return nil
}

// InvoiceLineExists checks if the InvoiceLine row exists.
func InvoiceLineExists(ctx context.Context, exec boil.ContextExecutor, invoiceLineId int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"InvoiceLine\" where \"InvoiceLineId\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, invoiceLineId)
	}
	row := exec.QueryRowContext(ctx, sql, invoiceLineId)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if InvoiceLine exists")
	}

	return exists, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *InvoiceLine) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no InvoiceLine provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(invoiceLineColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	invoiceLineUpsertCacheMut.RLock()
	cache, cached := invoiceLineUpsertCache[key]
	invoiceLineUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			invoiceLineAllColumns,
			invoiceLineColumnsWithDefault,
			invoiceLineColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			invoiceLineAllColumns,
			invoiceLinePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert InvoiceLine, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(invoiceLinePrimaryKeyColumns))
			copy(conflict, invoiceLinePrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"InvoiceLine\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(invoiceLineType, invoiceLineMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(invoiceLineType, invoiceLineMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert InvoiceLine")
	}

	if !cached {
		invoiceLineUpsertCacheMut.Lock()
		invoiceLineUpsertCache[key] = cache
		invoiceLineUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}
