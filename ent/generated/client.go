// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/kx-boutique/ent/generated/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kx-boutique/ent/generated/cart"
	"github.com/kx-boutique/ent/generated/cartitem"
	"github.com/kx-boutique/ent/generated/product"
	"github.com/kx-boutique/ent/generated/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Cart is the client for interacting with the Cart builders.
	Cart *CartClient
	// CartItem is the client for interacting with the CartItem builders.
	CartItem *CartItemClient
	// Product is the client for interacting with the Product builders.
	Product *ProductClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Cart = NewCartClient(c.config)
	c.CartItem = NewCartItemClient(c.config)
	c.Product = NewProductClient(c.config)
	c.User = NewUserClient(c.config)
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
var ErrTxStarted = errors.New("generated: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("generated: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:      ctx,
		config:   cfg,
		Cart:     NewCartClient(cfg),
		CartItem: NewCartItemClient(cfg),
		Product:  NewProductClient(cfg),
		User:     NewUserClient(cfg),
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
		Cart:     NewCartClient(cfg),
		CartItem: NewCartItemClient(cfg),
		Product:  NewProductClient(cfg),
		User:     NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Cart.
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
	c.Cart.Use(hooks...)
	c.CartItem.Use(hooks...)
	c.Product.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Cart.Intercept(interceptors...)
	c.CartItem.Intercept(interceptors...)
	c.Product.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *CartMutation:
		return c.Cart.mutate(ctx, m)
	case *CartItemMutation:
		return c.CartItem.mutate(ctx, m)
	case *ProductMutation:
		return c.Product.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("generated: unknown mutation type %T", m)
	}
}

// CartClient is a client for the Cart schema.
type CartClient struct {
	config
}

