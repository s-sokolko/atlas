// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"ariga.io/atlas/integration/entinteg/ent/migrate"

	"ariga.io/atlas/integration/entinteg/ent/activity"
	"ariga.io/atlas/integration/entinteg/ent/defaultcontainer"
	"ariga.io/atlas/integration/entinteg/ent/group"
	"ariga.io/atlas/integration/entinteg/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Activity is the client for interacting with the Activity builders.
	Activity *ActivityClient
	// DefaultContainer is the client for interacting with the DefaultContainer builders.
	DefaultContainer *DefaultContainerClient
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Activity = NewActivityClient(c.config)
	c.DefaultContainer = NewDefaultContainerClient(c.config)
	c.Group = NewGroupClient(c.config)
	c.User = NewUserClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		Activity:         NewActivityClient(cfg),
		DefaultContainer: NewDefaultContainerClient(cfg),
		Group:            NewGroupClient(cfg),
		User:             NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		config:           cfg,
		Activity:         NewActivityClient(cfg),
		DefaultContainer: NewDefaultContainerClient(cfg),
		Group:            NewGroupClient(cfg),
		User:             NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Activity.
//		Query().
//		Count(ctx)
//
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
	c.Activity.Use(hooks...)
	c.DefaultContainer.Use(hooks...)
	c.Group.Use(hooks...)
	c.User.Use(hooks...)
}

// ActivityClient is a client for the Activity schema.
type ActivityClient struct {
	config
}

