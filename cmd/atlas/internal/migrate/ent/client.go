// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/s-sokolko/atlas/cmd/atlas/internal/migrate/ent/migrate"

	"github.com/s-sokolko/atlas/cmd/atlas/internal/migrate/ent/revision"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"

	stdsql "database/sql"

	"github.com/s-sokolko/atlas/cmd/atlas/internal/migrate/ent/internal"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Revision is the client for interacting with the Revision builders.
	Revision *RevisionClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Revision = NewRevisionClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
		// schemaConfig contains alternative names for all tables.
		schemaConfig SchemaConfig
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Revision: NewRevisionClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Revision: NewRevisionClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Revision.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Revision.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Revision.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *RevisionMutation:
		return c.Revision.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// RevisionClient is a client for the Revision schema.
type RevisionClient struct {
	config
}

// NewRevisionClient returns a client for the Revision from the given config.
func NewRevisionClient(c config) *RevisionClient {
	return &RevisionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `revision.Hooks(f(g(h())))`.
func (c *RevisionClient) Use(hooks ...Hook) {
	c.hooks.Revision = append(c.hooks.Revision, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `revision.Intercept(f(g(h())))`.
func (c *RevisionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Revision = append(c.inters.Revision, interceptors...)
}

// Create returns a builder for creating a Revision entity.
func (c *RevisionClient) Create() *RevisionCreate {
	mutation := newRevisionMutation(c.config, OpCreate)
	return &RevisionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Revision entities.
func (c *RevisionClient) CreateBulk(builders ...*RevisionCreate) *RevisionCreateBulk {
	return &RevisionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RevisionClient) MapCreateBulk(slice any, setFunc func(*RevisionCreate, int)) *RevisionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RevisionCreateBulk{err: fmt.Errorf("calling to RevisionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RevisionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RevisionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Revision.
func (c *RevisionClient) Update() *RevisionUpdate {
	mutation := newRevisionMutation(c.config, OpUpdate)
	return &RevisionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RevisionClient) UpdateOne(r *Revision) *RevisionUpdateOne {
	mutation := newRevisionMutation(c.config, OpUpdateOne, withRevision(r))
	return &RevisionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RevisionClient) UpdateOneID(id string) *RevisionUpdateOne {
	mutation := newRevisionMutation(c.config, OpUpdateOne, withRevisionID(id))
	return &RevisionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Revision.
func (c *RevisionClient) Delete() *RevisionDelete {
	mutation := newRevisionMutation(c.config, OpDelete)
	return &RevisionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RevisionClient) DeleteOne(r *Revision) *RevisionDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RevisionClient) DeleteOneID(id string) *RevisionDeleteOne {
	builder := c.Delete().Where(revision.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RevisionDeleteOne{builder}
}

// Query returns a query builder for Revision.
func (c *RevisionClient) Query() *RevisionQuery {
	return &RevisionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRevision},
		inters: c.Interceptors(),
	}
}

// Get returns a Revision entity by its id.
func (c *RevisionClient) Get(ctx context.Context, id string) (*Revision, error) {
	return c.Query().Where(revision.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RevisionClient) GetX(ctx context.Context, id string) *Revision {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RevisionClient) Hooks() []Hook {
	return c.hooks.Revision
}

// Interceptors returns the client interceptors.
func (c *RevisionClient) Interceptors() []Interceptor {
	return c.inters.Revision
}

func (c *RevisionClient) mutate(ctx context.Context, m *RevisionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RevisionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RevisionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RevisionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RevisionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Revision mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Revision []ent.Hook
	}
	inters struct {
		Revision []ent.Interceptor
	}
)

// SchemaConfig represents alternative schema names for all tables
// that can be passed at runtime.
type SchemaConfig = internal.SchemaConfig

// AlternateSchemas allows alternate schema names to be
// passed into ent operations.
func AlternateSchema(schemaConfig SchemaConfig) Option {
	return func(c *config) {
		c.schemaConfig = schemaConfig
	}
}

// ExecContext allows calling the underlying ExecContext method of the driver if it is supported by it.
// See, database/sql#DB.ExecContext for more information.
func (c *config) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := c.driver.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the driver if it is supported by it.
// See, database/sql#DB.QueryContext for more information.
func (c *config) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := c.driver.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}
