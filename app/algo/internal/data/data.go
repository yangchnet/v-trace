package data

import (
	"context"
	"database/sql"
	"fmt"

	"gitee.com/qciip-icp/v-trace/app/algo/internal/conf"
	"gitee.com/qciip-icp/v-trace/app/algo/internal/data/db"
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

type Store interface {
	db.Querier
	ExecTx(ctx context.Context, fn func(*db.Queries) error) error
}

type DBStore struct {
	db *sql.DB
	*db.Queries
}

type Data struct {
	Store
	cache cache.Cache
}

var _ Store = (*Data)(nil)

// NewUserRepo creates a new data which is a biz.UserRepo.
func NewData(ctx context.Context, c *conf.DataConfig) *Data {
	db, queries, err := newDB(ctx, &c.Db)
	if err != nil {
		panic(err)
	}

	store := &DBStore{
		db:      db,
		Queries: queries,
	}

	cache := cache.NewRedisStore(&c.Redis)
	return &Data{
		Store: store,
		cache: cache,
	}
}

// ExecTx executes a function within a database transaction.
func (d *DBStore) ExecTx(ctx context.Context, fn func(*db.Queries) error) error {
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