// NewCartClient returns a client for the Cart from the given config.
func NewCartClient(c config) *CartClient {
	return &CartClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cart.Hooks(f(g(h())))`.
func (c *CartClient) Use(hooks ...Hook) {
	c.hooks.Cart = append(c.hooks.Cart, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `cart.Intercept(f(g(h())))`.
func (c *CartClient) Intercept(interceptors ...Interceptor) {
	c.inters.Cart = append(c.inters.Cart, interceptors...)
}

// Create returns a builder for creating a Cart entity.
func (c *CartClient) Create() *CartCreate {
	mutation := newCartMutation(c.config, OpCreate)
	return &CartCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Cart entities.
func (c *CartClient) CreateBulk(builders ...*CartCreate) *CartCreateBulk {
	return &CartCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CartClient) MapCreateBulk(slice any, setFunc func(*CartCreate, int)) *CartCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CartCreateBulk{err: fmt.Errorf("calling to CartClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CartCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CartCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Cart.
func (c *CartClient) Update() *CartUpdate {
	mutation := newCartMutation(c.config, OpUpdate)
	return &CartUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CartClient) UpdateOne(ca *Cart) *CartUpdateOne {
	mutation := newCartMutation(c.config, OpUpdateOne, withCart(ca))
	return &CartUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CartClient) UpdateOneID(id uuid.UUID) *CartUpdateOne {
	mutation := newCartMutation(c.config, OpUpdateOne, withCartID(id))
	return &CartUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Cart.
func (c *CartClient) Delete() *CartDelete {
	mutation := newCartMutation(c.config, OpDelete)
	return &CartDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CartClient) DeleteOne(ca *Cart) *CartDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CartClient) DeleteOneID(id uuid.UUID) *CartDeleteOne {
	builder := c.Delete().Where(cart.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CartDeleteOne{builder}
}

// Query returns a query builder for Cart.
func (c *CartClient) Query() *CartQuery {
	return &CartQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCart},
		inters: c.Interceptors(),
	}
}

// Get returns a Cart entity by its id.
func (c *CartClient) Get(ctx context.Context, id uuid.UUID) (*Cart, error) {
	return c.Query().Where(cart.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CartClient) GetX(ctx context.Context, id uuid.UUID) *Cart {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUserIDOwner queries the user_id_owner edge of a Cart.
func (c *CartClient) QueryUserIDOwner(ca *Cart) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cart.Table, cart.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, cart.UserIDOwnerTable, cart.UserIDOwnerColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCartItem queries the cart_item edge of a Cart.
func (c *CartClient) QueryCartItem(ca *Cart) *CartItemQuery {
	query := (&CartItemClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ca.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cart.Table, cart.FieldID, id),
			sqlgraph.To(cartitem.Table, cartitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, cart.CartItemTable, cart.CartItemColumn),
		)
		fromV = sqlgraph.Neighbors(ca.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CartClient) Hooks() []Hook {
	return c.hooks.Cart
}

// Interceptors returns the client interceptors.
func (c *CartClient) Interceptors() []Interceptor {
	return c.inters.Cart
}

func (c *CartClient) mutate(ctx context.Context, m *CartMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CartCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CartUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CartUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CartDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Cart mutation op: %q", m.Op())
	}
}

// CartItemClient is a client for the CartItem schema.
type CartItemClient struct {
	config
}

// NewCartItemClient returns a client for the CartItem from the given config.
func NewCartItemClient(c config) *CartItemClient {
	return &CartItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cartitem.Hooks(f(g(h())))`.
func (c *CartItemClient) Use(hooks ...Hook) {
	c.hooks.CartItem = append(c.hooks.CartItem, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `cartitem.Intercept(f(g(h())))`.
func (c *CartItemClient) Intercept(interceptors ...Interceptor) {
	c.inters.CartItem = append(c.inters.CartItem, interceptors...)
}

// Create returns a builder for creating a CartItem entity.
func (c *CartItemClient) Create() *CartItemCreate {
	mutation := newCartItemMutation(c.config, OpCreate)
	return &CartItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CartItem entities.
func (c *CartItemClient) CreateBulk(builders ...*CartItemCreate) *CartItemCreateBulk {
	return &CartItemCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *CartItemClient) MapCreateBulk(slice any, setFunc func(*CartItemCreate, int)) *CartItemCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &CartItemCreateBulk{err: fmt.Errorf("calling to CartItemClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*CartItemCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &CartItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CartItem.
func (c *CartItemClient) Update() *CartItemUpdate {
	mutation := newCartItemMutation(c.config, OpUpdate)
	return &CartItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CartItemClient) UpdateOne(ci *CartItem) *CartItemUpdateOne {
	mutation := newCartItemMutation(c.config, OpUpdateOne, withCartItem(ci))
	return &CartItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CartItemClient) UpdateOneID(id uuid.UUID) *CartItemUpdateOne {
	mutation := newCartItemMutation(c.config, OpUpdateOne, withCartItemID(id))
	return &CartItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CartItem.
func (c *CartItemClient) Delete() *CartItemDelete {
	mutation := newCartItemMutation(c.config, OpDelete)
	return &CartItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CartItemClient) DeleteOne(ci *CartItem) *CartItemDeleteOne {
	return c.DeleteOneID(ci.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CartItemClient) DeleteOneID(id uuid.UUID) *CartItemDeleteOne {
	builder := c.Delete().Where(cartitem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CartItemDeleteOne{builder}
}

// Query returns a query builder for CartItem.
func (c *CartItemClient) Query() *CartItemQuery {
	return &CartItemQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCartItem},
		inters: c.Interceptors(),
	}
}

// Get returns a CartItem entity by its id.
func (c *CartItemClient) Get(ctx context.Context, id uuid.UUID) (*CartItem, error) {
	return c.Query().Where(cartitem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CartItemClient) GetX(ctx context.Context, id uuid.UUID) *CartItem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCartIDOwner queries the cart_id_owner edge of a CartItem.
func (c *CartItemClient) QueryCartIDOwner(ci *CartItem) *CartQuery {
	query := (&CartClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ci.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cartitem.Table, cartitem.FieldID, id),
			sqlgraph.To(cart.Table, cart.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cartitem.CartIDOwnerTable, cartitem.CartIDOwnerColumn),
		)
		fromV = sqlgraph.Neighbors(ci.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProductIDOwner queries the product_id_owner edge of a CartItem.
func (c *CartItemClient) QueryProductIDOwner(ci *CartItem) *ProductQuery {
	query := (&ProductClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ci.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(cartitem.Table, cartitem.FieldID, id),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cartitem.ProductIDOwnerTable, cartitem.ProductIDOwnerColumn),
		)
		fromV = sqlgraph.Neighbors(ci.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CartItemClient) Hooks() []Hook {
	return c.hooks.CartItem
}

// Interceptors returns the client interceptors.
func (c *CartItemClient) Interceptors() []Interceptor {
	return c.inters.CartItem
}

func (c *CartItemClient) mutate(ctx context.Context, m *CartItemMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&CartItemCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&CartItemUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&CartItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&CartItemDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown CartItem mutation op: %q", m.Op())
	}
}

// ProductClient is a client for the Product schema.
type ProductClient struct {
	config
}

// NewProductClient returns a client for the Product from the given config.
func NewProductClient(c config) *ProductClient {
	return &ProductClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `product.Hooks(f(g(h())))`.
func (c *ProductClient) Use(hooks ...Hook) {
	c.hooks.Product = append(c.hooks.Product, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `product.Intercept(f(g(h())))`.
func (c *ProductClient) Intercept(interceptors ...Interceptor) {
	c.inters.Product = append(c.inters.Product, interceptors...)
}

// Create returns a builder for creating a Product entity.
func (c *ProductClient) Create() *ProductCreate {
	mutation := newProductMutation(c.config, OpCreate)
	return &ProductCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Product entities.
func (c *ProductClient) CreateBulk(builders ...*ProductCreate) *ProductCreateBulk {
	return &ProductCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ProductClient) MapCreateBulk(slice any, setFunc func(*ProductCreate, int)) *ProductCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ProductCreateBulk{err: fmt.Errorf("calling to ProductClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ProductCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ProductCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Product.
func (c *ProductClient) Update() *ProductUpdate {
	mutation := newProductMutation(c.config, OpUpdate)
	return &ProductUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProductClient) UpdateOne(pr *Product) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProduct(pr))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProductClient) UpdateOneID(id uuid.UUID) *ProductUpdateOne {
	mutation := newProductMutation(c.config, OpUpdateOne, withProductID(id))
	return &ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Product.
func (c *ProductClient) Delete() *ProductDelete {
	mutation := newProductMutation(c.config, OpDelete)
	return &ProductDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ProductClient) DeleteOne(pr *Product) *ProductDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ProductClient) DeleteOneID(id uuid.UUID) *ProductDeleteOne {
	builder := c.Delete().Where(product.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProductDeleteOne{builder}
}

// Query returns a query builder for Product.
func (c *ProductClient) Query() *ProductQuery {
	return &ProductQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeProduct},
		inters: c.Interceptors(),
	}
}

// Get returns a Product entity by its id.
func (c *ProductClient) Get(ctx context.Context, id uuid.UUID) (*Product, error) {
	return c.Query().Where(product.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProductClient) GetX(ctx context.Context, id uuid.UUID) *Product {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCartItem queries the cart_item edge of a Product.
func (c *ProductClient) QueryCartItem(pr *Product) *CartItemQuery {
	query := (&CartItemClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, id),
			sqlgraph.To(cartitem.Table, cartitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, product.CartItemTable, product.CartItemColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProductClient) Hooks() []Hook {
	return c.hooks.Product
}

// Interceptors returns the client interceptors.
func (c *ProductClient) Interceptors() []Interceptor {
	return c.inters.Product
}

func (c *ProductClient) mutate(ctx context.Context, m *ProductMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ProductCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ProductUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ProductUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ProductDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Product mutation op: %q", m.Op())
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

// QueryCart queries the cart edge of a User.
func (c *UserClient) QueryCart(u *User) *CartQuery {
	query := (&CartClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(cart.Table, cart.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.CartTable, user.CartColumn),
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
		return nil, fmt.Errorf("generated: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Cart, CartItem, Product, User []ent.Hook
	}
	inters struct {
		Cart, CartItem, Product, User []ent.Interceptor
	}
)