// NewActivityClient returns a client for the Activity from the given config.
func NewActivityClient(c config) *ActivityClient {
	return &ActivityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `activity.Hooks(f(g(h())))`.
func (c *ActivityClient) Use(hooks ...Hook) {
	c.hooks.Activity = append(c.hooks.Activity, hooks...)
}

// Create returns a create builder for Activity.
func (c *ActivityClient) Create() *ActivityCreate {
	mutation := newActivityMutation(c.config, OpCreate)
	return &ActivityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Activity entities.
func (c *ActivityClient) CreateBulk(builders ...*ActivityCreate) *ActivityCreateBulk {
	return &ActivityCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Activity.
func (c *ActivityClient) Update() *ActivityUpdate {
	mutation := newActivityMutation(c.config, OpUpdate)
	return &ActivityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ActivityClient) UpdateOne(a *Activity) *ActivityUpdateOne {
	mutation := newActivityMutation(c.config, OpUpdateOne, withActivity(a))
	return &ActivityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ActivityClient) UpdateOneID(id int) *ActivityUpdateOne {
	mutation := newActivityMutation(c.config, OpUpdateOne, withActivityID(id))
	return &ActivityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Activity.
func (c *ActivityClient) Delete() *ActivityDelete {
	mutation := newActivityMutation(c.config, OpDelete)
	return &ActivityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ActivityClient) DeleteOne(a *Activity) *ActivityDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ActivityClient) DeleteOneID(id int) *ActivityDeleteOne {
	builder := c.Delete().Where(activity.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ActivityDeleteOne{builder}
}

// Query returns a query builder for Activity.
func (c *ActivityClient) Query() *ActivityQuery {
	return &ActivityQuery{
		config: c.config,
	}
}

// Get returns a Activity entity by its id.
func (c *ActivityClient) Get(ctx context.Context, id int) (*Activity, error) {
	return c.Query().Where(activity.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ActivityClient) GetX(ctx context.Context, id int) *Activity {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Activity.
func (c *ActivityClient) QueryUsers(a *Activity) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(activity.Table, activity.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, activity.UsersTable, activity.UsersPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ActivityClient) Hooks() []Hook {
	return c.hooks.Activity
}

// DefaultContainerClient is a client for the DefaultContainer schema.
type DefaultContainerClient struct {
	config
}

// NewDefaultContainerClient returns a client for the DefaultContainer from the given config.
func NewDefaultContainerClient(c config) *DefaultContainerClient {
	return &DefaultContainerClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `defaultcontainer.Hooks(f(g(h())))`.
func (c *DefaultContainerClient) Use(hooks ...Hook) {
	c.hooks.DefaultContainer = append(c.hooks.DefaultContainer, hooks...)
}

// Create returns a create builder for DefaultContainer.
func (c *DefaultContainerClient) Create() *DefaultContainerCreate {
	mutation := newDefaultContainerMutation(c.config, OpCreate)
	return &DefaultContainerCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of DefaultContainer entities.
func (c *DefaultContainerClient) CreateBulk(builders ...*DefaultContainerCreate) *DefaultContainerCreateBulk {
	return &DefaultContainerCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DefaultContainer.
func (c *DefaultContainerClient) Update() *DefaultContainerUpdate {
	mutation := newDefaultContainerMutation(c.config, OpUpdate)
	return &DefaultContainerUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DefaultContainerClient) UpdateOne(dc *DefaultContainer) *DefaultContainerUpdateOne {
	mutation := newDefaultContainerMutation(c.config, OpUpdateOne, withDefaultContainer(dc))
	return &DefaultContainerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DefaultContainerClient) UpdateOneID(id int) *DefaultContainerUpdateOne {
	mutation := newDefaultContainerMutation(c.config, OpUpdateOne, withDefaultContainerID(id))
	return &DefaultContainerUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DefaultContainer.
func (c *DefaultContainerClient) Delete() *DefaultContainerDelete {
	mutation := newDefaultContainerMutation(c.config, OpDelete)
	return &DefaultContainerDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DefaultContainerClient) DeleteOne(dc *DefaultContainer) *DefaultContainerDeleteOne {
	return c.DeleteOneID(dc.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DefaultContainerClient) DeleteOneID(id int) *DefaultContainerDeleteOne {
	builder := c.Delete().Where(defaultcontainer.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DefaultContainerDeleteOne{builder}
}

// Query returns a query builder for DefaultContainer.
func (c *DefaultContainerClient) Query() *DefaultContainerQuery {
	return &DefaultContainerQuery{
		config: c.config,
	}
}

// Get returns a DefaultContainer entity by its id.
func (c *DefaultContainerClient) Get(ctx context.Context, id int) (*DefaultContainer, error) {
	return c.Query().Where(defaultcontainer.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DefaultContainerClient) GetX(ctx context.Context, id int) *DefaultContainer {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DefaultContainerClient) Hooks() []Hook {
	return c.hooks.DefaultContainer
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `group.Hooks(f(g(h())))`.
func (c *GroupClient) Use(hooks ...Hook) {
	c.hooks.Group = append(c.hooks.Group, hooks...)
}

// Create returns a create builder for Group.
func (c *GroupClient) Create() *GroupCreate {
	mutation := newGroupMutation(c.config, OpCreate)
	return &GroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Group entities.
func (c *GroupClient) CreateBulk(builders ...*GroupCreate) *GroupCreateBulk {
	return &GroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	mutation := newGroupMutation(c.config, OpUpdate)
	return &GroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroup(gr))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id int) *GroupUpdateOne {
	mutation := newGroupMutation(c.config, OpUpdateOne, withGroupID(id))
	return &GroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	mutation := newGroupMutation(c.config, OpDelete)
	return &GroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GroupClient) DeleteOneID(id int) *GroupDeleteOne {
	builder := c.Delete().Where(group.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GroupDeleteOne{builder}
}

// Query returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{
		config: c.config,
	}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id int) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id int) *Group {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GroupClient) Hooks() []Hook {
	return c.hooks.Group
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryGroup queries the group edge of a User.
func (c *UserClient) QueryGroup(u *User) *GroupQuery {
	query := &GroupQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, user.GroupTable, user.GroupColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryActivities queries the activities edge of a User.
func (c *UserClient) QueryActivities(u *User) *ActivityQuery {
	query := &ActivityQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(activity.Table, activity.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.ActivitiesTable, user.ActivitiesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
