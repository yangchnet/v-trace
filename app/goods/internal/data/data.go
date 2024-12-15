package data

import (
	"context"
	"database/sql"
	"fmt"

	"gitee.com/qciip-icp/v-trace/pkg/tools/gobtools"

	"gitee.com/qciip-icp/v-trace/app/goods/internal/conf"
	"gitee.com/qciip-icp/v-trace/app/goods/internal/data/db"
	"gitee.com/qciip-icp/v-trace/pkg/cache"
	"gitee.com/qciip-icp/v-trace/pkg/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

func init() {
	db.GobRegister()
}

var ProviderSet = wire.NewSet(
	NewData,
)

func newDB(ctx context.Context, c *conf.DBConfig) (*sql.DB, *db.Queries, error) {
	conn, err := sql.Open(c.Driver, c.Dsn)
	if err != nil {
		logger.Error("Error creating database connection: %v", err)

		return nil, nil, err
	}

	queries := db.New(conn)

	return conn, queries, nil
}

type Interface interface {
	db.Querier
	CacheOp
	ExecTx(ctx context.Context, fn func(*db.Queries) error) error
}

type Data struct {
	*db.Queries
	cache cache.Cache
	db    *sql.DB
}

var _ Interface = (*Data)(nil)

// NewUserRepo creates a new data which is a biz.UserRepo.
func NewData(ctx context.Context, c *conf.DataConfig) *Data {
	db, queries, err := newDB(ctx, &c.Db)
	if err != nil {
		panic(err)
	}

	cache := cache.NewRedisStore(&c.Redis)
	return &Data{
		Queries: queries,
		cache:   cache,
		db:      db,
	}
}

// ExecTx executes a function within a database transaction.
func (d *Data) ExecTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type CacheOp interface {
	CacheCreate(ctx context.Context, key string, fn func(db.Querier) (any, error)) (any, error)
	CacheDelete(ctx context.Context, key string, fn func(db.Querier) error) error
	CacheUpdate(ctx context.Context, key string, fn func(db.Querier) error) error
	CacheGet(ctx context.Context, key string, fn func(db.Querier) (any, error)) (any, error)
}

// CacheGet
func (d *Data) CacheGet(ctx context.Context, key string, fn func(db.Querier) (any, error)) (any, error) {
	vsi, ok := d.cache.Get(ctx, key) // vsi is value string interface
	if ok {
		return gobtools.Deserialize([]byte(vsi.(string)))
	}

	value, err := fn(d)
	if err != nil {
		return nil, err
	}

	valueBytes, err := gobtools.Serialize(value)
	if err != nil {
		logger.Error(err)
		return value, nil
	}

	if err := d.cache.Set(ctx, key, valueBytes, &cache.Options{}); err != nil {
		logger.Error(err)
	}

	return value, nil
}

func (d *Data) CacheUpdate(ctx context.Context, key string, fn func(db.Querier) error) error {
	if err := fn(d); err != nil {
		return err
	}

	if err := d.cache.Delete(ctx, key); err != nil {
		logger.Error(err)
	}

	return nil
}

func (d *Data) CacheCreate(ctx context.Context, key string, fn func(db.Querier) (any, error)) (any, error) {
	value, err := fn(d)
	if err != nil {
		return nil, err
	}

	valueBytes, err := gobtools.Serialize(value)
	if err != nil {
		logger.Error(err)
		return value, nil
	}

	if err := d.cache.Set(ctx, key, valueBytes, &cache.Options{}); err != nil {
		logger.Error(err)
	}

	return value, nil
}

func (d *Data) CacheDelete(ctx context.Context, key string, fn func(db.Querier) error) error {
	if err := fn(d); err != nil {
		return err
	}

	if err := d.cache.Delete(ctx, key); err != nil {
		logger.Error(err)
	}

	return nil
}
