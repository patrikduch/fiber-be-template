// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"fiber-be-template/ent/migrate"

	"fiber-be-template/ent/role"
	"fiber-be-template/ent/user"
	"fiber-be-template/ent/userrole"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Role is the client for interacting with the Role builders.
	Role *RoleClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// UserRole is the client for interacting with the UserRole builders.
	UserRole *UserRoleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Role = NewRoleClient(c.config)
	c.User = NewUserClient(c.config)
	c.UserRole = NewUserRoleClient(c.config)
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
		Role:     NewRoleClient(cfg),
		User:     NewUserClient(cfg),
		UserRole: NewUserRoleClient(cfg),
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
		Role:     NewRoleClient(cfg),
		User:     NewUserClient(cfg),
		UserRole: NewUserRoleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Role.
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
	c.Role.Use(hooks...)
	c.User.Use(hooks...)
	c.UserRole.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Role.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
	c.UserRole.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *RoleMutation:
		return c.Role.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	case *UserRoleMutation:
		return c.UserRole.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// RoleClient is a client for the Role schema.
type RoleClient struct {
	config
}

// NewRoleClient returns a client for the Role from the given config.
func NewRoleClient(c config) *RoleClient {
	return &RoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `role.Hooks(f(g(h())))`.
func (c *RoleClient) Use(hooks ...Hook) {
	c.hooks.Role = append(c.hooks.Role, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `role.Intercept(f(g(h())))`.
func (c *RoleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Role = append(c.inters.Role, interceptors...)
}

// Create returns a builder for creating a Role entity.
func (c *RoleClient) Create() *RoleCreate {
	mutation := newRoleMutation(c.config, OpCreate)
	return &RoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Role entities.
func (c *RoleClient) CreateBulk(builders ...*RoleCreate) *RoleCreateBulk {
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RoleClient) MapCreateBulk(slice any, setFunc func(*RoleCreate, int)) *RoleCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RoleCreateBulk{err: fmt.Errorf("calling to RoleClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RoleCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Role.
func (c *RoleClient) Update() *RoleUpdate {
	mutation := newRoleMutation(c.config, OpUpdate)
	return &RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoleClient) UpdateOne(r *Role) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRole(r))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoleClient) UpdateOneID(id uuid.UUID) *RoleUpdateOne {
	mutation := newRoleMutation(c.config, OpUpdateOne, withRoleID(id))
	return &RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Role.
func (c *RoleClient) Delete() *RoleDelete {
	mutation := newRoleMutation(c.config, OpDelete)
	return &RoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RoleClient) DeleteOne(r *Role) *RoleDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RoleClient) DeleteOneID(id uuid.UUID) *RoleDeleteOne {
	builder := c.Delete().Where(role.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoleDeleteOne{builder}
}

// Query returns a query builder for Role.
func (c *RoleClient) Query() *RoleQuery {
	return &RoleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRole},
		inters: c.Interceptors(),
	}
}

// Get returns a Role entity by its id.
func (c *RoleClient) Get(ctx context.Context, id uuid.UUID) (*Role, error) {
	return c.Query().Where(role.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoleClient) GetX(ctx context.Context, id uuid.UUID) *Role {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUserRoles queries the user_roles edge of a Role.
func (c *RoleClient) QueryUserRoles(r *Role) *UserRoleQuery {
	query := (&UserRoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(role.Table, role.FieldID, id),
			sqlgraph.To(userrole.Table, userrole.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, role.UserRolesTable, role.UserRolesColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoleClient) Hooks() []Hook {
	return c.hooks.Role
}

// Interceptors returns the client interceptors.
func (c *RoleClient) Interceptors() []Interceptor {
	return c.inters.Role
}

func (c *RoleClient) mutate(ctx context.Context, m *RoleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RoleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RoleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RoleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Role mutation op: %q", m.Op())
	}
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

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
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
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUserRoles queries the user_roles edge of a User.
func (c *UserClient) QueryUserRoles(u *User) *UserRoleQuery {
	query := (&UserRoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(userrole.Table, userrole.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.UserRolesTable, user.UserRolesColumn),
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

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// UserRoleClient is a client for the UserRole schema.
type UserRoleClient struct {
	config
}

// NewUserRoleClient returns a client for the UserRole from the given config.
func NewUserRoleClient(c config) *UserRoleClient {
	return &UserRoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `userrole.Hooks(f(g(h())))`.
func (c *UserRoleClient) Use(hooks ...Hook) {
	c.hooks.UserRole = append(c.hooks.UserRole, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `userrole.Intercept(f(g(h())))`.
func (c *UserRoleClient) Intercept(interceptors ...Interceptor) {
	c.inters.UserRole = append(c.inters.UserRole, interceptors...)
}

// Create returns a builder for creating a UserRole entity.
func (c *UserRoleClient) Create() *UserRoleCreate {
	mutation := newUserRoleMutation(c.config, OpCreate)
	return &UserRoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of UserRole entities.
func (c *UserRoleClient) CreateBulk(builders ...*UserRoleCreate) *UserRoleCreateBulk {
	return &UserRoleCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserRoleClient) MapCreateBulk(slice any, setFunc func(*UserRoleCreate, int)) *UserRoleCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserRoleCreateBulk{err: fmt.Errorf("calling to UserRoleClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserRoleCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserRoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for UserRole.
func (c *UserRoleClient) Update() *UserRoleUpdate {
	mutation := newUserRoleMutation(c.config, OpUpdate)
	return &UserRoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserRoleClient) UpdateOne(ur *UserRole) *UserRoleUpdateOne {
	mutation := newUserRoleMutation(c.config, OpUpdateOne, withUserRole(ur))
	return &UserRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserRoleClient) UpdateOneID(id uuid.UUID) *UserRoleUpdateOne {
	mutation := newUserRoleMutation(c.config, OpUpdateOne, withUserRoleID(id))
	return &UserRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for UserRole.
func (c *UserRoleClient) Delete() *UserRoleDelete {
	mutation := newUserRoleMutation(c.config, OpDelete)
	return &UserRoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserRoleClient) DeleteOne(ur *UserRole) *UserRoleDeleteOne {
	return c.DeleteOneID(ur.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserRoleClient) DeleteOneID(id uuid.UUID) *UserRoleDeleteOne {
	builder := c.Delete().Where(userrole.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserRoleDeleteOne{builder}
}

// Query returns a query builder for UserRole.
func (c *UserRoleClient) Query() *UserRoleQuery {
	return &UserRoleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUserRole},
		inters: c.Interceptors(),
	}
}

// Get returns a UserRole entity by its id.
func (c *UserRoleClient) Get(ctx context.Context, id uuid.UUID) (*UserRole, error) {
	return c.Query().Where(userrole.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserRoleClient) GetX(ctx context.Context, id uuid.UUID) *UserRole {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a UserRole.
func (c *UserRoleClient) QueryUser(ur *UserRole) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ur.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(userrole.Table, userrole.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userrole.UserTable, userrole.UserColumn),
		)
		fromV = sqlgraph.Neighbors(ur.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRole queries the role edge of a UserRole.
func (c *UserRoleClient) QueryRole(ur *UserRole) *RoleQuery {
	query := (&RoleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ur.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(userrole.Table, userrole.FieldID, id),
			sqlgraph.To(role.Table, role.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userrole.RoleTable, userrole.RoleColumn),
		)
		fromV = sqlgraph.Neighbors(ur.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserRoleClient) Hooks() []Hook {
	return c.hooks.UserRole
}

// Interceptors returns the client interceptors.
func (c *UserRoleClient) Interceptors() []Interceptor {
	return c.inters.UserRole
}

func (c *UserRoleClient) mutate(ctx context.Context, m *UserRoleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserRoleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserRoleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserRoleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown UserRole mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Role, User, UserRole []ent.Hook
	}
	inters struct {
		Role, User, UserRole []ent.Interceptor
	}
)
